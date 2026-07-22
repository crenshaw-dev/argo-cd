package configbus

import (
	"strconv"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"

	enginecache "github.com/argoproj/argo-cd/gitops-engine/v3/pkg/cache"
	appv1 "github.com/argoproj/argo-cd/v3/pkg/apis/application/v1alpha1"
	"github.com/argoproj/argo-cd/v3/util/settings"
)

func crdDur(p *metav1.Duration) (time.Duration, bool) {
	if p == nil {
		return 0, false
	}
	return p.Duration, true
}

func crdDurStr(p *metav1.Duration) (string, bool) {
	if p == nil {
		return "", false
	}
	return p.Duration.String(), true
}

func crdInt(p *int32) (int, bool) {
	if p == nil {
		return 0, false
	}
	return int(*p), true
}

func crdInt64(p *int32) (int64, bool) {
	if p == nil {
		return 0, false
	}
	return int64(*p), true
}

func crdInt64FromInt64(p *int64) (int64, bool) {
	if p == nil {
		return 0, false
	}
	return *p, true
}

func crdInt64FromQty(p *resource.Quantity) (int64, bool) {
	if p == nil {
		return 0, false
	}
	return p.Value(), true
}

func crdQty(p *resource.Quantity) (resource.Quantity, bool) {
	if p == nil {
		return resource.Quantity{}, false
	}
	return *p, true
}

func crdBool(p *bool) (bool, bool) {
	if p == nil {
		return false, false
	}
	return *p, true
}

func crdBoolNot(p *bool) (bool, bool) {
	if p == nil {
		return false, false
	}
	return !*p, true
}

func crdStr(s string) (string, bool) {
	if s == "" {
		return "", false
	}
	return s, true
}

func crdStrSlice(ss []string) ([]string, bool) {
	if len(ss) == 0 {
		return nil, false
	}
	return ss, true
}

func crdLogFormat(l *appv1.LogConfig) (string, bool) {
	if l == nil {
		return "", false
	}
	return crdStr(l.Format)
}

func crdLogLevel(l *appv1.LogConfig) (string, bool) {
	if l == nil {
		return "", false
	}
	return crdStr(l.Level)
}

func crdMapCommaKV(m map[string]string) (string, bool) {
	if len(m) == 0 {
		return "", false
	}
	parts := make([]string, 0, len(m))
	for k, v := range m {
		parts = append(parts, k+"="+v)
	}
	return strings.Join(parts, ","), true
}

func crdMapColonKV(m map[string]string) (string, bool) {
	if len(m) == 0 {
		return "", false
	}
	parts := make([]string, 0, len(m))
	for k, v := range m {
		parts = append(parts, k+":"+v)
	}
	return strings.Join(parts, ","), true
}

func crdRespectRBAC(s string) (int, bool) {
	if s == "" {
		return int(enginecache.RespectRbacDisabled), true
	}
	switch s {
	case "normal":
		return int(enginecache.RespectRbacNormal), true
	case "strict":
		return int(enginecache.RespectRbacStrict), true
	default:
		return 0, false
	}
}

func crdSensitiveMaskMap(keys []string) (map[string]bool, bool) {
	if len(keys) == 0 {
		return nil, false
	}
	out := make(map[string]bool, len(keys))
	for _, k := range keys {
		out[k] = true
	}
	return out, true
}

func crdCompareOptions(co *appv1.CompareOptions) (settings.ArgoCDDiffOptions, bool) {
	if co == nil {
		return settings.ArgoCDDiffOptions{}, false
	}
	return settings.ArgoCDDiffOptions{
		IgnoreAggregatedRoles:              co.IgnoreAggregatedRoles,
		IgnoreResourceStatusField:          settings.IgnoreStatus(co.IgnoreResourceStatusField),
		IgnoreDifferencesOnResourceUpdates: co.IgnoreDifferencesOnResourceUpdates,
	}, true
}

func crdHelmOptions(h *appv1.HelmConfig) (*appv1.HelmOptions, bool) {
	if h == nil || len(h.ValuesFileSchemes) == 0 {
		return nil, false
	}
	return &appv1.HelmOptions{ValuesFileSchemes: h.ValuesFileSchemes}, true
}

func crdKustomizeOptions(k *appv1.KustomizeConfig) (*appv1.KustomizeOptions, bool) {
	if k == nil {
		return nil, false
	}
	if k.BuildOptions == "" && len(k.Versions) == 0 {
		return nil, false
	}
	return &appv1.KustomizeOptions{BuildOptions: k.BuildOptions, Versions: k.Versions}, true
}

func crdImpersonationEnabled(mode string) (bool, bool) {
	if mode == "" {
		return false, false
	}
	return mode != "disabled", true
}

func crdImpersonationEnforced(mode string) (bool, bool) {
	if mode == "" {
		return false, false
	}
	return mode == "required", true
}

func crdRBACOverlayCSV(overlays []appv1.RBACPolicyOverlay) (string, bool) {
	if len(overlays) == 0 {
		return "", false
	}
	var b strings.Builder
	for i, o := range overlays {
		if i > 0 {
			b.WriteString("\n")
		}
		b.WriteString(o.CSV)
	}
	return b.String(), true
}

func crdServerURLs(cfg *appv1.ArgoCDConfiguration) []string {
	if cfg == nil || cfg.Spec.Server == nil {
		return nil
	}
	return cfg.Spec.Server.URLs
}

func crdRepoClient(cfg *appv1.ArgoCDConfiguration) *appv1.RepoServerClientConfig {
	if cfg == nil || cfg.Spec.RepoServer == nil {
		return nil
	}
	return cfg.Spec.RepoServer.Client
}

func crdControllerK8s(cfg *appv1.ArgoCDConfiguration) *appv1.K8sClientConfig {
	if cfg == nil || cfg.Spec.Controller == nil {
		return nil
	}
	return cfg.Spec.Controller.K8sClient
}

