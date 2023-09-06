package notify

import (
	"context"
	"errors"
	"github.com/Rustixir/skyview"
	"github.com/jfyne/live"
)

var (
	ErrTimeout = errors.New("timeout")
)

type Local struct {
	evtChan chan live.Params
}

// NewLocal creates a new Local engine.
func NewLocal(bufSize int) *Local {
	return &Local{
		evtChan: make(chan live.Params, bufSize),
	}
}

// Send an event on a topic.
func (p *Local) Send(ctx context.Context, evt live.Params) error {
	select {
	case p.evtChan <- evt:
		return nil
	case <-ctx.Done():
		return ErrTimeout
	}
}

func (p *Local) StartListener(ctx context.Context, s live.Socket) {
	go func(chn chan live.Params) {
		for {
			select {
			case evt := <-chn:
				s.Self(ctx, skyview.NotifyEvent, evt)
			case <-ctx.Done():
				return
			}
		}
	}(p.evtChan)
}
