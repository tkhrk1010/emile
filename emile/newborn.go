package emile

import (
	"github.com/asynkron/protoactor-go/actor"
)

type NewbornEmile struct{}

func NewNewbornEmile() *NewbornEmile {
	return &NewbornEmile{}
}

type HelloRequest struct {
	Sender *actor.PID
	Word  string
}

type HelloResponse struct {
	Word string
}

func (e *NewbornEmile) Receive(context actor.Context) {
	switch context.Message().(type) {
	case *HelloRequest:
		context.Respond(&HelloResponse{Word: "ぅああ!"})
	}
}