func crdServerK8s(cfg *appv1.ArgoCDConfiguration) *appv1.K8sClientConfig {
	if cfg == nil || cfg.Spec.Server == nil {
		return nil
	}
	return cfg.Spec.Server.K8sClient
}

func crdAppSetK8s(cfg *appv1.ArgoCDConfiguration) *appv1.K8sClientConfig {
	if cfg == nil || cfg.Spec.ApplicationSet == nil {
		return nil
	}
	return cfg.Spec.ApplicationSet.K8sClient
}

func crdK8sQPS(k *appv1.K8sClientConfig) (string, bool) {
	if k == nil {
		return "", false
	}
	return crdStr(k.QPS)
}

func crdK8sBurst(k *appv1.K8sClientConfig) (string, bool) {
	if k == nil || k.Burst == nil {
		return "", false
	}
	return strconv.Itoa(int(*k.Burst)), true
}

func crdK8sMaxIdle(k *appv1.K8sClientConfig) (string, bool) {
	if k == nil || k.MaxIdleConnections == nil {
		return "", false
	}
	return strconv.Itoa(int(*k.MaxIdleConnections)), true
}

func crdK8sTCPTimeout(k *appv1.K8sClientConfig) (string, bool) {
	if k == nil || k.TCP == nil {
		return "", false
	}
	return crdDurStr(k.TCP.Timeout)
}

func crdK8sTCPKeepalive(k *appv1.K8sClientConfig) (string, bool) {
	if k == nil || k.TCP == nil {
		return "", false
	}
	return crdDurStr(k.TCP.KeepAlive)
}

func crdK8sTCPIdle(k *appv1.K8sClientConfig) (string, bool) {
	if k == nil || k.TCP == nil {
		return "", false
	}
	return crdDurStr(k.TCP.IdleTimeout)
}

func crdK8sTLSHandshake(k *appv1.K8sClientConfig) (string, bool) {
	if k == nil {
		return "", false
	}
	return crdDurStr(k.TLSHandshakeTimeout)
}

func crdK8sRetryMax(k *appv1.K8sClientConfig) (string, bool) {
	if k == nil || k.Retry == nil || k.Retry.Max == nil {
		return "", false
	}
	return strconv.Itoa(int(*k.Retry.Max)), true
}

func crdK8sRetryBackoff(k *appv1.K8sClientConfig) (string, bool) {
	if k == nil || k.Retry == nil || k.Retry.Backoff == nil {
		return "", false
	}
	return crdDurStr(k.Retry.Backoff.Duration)
}

func crdTLSMin(t *appv1.TLSVersionConfig) (string, bool) {
	if t == nil {
		return "", false
	}
	return crdStr(t.MinVersion)
}

func crdTLSMax(t *appv1.TLSVersionConfig) (string, bool) {
	if t == nil {
		return "", false
	}
	return crdStr(t.MaxVersion)
}

func crdTLSCiphers(t *appv1.TLSVersionConfig) (string, bool) {
	if t == nil || len(t.Ciphers) == 0 {
		return "", false
	}
	return strings.Join(t.Ciphers, ","), true
}

func crdMTLSCA(m *appv1.MTLSCertConfig) (string, bool) {
	if m == nil {
		return "", false
	}
	return crdStr(m.CACertPath)
}

func crdMTLSCert(m *appv1.MTLSCertConfig) (string, bool) {
	if m == nil {
		return "", false
	}
	return crdStr(m.ClientCertPath)
}

func crdMTLSCertKey(m *appv1.MTLSCertConfig) (string, bool) {
	if m == nil {
		return "", false
	}
	return crdStr(m.ClientCertKeyPath)
}

func crdAllowedNodeLabels(cfg *appv1.ArgoCDConfiguration) ([]string, bool) {
	if cfg.Spec.Controller != nil {
		// Empty list is a valid CR value (no allowed node labels).
		return cfg.Spec.Controller.AllowedNodeLabelKeys, true
	}
	return nil, false
}

func crdAppInstanceLabelKey(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.Controller != nil {
		return crdStr(cfg.Spec.Controller.InstanceLabelKey)
	}
	return "", false
}

func crdApplicationNamespaces(cfg *appv1.ArgoCDConfiguration) ([]string, bool) {
	// Empty / omitted means control-plane namespace only — still a valid resolved value.
	if cfg.Spec.ApplicationNamespaceGlobs == nil {
		return []string{}, true
	}
	return append([]string(nil), cfg.Spec.ApplicationNamespaceGlobs...), true
}

func crdApplicationsetAllowedScmProviders(cfg *appv1.ArgoCDConfiguration) ([]string, bool) {
	if cfg.Spec.ApplicationSet != nil {
		return append([]string(nil), cfg.Spec.ApplicationSet.AllowedSCMProviderURLs...), true
	}
	return nil, false
}

func crdApplicationsetEnableGitSubmodule(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.ApplicationSet != nil {
		return crdBool(cfg.Spec.ApplicationSet.GitSubmoduleEnabled)
	}
	return false, false
}

func crdApplicationsetEnableGithubApiMetrics(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.ApplicationSet != nil {
		return crdBool(cfg.Spec.ApplicationSet.GitHubAPIMetricsEnabled)
	}
	return false, false
}

func crdApplicationsetEnableNewGitFileGlobbing(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.ApplicationSet != nil {
		return crdBool(cfg.Spec.ApplicationSet.NewGitFileGlobbingEnabled)
	}
	return false, false
}

func crdApplicationsetEnablePolicyOverride(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.ApplicationSet != nil {
		return crdBool(cfg.Spec.ApplicationSet.PolicyOverrideEnabled)
	}
	return false, false
}

func crdApplicationsetEnableProgressiveSyncs(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.ApplicationSet != nil && cfg.Spec.ApplicationSet.ProgressiveSyncs != nil {
		return crdBool(cfg.Spec.ApplicationSet.ProgressiveSyncs.Enabled)
	}
	return false, false
}

