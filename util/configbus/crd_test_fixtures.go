package configbus

import (
	"time"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	appv1 "github.com/argoproj/argo-cd/v3/pkg/apis/application/v1alpha1"
)

func testControllerCRD() *appv1.ArgoCDConfiguration {
	trueVal := true
	falseVal := false
	statusProcessors := int32(20)
	operationProcessors := int32(10)
	hydrationProcessors := int32(5)
	return &appv1.ArgoCDConfiguration{
		Spec: appv1.ArgoCDConfigurationSpec{
			Controller: &appv1.ControllerConfig{
				Reconciliation: &appv1.ReconciliationConfig{
					Timeout:     &metav1.Duration{Duration: 120 * time.Second},
					HardTimeout: &metav1.Duration{Duration: 300 * time.Second},
					Jitter:      &metav1.Duration{Duration: 60 * time.Second},
				},
				SelfHeal: &appv1.SelfHealConfig{
					Timeout: &metav1.Duration{Duration: 30 * time.Second},
				},
				Sync: &appv1.ApplicationSyncConfig{
					Timeout: &metav1.Duration{Duration: 5 * time.Minute},
					Impersonation: &appv1.SyncImpersonationConfig{
						Mode: "disabled",
					},
				},
				ResourceTrackingMethod: "annotation",
				InstanceLabelKey:       "app.kubernetes.io/instance",
				Processors: &appv1.ControllerProcessorsConfig{
					Status:    &statusProcessors,
					Operation: &operationProcessors,
					Hydration: &hydrationProcessors,
				},
				RepoErrorGracePeriod:  &metav1.Duration{Duration: 90 * time.Second},
				ResourceHealthPersist: &trueVal,
				// Keep server-side diff off in the default test fixture so unit
				// tests do not attempt live API dry-runs against a fake cluster.
				Diff: &appv1.ControllerDiffConfig{
					ServerSide:                &appv1.DiffServerSideConfig{Enabled: &falseVal},
					IgnoreNormalizerJQTimeout: &metav1.Duration{Duration: 2 * time.Second},
				},
				Metrics: &appv1.ControllerMetricsConfig{
					Cluster: &appv1.ControllerMetricsClusterConfig{LabelKeys: []string{"team"}},
				},
			},
		},
	}
}

// TestControllerCRDSource returns a StaticCRDSource with controller fields
// populated for unit tests under the CRD-only cutover.
func TestControllerCRDSource() CRDSource {
	return StaticCRDSource{Object: testControllerCRD()}
}

// TestControllerCRDOptions customizes TestControllerCRDSourceFor params.
type TestControllerCRDOptions struct {
	ReconciliationTimeout time.Duration
	HardTimeout           time.Duration
	Jitter                time.Duration
	SelfHealTimeout       time.Duration
	SyncTimeout           time.Duration
	RepoErrorGracePeriod  time.Duration
	PersistResourceHealth *bool
}

// TestControllerCRDSourceFor returns a StaticCRDSource with timeouts matching
// ApplicationController constructor args used by unit tests.
func TestControllerCRDSourceFor(opts TestControllerCRDOptions) CRDSource {
	cfg := testControllerCRD()
	if opts.ReconciliationTimeout > 0 {
		cfg.Spec.Controller.Reconciliation.Timeout = &metav1.Duration{Duration: opts.ReconciliationTimeout}
	}
	if opts.HardTimeout > 0 {
		cfg.Spec.Controller.Reconciliation.HardTimeout = &metav1.Duration{Duration: opts.HardTimeout}
	}
	if opts.Jitter > 0 {
		cfg.Spec.Controller.Reconciliation.Jitter = &metav1.Duration{Duration: opts.Jitter}
	}
	if opts.SelfHealTimeout > 0 {
		cfg.Spec.Controller.SelfHeal.Timeout = &metav1.Duration{Duration: opts.SelfHealTimeout}
	}
	if opts.SyncTimeout > 0 {
		cfg.Spec.Controller.Sync.Timeout = &metav1.Duration{Duration: opts.SyncTimeout}
	} else if opts.SyncTimeout == 0 {
		// Explicit zero means "no sync timeout" for operation tests.
		cfg.Spec.Controller.Sync.Timeout = nil
	}
	if opts.RepoErrorGracePeriod > 0 {
		cfg.Spec.Controller.RepoErrorGracePeriod = &metav1.Duration{Duration: opts.RepoErrorGracePeriod}
	}
	if opts.PersistResourceHealth != nil {
		cfg.Spec.Controller.ResourceHealthPersist = opts.PersistResourceHealth
	}
	return StaticCRDSource{Object: cfg}
}

