package repository

import (
	"time"

	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/argoproj/argo-cd/v3/util/configbus"
)

// Ensure Service satisfies configbus.ReposerverLegacy.
var _ configbus.ReposerverLegacy = (*Service)(nil)

// Legacy* accessors implement configbus.ReposerverLegacy for the configbus
// Provider only. Product code and tests must read via configProvider.*; these
// methods are the sole allowed readers of the deprecated struct fields.

//nolint:staticcheck // SA1019: sole allowed reader of deprecated OCIMediaTypes
func (s *Service) LegacyOCIMediaTypes() []string {
	return s.initConstants.OCIMediaTypes
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated ParallelismLimit
func (s *Service) LegacyParallelismLimit() int64 {
	return s.initConstants.ParallelismLimit
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated PauseGenerationAfterFailedGenerationAttempts
func (s *Service) LegacyPauseGenerationAfterFailedGenerationAttempts() int {
	return s.initConstants.PauseGenerationAfterFailedGenerationAttempts
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated PauseGenerationOnFailureForMinutes
func (s *Service) LegacyPauseGenerationOnFailureForMinutes() int {
	return s.initConstants.PauseGenerationOnFailureForMinutes
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated PauseGenerationOnFailureForRequests
func (s *Service) LegacyPauseGenerationOnFailureForRequests() int {
	return s.initConstants.PauseGenerationOnFailureForRequests
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated SubmoduleEnabled
func (s *Service) LegacySubmoduleEnabled() bool {
	return s.initConstants.SubmoduleEnabled
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated MaxCombinedDirectoryManifestsSize
func (s *Service) LegacyMaxCombinedDirectoryManifestsSize() resource.Quantity {
	return s.initConstants.MaxCombinedDirectoryManifestsSize
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated CMPTarExcludedGlobs
func (s *Service) LegacyCMPTarExcludedGlobs() []string {
	return s.initConstants.CMPTarExcludedGlobs
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated AllowOutOfBoundsSymlinks
func (s *Service) LegacyAllowOutOfBoundsSymlinks() bool {
	return s.initConstants.AllowOutOfBoundsSymlinks
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated StreamedManifestMaxExtractedSize
func (s *Service) LegacyStreamedManifestMaxExtractedSize() int64 {
	return s.initConstants.StreamedManifestMaxExtractedSize
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated StreamedManifestMaxTarSize
func (s *Service) LegacyStreamedManifestMaxTarSize() int64 {
	return s.initConstants.StreamedManifestMaxTarSize
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated HelmManifestMaxExtractedSize
func (s *Service) LegacyHelmManifestMaxExtractedSize() int64 {
	return s.initConstants.HelmManifestMaxExtractedSize
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated HelmRegistryMaxIndexSize
func (s *Service) LegacyHelmRegistryMaxIndexSize() int64 {
	return s.initConstants.HelmRegistryMaxIndexSize
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated OCIManifestMaxExtractedSize
func (s *Service) LegacyOCIManifestMaxExtractedSize() int64 {
	return s.initConstants.OCIManifestMaxExtractedSize
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated DisableOCIManifestMaxExtractedSize
func (s *Service) LegacyDisableOCIManifestMaxExtractedSize() bool {
	return s.initConstants.DisableOCIManifestMaxExtractedSize
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated DisableHelmManifestMaxExtractedSize
func (s *Service) LegacyDisableHelmManifestMaxExtractedSize() bool {
	return s.initConstants.DisableHelmManifestMaxExtractedSize
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated IncludeHiddenDirectories
func (s *Service) LegacyIncludeHiddenDirectories() bool {
	return s.initConstants.IncludeHiddenDirectories
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated CMPUseManifestGeneratePaths
func (s *Service) LegacyCMPUseManifestGeneratePaths() bool {
	return s.initConstants.CMPUseManifestGeneratePaths
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated EnableBuiltinGitConfig
func (s *Service) LegacyEnableBuiltinGitConfig() bool {
	return s.initConstants.EnableBuiltinGitConfig
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated HelmUserAgent
func (s *Service) LegacyHelmUserAgent() string {
	return s.initConstants.HelmUserAgent
}

//nolint:staticcheck // SA1019: sole allowed reader of deprecated HelmChartCacheExpiration
func (s *Service) LegacyHelmChartCacheExpiration() time.Duration {
	return s.initConstants.HelmChartCacheExpiration
}

func (s *Service) LegacyRepoCacheExpiration() time.Duration {
	return s.cache.LegacyRepoCacheExpiration()
}

func (s *Service) LegacyRevisionCacheExpiration() time.Duration {
	return s.cache.LegacyRevisionCacheExpiration()
}

func (s *Service) LegacyRevisionCacheLockTimeout() time.Duration {
	return s.cache.LegacyRevisionCacheLockTimeout()
}