func crdApplicationsetEnableScmProviders(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.ApplicationSet != nil {
		return crdBool(cfg.Spec.ApplicationSet.SCMProvidersEnabled)
	}
	return false, false
}

func crdApplicationsetEnableTokenrefStrictMode(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.ApplicationSet != nil {
		return crdBool(cfg.Spec.ApplicationSet.TokenRefStrictModeEnabled)
	}
	return false, false
}

func crdApplicationsetMetricsAddr(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.ApplicationSet != nil && cfg.Spec.ApplicationSet.Metrics != nil {
		return crdStr(cfg.Spec.ApplicationSet.Metrics.Address)
	}
	return "", false
}

func crdApplicationsetMetricsApplicationsetLabels(cfg *appv1.ArgoCDConfiguration) ([]string, bool) {
	if cfg.Spec.ApplicationSet != nil && cfg.Spec.ApplicationSet.Metrics != nil {
		return crdStrSlice(cfg.Spec.ApplicationSet.Metrics.ApplicationSetLabels)
	}
	return nil, false
}

func crdApplicationsetPolicy(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.ApplicationSet != nil {
		return crdStr(cfg.Spec.ApplicationSet.Policy)
	}
	return "", false
}

func crdApplicationsetProbeAddr(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.ApplicationSet != nil {
		return crdStr(cfg.Spec.ApplicationSet.ProbeAddr)
	}
	return "", false
}

func crdApplicationsetRequeueAfter(cfg *appv1.ArgoCDConfiguration) (time.Duration, bool) {
	if cfg.Spec.ApplicationSet != nil {
		return crdDur(cfg.Spec.ApplicationSet.RequeueAfter)
	}
	return 0, false
}

func crdApplicationsetScmRootCaPath(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.ApplicationSet != nil {
		// Empty SCM root CA path is valid (use system CAs).
		return cfg.Spec.ApplicationSet.SCMRootCAPath, true
	}
	return "", false
}

func crdApplicationsetStatusMaxResourcesCount(cfg *appv1.ArgoCDConfiguration) (int, bool) {
	if cfg.Spec.ApplicationSet != nil {
		return crdInt(cfg.Spec.ApplicationSet.StatusMaxResourcesCount)
	}
	return 0, false
}

func crdApplicationsetWebhookAddr(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.ApplicationSet != nil {
		return crdStr(cfg.Spec.ApplicationSet.WebhookAddr)
	}
	return "", false
}

func crdCommitAuthorEmail(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.CommitServer == nil || cfg.Spec.CommitServer.Commit == nil || cfg.Spec.CommitServer.Commit.Author == nil {
		return "", true
	}
	return cfg.Spec.CommitServer.Commit.Author.Email, true
}

func crdCommitAuthorName(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.CommitServer == nil || cfg.Spec.CommitServer.Commit == nil || cfg.Spec.CommitServer.Commit.Author == nil {
		return "", true
	}
	return cfg.Spec.CommitServer.Commit.Author.Name, true
}

func crdCommitserverGrpcEnableTxtServiceConfig(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.CommitServer != nil {
		return crdBool(cfg.Spec.CommitServer.GRPCTXTServiceConfigEnabled)
	}
	return false, false
}

func crdCommitserverListenAddress(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.CommitServer != nil && cfg.Spec.CommitServer.Listen != nil {
		return crdStr(cfg.Spec.CommitServer.Listen.Address)
	}
	return "", false
}

func crdCommitserverLogFormat(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.CommitServer != nil {
		return crdLogFormat(cfg.Spec.CommitServer.Log)
	}
	return "", false
}

func crdCommitserverLogLevel(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.CommitServer != nil {
		return crdLogLevel(cfg.Spec.CommitServer.Log)
	}
	return "", false
}

func crdCommitserverMetricsListenAddress(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.CommitServer != nil && cfg.Spec.CommitServer.Listen != nil {
		return crdStr(cfg.Spec.CommitServer.Listen.MetricsAddress)
	}
	return "", false
}

func crdControllerDiffServerSide(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.Controller != nil && cfg.Spec.Controller.Diff != nil && cfg.Spec.Controller.Diff.ServerSide != nil {
		return crdBool(cfg.Spec.Controller.Diff.ServerSide.Enabled)
	}
	return false, false
}

func crdControllerHydrationProcessors(cfg *appv1.ArgoCDConfiguration) (int, bool) {
	if cfg.Spec.Controller != nil && cfg.Spec.Controller.Processors != nil && cfg.Spec.Controller.Processors.Hydration != nil {
		return int(*cfg.Spec.Controller.Processors.Hydration), true
	}
	return 0, false
}

func crdControllerIgnoreNormalizerJqTimeout(cfg *appv1.ArgoCDConfiguration) (time.Duration, bool) {
	if cfg.Spec.Controller != nil && cfg.Spec.Controller.Diff != nil {
		return crdDur(cfg.Spec.Controller.Diff.IgnoreNormalizerJQTimeout)
	}
	return 0, false
}

func crdControllerOperationProcessors(cfg *appv1.ArgoCDConfiguration) (int, bool) {
	if cfg.Spec.Controller != nil && cfg.Spec.Controller.Processors != nil && cfg.Spec.Controller.Processors.Operation != nil {
		return int(*cfg.Spec.Controller.Processors.Operation), true
	}
	return 0, false
}

func crdControllerRepoErrorGracePeriodSeconds(cfg *appv1.ArgoCDConfiguration) (time.Duration, bool) {
	if cfg.Spec.Controller != nil {
		return crdDur(cfg.Spec.Controller.RepoErrorGracePeriod)
	}
	return 0, false
}

func crdControllerResourceHealthPersist(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.Controller != nil {
		return crdBool(cfg.Spec.Controller.ResourceHealthPersist)
	}
	return false, false
}