func testServerCRD() *appv1.ArgoCDConfiguration {
	return testServerCRDFor(TestServerCRDOptions{})
}

// TestServerCRDSource returns a StaticCRDSource with server fields populated
// for unit tests under the CRD-only cutover.
func TestServerCRDSource() CRDSource {
	return StaticCRDSource{Object: testServerCRD()}
}

// TestServerCRDOptions customizes TestServerCRDSourceFor for server unit tests.
type TestServerCRDOptions struct {
	ListenPort                int
	ListenAddress             string
	MetricsPort               int
	MetricsAddress            string
	Insecure                  bool
	DisableAuth               bool
	RootPath                  string
	BaseHref                  string
	StaticAssetsPath          string
	XFrameOptions             *string
	ContentSecurityPolicy     *string
	ApplicationNamespaceGlobs []string
	EnableProxyExtension      bool
	SyncWithReplaceAllowed    bool
	WebhookParallelism        int
	WebhookRefreshWorkers     int
	DexServerAddr             string
}

// TestServerCRDSourceFor returns a StaticCRDSource with server fields matching opts.
func TestServerCRDSourceFor(opts TestServerCRDOptions) CRDSource {
	return StaticCRDSource{Object: testServerCRDFor(opts)}
}

func testServerCRDFor(opts TestServerCRDOptions) *appv1.ArgoCDConfiguration {
	trueVal := true
	falseVal := false
	port := int32(8080)
	if opts.ListenPort > 0 {
		port = int32(opts.ListenPort)
	}
	metricsPort := int32(8083)
	if opts.MetricsPort > 0 {
		metricsPort = int32(opts.MetricsPort)
	}
	listenAddr := "0.0.0.0"
	if opts.ListenAddress != "" {
		listenAddr = opts.ListenAddress
	}
	metricsAddr := "0.0.0.0"
	if opts.MetricsAddress != "" {
		metricsAddr = opts.MetricsAddress
	}
	tlsEnabled := !opts.Insecure
	authEnabled := !opts.DisableAuth
	if opts.ListenPort == 0 && !opts.DisableAuth {
		authEnabled = true
	}
	syncReplaceAllowed := true
	if opts.SyncWithReplaceAllowed {
		syncReplaceAllowed = true
	}
	webhookParallelism := int32(50)
	if opts.WebhookParallelism > 0 {
		webhookParallelism = int32(opts.WebhookParallelism)
	}
	webhookWorkers := int32(5)
	if opts.WebhookRefreshWorkers > 0 {
		webhookWorkers = int32(opts.WebhookRefreshWorkers)
	}
	dexAddr := "argocd-dex-server:5556"
	if opts.DexServerAddr != "" {
		dexAddr = opts.DexServerAddr
	}
	appNS := []string{"app-ns"}
	if opts.ApplicationNamespaceGlobs != nil {
		appNS = opts.ApplicationNamespaceGlobs
	}
	xfo := "sameorigin"
	if opts.XFrameOptions != nil {
		xfo = *opts.XFrameOptions
	}
	csp := "frame-ancestors 'self';"
	if opts.ContentSecurityPolicy != nil {
		csp = *opts.ContentSecurityPolicy
	}
	baseHref := "/"
	if opts.BaseHref != "" {
		baseHref = opts.BaseHref
	}
	staticAssets := "/tmp/static"
	if opts.StaticAssetsPath != "" {
		staticAssets = opts.StaticAssetsPath
	}
	proxyExt := opts.EnableProxyExtension
	return &appv1.ArgoCDConfiguration{
		Spec: appv1.ArgoCDConfigurationSpec{
			Server: &appv1.ServerConfig{
				TLSEnabled:            &tlsEnabled,
				AuthEnabled:           &authEnabled,
				Compression:           "gzip",
				BaseHref:              baseHref,
				RootPath:              opts.RootPath,
				StaticAssetsPath:      staticAssets,
				XFrameOptions:         xfo,
				ContentSecurityPolicy: csp,
				ProxyExtensionEnabled: &proxyExt,
				SyncReplaceAllowed:    &syncReplaceAllowed,
				APIContentTypes:       []string{"application/json"},
				Listen: &appv1.ServerListenConfig{
					Address:        listenAddr,
					Port:           &port,
					MetricsAddress: metricsAddr,
					MetricsPort:    &metricsPort,
				},
				DexConnection: &appv1.DexConnectionConfig{
					Address: dexAddr,
				},
				Webhook: &appv1.WebhookConfig{
					ParallelismLimit: &webhookParallelism,
					Refresh: &appv1.WebhookRefreshConfig{
						Workers: &webhookWorkers,
					},
				},
			},
			ApplicationNamespaceGlobs: appNS,
			Controller: &appv1.ControllerConfig{
				InstanceLabelKey:       "app.kubernetes.io/instance",
				ResourceTrackingMethod: "annotation",
				SourceHydrator:         &appv1.SourceHydratorConfig{Enabled: &trueVal},
				Sync: &appv1.ApplicationSyncConfig{
					Impersonation: &appv1.SyncImpersonationConfig{Mode: "disabled"},
				},
			},
			// Server wires ApplicationSet SCM settings via shared ApplicationSet CR fields.
			ApplicationSet: &appv1.ApplicationSetConfig{
				AllowedSCMProviderURLs:    []string{},
				SCMProvidersEnabled:       &trueVal,
				GitHubAPIMetricsEnabled:   &falseVal,
				GitSubmoduleEnabled:       &trueVal,
				NewGitFileGlobbingEnabled: &falseVal,
				SCMRootCAPath:             "",
			},
		},
	}
}

