package configbus

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	appv1 "github.com/argoproj/argo-cd/v3/pkg/apis/application/v1alpha1"
)

// Tests apply an ArgoCDConfiguration only in fixtures — production requires the singleton CR.

func TestProvider_CRDRequiredForMappedSettings(t *testing.T) {
	crdObj := &appv1.ArgoCDConfiguration{
		Spec: appv1.ArgoCDConfigurationSpec{
			Controller: &appv1.ControllerConfig{
				Reconciliation: &appv1.ReconciliationConfig{
					Timeout:     &metav1.Duration{Duration: 45 * time.Second},
					HardTimeout: &metav1.Duration{Duration: 10 * time.Minute},
					Jitter:      &metav1.Duration{Duration: 15 * time.Second},
				},
				SelfHeal: &appv1.SelfHealConfig{
					Timeout: &metav1.Duration{Duration: 90 * time.Second},
				},
				ResourceHealthPersist: ptrBool(true),
			},
			InstallationID: "from-crd",
		},
	}
	crd := StaticCRDSource{Object: crdObj}
	p := NewCRDProvider(crd)

	timeout, err := p.ReconciliationTimeout(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 45*time.Second, timeout)

	hard, err := p.HardReconciliationTimeout(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 10*time.Minute, hard)

	jitter, err := p.ReconciliationJitter(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 15*time.Second, jitter)

	selfHeal, err := p.SelfHealTimeout(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 90*time.Second, selfHeal)

	persist, err := p.PersistResourceHealth(context.Background())
	require.NoError(t, err)
	assert.True(t, persist)

	installID, err := p.InstallationID(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "from-crd", installID)
}

func TestProvider_AbsentCRDErrorsForMappedSettings(t *testing.T) {
	p := NewCRDProvider(StaticCRDSource{})

	_, err := p.ReconciliationTimeout(context.Background())
	require.ErrorIs(t, err, ErrNotConfigured)

	_, err = p.SelfHealTimeout(context.Background())
	require.ErrorIs(t, err, ErrNotConfigured)
}

func TestProvider_PartialCRDOnlyOverridesSetFields(t *testing.T) {
	crdObj := &appv1.ArgoCDConfiguration{
		Spec: appv1.ArgoCDConfigurationSpec{
			Controller: &appv1.ControllerConfig{
				Reconciliation: &appv1.ReconciliationConfig{
					Timeout: &metav1.Duration{Duration: 30 * time.Second},
				},
			},
		},
	}
	p := NewCRDProvider(StaticCRDSource{Object: crdObj})

	timeout, err := p.ReconciliationTimeout(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 30*time.Second, timeout)

	_, err = p.HardReconciliationTimeout(context.Background())
	require.ErrorIs(t, err, ErrNotConfigured)
}

func ptrBool(b bool) *bool { return &b }