func crdControllerSelfHealTimeoutSeconds(cfg *appv1.ArgoCDConfiguration) (time.Duration, bool) {
	if cfg.Spec.Controller != nil && cfg.Spec.Controller.SelfHeal != nil {
		return crdDur(cfg.Spec.Controller.SelfHeal.Timeout)
	}
	return 0, false
}

func crdControllerStatusProcessors(cfg *appv1.ArgoCDConfiguration) (int, bool) {
	if cfg.Spec.Controller != nil && cfg.Spec.Controller.Processors != nil && cfg.Spec.Controller.Processors.Status != nil {
		return int(*cfg.Spec.Controller.Processors.Status), true
	}
	return 0, false
}

func crdControllerSyncTimeoutSeconds(cfg *appv1.ArgoCDConfiguration) (time.Duration, bool) {
	if cfg.Spec.Controller == nil || cfg.Spec.Controller.Sync == nil {
		return 0, false
	}
	if cfg.Spec.Controller.Sync.Timeout == nil {
		return 0, true
	}
	return cfg.Spec.Controller.Sync.Timeout.Duration, true
}

func crdHelmSettings(cfg *appv1.ArgoCDConfiguration) (*appv1.HelmOptions, bool, error) {
	if cfg.Spec.RepoServer == nil || cfg.Spec.RepoServer.Helm == nil {
		return &appv1.HelmOptions{}, true, nil
	}
	v, ok := crdHelmOptions(cfg.Spec.RepoServer.Helm)
	if !ok {
		return &appv1.HelmOptions{}, true, nil
	}
	return v, true, nil
}

func crdHydratorEnabled(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.Controller != nil && cfg.Spec.Controller.SourceHydrator != nil && cfg.Spec.Controller.SourceHydrator.Enabled != nil {
		return crdBool(cfg.Spec.Controller.SourceHydrator.Enabled)
	}
	return false, false
}

func crdIgnoreResourceUpdatesEnabled(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.Controller == nil {
		return false, false
	}
	if cfg.Spec.Controller.Diff == nil || cfg.Spec.Controller.Diff.IgnoreResourceUpdatesEnabled == nil {
		return false, true
	}
	return *cfg.Spec.Controller.Diff.IgnoreResourceUpdatesEnabled, true
}

func crdCfgImpersonationEnabled(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.Controller != nil && cfg.Spec.Controller.Sync != nil && cfg.Spec.Controller.Sync.Impersonation != nil {
		return crdImpersonationEnabled(cfg.Spec.Controller.Sync.Impersonation.Mode)
	}
	return false, false
}

func crdCfgImpersonationEnforced(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.Controller != nil && cfg.Spec.Controller.Sync != nil && cfg.Spec.Controller.Sync.Impersonation != nil {
		return crdImpersonationEnforced(cfg.Spec.Controller.Sync.Impersonation.Mode)
	}
	return false, false
}

func crdInstallationID(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	// Empty installation ID is valid (disables multi-install tracking filters).
	return cfg.Spec.InstallationID, true
}

func crdKustomizeBuildOptions(cfg *appv1.ArgoCDConfiguration) (*appv1.KustomizeOptions, bool, error) {
	if cfg.Spec.RepoServer == nil || cfg.Spec.RepoServer.Kustomize == nil {
		return &appv1.KustomizeOptions{}, true, nil
	}
	v, ok := crdKustomizeOptions(cfg.Spec.RepoServer.Kustomize)
	if !ok {
		return &appv1.KustomizeOptions{}, true, nil
	}
	return v, true, nil
}

func crdNotificationsAppLabelSelector(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.Notifications == nil {
		return "", false
	}
	// Empty selector is valid (match all); Notifications presence marks the field configured.
	return cfg.Spec.Notifications.AppLabelSelector, true
}

func crdNotificationsConfigMapName(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.Notifications == nil {
		return "", false
	}
	return cfg.Spec.Notifications.ConfigMapName, true
}

func crdNotificationsSecretName(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.Notifications == nil {
		return "", false
	}
	return cfg.Spec.Notifications.SecretName, true
}

func crdNotificationsSelfserviceEnabled(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.Notifications != nil {
		return crdBool(cfg.Spec.Notifications.SelfServiceEnabled)
	}
	return false, false
}

func crdNotificationscontrollerProcessorsCount(cfg *appv1.ArgoCDConfiguration) (int, bool) {
	if cfg.Spec.Notifications != nil && cfg.Spec.Notifications.ProcessorsCount != nil {
		return int(*cfg.Spec.Notifications.ProcessorsCount), true
	}
	return 0, false
}

func crdReposerverAllowOobSymlinks(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.RepoServer != nil {
		return crdBool(cfg.Spec.RepoServer.AllowOOBSymlinks)
	}
	return false, false
}

func crdReposerverDisableHelmManifestMaxExtractedSize(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.RepoServer != nil && cfg.Spec.RepoServer.Helm != nil && cfg.Spec.RepoServer.Helm.Manifest != nil {
		return crdBoolNot(cfg.Spec.RepoServer.Helm.Manifest.MaxExtractedSizeEnabled)
	}
	return false, false
}

func crdReposerverDisableOciManifestMaxExtractedSize(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.RepoServer != nil && cfg.Spec.RepoServer.OCI != nil && cfg.Spec.RepoServer.OCI.Manifest != nil {
		return crdBoolNot(cfg.Spec.RepoServer.OCI.Manifest.MaxExtractedSizeEnabled)
	}
	return false, false
}

func crdReposerverEnableBuiltinGitConfig(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.RepoServer != nil && cfg.Spec.RepoServer.Git != nil {
		return crdBool(cfg.Spec.RepoServer.Git.BuiltinConfigEnabled)
	}
	return false, false
}

func crdReposerverEnableGitSubmodule(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.RepoServer != nil && cfg.Spec.RepoServer.Git != nil {
		return crdBool(cfg.Spec.RepoServer.Git.SubmoduleEnabled)
	}
	return false, false
}

