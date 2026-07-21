package controllers

import (
	"github.com/argoproj/argo-cd/v3/applicationset/generators"
	"github.com/argoproj/argo-cd/v3/applicationset/services"
	"github.com/argoproj/argo-cd/v3/util/configbus"
	"github.com/argoproj/argo-cd/v3/util/settings"
)

// InitConfigProvider wires the CRD-backed configbus provider after the
// reconciler struct is built.
func (r *ApplicationSetReconciler) InitConfigProvider(_ *settings.SettingsManager, scmConfig *generators.SCMConfig, argoCDService *services.ArgoCDService, crd configbus.CRDSource) {
	r.scmConfig = scmConfig
	r.argoCDService = argoCDService
	r.configProvider = configbus.NewCRDProvider(crd)
	generators.SetDefaultRequeueProvider(r.configProvider)
	if scmConfig != nil {
		scmConfig.SetConfigProvider(r.configProvider)
	}
	if argoCDService != nil {
		argoCDService.SetConfigProvider(r.configProvider)
	}
}

// ensureConfigProvider lazily wires a CRD test fixture. Production always
// calls InitConfigProvider; unit tests that construct the reconciler directly
// hit this path so configProvider getters remain usable.
func (r *ApplicationSetReconciler) ensureConfigProvider() {
	if r.configProvider == nil {
		//nolint:staticcheck // SA1019: capture reconciler fields into CRD test fixture
		r.configProvider = configbus.NewCRDProvider(configbus.TestApplicationsetCRDSourceFromFields(configbus.ApplicationsetCRDLegacyFields{
			Policy:                      string(r.Policy),
			EnablePolicyOverride:        r.EnablePolicyOverride,
			ApplicationSetNamespaces:    r.ApplicationSetNamespaces,
			EnableProgressiveSyncs:      r.EnableProgressiveSyncs,
			GlobalPreservedAnnotations:  r.GlobalPreservedAnnotations,
			GlobalPreservedLabels:       r.GlobalPreservedLabels,
			MaxResourcesStatusCount:     r.MaxResourcesStatusCount,
			MetricsAddr:                 r.MetricsAddr,
			MetricsApplicationsetLabels: r.MetricsApplicationsetLabels,
			ProbeAddr:                   r.ProbeAddr,
			WebhookAddr:                 r.WebhookAddr,
		}))
		generators.SetDefaultRequeueProvider(r.configProvider)
		if r.scmConfig != nil {
			r.scmConfig.SetConfigProvider(r.configProvider)
		}
		if r.argoCDService != nil {
			r.argoCDService.SetConfigProvider(r.configProvider)
		}
	}
}
