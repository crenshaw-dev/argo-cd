package fixture

import (
	_ "embed"
	"context"
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"

	"github.com/argoproj/argo-cd/v3/pkg/apis/application"
	"github.com/argoproj/argo-cd/v3/pkg/apis/application/v1alpha1"
	"github.com/argoproj/argo-cd/v3/util/settings"
)

//go:embed e2e-argocd-config.yaml
var e2eArgoCDConfigYAML []byte

const e2eOIDCClientSecretName = "argocd-e2e-oidc-client"
const e2eOIDCClientSecretKey = "clientSecret"

// updateArgoCDConfiguration get-or-creates the argocd-config singleton and applies updater.
func updateArgoCDConfiguration(updater func(cfg *v1alpha1.ArgoCDConfiguration) error) error {
	ctx := context.Background()
	ns := TestNamespace()
	client := AppClientset.ArgoprojV1alpha1().ArgoCDConfigurations(ns)

	cfg, err := client.Get(ctx, application.ArgoCDConfigurationName, metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		cfg = &v1alpha1.ArgoCDConfiguration{
			ObjectMeta: metav1.ObjectMeta{Name: application.ArgoCDConfigurationName},
		}
		if err := updater(cfg); err != nil {
			return err
		}
		_, err = client.Create(ctx, cfg, metav1.CreateOptions{})
		return err
	}
	if err != nil {
		return err
	}

	if err := updater(cfg); err != nil {
		return err
	}
	_, err = client.Update(ctx, cfg, metav1.UpdateOptions{})
	return err
}

// resetArgoCDConfiguration restores the e2e baseline singleton (not an empty
// spec — CRD-only mode requires mapped fields to be present).
func resetArgoCDConfiguration() error {
	return updateArgoCDConfiguration(func(cfg *v1alpha1.ArgoCDConfiguration) error {
		baseline, err := e2eBaselineArgoCDConfiguration()
		if err != nil {
			return err
		}
		cfg.Spec = baseline.Spec
		return nil
	})
}

func e2eBaselineArgoCDConfiguration() (*v1alpha1.ArgoCDConfiguration, error) {
	var cfg v1alpha1.ArgoCDConfiguration
	if err := yaml.Unmarshal(e2eArgoCDConfigYAML, &cfg); err != nil {
		return nil, fmt.Errorf("parse e2e baseline ArgoCDConfiguration: %w", err)
	}
	return &cfg, nil
}

func ensureController(cfg *v1alpha1.ArgoCDConfiguration) *v1alpha1.ControllerConfig {
	if cfg.Spec.Controller == nil {
		cfg.Spec.Controller = &v1alpha1.ControllerConfig{}
	}
	return cfg.Spec.Controller
}

func ensureControllerResource(cfg *v1alpha1.ArgoCDConfiguration) *v1alpha1.ResourceConfig {
	c := ensureController(cfg)
	if c.Resource == nil {
		c.Resource = &v1alpha1.ResourceConfig{}
	}
	return c.Resource
}

func ensureServer(cfg *v1alpha1.ArgoCDConfiguration) *v1alpha1.ServerConfig {
	if cfg.Spec.Server == nil {
		cfg.Spec.Server = &v1alpha1.ServerConfig{}
	}
	return cfg.Spec.Server
}

func ensureServerRBAC(cfg *v1alpha1.ArgoCDConfiguration) *v1alpha1.RBACConfig {
	s := ensureServer(cfg)
	if s.RBAC == nil {
		s.RBAC = &v1alpha1.RBACConfig{}
	}
	return s.RBAC
}

func ensureRepoServer(cfg *v1alpha1.ArgoCDConfiguration) *v1alpha1.RepoServerConfig {
	if cfg.Spec.RepoServer == nil {
		cfg.Spec.RepoServer = &v1alpha1.RepoServerConfig{}
	}
	return cfg.Spec.RepoServer
}

func ensureCluster(cfg *v1alpha1.ArgoCDConfiguration) *v1alpha1.ClusterRegistrationConfig {
	if cfg.Spec.Cluster == nil {
		cfg.Spec.Cluster = &v1alpha1.ClusterRegistrationConfig{}
	}
	return cfg.Spec.Cluster
}

