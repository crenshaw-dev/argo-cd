package configbus

func (p *CRDProvider) AllowedScmProviders() ([]string, error) {
	return nil, ErrNotConfigured
}

func (p *CRDProvider) ApplicationNamespaces() ([]string, error) {
	return nil, ErrNotConfigured
}

func (p *CRDProvider) BaseHRef() (string, error) {
	return "", ErrNotConfigured
}

func (p *CRDProvider) ContentSecurityPolicy() (string, error) {
	return "", ErrNotConfigured
}

func (p *CRDProvider) ContentTypes() ([]string, error) {
	return nil, ErrNotConfigured
}

func (p *CRDProvider) DexServerAddr() (string, error) {
	return "", ErrNotConfigured
}

func (p *CRDProvider) DexServerPlaintext() (bool, error) {
	return false, ErrNotConfigured
}

func (p *CRDProvider) DexServerStrictTLS() (bool, error) {
	return false, ErrNotConfigured
}

func (p *CRDProvider) DisableAuth() (bool, error) {
	return false, ErrNotConfigured
}

func (p *CRDProvider) EnableGZip() (bool, error) {
	return false, ErrNotConfigured
}

func (p *CRDProvider) EnableGitHubAPIMetrics() (bool, error) {
	return false, ErrNotConfigured
}

func (p *CRDProvider) EnableK8sEvent() ([]string, error) {
	return nil, ErrNotConfigured
}

func (p *CRDProvider) EnableNewGitFileGlobbing() (bool, error) {
	return false, ErrNotConfigured
}

func (p *CRDProvider) EnableProxyExtension() (bool, error) {
	return false, ErrNotConfigured
}

func (p *CRDProvider) EnableScmProviders() (bool, error) {
	return false, ErrNotConfigured
}

func (p *CRDProvider) GitSubmoduleEnabled() (bool, error) {
	return false, ErrNotConfigured
}

func (p *CRDProvider) HydratorEnabled() (bool, error) {
	return false, ErrNotConfigured
}

func (p *CRDProvider) Insecure() (bool, error) {
	return false, ErrNotConfigured
}

func (p *CRDProvider) ListenHost() (string, error) {
	return "", ErrNotConfigured
}

func (p *CRDProvider) ListenPort() (int, error) {
	return 0, ErrNotConfigured
}

func (p *CRDProvider) MetricsHost() (string, error) {
	return "", ErrNotConfigured
}

func (p *CRDProvider) MetricsPort() (int, error) {
	return 0, ErrNotConfigured
}

func (p *CRDProvider) RootPath() (string, error) {
	return "", ErrNotConfigured
}

func (p *CRDProvider) ScmRootCAPath() (string, error) {
	return "", ErrNotConfigured
}

func (p *CRDProvider) StaticAssetsDir() (string, error) {
	return "", ErrNotConfigured
}

func (p *CRDProvider) SyncWithReplaceAllowed() (bool, error) {
	return false, ErrNotConfigured
}

func (p *CRDProvider) WebhookParallelism() (int, error) {
	return 0, ErrNotConfigured
}

func (p *CRDProvider) WebhookRefreshWorkers() (int, error) {
	return 0, ErrNotConfigured
}

func (p *CRDProvider) XFrameOptions() (string, error) {
	return "", ErrNotConfigured
}
