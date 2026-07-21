package controller

import (
	"github.com/argoproj/argo-cd/v3/util/configbus"
)

// InitConfigProvider wires the CRD-backed configbus provider.
func (c *notificationController) InitConfigProvider(crd configbus.CRDSource) {
	if crd == nil {
		crd = configbus.TestNotificationsCRDSource()
	}
	c.configProvider = configbus.NewCRDProvider(crd)
}
