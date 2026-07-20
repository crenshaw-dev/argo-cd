package configbus

func (h *HybridProvider) AllowedScmProviders() ([]string, error) {
	return configured(h.crd.AllowedScmProviders, h.legacy.AllowedScmProviders)
}

func (h *HybridProvider) ApplicationNamespaces() ([]string, error) {
	return configured(h.crd.ApplicationNamespaces, h.legacy.ApplicationNamespaces)
}

func (h *HybridProvider) BaseHRef() (string, error) {
	return configured(h.crd.BaseHRef, h.legacy.BaseHRef)
}

func (h *HybridProvider) ContentSecurityPolicy() (string, error) {
	return configured(h.crd.ContentSecurityPolicy, h.legacy.ContentSecurityPolicy)
}

func (h *HybridProvider) ContentTypes() ([]string, error) {
	return configured(h.crd.ContentTypes, h.legacy.ContentTypes)
}

func (h *HybridProvider) DexServerAddr() (string, error) {
	return configured(h.crd.DexServerAddr, h.legacy.DexServerAddr)
}

func (h *HybridProvider) DexServerPlaintext() (bool, error) {
	return configured(h.crd.DexServerPlaintext, h.legacy.DexServerPlaintext)
}

func (h *HybridProvider) DexServerStrictTLS() (bool, error) {
	return configured(h.crd.DexServerStrictTLS, h.legacy.DexServerStrictTLS)
}

func (h *HybridProvider) DisableAuth() (bool, error) {
	return configured(h.crd.DisableAuth, h.legacy.DisableAuth)
}

func (h *HybridProvider) EnableGZip() (bool, error) {
	return configured(h.crd.EnableGZip, h.legacy.EnableGZip)
}

func (h *HybridProvider) EnableGitHubAPIMetrics() (bool, error) {
	return configured(h.crd.EnableGitHubAPIMetrics, h.legacy.EnableGitHubAPIMetrics)
}

func (h *HybridProvider) EnableK8sEvent() ([]string, error) {
	return configured(h.crd.EnableK8sEvent, h.legacy.EnableK8sEvent)
}

func (h *HybridProvider) EnableNewGitFileGlobbing() (bool, error) {
	return configured(h.crd.EnableNewGitFileGlobbing, h.legacy.EnableNewGitFileGlobbing)
}

func (h *HybridProvider) EnableProxyExtension() (bool, error) {
	return configured(h.crd.EnableProxyExtension, h.legacy.EnableProxyExtension)
}

func (h *HybridProvider) EnableScmProviders() (bool, error) {
	return configured(h.crd.EnableScmProviders, h.legacy.EnableScmProviders)
}

func (h *HybridProvider) GitSubmoduleEnabled() (bool, error) {
	return configured(h.crd.GitSubmoduleEnabled, h.legacy.GitSubmoduleEnabled)
}

func (h *HybridProvider) HydratorEnabled() (bool, error) {
	return configured(h.crd.HydratorEnabled, h.legacy.HydratorEnabled)
}

func (h *HybridProvider) Insecure() (bool, error) {
	return configured(h.crd.Insecure, h.legacy.Insecure)
}

func (h *HybridProvider) ListenHost() (string, error) {
	return configured(h.crd.ListenHost, h.legacy.ListenHost)
}

func (h *HybridProvider) ListenPort() (int, error) {
	return configured(h.crd.ListenPort, h.legacy.ListenPort)
}

func (h *HybridProvider) MetricsHost() (string, error) {
	return configured(h.crd.MetricsHost, h.legacy.MetricsHost)
}

func (h *HybridProvider) MetricsPort() (int, error) {
	return configured(h.crd.MetricsPort, h.legacy.MetricsPort)
}

func (h *HybridProvider) RootPath() (string, error) {
	return configured(h.crd.RootPath, h.legacy.RootPath)
}

func (h *HybridProvider) ScmRootCAPath() (string, error) {
	return configured(h.crd.ScmRootCAPath, h.legacy.ScmRootCAPath)
}

func (h *HybridProvider) StaticAssetsDir() (string, error) {
	return configured(h.crd.StaticAssetsDir, h.legacy.StaticAssetsDir)
}

func (h *HybridProvider) SyncWithReplaceAllowed() (bool, error) {
	return configured(h.crd.SyncWithReplaceAllowed, h.legacy.SyncWithReplaceAllowed)
}

func (h *HybridProvider) WebhookParallelism() (int, error) {
	return configured(h.crd.WebhookParallelism, h.legacy.WebhookParallelism)
}

func (h *HybridProvider) WebhookRefreshWorkers() (int, error) {
	return configured(h.crd.WebhookRefreshWorkers, h.legacy.WebhookRefreshWorkers)
}

func (h *HybridProvider) XFrameOptions() (string, error) {
	return configured(h.crd.XFrameOptions, h.legacy.XFrameOptions)
}