func crdReposerverHelmManifestMaxExtractedSize(cfg *appv1.ArgoCDConfiguration) (int64, bool) {
	if cfg.Spec.RepoServer != nil && cfg.Spec.RepoServer.Helm != nil && cfg.Spec.RepoServer.Helm.Manifest != nil {
		return crdInt64FromQty(cfg.Spec.RepoServer.Helm.Manifest.MaxExtractedSize)
	}
	return 0, false
}

func crdReposerverHelmUserAgent(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.RepoServer != nil && cfg.Spec.RepoServer.Helm != nil {
		return crdStr(cfg.Spec.RepoServer.Helm.UserAgent)
	}
	return "", false
}

func crdReposerverIncludeHiddenDirectories(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.RepoServer != nil {
		return crdBool(cfg.Spec.RepoServer.IncludeHiddenDirectories)
	}
	return false, false
}

func crdReposerverMaxCombinedDirectoryManifestsSize(cfg *appv1.ArgoCDConfiguration) (resource.Quantity, bool) {
	if cfg.Spec.RepoServer != nil {
		return crdQty(cfg.Spec.RepoServer.MaxCombinedDirectoryManifestsSize)
	}
	return resource.Quantity{}, false
}

func crdReposerverOciManifestMaxExtractedSize(cfg *appv1.ArgoCDConfiguration) (int64, bool) {
	if cfg.Spec.RepoServer != nil && cfg.Spec.RepoServer.OCI != nil && cfg.Spec.RepoServer.OCI.Manifest != nil {
		return crdInt64FromQty(cfg.Spec.RepoServer.OCI.Manifest.MaxExtractedSize)
	}
	return 0, false
}

func crdReposerverParallelismLimit(cfg *appv1.ArgoCDConfiguration) (int64, bool) {
	if cfg.Spec.RepoServer != nil {
		return crdInt64(cfg.Spec.RepoServer.ParallelismLimit)
	}
	return 0, false
}

func crdReposerverPluginUseManifestGeneratePaths(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.RepoServer != nil && cfg.Spec.RepoServer.Plugin != nil {
		return crdBool(cfg.Spec.RepoServer.Plugin.UseManifestGeneratePaths)
	}
	return false, false
}

func crdReposerverRepoCacheExpiration(cfg *appv1.ArgoCDConfiguration) (time.Duration, bool) {
	if cfg.Spec.RepoServer != nil && cfg.Spec.RepoServer.Cache != nil {
		return crdDur(cfg.Spec.RepoServer.Cache.RepoExpiration)
	}
	return 0, false
}

func crdReposerverRevisionCacheLockTimeout(cfg *appv1.ArgoCDConfiguration) (time.Duration, bool) {
	if cfg.Spec.RepoServer != nil {
		return crdDur(cfg.Spec.RepoServer.RevisionCacheLockTimeout)
	}
	return 0, false
}

func crdReposerverStreamedManifestMaxExtractedSize(cfg *appv1.ArgoCDConfiguration) (int64, bool) {
	if cfg.Spec.RepoServer != nil && cfg.Spec.RepoServer.StreamedManifest != nil {
		return crdInt64FromQty(cfg.Spec.RepoServer.StreamedManifest.MaxExtractedSize)
	}
	return 0, false
}

func crdReposerverStreamedManifestMaxTarSize(cfg *appv1.ArgoCDConfiguration) (int64, bool) {
	if cfg.Spec.RepoServer != nil && cfg.Spec.RepoServer.StreamedManifest != nil {
		return crdInt64FromQty(cfg.Spec.RepoServer.StreamedManifest.MaxTarSize)
	}
	return 0, false
}

func crdResourceCompareOptions(cfg *appv1.ArgoCDConfiguration) (settings.ArgoCDDiffOptions, bool, error) {
	if cfg.Spec.Controller == nil {
		return settings.ArgoCDDiffOptions{}, false, nil
	}
	if cfg.Spec.Controller.Diff == nil || cfg.Spec.Controller.Diff.CompareOptions == nil {
		return settings.ArgoCDDiffOptions{}, true, nil
	}
	v, ok := crdCompareOptions(cfg.Spec.Controller.Diff.CompareOptions)
	if !ok {
		return settings.ArgoCDDiffOptions{}, true, nil
	}
	return v, true, nil
}

func crdResourceTrackingMethod(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.Controller != nil {
		return crdStr(cfg.Spec.Controller.ResourceTrackingMethod)
	}
	return "", false
}

func crdCfgResourcesFilter(cfg *appv1.ArgoCDConfiguration) (*settings.ResourcesFilter, bool) {
	return crdResourcesFilterExclusions(cfg)
}

func crdCfgRespectRBAC(cfg *appv1.ArgoCDConfiguration) (int, bool) {
	if cfg.Spec.Controller == nil {
		return 0, false
	}
	if cfg.Spec.Controller.Resource == nil {
		return crdRespectRBAC("")
	}
	return crdRespectRBAC(cfg.Spec.Controller.Resource.RespectRBAC)
}

func crdServerBasehref(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.Server != nil {
		return cfg.Spec.Server.BaseHref, true
	}
	return "", false
}

func crdServerContentSecurityPolicy(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.Server != nil {
		return cfg.Spec.Server.ContentSecurityPolicy, true
	}
	return "", false
}

func crdServerDexServer(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.Server != nil && cfg.Spec.Server.DexConnection != nil {
		return crdStr(cfg.Spec.Server.DexConnection.Address)
	}
	return "", false
}

func crdServerDexServerPlaintext(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.Server != nil && cfg.Spec.Server.DexConnection != nil {
		return crdBoolNot(cfg.Spec.Server.DexConnection.TLSEnabled)
	}
	return false, false
}

func crdServerDexServerStrictTls(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.Server != nil && cfg.Spec.Server.DexConnection != nil && cfg.Spec.Server.DexConnection.InsecureSkipVerify != nil {
		return !*cfg.Spec.Server.DexConnection.InsecureSkipVerify, true
	}
	return false, false
}

