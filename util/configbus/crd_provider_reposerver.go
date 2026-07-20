package configbus

import (
	"time"

	"k8s.io/apimachinery/pkg/api/resource"
)

func (p *CRDProvider) AllowOutOfBoundsSymlinks() (bool, error) {
	return false, ErrNotConfigured
}

func (p *CRDProvider) CMPTarExcludedGlobs() ([]string, error) {
	return nil, ErrNotConfigured
}

func (p *CRDProvider) CMPUseManifestGeneratePaths() (bool, error) {
	return false, ErrNotConfigured
}

func (p *CRDProvider) DisableHelmManifestMaxExtractedSize() (bool, error) {
	return false, ErrNotConfigured
}

func (p *CRDProvider) DisableOCIManifestMaxExtractedSize() (bool, error) {
	return false, ErrNotConfigured
}

func (p *CRDProvider) EnableBuiltinGitConfig() (bool, error) {
	return false, ErrNotConfigured
}

func (p *CRDProvider) HelmChartCacheExpiration() (time.Duration, error) {
	return 0, ErrNotConfigured
}

func (p *CRDProvider) HelmManifestMaxExtractedSize() (int64, error) {
	return 0, ErrNotConfigured
}

func (p *CRDProvider) HelmRegistryMaxIndexSize() (int64, error) {
	return 0, ErrNotConfigured
}

func (p *CRDProvider) HelmUserAgent() (string, error) {
	return "", ErrNotConfigured
}

func (p *CRDProvider) IncludeHiddenDirectories() (bool, error) {
	return false, ErrNotConfigured
}

func (p *CRDProvider) MaxCombinedDirectoryManifestsSize() (resource.Quantity, error) {
	return resource.Quantity{}, ErrNotConfigured
}

func (p *CRDProvider) OCIManifestMaxExtractedSize() (int64, error) {
	return 0, ErrNotConfigured
}

func (p *CRDProvider) OCIMediaTypes() ([]string, error) {
	return nil, ErrNotConfigured
}

func (p *CRDProvider) ParallelismLimit() (int64, error) {
	return 0, ErrNotConfigured
}

func (p *CRDProvider) PauseGenerationAfterFailedGenerationAttempts() (int, error) {
	return 0, ErrNotConfigured
}

func (p *CRDProvider) PauseGenerationOnFailureForMinutes() (int, error) {
	return 0, ErrNotConfigured
}

func (p *CRDProvider) PauseGenerationOnFailureForRequests() (int, error) {
	return 0, ErrNotConfigured
}

func (p *CRDProvider) RepoCacheExpiration() (time.Duration, error) {
	return 0, ErrNotConfigured
}

func (p *CRDProvider) RevisionCacheExpiration() (time.Duration, error) {
	return 0, ErrNotConfigured
}

func (p *CRDProvider) RevisionCacheLockTimeout() (time.Duration, error) {
	return 0, ErrNotConfigured
}

func (p *CRDProvider) StreamedManifestMaxExtractedSize() (int64, error) {
	return 0, ErrNotConfigured
}

func (p *CRDProvider) StreamedManifestMaxTarSize() (int64, error) {
	return 0, ErrNotConfigured
}

func (p *CRDProvider) SubmoduleEnabled() (bool, error) {
	return false, ErrNotConfigured
}
