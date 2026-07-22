package configbus_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/argoproj/argo-cd/v3/util/configbus"
)

// controllerResolvedMethods are Provider getters the application-controller
// CRD fixture must resolve after the CRD-only cutover.
var controllerResolvedMethods = map[string]struct{}{
	"AppInstanceLabelKey":           {},
	"ControllerHydrationProcessors": {},
	"ControllerOperationProcessors": {},
	"ControllerStatusProcessors":    {},
	"HardReconciliationTimeout":     {},
	"IgnoreNormalizerJQTimeout":     {},
	"IsImpersonationEnabled":        {},
	"IsImpersonationEnforced":       {},
	"MetricsClusterLabels":          {},
	"PersistResourceHealth":         {},
	"ReconciliationJitter":          {},
	"ReconciliationTimeout":         {},
	"RepoErrorGracePeriod":          {},
	"SelfHealRetry":                 {},
	"SelfHealTimeout":               {},
	"ServerSideDiff":                {},
	"SyncTimeout":                   {},
	"TrackingMethod":                {},
}

// TestControllerChainResolvesAllFields asserts the application-controller
// production Provider (CRD-only) resolves every controller-owned field getter
// without leaking ErrNotConfigured when the test fixture CR is present.
func TestControllerChainResolvesAllFields(t *testing.T) {
	t.Parallel()
	provider := configbus.NewCRDProvider(configbus.TestControllerCRDSource())

	pv := reflect.ValueOf(provider)
	pt := pv.Type()
	ctx := context.Background()
	for i := 0; i < pt.NumMethod(); i++ {
		m := pt.Method(i)
		if _, ok := controllerResolvedMethods[m.Name]; !ok {
			continue
		}
		// Method.Type includes the receiver; getters are (recv, ctx) -> (T, error).
		if m.Type.NumIn() != 2 || m.Type.NumOut() != 2 {
			continue
		}
		method := pv.Method(i)
		in := make([]reflect.Value, method.Type().NumIn())
		for j := range method.Type().NumIn() {
			if method.Type().In(j).String() == "context.Context" {
				in[j] = reflect.ValueOf(ctx)
				continue
			}
			in[j] = reflect.Zero(method.Type().In(j))
		}
		results := method.Call(in)
		err, _ := results[1].Interface().(error)
		require.NoError(t, err, "controller CRD fixture must resolve %s", m.Name)
		require.NotErrorIs(t, err, configbus.ErrNotConfigured, m.Name)
	}
}