func crdServerDisableAuth(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.Server != nil {
		return crdBoolNot(cfg.Spec.Server.AuthEnabled)
	}
	return false, false
}

func crdAnonymousUserEnabled(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.Server != nil && cfg.Spec.Server.Users != nil {
		return crdBool(cfg.Spec.Server.Users.AnonymousEnabled)
	}
	return false, false
}

func crdExecEnabled(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.Server != nil && cfg.Spec.Server.Exec != nil {
		return cfg.Spec.Server.Exec.Enabled, true
	}
	return false, false
}

func crdExecShells(cfg *appv1.ArgoCDConfiguration) ([]string, bool) {
	if cfg.Spec.Server != nil && cfg.Spec.Server.Exec != nil && cfg.Spec.Server.Exec.Shells != nil {
		return append([]string(nil), cfg.Spec.Server.Exec.Shells...), true
	}
	return nil, false
}

func crdStatusBadgeEnabled(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.Server != nil && cfg.Spec.Server.StatusBadge != nil {
		return cfg.Spec.Server.StatusBadge.Enabled, true
	}
	return false, false
}

func crdServerEnableGzip(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.Server != nil && cfg.Spec.Server.Compression != "" {
		return cfg.Spec.Server.Compression == "gzip", true
	}
	return false, false
}

func crdServerEnableProxyExtension(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.Server != nil {
		return crdBool(cfg.Spec.Server.ProxyExtensionEnabled)
	}
	return false, false
}

func crdServerInsecure(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.Server != nil {
		return crdBoolNot(cfg.Spec.Server.TLSEnabled)
	}
	return false, false
}

func crdServerListenAddress(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.Server != nil && cfg.Spec.Server.Listen != nil {
		return crdStr(cfg.Spec.Server.Listen.Address)
	}
	return "", false
}

func crdServerListenPort(cfg *appv1.ArgoCDConfiguration) (int, bool) {
	if cfg.Spec.Server != nil && cfg.Spec.Server.Listen != nil {
		return crdInt(cfg.Spec.Server.Listen.Port)
	}
	return 0, false
}

func crdServerMetricsListenAddress(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.Server != nil && cfg.Spec.Server.Listen != nil {
		return crdStr(cfg.Spec.Server.Listen.MetricsAddress)
	}
	return "", false
}

func crdServerMetricsPort(cfg *appv1.ArgoCDConfiguration) (int, bool) {
	if cfg.Spec.Server != nil && cfg.Spec.Server.Listen != nil {
		return crdInt(cfg.Spec.Server.Listen.MetricsPort)
	}
	return 0, false
}

func crdServerRootpath(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.Server != nil {
		return cfg.Spec.Server.RootPath, true
	}
	return "", false
}

func crdServerStaticassets(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.Server != nil {
		return cfg.Spec.Server.StaticAssetsPath, true
	}
	return "", false
}

func crdServerSyncReplaceAllowed(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.Server != nil {
		return crdBool(cfg.Spec.Server.SyncReplaceAllowed)
	}
	return false, false
}

func crdServerWebhookParallelismLimit(cfg *appv1.ArgoCDConfiguration) (int, bool) {
	if cfg.Spec.Server != nil && cfg.Spec.Server.Webhook != nil {
		return crdInt(cfg.Spec.Server.Webhook.ParallelismLimit)
	}
	return 0, false
}

func crdServerWebhookRefreshWorkers(cfg *appv1.ArgoCDConfiguration) (int, bool) {
	if cfg.Spec.Server != nil && cfg.Spec.Server.Webhook != nil && cfg.Spec.Server.Webhook.Refresh != nil {
		return crdInt(cfg.Spec.Server.Webhook.Refresh.Workers)
	}
	return 0, false
}

func crdServerXFrameOptions(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.Server != nil {
		return cfg.Spec.Server.XFrameOptions, true
	}
	return "", false
}

func crdSourceHydratorCommitMessageTemplate(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.Controller == nil {
		return "", false
	}
	if cfg.Spec.Controller.SourceHydrator == nil {
		return "", true
	}
	// Empty template is valid; callers fall back to the built-in default.
	return cfg.Spec.Controller.SourceHydrator.CommitMessageTemplate, true
}

func crdSourceHydratorReadmeMessageTemplate(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.Controller == nil {
		return "", false
	}
	if cfg.Spec.Controller.SourceHydrator == nil {
		return "", true
	}
	return cfg.Spec.Controller.SourceHydrator.ReadmeMessageTemplate, true
}

func crdControllerMetricsClusterLabels(cfg *appv1.ArgoCDConfiguration) ([]string, bool) {
	if cfg.Spec.Controller != nil && cfg.Spec.Controller.Metrics != nil && cfg.Spec.Controller.Metrics.Cluster != nil {
		return cfg.Spec.Controller.Metrics.Cluster.LabelKeys, true
	}
	return nil, false
}

func crdReconciliationTimeout(cfg *appv1.ArgoCDConfiguration) (time.Duration, bool) {
	if cfg.Spec.Controller != nil && cfg.Spec.Controller.Reconciliation != nil {
		return crdDur(cfg.Spec.Controller.Reconciliation.Timeout)
	}
	return 0, false
}

func crdHardReconciliationTimeout(cfg *appv1.ArgoCDConfiguration) (time.Duration, bool) {
	if cfg.Spec.Controller != nil && cfg.Spec.Controller.Reconciliation != nil {
		return crdDur(cfg.Spec.Controller.Reconciliation.HardTimeout)
	}
	return 0, false
}

func crdReconciliationJitter(cfg *appv1.ArgoCDConfiguration) (time.Duration, bool) {
	if cfg.Spec.Controller != nil && cfg.Spec.Controller.Reconciliation != nil {
		return crdDur(cfg.Spec.Controller.Reconciliation.Jitter)
	}
	return 0, false
}

