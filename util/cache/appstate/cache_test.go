package appstate

import (
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	. "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	cacheutil "github.com/argoproj/argo-cd/v2/util/cache"
)

type fixtures struct {
	*Cache
}

func newFixtures() *fixtures {
	return &fixtures{NewCache(
		cacheutil.NewCache(cacheutil.NewInMemoryCache(1*time.Hour)),
		1*time.Minute,
	)}
}

func TestCache_GetAppManagedResources(t *testing.T) {
	appNameID := cacheutil.NewAppID("my-appname", "argocd")
	otherAppID := cacheutil.NewAppID("other-appname", "argocd")

	cache := newFixtures().Cache
	// cache miss
	value := &[]*ResourceDiff{}
	err := cache.GetAppManagedResources(appNameID, value)
	assert.Equal(t, ErrCacheMiss, err)
	// populate cache
	err = cache.SetAppManagedResources(appNameID, []*ResourceDiff{{Name: "my-name"}})
	assert.NoError(t, err)
	// cache miss
	err = cache.GetAppManagedResources(otherAppID, value)
	assert.Equal(t, ErrCacheMiss, err)
	// cache hit
	err = cache.GetAppManagedResources(appNameID, value)
	assert.NoError(t, err)
	assert.Equal(t, &[]*ResourceDiff{{Name: "my-name"}}, value)
}

func TestCache_GetAppResourcesTree(t *testing.T) {
	appNameID := cacheutil.NewAppID("my-appname", "argocd")
	otherAppID := cacheutil.NewAppID("other-appname", "argocd")

	cache := newFixtures().Cache
	// cache miss
	value := &ApplicationTree{}
	err := cache.GetAppResourcesTree(appNameID, value)
	assert.Equal(t, ErrCacheMiss, err)
	// populate cache
	err = cache.SetAppResourcesTree(appNameID, &ApplicationTree{Nodes: []ResourceNode{{}}})
	assert.NoError(t, err)
	// cache miss
	err = cache.GetAppResourcesTree(otherAppID, value)
	assert.Equal(t, ErrCacheMiss, err)
	// cache hit
	err = cache.GetAppResourcesTree(appNameID, value)
	assert.NoError(t, err)
	assert.Equal(t, &ApplicationTree{Nodes: []ResourceNode{{}}}, value)
}

func TestAddCacheFlagsToCmd(t *testing.T) {
	cache, err := AddCacheFlagsToCmd(&cobra.Command{})()
	assert.NoError(t, err)
	assert.Equal(t, 1*time.Hour, cache.appStateCacheExpiration)
}
