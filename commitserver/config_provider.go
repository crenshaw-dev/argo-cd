package commitserver

import (
	"github.com/argoproj/argo-cd/v3/util/configbus"
)

// InitConfigProvider wires the CRD-backed configbus provider.
func (a *ArgoCDCommitServer) InitConfigProvider(crd configbus.CRDSource) {
	if crd == nil {
		crd = configbus.TestCommitserverCRDSource()
	}
	a.configProvider = configbus.NewCRDProvider(crd)
}
