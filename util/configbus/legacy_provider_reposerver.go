package configbus

import (
	"errors"
	"time"

	"k8s.io/apimachinery/pkg/api/resource"
)

// ReposerverLegacy is implemented by *repository.Service.
// Methods return component-resolved flag/env values already stored on the
// service (or structs it owns). Setting.Get callbacks read only through
// this interface; runtime code uses these Legacy* methods (or Provider
// getters that Resolve to them).
type ReposerverLegacy interface {
	LegacyOCIMediaTypes() []string
	LegacyParallelismLimit() int64
	LegacyPauseGenerationAfterFailedGenerationAttempts() int
	LegacyPauseGenerationOnFailureForMinutes() int
	LegacyPauseGenerationOnFailureForRequests() int
	LegacySubmoduleEnabled() bool
	LegacyMaxCombinedDirectoryManifestsSize() resource.Quantity
	LegacyCMPTarExcludedGlobs() []string
	LegacyAllowOutOfBoundsSymlinks() bool
	LegacyStreamedManifestMaxExtractedSize() int64
	LegacyStreamedManifestMaxTarSize() int64
	LegacyHelmManifestMaxExtractedSize() int64
	LegacyHelmRegistryMaxIndexSize() int64
	LegacyOCIManifestMaxExtractedSize() int64
	LegacyDisableOCIManifestMaxExtractedSize() bool
	LegacyDisableHelmManifestMaxExtractedSize() bool
	LegacyIncludeHiddenDirectories() bool
	LegacyCMPUseManifestGeneratePaths() bool
	LegacyEnableBuiltinGitConfig() bool
	LegacyHelmUserAgent() string
	LegacyHelmChartCacheExpiration() time.Duration
	LegacyRepoCacheExpiration() time.Duration
	LegacyRevisionCacheExpiration() time.Duration
	LegacyRevisionCacheLockTimeout() time.Duration
}

func (p *LegacyProvider) requireReposerverLegacy() (ReposerverLegacy, error) {
	if p == nil || p.legacy == nil || p.legacy.Reposerver == nil {
		return nil, errors.New("config: ReposerverLegacy not supplied by component")
	}
	return p.legacy.Reposerver, nil
}

func (p *LegacyProvider) AllowOutOfBoundsSymlinks() (bool, error) {
	s, err := p.requireReposerverLegacy()
	if err != nil {
		return false, err
	}
	return s.LegacyAllowOutOfBoundsSymlinks(), nil
}

func (p *LegacyProvider) CMPTarExcludedGlobs() ([]string, error) {
	s, err := p.requireReposerverLegacy()
	if err != nil {
		return nil, err
	}
	return s.LegacyCMPTarExcludedGlobs(), nil
}

func (p *LegacyProvider) CMPUseManifestGeneratePaths() (bool, error) {
	s, err := p.requireReposerverLegacy()
	if err != nil {
		return false, err
	}
	return s.LegacyCMPUseManifestGeneratePaths(), nil
}

func (p *LegacyProvider) DisableHelmManifestMaxExtractedSize() (bool, error) {
	s, err := p.requireReposerverLegacy()
	if err != nil {
		return false, err
	}
	return s.LegacyDisableHelmManifestMaxExtractedSize(), nil
}

func (p *LegacyProvider) DisableOCIManifestMaxExtractedSize() (bool, error) {
	s, err := p.requireReposerverLegacy()
	if err != nil {
		return false, err
	}
	return s.LegacyDisableOCIManifestMaxExtractedSize(), nil
}

func (p *LegacyProvider) EnableBuiltinGitConfig() (bool, error) {
	s, err := p.requireReposerverLegacy()
	if err != nil {
		return false, err
	}
	return s.LegacyEnableBuiltinGitConfig(), nil
}

func (p *LegacyProvider) HelmChartCacheExpiration() (time.Duration, error) {
	s, err := p.requireReposerverLegacy()
	if err != nil {
		return 0, err
	}
	return s.LegacyHelmChartCacheExpiration(), nil
}

func (p *LegacyProvider) HelmManifestMaxExtractedSize() (int64, error) {
	s, err := p.requireReposerverLegacy()
	if err != nil {
		return 0, err
	}
	return s.LegacyHelmManifestMaxExtractedSize(), nil
}

