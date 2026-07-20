package configbus

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/util/wait"

	"github.com/argoproj/argo-cd/v3/util/argo/normalizers"
)

func TestLegacyProviderResourceOverridesRequiresSettingsMgr(t *testing.T) {
	p := NewLegacyProvider(nil, nil)
	_, err := p.ResourceOverrides()
	require.Error(t, err)
	assert.NotErrorIs(t, err, ErrNotConfigured)
}

func TestLegacyProviderTimeoutsWithoutControllerError(t *testing.T) {
	p := NewLegacyProvider(nil, nil)

	_, err := p.ReconciliationTimeout()
	require.Error(t, err)
	assert.NotErrorIs(t, err, ErrNotConfigured)

	_, err = p.HardReconciliationTimeout()
	require.Error(t, err)
	assert.NotErrorIs(t, err, ErrNotConfigured)

	_, err = p.ReconciliationJitter()
	require.Error(t, err)
	assert.NotErrorIs(t, err, ErrNotConfigured)
}

func TestLegacyProviderGitRequestTimeoutDefault(t *testing.T) {
	t.Setenv("ARGOCD_GIT_REQUEST_TIMEOUT", "")
	p := NewLegacyProvider(nil, nil)
	d, err := p.GitRequestTimeout()
	require.NoError(t, err)
	assert.Equal(t, 15*time.Second, d)

	t.Setenv("ARGOCD_GIT_REQUEST_TIMEOUT", "30s")
	d, err = p.GitRequestTimeout()
	require.NoError(t, err)
	assert.Equal(t, 30*time.Second, d)
}

func TestLegacyProviderControllerLegacyRoundTrip(t *testing.T) {
	stub := &stubControllerLegacy{
		statusRefresh:    120 * time.Second,
		syncTimeout:      5 * time.Minute,
		serverSideDiff:   true,
		metricsLabels:    []string{"team"},
		selfHealTimeout:  30 * time.Second,
		persistHealth:    true,
		repoErrorGrace:   90 * time.Second,
		ignoreNormalizer: normalizers.IgnoreNormalizerOpts{JQExecutionTimeout: 2 * time.Second},
	}
	p := NewLegacyProvider(nil, &LegacyValues{Controller: stub})

	d, err := p.ReconciliationTimeout()
	require.NoError(t, err)
	assert.Equal(t, 120*time.Second, d)

	d, err = p.SyncTimeout()
	require.NoError(t, err)
	assert.Equal(t, 5*time.Minute, d)

	b, err := p.ServerSideDiff()
	require.NoError(t, err)
	assert.True(t, b)

	labels, err := p.MetricsClusterLabels()
	require.NoError(t, err)
	assert.Equal(t, []string{"team"}, labels)

	jq, err := p.IgnoreNormalizerJQTimeout()
	require.NoError(t, err)
	assert.Equal(t, 2*time.Second, jq)
}

func TestLegacyProviderServerLegacyRoundTrip(t *testing.T) {
	stub := &stubServerLegacy{
		insecure:   true,
		enableGZip: true,
		listenPort: 8080,
		baseHRef:   "/argo/",
	}
	p := NewLegacyProvider(nil, &LegacyValues{Server: stub})

	b, err := p.Insecure()
	require.NoError(t, err)
	assert.True(t, b)

	g, err := p.EnableGZip()
	require.NoError(t, err)
	assert.True(t, g)

	port, err := p.ListenPort()
	require.NoError(t, err)
	assert.Equal(t, 8080, port)

	href, err := p.BaseHRef()
	require.NoError(t, err)
	assert.Equal(t, "/argo/", href)
}

