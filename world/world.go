package world

import (
	"log"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/tkhrk1010/emile/emile"
)

type World struct{}

type HelloRequest struct {
	Destination *actor.PID
}

type HelloResponse struct{
	Word string
}

func NewWorld() *World {
	return &World{}
}

func (w *World) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *HelloRequest:

		// Emileに挨拶する
		worldSayHelloToEmile := &emile.HelloRequest{Word: "Hello, Emile!"}
		log.Printf("World: %v say hello to Emile\n", context.Self().Id)
		
		destPID := msg.Destination
		response, err := context.RequestFuture(destPID, worldSayHelloToEmile, 5*time.Second).Result()

		if err != nil {
			log.Println("World Hello request timeout or error: ", err)
			return
		}

		if resp, ok := response.(*emile.HelloResponse); ok {
			log.Printf("Response from Emile: %s\n", resp.Word)
		}
		context.Respond(&HelloResponse{ Word: "He is fine."})
	}
}