// TestReposerverCRDSource returns a StaticCRDSource with repo-server fields
// populated for unit tests under the CRD-only cutover.
func TestReposerverCRDSource() CRDSource {
	return StaticCRDSource{Object: testReposerverCRD()}
}

// TestReposerverCRDOptions customizes TestReposerverCRDSourceFor for repo-server unit tests.
type TestReposerverCRDOptions struct {
	IncludeHiddenDirectories *bool
	AllowOOBSymlinks         *bool
	SubmoduleEnabled         *bool
}

// TestReposerverCRDSourceFor returns a StaticCRDSource with repo-server fields matching opts.
func TestReposerverCRDSourceFor(opts TestReposerverCRDOptions) CRDSource {
	cfg := testReposerverCRD()
	if opts.IncludeHiddenDirectories != nil {
		cfg.Spec.RepoServer.IncludeHiddenDirectories = opts.IncludeHiddenDirectories
	}
	if opts.AllowOOBSymlinks != nil {
		cfg.Spec.RepoServer.AllowOOBSymlinks = opts.AllowOOBSymlinks
	}
	if opts.SubmoduleEnabled != nil && cfg.Spec.RepoServer.Git != nil {
		cfg.Spec.RepoServer.Git.SubmoduleEnabled = opts.SubmoduleEnabled
	}
	return StaticCRDSource{Object: cfg}
}