func splitCSV(value string) []string {
	if value == "" {
		return nil
	}
	parts := strings.Split(value, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func splitGroupKind(key string) (group, kind string) {
	if key == "*/*" {
		return "*", "*"
	}
	parts := strings.SplitN(key, "/", 2)
	if len(parts) == 1 {
		return "", parts[0]
	}
	return parts[0], parts[1]
}

func resourceOverridesToConfig(overrides map[string]v1alpha1.ResourceOverride) (*v1alpha1.ResourceConfig, error) {
	rc := &v1alpha1.ResourceConfig{}
	for key, v := range overrides {
		group, kind := splitGroupKind(key)
		if v.HealthLua != "" || v.UseOpenLibs {
			rc.Health = append(rc.Health, v1alpha1.ResourceHealthCustomization{
				Group:       group,
				Kind:        kind,
				HealthLua:   v.HealthLua,
				UseOpenLibs: v.UseOpenLibs,
			})
		}
		if v.Actions != "" {
			actions, err := parseResourceActions(group, kind, v.Actions)
			if err != nil {
				return nil, err
			}
			rc.Actions = append(rc.Actions, actions)
		}
		if len(v.IgnoreDifferences.JSONPointers) > 0 ||
			len(v.IgnoreDifferences.JQPathExpressions) > 0 ||
			len(v.IgnoreDifferences.ManagedFieldsManagers) > 0 {
			rc.IgnoreDifferences = append(rc.IgnoreDifferences, v1alpha1.ResourceIgnoreCustomization{
				Group:                 group,
				Kind:                  kind,
				JSONPointers:          append([]string(nil), v.IgnoreDifferences.JSONPointers...),
				JQPathExpressions:     append([]string(nil), v.IgnoreDifferences.JQPathExpressions...),
				ManagedFieldsManagers: append([]string(nil), v.IgnoreDifferences.ManagedFieldsManagers...),
			})
		}
		if len(v.IgnoreResourceUpdates.JSONPointers) > 0 ||
			len(v.IgnoreResourceUpdates.JQPathExpressions) > 0 ||
			len(v.IgnoreResourceUpdates.ManagedFieldsManagers) > 0 {
			rc.IgnoreResourceUpdates = append(rc.IgnoreResourceUpdates, v1alpha1.ResourceIgnoreCustomization{
				Group:                 group,
				Kind:                  kind,
				JSONPointers:          append([]string(nil), v.IgnoreResourceUpdates.JSONPointers...),
				JQPathExpressions:     append([]string(nil), v.IgnoreResourceUpdates.JQPathExpressions...),
				ManagedFieldsManagers: append([]string(nil), v.IgnoreResourceUpdates.ManagedFieldsManagers...),
			})
		}
		if len(v.KnownTypeFields) > 0 {
			fields := make([]v1alpha1.KnownTypeField, 0, len(v.KnownTypeFields))
			fields = append(fields, v.KnownTypeFields...)
			rc.KnownTypeFields = append(rc.KnownTypeFields, v1alpha1.ResourceKnownTypesCustomization{
				Group:  group,
				Kind:   kind,
				Fields: fields,
			})
		}
	}
	return rc, nil
}

func parseResourceActions(group, kind, blob string) (v1alpha1.ResourceActionsCustomization, error) {
	out := v1alpha1.ResourceActionsCustomization{Group: group, Kind: kind}
	var raw map[string]any
	if err := yaml.Unmarshal([]byte(blob), &raw); err != nil {
		return out, fmt.Errorf("parse resource actions for %s/%s: %w", group, kind, err)
	}
	if v, ok := raw["discovery.lua"].(string); ok {
		out.DiscoveryLua = v
	}
	if v, ok := raw["mergeBuiltinActions"].(bool); ok {
		out.MergeBuiltinActions = v
	}
	if defs, ok := raw["definitions"].([]any); ok {
		for _, d := range defs {
			m, ok := d.(map[string]any)
			if !ok {
				continue
			}
			def := v1alpha1.ResourceActionDefinition{}
			if name, ok := m["name"].(string); ok {
				def.Name = name
			}
			if lua, ok := m["action.lua"].(string); ok {
				def.ActionLua = lua
			}
			out.Definitions = append(out.Definitions, def)
		}
	}
	return out, nil
}

func applyResourceCustomizationLists(dst *v1alpha1.ResourceConfig, src *v1alpha1.ResourceConfig) {
	dst.Health = src.Health
	dst.Actions = src.Actions
	dst.IgnoreDifferences = src.IgnoreDifferences
	dst.IgnoreResourceUpdates = src.IgnoreResourceUpdates
	dst.KnownTypeFields = src.KnownTypeFields
}

func filteredFromSettings(in []settings.FilteredResource) []v1alpha1.FilteredResource {
	if in == nil {
		return nil
	}
	out := make([]v1alpha1.FilteredResource, 0, len(in))
	for _, r := range in {
		out = append(out, v1alpha1.FilteredResource{
			APIGroups: append([]string(nil), r.APIGroups...),
			Kinds:     append([]string(nil), r.Kinds...),
			Clusters:  append([]string(nil), r.Clusters...),
		})
	}
	return out
}

func upsertOIDCClientSecret(secret string) error {
	ctx := context.Background()
	ns := TestNamespace()
	sec := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{Name: e2eOIDCClientSecretName},
		Type:       corev1.SecretTypeOpaque,
		StringData: map[string]string{e2eOIDCClientSecretKey: secret},
	}
	_, err := KubeClientset.CoreV1().Secrets(ns).Create(ctx, sec, metav1.CreateOptions{})
	if apierrors.IsAlreadyExists(err) {
		existing, getErr := KubeClientset.CoreV1().Secrets(ns).Get(ctx, e2eOIDCClientSecretName, metav1.GetOptions{})
		if getErr != nil {
			return getErr
		}
		if existing.StringData == nil {
			existing.StringData = map[string]string{}
		}
		existing.StringData[e2eOIDCClientSecretKey] = secret
		_, err = KubeClientset.CoreV1().Secrets(ns).Update(ctx, existing, metav1.UpdateOptions{})
	}
	return err
}
