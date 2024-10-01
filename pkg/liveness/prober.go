package liveness

import (
	"context"
	"time"
)

// Prober can handle services/applications.
type Prober interface {
	// Watch starts a goroutine which watches for new messages in the liveness channel and respond.
	// This method must be called on main service thread before the lock action will be called.
	// Example:
	//        livenessProbe := liveness.NewProbe(ctx)
	// 	      usefulService := useful.NewService(ctx)
	//        livenessProbe.Watch(usefulService)
	//        wg.Add(1)
	//        go usefulService.DoWork(wg)
	//        // service is alive and probe will be report it
	// 		  isAlive := livenessProbe.IsAlive() // (bool) true
	//		  wg.Wait()
	// 		  isAlive := livenessProbe.IsAlive() // (bool) false (but depends on service which  respond on IsAlive question.)
	Watch(services ...Servicer)
	// IsAlive checks wether is target service (a service which called Watch method) alive.
	IsAlive() bool
}

type Prob struct {
	askCh   chan context.Context
	respCh  chan bool
	timeout time.Duration
}

func NewProbe(timeout time.Duration) *Prob {
	return &Prob{
		askCh:   make(chan context.Context),
		respCh:  make(chan bool),
		timeout: timeout,
	}
}

func (p *Prob) Watch(services ...Servicer) {
	go func() {
		for ctx := range p.askCh {
			isAlive := true
			for _, service := range services {
				isAlive = isAlive && service.IsAlive(ctx)
			}
			p.respCh <- isAlive
		}
	}()
}

func (p *Prob) IsAlive() bool {
	ctx, cancel := context.WithTimeout(context.Background(), p.timeout)
	defer cancel()

	go func() {
		p.askCh <- ctx
	}()

	select {
	case <-ctx.Done():
		return false
	case r := <-p.respCh:
		return r
	}
}
