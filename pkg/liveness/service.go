package liveness

import "context"

type Servicer interface {
	IsAlive(ctx context.Context) bool
}
