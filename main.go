package main

import (
	"log"
	"sync"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/tkhrk1010/emile/emile"
	"github.com/tkhrk1010/emile/world"
)

func main() {
	// actor systemの起動
	system := actor.NewActorSystem()
	rootContext := system.Root

	// 世界の誕生
	worldActor := system.Root.Spawn(actor.PropsFromProducer(func() actor.Actor {
		return world.NewWorld()
	}))
	log.Print("world: ", worldActor.Id)

	// Emileの誕生
	emileActor := actor.PropsFromProducer(func() actor.Actor {
		return emile.NewNewbornEmile()
	})
	emilePID := rootContext.Spawn(emileActor)

	//
	// 世界がEmileに挨拶するシーン(原典未記載)
	var wg sync.WaitGroup
	wg.Add(1)
	
	go func() {
		defer wg.Done()
		
		worldHelloRequest := &world.HelloRequest{ Destination: emilePID }
		log.Printf("Say Hello request to world: %v\n", worldActor.Id)

		response, err := rootContext.RequestFuture(worldActor, worldHelloRequest, 5*time.Second).Result()

		if err != nil {
			log.Println("Hello request timeout or error:", err)
			return
		}

		log.Printf("Response from world: %v\n", response.(*world.HelloResponse).Word)

	}()

	wg.Wait()
}
