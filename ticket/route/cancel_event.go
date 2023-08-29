package route

import (
	"github.com/asynkron/protoactor-go/actor"
	"github.com/labstack/echo/v4"
	"github.com/ytake/protoactor-go-http/ticket/message"
	"github.com/ytake/protoactor-go-http/ticket/ticket_actor"
	"net/http"
	"time"
)

// CancelEvent is a router
type CancelEvent struct {
	actor *ticket_actor.Root
}

// NewCancelEvent create new instance
func NewCancelEvent(actor *ticket_actor.Root) *CancelEvent {
	return &CancelEvent{
		actor: actor,
	}
}

// retrievePID is a router
func (ce *CancelEvent) retrievePID() *actor.PID {
	return ce.actor.PID()
}

func (ce *CancelEvent) Handle(c echo.Context) error {
	ev := &message.CancelEvent{Name: c.Param("name")}
	future := ce.actor.ActorSystem().Root.RequestFuture(ce.retrievePID(), ev, 2*time.Second)
	r, err := future.Result()
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	switch r.(type) {
	case *message.Cancel:
		return c.NoContent(http.StatusOK)
	default:
		return c.NoContent(http.StatusBadRequest)
	}
}