func TestCRDProviderReturnsErrNotConfigured(t *testing.T) {
	p := NewCRDProvider(nil)

	_, err := p.ReconciliationTimeout()
	assert.ErrorIs(t, err, ErrNotConfigured)

	_, err = p.ResourceOverrides()
	assert.ErrorIs(t, err, ErrNotConfigured)

	_, err = p.GitRequestTimeout()
	assert.ErrorIs(t, err, ErrNotConfigured)

	_, err = p.Insecure()
	assert.ErrorIs(t, err, ErrNotConfigured)

	_, err = p.ListenPort()
	assert.ErrorIs(t, err, ErrNotConfigured)
}

func TestHybridProviderFallsBackToLegacyOnErrNotConfigured(t *testing.T) {
	stub := &stubControllerLegacy{
		statusRefresh:  90 * time.Second,
		syncTimeout:    2 * time.Minute,
		serverSideDiff: true,
		metricsLabels:  []string{"env"},
	}
	p := NewHybridProvider(
		NewCRDProvider(nil),
		NewLegacyProvider(nil, &LegacyValues{Controller: stub}),
	)

	d, err := p.ReconciliationTimeout()
	require.NoError(t, err)
	assert.Equal(t, 90*time.Second, d)

	d, err = p.SyncTimeout()
	require.NoError(t, err)
	assert.Equal(t, 2*time.Minute, d)

	b, err := p.ServerSideDiff()
	require.NoError(t, err)
	assert.True(t, b)

	labels, err := p.MetricsClusterLabels()
	require.NoError(t, err)
	assert.Equal(t, []string{"env"}, labels)
}

func TestConfiguredFallback(t *testing.T) {
	t.Run("ErrNotConfigured falls back to legacy", func(t *testing.T) {
		v, err := configured(
			func() (string, error) { return "", ErrNotConfigured },
			func() (string, error) { return "legacy", nil },
		)
		require.NoError(t, err)
		assert.Equal(t, "legacy", v)
	})

	t.Run("non-ErrNotConfigured CRD error does not fall back", func(t *testing.T) {
		crdErr := errors.New("crd boom")
		legacyCalled := false
		v, err := configured(
			func() (string, error) { return "", crdErr },
			func() (string, error) {
				legacyCalled = true
				return "legacy", nil
			},
		)
		assert.ErrorIs(t, err, crdErr)
		assert.Empty(t, v)
		assert.False(t, legacyCalled)
	})

	t.Run("CRD success skips legacy", func(t *testing.T) {
		legacyCalled := false
		v, err := configured(
			func() (string, error) { return "crd", nil },
			func() (string, error) {
				legacyCalled = true
				return "legacy", nil
			},
		)
		require.NoError(t, err)
		assert.Equal(t, "crd", v)
		assert.False(t, legacyCalled)
	})
}

type stubControllerLegacy struct {
	statusRefresh, statusHard, statusJitter time.Duration
	syncTimeout, selfHealTimeout            time.Duration
	selfHealBackoff                         *wait.Backoff
	ignoreNormalizer                        normalizers.IgnoreNormalizerOpts
	metricsLabels                           []string
	serverSideDiff, persistHealth           bool
	repoErrorGrace                          time.Duration
}

func (s *stubControllerLegacy) LegacyStatusRefreshTimeout() time.Duration {
	return s.statusRefresh
}

func (s *stubControllerLegacy) LegacyStatusHardRefreshTimeout() time.Duration {
	return s.statusHard
}
func (s *stubControllerLegacy) LegacyStatusRefreshJitter() time.Duration { return s.statusJitter }
func (s *stubControllerLegacy) LegacySyncTimeout() time.Duration         { return s.syncTimeout }
func (s *stubControllerLegacy) LegacySelfHealTimeout() time.Duration     { return s.selfHealTimeout }

func (s *stubControllerLegacy) LegacySelfHealBackoff() *wait.Backoff { return s.selfHealBackoff }

