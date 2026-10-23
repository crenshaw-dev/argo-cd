package repository

import (
	"context"
	"errors"
)

// ErrShuttingDown is the context cause on a request cancelled because the
// repo server is draining and then stopping.
var ErrShuttingDown = errors.New("argocd-repo-server is shutting down")

// manifestGenHookKey is a test seam. A hook stored under this key runs once
// checkout has finished and manifest generation is about to start.
type manifestGenHookKey struct{}

func withManifestGenHook(ctx context.Context, hook func()) context.Context {
	return context.WithValue(ctx, manifestGenHookKey{}, hook)
}
