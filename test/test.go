package test

import (
	"fmt"
	tamber "github.com/tamber/tamber-go"
	"github.com/tamber/tamber-go/actor"
	"github.com/tamber/tamber-go/engine"
)

func Test() {
	tamber.DefaultKey = "UJVoOJrSoU4FfXpmM9R6"
	a, err := actor.Create(&tamber.ActorParams{
		Id: "68753A444D6F",
		Behaviors: &[]tamber.ActorBehavior{
			tamber.ActorBehavior{
				Behavior: "like",
				Item:     "9F45B8EK",
				Value:    0.23,
				Created:  1446417346,
			},
		},
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v", a)
	}
}

func EngieVarTest() {
	e := &engine.API{}
	e.Init("UJVoOJrSoU4FfXpmM9R6", nil)
	a, err := e.Actors.Create(&tamber.ActorParams{
		Id: "68753A444D6F",
		Behaviors: &[]tamber.ActorBehavior{
			tamber.ActorBehavior{
				Behavior: "like",
				Item:     "9F45B8EK",
				Value:    0.23,
				Created:  1446417346,
			},
		},
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v", a)
	}
}