func crdResourceOverrides(cfg *appv1.ArgoCDConfiguration) (map[string]appv1.ResourceOverride, bool, error) {
	if cfg.Spec.Controller == nil {
		return nil, false, nil
	}
	// Absent or empty Resource subgroup means no customizations (empty map).
	if cfg.Spec.Controller.Resource == nil {
		return map[string]appv1.ResourceOverride{}, true, nil
	}
	r := cfg.Spec.Controller.Resource
	if len(r.Health) == 0 &&
		len(r.Actions) == 0 &&
		len(r.IgnoreDifferences) == 0 &&
		len(r.IgnoreResourceUpdates) == 0 &&
		len(r.KnownTypeFields) == 0 {
		return map[string]appv1.ResourceOverride{}, true, nil
	}
	out, err := mergeResourceOverrides(r)
	if err != nil {
		return nil, false, err
	}
	return out, true, nil
}

func crdApplicationsetNamespaces(cfg *appv1.ArgoCDConfiguration) ([]string, bool) {
	if cfg.Spec.ApplicationSet != nil {
		return crdStrSlice(cfg.Spec.ApplicationSet.NamespaceGlobs)
	}
	return nil, false
}

func crdApplicationsetGlobalPreservedAnnotations(cfg *appv1.ArgoCDConfiguration) ([]string, bool) {
	if cfg.Spec.ApplicationSet == nil {
		return nil, false
	}
	if cfg.Spec.ApplicationSet.GlobalPreserved == nil {
		// Empty preserved keys is valid when ApplicationSet is configured.
		return nil, true
	}
	return cfg.Spec.ApplicationSet.GlobalPreserved.AnnotationKeys, true
}

func crdApplicationsetGlobalPreservedLabels(cfg *appv1.ArgoCDConfiguration) ([]string, bool) {
	if cfg.Spec.ApplicationSet == nil {
		return nil, false
	}
	if cfg.Spec.ApplicationSet.GlobalPreserved == nil {
		return nil, true
	}
	return cfg.Spec.ApplicationSet.GlobalPreserved.LabelKeys, true
}

func crdSensitiveAnnotations(cfg *appv1.ArgoCDConfiguration) (map[string]bool, bool) {
	if cfg.Spec.Controller == nil {
		return nil, false
	}
	out := map[string]bool{}
	if cfg.Spec.Controller.Resource != nil {
		for _, k := range cfg.Spec.Controller.Resource.SensitiveMaskAnnotationKeys {
			if k == "" {
				continue
			}
			out[k] = true
		}
	}
	return out, true
}

func crdResourceCustomLabels(cfg *appv1.ArgoCDConfiguration) ([]string, bool) {
	if cfg.Spec.Controller == nil {
		return nil, false
	}
	if cfg.Spec.Controller.Resource == nil {
		return []string{}, true
	}
	return append([]string(nil), cfg.Spec.Controller.Resource.CustomLabelKeys...), true
}

func crdEnabledSourceTypes(cfg *appv1.ArgoCDConfiguration) (map[string]bool, bool) {
	// Defaults match argocd-cm: unset enable keys mean enabled. Plugin cannot be disabled.
	res := map[string]bool{
		string(appv1.ApplicationSourceTypeKustomize): true,
		string(appv1.ApplicationSourceTypeHelm):      true,
		string(appv1.ApplicationSourceTypeDirectory): true,
		string(appv1.ApplicationSourceTypePlugin):    true,
	}
	if cfg.Spec.RepoServer == nil {
		return res, true
	}
	if cfg.Spec.RepoServer.Kustomize != nil && cfg.Spec.RepoServer.Kustomize.Enabled != nil {
		res[string(appv1.ApplicationSourceTypeKustomize)] = *cfg.Spec.RepoServer.Kustomize.Enabled
	}
	if cfg.Spec.RepoServer.Helm != nil && cfg.Spec.RepoServer.Helm.Enabled != nil {
		res[string(appv1.ApplicationSourceTypeHelm)] = *cfg.Spec.RepoServer.Helm.Enabled
	}
	if cfg.Spec.RepoServer.Jsonnet != nil && cfg.Spec.RepoServer.Jsonnet.Enabled != nil {
		res[string(appv1.ApplicationSourceTypeDirectory)] = *cfg.Spec.RepoServer.Jsonnet.Enabled
	}
	return res, true
}

func crdServerContentTypes(cfg *appv1.ArgoCDConfiguration) ([]string, bool) {
	if cfg.Spec.Server == nil {
		return nil, false
	}
	return append([]string(nil), cfg.Spec.Server.APIContentTypes...), true
}

func crdSelfHealRetry(cfg *appv1.ArgoCDConfiguration) (SelfHealRetry, bool) {
	if cfg.Spec.Controller == nil {
		return SelfHealRetry{}, false
	}
	if cfg.Spec.Controller.SelfHeal == nil || cfg.Spec.Controller.SelfHeal.Backoff == nil {
		// Absent backoff subgroup means flat SelfHealTimeout (no exponential backoff).
		return SelfHealRetry{Backoff: nil}, true
	}
	b := cfg.Spec.Controller.SelfHeal.Backoff
	out := &wait.Backoff{
		Duration: 2 * time.Second,
		Factor:   3,
		Cap:      300 * time.Second,
	}
	if b.Duration != nil {
		out.Duration = b.Duration.Duration
	}
	if b.Factor != nil {
		out.Factor = float64(*b.Factor)
	}
	if b.MaxDuration != nil {
		out.Cap = b.MaxDuration.Duration
	}
	return SelfHealRetry{Backoff: out}, true
}

func crdReposerverOCIMediaTypes(cfg *appv1.ArgoCDConfiguration) ([]string, bool) {
	if cfg.Spec.RepoServer == nil {
		return nil, false
	}
	if cfg.Spec.RepoServer.OCI == nil {
		return []string{}, true
	}
	return append([]string(nil), cfg.Spec.RepoServer.OCI.LayerMediaTypes...), true
}