func (s *stubControllerLegacy) LegacyIgnoreNormalizerOpts() normalizers.IgnoreNormalizerOpts {
	return s.ignoreNormalizer
}
func (s *stubControllerLegacy) LegacyMetricsClusterLabels() []string { return s.metricsLabels }
func (s *stubControllerLegacy) LegacyServerSideDiff() bool           { return s.serverSideDiff }
func (s *stubControllerLegacy) LegacyPersistResourceHealth() bool    { return s.persistHealth }
func (s *stubControllerLegacy) LegacyRepoErrorGracePeriod() time.Duration {
	return s.repoErrorGrace
}

type stubServerLegacy struct {
	insecure, enableGZip, disableAuth, enableProxy, syncReplace     bool
	hydratorEnabled, gitSubmodule, newGitGlob, enableScm, ghMetrics bool
	dexPlaintext, dexStrictTLS                                      bool
	listenPort, metricsPort, webhookParallelism, webhookWorkers     int
	listenHost, metricsHost, staticAssets, dexAddr, baseHRef        string
	rootPath, xFrame, csp, scmRootCA                                string
	contentTypes, appNamespaces, k8sEvents, allowedScm              []string
}

func (s *stubServerLegacy) LegacyInsecure() bool                  { return s.insecure }
func (s *stubServerLegacy) LegacyContentTypes() []string          { return s.contentTypes }
func (s *stubServerLegacy) LegacyEnableGZip() bool                { return s.enableGZip }
func (s *stubServerLegacy) LegacyStaticAssetsDir() string         { return s.staticAssets }
func (s *stubServerLegacy) LegacyListenHost() string              { return s.listenHost }
func (s *stubServerLegacy) LegacyListenPort() int                 { return s.listenPort }
func (s *stubServerLegacy) LegacyMetricsHost() string             { return s.metricsHost }
func (s *stubServerLegacy) LegacyMetricsPort() int                { return s.metricsPort }
func (s *stubServerLegacy) LegacyDexServerAddr() string           { return s.dexAddr }
func (s *stubServerLegacy) LegacyDexServerPlaintext() bool        { return s.dexPlaintext }
func (s *stubServerLegacy) LegacyDexServerStrictTLS() bool        { return s.dexStrictTLS }
func (s *stubServerLegacy) LegacyBaseHRef() string                { return s.baseHRef }
func (s *stubServerLegacy) LegacyRootPath() string                { return s.rootPath }
func (s *stubServerLegacy) LegacyApplicationNamespaces() []string { return s.appNamespaces }
func (s *stubServerLegacy) LegacyHydratorEnabled() bool           { return s.hydratorEnabled }
func (s *stubServerLegacy) LegacyDisableAuth() bool               { return s.disableAuth }
func (s *stubServerLegacy) LegacyEnableProxyExtension() bool      { return s.enableProxy }
func (s *stubServerLegacy) LegacyWebhookParallelism() int         { return s.webhookParallelism }
func (s *stubServerLegacy) LegacyWebhookRefreshWorkers() int      { return s.webhookWorkers }
func (s *stubServerLegacy) LegacyEnableK8sEvent() []string        { return s.k8sEvents }
func (s *stubServerLegacy) LegacySyncWithReplaceAllowed() bool    { return s.syncReplace }
func (s *stubServerLegacy) LegacyXFrameOptions() string           { return s.xFrame }
func (s *stubServerLegacy) LegacyContentSecurityPolicy() string   { return s.csp }
func (s *stubServerLegacy) LegacyGitSubmoduleEnabled() bool       { return s.gitSubmodule }
func (s *stubServerLegacy) LegacyEnableNewGitFileGlobbing() bool  { return s.newGitGlob }
func (s *stubServerLegacy) LegacyScmRootCAPath() string           { return s.scmRootCA }
func (s *stubServerLegacy) LegacyAllowedScmProviders() []string   { return s.allowedScm }
func (s *stubServerLegacy) LegacyEnableScmProviders() bool        { return s.enableScm }
func (s *stubServerLegacy) LegacyEnableGitHubAPIMetrics() bool    { return s.ghMetrics }
