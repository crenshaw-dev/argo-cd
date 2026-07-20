package server

import (
	"github.com/argoproj/argo-cd/v3/util/configbus"
)

// Ensure ArgoCDServer satisfies configbus.ServerLegacy.
var _ configbus.ServerLegacy = (*ArgoCDServer)(nil)

// Legacy* accessors implement configbus.ServerLegacy for the configbus
// Provider only. Product code and tests must read via configProvider.*; these
// methods are the sole allowed readers of the deprecated opts fields.

//nolint:staticcheck // SA1019: sole allowed reader of deprecated Insecure
func (a *ArgoCDServer) LegacyInsecure() bool {
	return a.Insecure
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated ContentTypes
func (a *ArgoCDServer) LegacyContentTypes() []string {
	return a.ContentTypes
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated EnableGZip
func (a *ArgoCDServer) LegacyEnableGZip() bool {
	return a.EnableGZip
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated StaticAssetsDir
func (a *ArgoCDServer) LegacyStaticAssetsDir() string {
	return a.StaticAssetsDir
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated ListenHost
func (a *ArgoCDServer) LegacyListenHost() string {
	return a.ListenHost
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated ListenPort
func (a *ArgoCDServer) LegacyListenPort() int {
	return a.ListenPort
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated MetricsHost
func (a *ArgoCDServer) LegacyMetricsHost() string {
	return a.MetricsHost
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated MetricsPort
func (a *ArgoCDServer) LegacyMetricsPort() int {
	return a.MetricsPort
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated DexServerAddr
func (a *ArgoCDServer) LegacyDexServerAddr() string {
	return a.DexServerAddr
}

func (a *ArgoCDServer) LegacyDexServerPlaintext() bool {
	if a.DexTLSConfig == nil {
		return false
	}
	return a.DexTLSConfig.DisableTLS
}

func (a *ArgoCDServer) LegacyDexServerStrictTLS() bool {
	if a.DexTLSConfig == nil {
		return false
	}
	return a.DexTLSConfig.StrictValidation
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated BaseHRef
func (a *ArgoCDServer) LegacyBaseHRef() string {
	return a.BaseHRef
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated RootPath
func (a *ArgoCDServer) LegacyRootPath() string {
	return a.RootPath
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated ApplicationNamespaces
func (a *ArgoCDServer) LegacyApplicationNamespaces() []string {
	return a.ApplicationNamespaces
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated HydratorEnabled
func (a *ArgoCDServer) LegacyHydratorEnabled() bool {
	return a.HydratorEnabled
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated DisableAuth
func (a *ArgoCDServer) LegacyDisableAuth() bool {
	return a.DisableAuth
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated EnableProxyExtension
func (a *ArgoCDServer) LegacyEnableProxyExtension() bool {
	return a.EnableProxyExtension
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated WebhookParallelism
func (a *ArgoCDServer) LegacyWebhookParallelism() int {
	return a.WebhookParallelism
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated WebhookRefreshWorkers
func (a *ArgoCDServer) LegacyWebhookRefreshWorkers() int {
	return a.WebhookRefreshWorkers
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated EnableK8sEvent
func (a *ArgoCDServer) LegacyEnableK8sEvent() []string {
	return a.EnableK8sEvent
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated SyncWithReplaceAllowed
func (a *ArgoCDServer) LegacySyncWithReplaceAllowed() bool {
	return a.SyncWithReplaceAllowed
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated XFrameOptions
func (a *ArgoCDServer) LegacyXFrameOptions() string {
	return a.XFrameOptions
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated ContentSecurityPolicy
func (a *ArgoCDServer) LegacyContentSecurityPolicy() string {
	return a.ContentSecurityPolicy
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated GitSubmoduleEnabled
func (a *ArgoCDServer) LegacyGitSubmoduleEnabled() bool {
	return a.GitSubmoduleEnabled
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated EnableNewGitFileGlobbing
func (a *ArgoCDServer) LegacyEnableNewGitFileGlobbing() bool {
	return a.EnableNewGitFileGlobbing
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated ScmRootCAPath
func (a *ArgoCDServer) LegacyScmRootCAPath() string {
	return a.ScmRootCAPath
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated AllowedScmProviders
func (a *ArgoCDServer) LegacyAllowedScmProviders() []string {
	return a.AllowedScmProviders
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated EnableScmProviders
func (a *ArgoCDServer) LegacyEnableScmProviders() bool {
	return a.EnableScmProviders
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated EnableGitHubAPIMetrics
func (a *ArgoCDServer) LegacyEnableGitHubAPIMetrics() bool {
	return a.EnableGitHubAPIMetrics
}