func crdReposerverCMPTarExcludedGlobs(cfg *appv1.ArgoCDConfiguration) ([]string, bool) {
	if cfg.Spec.RepoServer == nil {
		return nil, false
	}
	if cfg.Spec.RepoServer.Plugin == nil {
		return []string{}, true
	}
	return append([]string(nil), cfg.Spec.RepoServer.Plugin.TarExclusionGlobs...), true
}

func crdServerURL(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	urls := crdServerURLs(cfg)
	if urls == nil {
		return "", false
	}
	if len(urls) == 0 {
		return "", true
	}
	return urls[0], true
}

func crdAdditionalURLs(cfg *appv1.ArgoCDConfiguration) ([]string, bool) {
	urls := crdServerURLs(cfg)
	if urls == nil {
		return nil, false
	}
	if len(urls) <= 1 {
		return []string{}, true
	}
	return append([]string(nil), urls[1:]...), true
}

func crdApplicationFineGrainedRBACInheritanceDisabled(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.Server != nil && cfg.Spec.Server.RBAC != nil {
		return crdBoolNot(cfg.Spec.Server.RBAC.ApplicationFineGrainedInheritanceEnabled)
	}
	return false, false
}

func crdIncludeEventLabelKeys(cfg *appv1.ArgoCDConfiguration) ([]string, bool) {
	if cfg.Spec.Controller == nil || cfg.Spec.Controller.Resource == nil || cfg.Spec.Controller.Resource.EventLabels == nil {
		return nil, false
	}
	return append([]string(nil), cfg.Spec.Controller.Resource.EventLabels.IncludeKeyGlobs...), true
}

func crdExcludeEventLabelKeys(cfg *appv1.ArgoCDConfiguration) ([]string, bool) {
	if cfg.Spec.Controller == nil || cfg.Spec.Controller.Resource == nil || cfg.Spec.Controller.Resource.EventLabels == nil {
		return nil, false
	}
	return append([]string(nil), cfg.Spec.Controller.Resource.EventLabels.ExcludeKeyGlobs...), true
}

func crdGoogleAnalytics(cfg *appv1.ArgoCDConfiguration) (*settings.GoogleAnalytics, bool) {
	if cfg.Spec.Server == nil || cfg.Spec.Server.GoogleAnalytics == nil {
		return nil, false
	}
	ga := cfg.Spec.Server.GoogleAnalytics
	return &settings.GoogleAnalytics{
		TrackingID:     ga.TrackingID,
		AnonymizeUsers: ga.AnonymizeUsers,
	}, true
}

func crdHelp(cfg *appv1.ArgoCDConfiguration) (*settings.Help, bool) {
	if cfg.Spec.Server == nil || cfg.Spec.Server.Help == nil {
		return nil, false
	}
	h := cfg.Spec.Server.Help
	out := &settings.Help{BinaryURLs: copyStringMap(h.BinaryURLs)}
	if h.Chat != nil {
		out.ChatURL = h.Chat.URL
		out.ChatText = h.Chat.Text
	}
	return out, true
}

func crdInClusterEnabled(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.Cluster != nil {
		return crdBool(cfg.Spec.Cluster.InClusterEnabled)
	}
	return false, false
}

func crdMaxPodLogsToRender(cfg *appv1.ArgoCDConfiguration) (int64, bool) {
	if cfg.Spec.Server != nil && cfg.Spec.Server.Logs != nil {
		return crdInt64FromInt64(cfg.Spec.Server.Logs.MaxPodsToRender)
	}
	return 0, false
}

func crdMaxWebhookPayloadSize(cfg *appv1.ArgoCDConfiguration) (int64, bool) {
	if cfg.Spec.Server != nil && cfg.Spec.Server.Webhook != nil {
		return crdInt64FromQty(cfg.Spec.Server.Webhook.MaxPayloadSize)
	}
	return 0, false
}

func crdOIDCLogoutURL(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.Server == nil || cfg.Spec.Server.OIDC == nil {
		return "", false
	}
	return cfg.Spec.Server.OIDC.LogoutURL, true
}

func crdPasswordPattern(cfg *appv1.ArgoCDConfiguration) (string, bool) {
	if cfg.Spec.Server != nil && cfg.Spec.Server.Users != nil {
		return crdStr(cfg.Spec.Server.Users.PasswordRegex)
	}
	return "", false
}

func crdRequireOverridePrivilegeForRevisionSync(cfg *appv1.ArgoCDConfiguration) (bool, bool) {
	if cfg.Spec.Controller != nil && cfg.Spec.Controller.Sync != nil {
		return crdBool(cfg.Spec.Controller.Sync.RequireOverridePrivilegeForRevisionSync)
	}
	return false, false
}

func crdUserSessionDuration(cfg *appv1.ArgoCDConfiguration) (time.Duration, bool) {
	if cfg.Spec.Server != nil && cfg.Spec.Server.Users != nil {
		return crdDur(cfg.Spec.Server.Users.SessionDuration)
	}
	return 0, false
}

func crdWebhookRefreshJitter(cfg *appv1.ArgoCDConfiguration) (time.Duration, bool) {
	if cfg.Spec.Server != nil && cfg.Spec.Server.Webhook != nil && cfg.Spec.Server.Webhook.Refresh != nil {
		return crdDur(cfg.Spec.Server.Webhook.Refresh.Jitter)
	}
	return 0, false
}

func crdWebhookRefreshJitterThreshold(cfg *appv1.ArgoCDConfiguration) (int, bool) {
	if cfg.Spec.Server != nil && cfg.Spec.Server.Webhook != nil && cfg.Spec.Server.Webhook.Refresh != nil {
		return crdInt(cfg.Spec.Server.Webhook.Refresh.JitterThreshold)
	}
	return 0, false
}