func (p *LegacyProvider) HelmRegistryMaxIndexSize() (int64, error) {
	s, err := p.requireReposerverLegacy()
	if err != nil {
		return 0, err
	}
	return s.LegacyHelmRegistryMaxIndexSize(), nil
}

func (p *LegacyProvider) HelmUserAgent() (string, error) {
	s, err := p.requireReposerverLegacy()
	if err != nil {
		return "", err
	}
	return s.LegacyHelmUserAgent(), nil
}

func (p *LegacyProvider) IncludeHiddenDirectories() (bool, error) {
	s, err := p.requireReposerverLegacy()
	if err != nil {
		return false, err
	}
	return s.LegacyIncludeHiddenDirectories(), nil
}

func (p *LegacyProvider) MaxCombinedDirectoryManifestsSize() (resource.Quantity, error) {
	s, err := p.requireReposerverLegacy()
	if err != nil {
		return resource.Quantity{}, err
	}
	return s.LegacyMaxCombinedDirectoryManifestsSize(), nil
}

func (p *LegacyProvider) OCIManifestMaxExtractedSize() (int64, error) {
	s, err := p.requireReposerverLegacy()
	if err != nil {
		return 0, err
	}
	return s.LegacyOCIManifestMaxExtractedSize(), nil
}

func (p *LegacyProvider) OCIMediaTypes() ([]string, error) {
	s, err := p.requireReposerverLegacy()
	if err != nil {
		return nil, err
	}
	return s.LegacyOCIMediaTypes(), nil
}

func (p *LegacyProvider) ParallelismLimit() (int64, error) {
	s, err := p.requireReposerverLegacy()
	if err != nil {
		return 0, err
	}
	return s.LegacyParallelismLimit(), nil
}

func (p *LegacyProvider) PauseGenerationAfterFailedGenerationAttempts() (int, error) {
	s, err := p.requireReposerverLegacy()
	if err != nil {
		return 0, err
	}
	return s.LegacyPauseGenerationAfterFailedGenerationAttempts(), nil
}

func (p *LegacyProvider) PauseGenerationOnFailureForMinutes() (int, error) {
	s, err := p.requireReposerverLegacy()
	if err != nil {
		return 0, err
	}
	return s.LegacyPauseGenerationOnFailureForMinutes(), nil
}

func (p *LegacyProvider) PauseGenerationOnFailureForRequests() (int, error) {
	s, err := p.requireReposerverLegacy()
	if err != nil {
		return 0, err
	}
	return s.LegacyPauseGenerationOnFailureForRequests(), nil
}

func (p *LegacyProvider) RepoCacheExpiration() (time.Duration, error) {
	s, err := p.requireReposerverLegacy()
	if err != nil {
		return 0, err
	}
	return s.LegacyRepoCacheExpiration(), nil
}

func (p *LegacyProvider) RevisionCacheExpiration() (time.Duration, error) {
	s, err := p.requireReposerverLegacy()
	if err != nil {
		return 0, err
	}
	return s.LegacyRevisionCacheExpiration(), nil
}

func (p *LegacyProvider) RevisionCacheLockTimeout() (time.Duration, error) {
	s, err := p.requireReposerverLegacy()
	if err != nil {
		return 0, err
	}
	return s.LegacyRevisionCacheLockTimeout(), nil
}

func (p *LegacyProvider) StreamedManifestMaxExtractedSize() (int64, error) {
	s, err := p.requireReposerverLegacy()
	if err != nil {
		return 0, err
	}
	return s.LegacyStreamedManifestMaxExtractedSize(), nil
}

func (p *LegacyProvider) StreamedManifestMaxTarSize() (int64, error) {
	s, err := p.requireReposerverLegacy()
	if err != nil {
		return 0, err
	}
	return s.LegacyStreamedManifestMaxTarSize(), nil
}

func (p *LegacyProvider) SubmoduleEnabled() (bool, error) {
	s, err := p.requireReposerverLegacy()
	if err != nil {
		return false, err
	}
	return s.LegacySubmoduleEnabled(), nil
}
