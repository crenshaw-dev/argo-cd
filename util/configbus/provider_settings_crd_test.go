package configbus

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	appv1 "github.com/argoproj/argo-cd/v3/pkg/apis/application/v1alpha1"
	"github.com/argoproj/argo-cd/v3/util/settings"
)

func TestProvider_SensitiveAnnotationsFromCRD(t *testing.T) {
	cfg := testControllerCRD()
	cfg.Spec.Controller.Resource = &appv1.ResourceConfig{
		SensitiveMaskAnnotationKeys: []string{"kubectl.kubernetes.io/last-applied-configuration", "custom/secret"},
	}
	p := NewCRDProvider(StaticCRDSource{Object: cfg})

	got, err := p.SensitiveAnnotations(context.Background())
	require.NoError(t, err)
	assert.Equal(t, map[string]bool{
		"kubectl.kubernetes.io/last-applied-configuration": true,
		"custom/secret": true,
	}, got)
}

func TestProvider_ResourceCustomLabelsFromCRD(t *testing.T) {
	cfg := testControllerCRD()
	cfg.Spec.Controller.Resource = &appv1.ResourceConfig{
		CustomLabelKeys: []string{"team", "env"},
	}
	p := NewCRDProvider(StaticCRDSource{Object: cfg})

	got, err := p.ResourceCustomLabels(context.Background())
	require.NoError(t, err)
	assert.Equal(t, []string{"team", "env"}, got)
}

func TestProvider_EnabledSourceTypesFromCRD(t *testing.T) {
	falseVal := false
	cfg := testControllerCRD()
	cfg.Spec.RepoServer = &appv1.RepoServerConfig{
		Helm:      &appv1.HelmConfig{Enabled: &falseVal},
		Kustomize: &appv1.KustomizeConfig{Enabled: &falseVal},
		Jsonnet:   &appv1.JsonnetConfig{},
	}
	p := NewCRDProvider(StaticCRDSource{Object: cfg})

	got, err := p.EnabledSourceTypes(context.Background())
	require.NoError(t, err)
	assert.False(t, got[string(appv1.ApplicationSourceTypeHelm)])
	assert.False(t, got[string(appv1.ApplicationSourceTypeKustomize)])
	assert.True(t, got[string(appv1.ApplicationSourceTypeDirectory)])
	assert.True(t, got[string(appv1.ApplicationSourceTypePlugin)])
}

func TestProvider_IgnoreResourceUpdatesOverridesFromCRD(t *testing.T) {
	cfg := testControllerCRD()
	cfg.Spec.Controller.Diff.CompareOptions = &appv1.CompareOptions{
		IgnoreDifferencesOnResourceUpdates: true,
	}
	cfg.Spec.Controller.Resource = &appv1.ResourceConfig{
		IgnoreDifferences: []appv1.ResourceIgnoreCustomization{{
			Group:        "apps",
			Kind:         "Deployment",
			JSONPointers: []string{"/spec/replicas"},
		}},
		IgnoreResourceUpdates: []appv1.ResourceIgnoreCustomization{{
			Group:        "apps",
			Kind:         "Deployment",
			JSONPointers: []string{"/status"},
		}},
	}
	p := NewCRDProvider(StaticCRDSource{Object: cfg})

	got, err := p.IgnoreResourceUpdatesOverrides(context.Background())
	require.NoError(t, err)

	dep := got["apps/Deployment"]
	assert.Contains(t, dep.IgnoreDifferences.JSONPointers, "/status")
	assert.Contains(t, dep.IgnoreDifferences.JSONPointers, "/spec/replicas")
	assert.Empty(t, dep.IgnoreResourceUpdates.JSONPointers)

	all := got["*/*"]
	assert.Contains(t, all.IgnoreDifferences.JSONPointers, "/metadata/resourceVersion")
	assert.Contains(t, all.IgnoreDifferences.JSONPointers, "/metadata/generation")
	assert.Contains(t, all.IgnoreDifferences.JSONPointers, "/metadata/managedFields")
}

func TestProvider_ContentTypesFromCRD(t *testing.T) {
	p := NewCRDProvider(TestServerCRDSource())
	got, err := p.ContentTypes(context.Background())
	require.NoError(t, err)
	assert.Equal(t, []string{"application/json"}, got)
}

func TestProvider_SelfHealBackoffFromCRD(t *testing.T) {
	cfg := testControllerCRD()
	factor := int32(3)
	cfg.Spec.Controller.SelfHeal.Backoff = &appv1.BackoffConfig{
		Duration:    &metav1.Duration{Duration: 2 * time.Second},
		Factor:      &factor,
		MaxDuration: &metav1.Duration{Duration: 5 * time.Minute},
	}
	p := NewCRDProvider(StaticCRDSource{Object: cfg})

	got, err := p.SelfHealBackoff(context.Background())
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.Equal(t, 2*time.Second, got.Duration)
	assert.Equal(t, float64(3), got.Factor)
	assert.Equal(t, 5*time.Minute, got.Cap)
}

func TestProvider_IgnoreNormalizerJQTimeoutFromCRD(t *testing.T) {
	p := NewCRDProvider(TestControllerCRDSource())
	timeout, err := p.IgnoreNormalizerJQTimeout(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 2*time.Second, timeout)
}

func TestProvider_MappedSettingsRequireCRD(t *testing.T) {
	p := NewCRDProvider(nil)

	_, err := p.SensitiveAnnotations(context.Background())
	require.Error(t, err)
	_, err = p.ResourceCustomLabels(context.Background())
	require.Error(t, err)
	_, err = p.EnabledSourceTypes(context.Background())
	require.Error(t, err)
	_, err = p.IgnoreResourceUpdatesOverrides(context.Background())
	require.Error(t, err)
	_, err = p.SelfHealBackoff(context.Background())
	require.Error(t, err)
}

func TestBuildIgnoreResourceUpdatesOverrides_UsesSharedHelper(t *testing.T) {
	overrides := settings.BuildIgnoreResourceUpdatesOverrides(
		settings.ArgoCDDiffOptions{IgnoreDifferencesOnResourceUpdates: false},
		map[string]appv1.ResourceOverride{},
	)
	require.Contains(t, overrides, "*/*")
}
