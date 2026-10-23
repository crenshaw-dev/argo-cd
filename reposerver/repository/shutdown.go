package repository

import "errors"

// ErrShuttingDown is the context cause on a request cancelled because the
// repo server is draining and then stopping.
var ErrShuttingDown = errors.New("argocd-repo-server is shutting down")