func testReposerverCRD() *appv1.ArgoCDConfiguration {
	trueVal := true
	falseVal := false
	par := int32(1)
	maxCombined := resource.MustParse("10M")
	streamedTar := resource.MustParse("0")
	streamedExtracted := resource.MustParse("0")
	helmExtracted := resource.MustParse("0")
	ociExtracted := resource.MustParse("0")
	return &appv1.ArgoCDConfiguration{
		Spec: appv1.ArgoCDConfigurationSpec{
			RepoServer: &appv1.RepoServerConfig{
				ParallelismLimit:                  &par,
				AllowOOBSymlinks:                  &falseVal,
				IncludeHiddenDirectories:          &falseVal,
				MaxCombinedDirectoryManifestsSize: &maxCombined,
				Cache: &appv1.RepoServerCacheConfig{
					RepoExpiration: &metav1.Duration{Duration: 2 * time.Hour},
				},
				Git: &appv1.RepoServerGitConfig{
					SubmoduleEnabled:     &falseVal,
					BuiltinConfigEnabled: &falseVal,
				},
				Helm: &appv1.HelmConfig{
					UserAgent: "test-agent",
					Manifest: &appv1.HelmManifestConfig{
						MaxExtractedSize:        &helmExtracted,
						MaxExtractedSizeEnabled: &trueVal,
					},
				},
				OCI: &appv1.RepoServerOCIConfig{
					LayerMediaTypes: []string{"application/vnd.oci.image.layer.v1.tar"},
					Manifest: &appv1.OCIManifestConfig{
						MaxExtractedSize:        &ociExtracted,
						MaxExtractedSizeEnabled: &trueVal,
					},
				},
				StreamedManifest: &appv1.StreamedManifestConfig{
					MaxTarSize:       &streamedTar,
					MaxExtractedSize: &streamedExtracted,
				},
				Plugin: &appv1.RepoServerPluginConfig{
					UseManifestGeneratePaths: &falseVal,
				},
				RevisionCacheLockTimeout: &metav1.Duration{Duration: 15 * time.Second},
			},
		},
	}
}

// TestApplicationsetCRDSource returns a StaticCRDSource with ApplicationSet
// fields populated for unit tests under the CRD-only cutover.
func TestApplicationsetCRDSource() CRDSource {
	return StaticCRDSource{Object: testApplicationsetCRD()}
}

func testApplicationsetCRD() *appv1.ArgoCDConfiguration {
	trueVal := true
	falseVal := false
	maxRes := int32(100)
	return &appv1.ArgoCDConfiguration{
		Spec: appv1.ArgoCDConfigurationSpec{
			ApplicationSet: &appv1.ApplicationSetConfig{
				Policy:                    "sync",
				PolicyOverrideEnabled:     &trueVal,
				NamespaceGlobs:            []string{"argocd"},
				GlobalPreserved:           &appv1.GlobalPreservedKeysConfig{},
				ProgressiveSyncs:          &appv1.ProgressiveSyncsConfig{Enabled: &trueVal},
				StatusMaxResourcesCount:   &maxRes,
				AllowedSCMProviderURLs:    []string{"https://github.com"},
				SCMProvidersEnabled:       &trueVal,
				GitHubAPIMetricsEnabled:   &trueVal,
				TokenRefStrictModeEnabled: &falseVal,
				SCMRootCAPath:             "/ca.pem",
				GitSubmoduleEnabled:       &trueVal,
				NewGitFileGlobbingEnabled: &trueVal,
				RequeueAfter:              &metav1.Duration{Duration: 3 * time.Minute},
				Metrics: &appv1.ApplicationSetMetricsConfig{
					Address:              ":8080",
					ApplicationSetLabels: []string{"team"},
				},
				ProbeAddr:   ":8081",
				WebhookAddr: ":7000",
			},
		},
	}
}

// ApplicationsetCRDLegacyFields is the subset of ApplicationSet reconciler
// fields unit tests set before ensureConfigProvider. Avoids a Legacy interface.
type ApplicationsetCRDLegacyFields struct {
	Policy                      string
	EnablePolicyOverride        bool
	ApplicationSetNamespaces    []string
	EnableProgressiveSyncs      bool
	GlobalPreservedAnnotations  []string
	GlobalPreservedLabels       []string
	MaxResourcesStatusCount     int
	MetricsAddr                 string
	MetricsApplicationsetLabels []string
	ProbeAddr                   string
	WebhookAddr                 string
}

