package configbus

import "errors"

// ServerLegacy is implemented by *server.ArgoCDServer.
// Methods return component-resolved flag/env values already stored on the
// server (or structs it owns).
type ServerLegacy interface {
	LegacyInsecure() bool
	LegacyContentTypes() []string
	LegacyEnableGZip() bool
	LegacyStaticAssetsDir() string
	LegacyListenHost() string
	LegacyListenPort() int
	LegacyMetricsHost() string
	LegacyMetricsPort() int
	LegacyDexServerAddr() string
	LegacyDexServerPlaintext() bool
	LegacyDexServerStrictTLS() bool
	LegacyBaseHRef() string
	LegacyRootPath() string
	LegacyApplicationNamespaces() []string
	LegacyHydratorEnabled() bool
	LegacyDisableAuth() bool
	LegacyEnableProxyExtension() bool
	LegacyWebhookParallelism() int
	LegacyWebhookRefreshWorkers() int
	LegacyEnableK8sEvent() []string
	LegacySyncWithReplaceAllowed() bool
	LegacyXFrameOptions() string
	LegacyContentSecurityPolicy() string
	LegacyGitSubmoduleEnabled() bool
	LegacyEnableNewGitFileGlobbing() bool
	LegacyScmRootCAPath() string
	LegacyAllowedScmProviders() []string
	LegacyEnableScmProviders() bool
	LegacyEnableGitHubAPIMetrics() bool
}

func (p *LegacyProvider) requireServerLegacy() (ServerLegacy, error) {
	if p == nil || p.legacy == nil || p.legacy.Server == nil {
		return nil, errors.New("config: ServerLegacy not supplied by component")
	}
	return p.legacy.Server, nil
}

func (p *LegacyProvider) AllowedScmProviders() ([]string, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return nil, err
	}
	return s.LegacyAllowedScmProviders(), nil
}

func (p *LegacyProvider) ApplicationNamespaces() ([]string, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return nil, err
	}
	return s.LegacyApplicationNamespaces(), nil
}

func (p *LegacyProvider) BaseHRef() (string, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return "", err
	}
	return s.LegacyBaseHRef(), nil
}

func (p *LegacyProvider) ContentSecurityPolicy() (string, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return "", err
	}
	return s.LegacyContentSecurityPolicy(), nil
}

func (p *LegacyProvider) ContentTypes() ([]string, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return nil, err
	}
	return s.LegacyContentTypes(), nil
}

func (p *LegacyProvider) DexServerAddr() (string, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return "", err
	}
	return s.LegacyDexServerAddr(), nil
}

func (p *LegacyProvider) DexServerPlaintext() (bool, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return false, err
	}
	return s.LegacyDexServerPlaintext(), nil
}

func (p *LegacyProvider) DexServerStrictTLS() (bool, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return false, err
	}
	return s.LegacyDexServerStrictTLS(), nil
}

func (p *LegacyProvider) DisableAuth() (bool, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return false, err
	}
	return s.LegacyDisableAuth(), nil
}

func (p *LegacyProvider) EnableGZip() (bool, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return false, err
	}
	return s.LegacyEnableGZip(), nil
}

func (p *LegacyProvider) EnableGitHubAPIMetrics() (bool, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return false, err
	}
	return s.LegacyEnableGitHubAPIMetrics(), nil
}

func (p *LegacyProvider) EnableK8sEvent() ([]string, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return nil, err
	}
	return s.LegacyEnableK8sEvent(), nil
}

func (p *LegacyProvider) EnableNewGitFileGlobbing() (bool, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return false, err
	}
	return s.LegacyEnableNewGitFileGlobbing(), nil
}

func (p *LegacyProvider) EnableProxyExtension() (bool, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return false, err
	}
	return s.LegacyEnableProxyExtension(), nil
}

func (p *LegacyProvider) EnableScmProviders() (bool, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return false, err
	}
	return s.LegacyEnableScmProviders(), nil
}

func (p *LegacyProvider) GitSubmoduleEnabled() (bool, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return false, err
	}
	return s.LegacyGitSubmoduleEnabled(), nil
}

func (p *LegacyProvider) HydratorEnabled() (bool, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return false, err
	}
	return s.LegacyHydratorEnabled(), nil
}

func (p *LegacyProvider) Insecure() (bool, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return false, err
	}
	return s.LegacyInsecure(), nil
}

func (p *LegacyProvider) ListenHost() (string, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return "", err
	}
	return s.LegacyListenHost(), nil
}

func (p *LegacyProvider) ListenPort() (int, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return 0, err
	}
	return s.LegacyListenPort(), nil
}

func (p *LegacyProvider) MetricsHost() (string, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return "", err
	}
	return s.LegacyMetricsHost(), nil
}

func (p *LegacyProvider) MetricsPort() (int, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return 0, err
	}
	return s.LegacyMetricsPort(), nil
}

func (p *LegacyProvider) RootPath() (string, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return "", err
	}
	return s.LegacyRootPath(), nil
}

func (p *LegacyProvider) ScmRootCAPath() (string, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return "", err
	}
	return s.LegacyScmRootCAPath(), nil
}

func (p *LegacyProvider) StaticAssetsDir() (string, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return "", err
	}
	return s.LegacyStaticAssetsDir(), nil
}

func (p *LegacyProvider) SyncWithReplaceAllowed() (bool, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return false, err
	}
	return s.LegacySyncWithReplaceAllowed(), nil
}

func (p *LegacyProvider) WebhookParallelism() (int, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return 0, err
	}
	return s.LegacyWebhookParallelism(), nil
}

func (p *LegacyProvider) WebhookRefreshWorkers() (int, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return 0, err
	}
	return s.LegacyWebhookRefreshWorkers(), nil
}

func (p *LegacyProvider) XFrameOptions() (string, error) {
	s, err := p.requireServerLegacy()
	if err != nil {
		return "", err
	}
	return s.LegacyXFrameOptions(), nil
}
