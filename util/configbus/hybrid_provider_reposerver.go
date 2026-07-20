package configbus

import (
	"time"

	"k8s.io/apimachinery/pkg/api/resource"
)

func (h *HybridProvider) AllowOutOfBoundsSymlinks() (bool, error) {
	return configured(h.crd.AllowOutOfBoundsSymlinks, h.legacy.AllowOutOfBoundsSymlinks)
}

func (h *HybridProvider) CMPTarExcludedGlobs() ([]string, error) {
	return configured(h.crd.CMPTarExcludedGlobs, h.legacy.CMPTarExcludedGlobs)
}

func (h *HybridProvider) CMPUseManifestGeneratePaths() (bool, error) {
	return configured(h.crd.CMPUseManifestGeneratePaths, h.legacy.CMPUseManifestGeneratePaths)
}

func (h *HybridProvider) DisableHelmManifestMaxExtractedSize() (bool, error) {
	return configured(h.crd.DisableHelmManifestMaxExtractedSize, h.legacy.DisableHelmManifestMaxExtractedSize)
}

func (h *HybridProvider) DisableOCIManifestMaxExtractedSize() (bool, error) {
	return configured(h.crd.DisableOCIManifestMaxExtractedSize, h.legacy.DisableOCIManifestMaxExtractedSize)
}

func (h *HybridProvider) EnableBuiltinGitConfig() (bool, error) {
	return configured(h.crd.EnableBuiltinGitConfig, h.legacy.EnableBuiltinGitConfig)
}

func (h *HybridProvider) HelmChartCacheExpiration() (time.Duration, error) {
	return configured(h.crd.HelmChartCacheExpiration, h.legacy.HelmChartCacheExpiration)
}

func (h *HybridProvider) HelmManifestMaxExtractedSize() (int64, error) {
	return configured(h.crd.HelmManifestMaxExtractedSize, h.legacy.HelmManifestMaxExtractedSize)
}

func (h *HybridProvider) HelmRegistryMaxIndexSize() (int64, error) {
	return configured(h.crd.HelmRegistryMaxIndexSize, h.legacy.HelmRegistryMaxIndexSize)
}

func (h *HybridProvider) HelmUserAgent() (string, error) {
	return configured(h.crd.HelmUserAgent, h.legacy.HelmUserAgent)
}

func (h *HybridProvider) IncludeHiddenDirectories() (bool, error) {
	return configured(h.crd.IncludeHiddenDirectories, h.legacy.IncludeHiddenDirectories)
}

func (h *HybridProvider) MaxCombinedDirectoryManifestsSize() (resource.Quantity, error) {
	return configured(h.crd.MaxCombinedDirectoryManifestsSize, h.legacy.MaxCombinedDirectoryManifestsSize)
}

func (h *HybridProvider) OCIManifestMaxExtractedSize() (int64, error) {
	return configured(h.crd.OCIManifestMaxExtractedSize, h.legacy.OCIManifestMaxExtractedSize)
}

func (h *HybridProvider) OCIMediaTypes() ([]string, error) {
	return configured(h.crd.OCIMediaTypes, h.legacy.OCIMediaTypes)
}

func (h *HybridProvider) ParallelismLimit() (int64, error) {
	return configured(h.crd.ParallelismLimit, h.legacy.ParallelismLimit)
}

func (h *HybridProvider) PauseGenerationAfterFailedGenerationAttempts() (int, error) {
	return configured(h.crd.PauseGenerationAfterFailedGenerationAttempts, h.legacy.PauseGenerationAfterFailedGenerationAttempts)
}

func (h *HybridProvider) PauseGenerationOnFailureForMinutes() (int, error) {
	return configured(h.crd.PauseGenerationOnFailureForMinutes, h.legacy.PauseGenerationOnFailureForMinutes)
}

func (h *HybridProvider) PauseGenerationOnFailureForRequests() (int, error) {
	return configured(h.crd.PauseGenerationOnFailureForRequests, h.legacy.PauseGenerationOnFailureForRequests)
}

func (h *HybridProvider) RepoCacheExpiration() (time.Duration, error) {
	return configured(h.crd.RepoCacheExpiration, h.legacy.RepoCacheExpiration)
}

func (h *HybridProvider) RevisionCacheExpiration() (time.Duration, error) {
	return configured(h.crd.RevisionCacheExpiration, h.legacy.RevisionCacheExpiration)
}

func (h *HybridProvider) RevisionCacheLockTimeout() (time.Duration, error) {
	return configured(h.crd.RevisionCacheLockTimeout, h.legacy.RevisionCacheLockTimeout)
}

func (h *HybridProvider) StreamedManifestMaxExtractedSize() (int64, error) {
	return configured(h.crd.StreamedManifestMaxExtractedSize, h.legacy.StreamedManifestMaxExtractedSize)
}

func (h *HybridProvider) StreamedManifestMaxTarSize() (int64, error) {
	return configured(h.crd.StreamedManifestMaxTarSize, h.legacy.StreamedManifestMaxTarSize)
}

func (h *HybridProvider) SubmoduleEnabled() (bool, error) {
	return configured(h.crd.SubmoduleEnabled, h.legacy.SubmoduleEnabled)
}