// TestApplicationsetCRDSourceFromFields returns a StaticCRDSource reflecting
// reconciler fields set by unit tests under the CRD-only cutover.
func TestApplicationsetCRDSourceFromFields(f ApplicationsetCRDLegacyFields) CRDSource {
	cfg := testApplicationsetCRD()
	appSet := cfg.Spec.ApplicationSet
	if f.Policy != "" {
		appSet.Policy = f.Policy
	}
	enableOverride := f.EnablePolicyOverride
	appSet.PolicyOverrideEnabled = &enableOverride
	if len(f.ApplicationSetNamespaces) > 0 {
		appSet.NamespaceGlobs = f.ApplicationSetNamespaces
	}
	enableProgressive := f.EnableProgressiveSyncs
	appSet.ProgressiveSyncs = &appv1.ProgressiveSyncsConfig{Enabled: &enableProgressive}
	appSet.GlobalPreserved = &appv1.GlobalPreservedKeysConfig{
		AnnotationKeys: f.GlobalPreservedAnnotations,
		LabelKeys:      f.GlobalPreservedLabels,
	}
	if f.MaxResourcesStatusCount > 0 {
		v := int32(f.MaxResourcesStatusCount)
		appSet.StatusMaxResourcesCount = &v
	}
	if f.MetricsAddr != "" {
		if appSet.Metrics == nil {
			appSet.Metrics = &appv1.ApplicationSetMetricsConfig{}
		}
		appSet.Metrics.Address = f.MetricsAddr
	}
	if len(f.MetricsApplicationsetLabels) > 0 {
		if appSet.Metrics == nil {
			appSet.Metrics = &appv1.ApplicationSetMetricsConfig{}
		}
		appSet.Metrics.ApplicationSetLabels = f.MetricsApplicationsetLabels
	}
	if f.ProbeAddr != "" {
		appSet.ProbeAddr = f.ProbeAddr
	}
	if f.WebhookAddr != "" {
		appSet.WebhookAddr = f.WebhookAddr
	}
	return StaticCRDSource{Object: cfg}
}

// TestNotificationsCRDSource returns a StaticCRDSource with notifications
// fields populated for unit tests under the CRD-only cutover.
func TestNotificationsCRDSource() CRDSource {
	return StaticCRDSource{Object: testNotificationsCRD()}
}

func testNotificationsCRD() *appv1.ArgoCDConfiguration {
	trueVal := true
	return &appv1.ArgoCDConfiguration{
		Spec: appv1.ArgoCDConfigurationSpec{
			Notifications: &appv1.NotificationsConfig{
				SelfServiceEnabled: &trueVal,
				ConfigMapName:      "argocd-notifications-cm",
				SecretName:         "argocd-notifications-secret",
				AppLabelSelector:   "app=demo",
			},
			ApplicationNamespaceGlobs: []string{"team-a"},
		},
	}
}

func testCommitserverCRD() *appv1.ArgoCDConfiguration {
	trueVal := true
	return &appv1.ArgoCDConfiguration{
		Spec: appv1.ArgoCDConfigurationSpec{
			CommitServer: &appv1.CommitServerConfig{
				Listen: &appv1.ListenConfig{
					Address:        "0.0.0.0",
					MetricsAddress: "0.0.0.0",
				},
				Log: &appv1.LogConfig{
					Format: "json",
					Level:  "debug",
				},
				GRPCTXTServiceConfigEnabled: &trueVal,
			},
		},
	}
}

// TestCommitserverCRDSource returns a StaticCRDSource with commit-server fields
// populated for unit tests under the CRD-only cutover.
func TestCommitserverCRDSource() CRDSource {
	return StaticCRDSource{Object: testCommitserverCRD()}
}
