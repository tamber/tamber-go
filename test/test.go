package test

import (
	"fmt"
	tamber "github.com/tamber/tamber-go"
	"github.com/tamber/tamber-go/actor"
)

func Test() {
	tamber.DefaultKey = "UJVoOJrSoU4FfXpmM9R6"
	a, err := actor.Create(&tamber.ActorParams{
		Id: "68753A444D6F",
		Behaviors: &[]tamber.ActorBehavior{
			ActorBehavior{
				Behavior: "like",
				Item:     "9F45B8EK",
				Value:    0.23,
				Created:  1446417346,
			},
		},
	})
	if err != nil {
		fmt.Println(err)
	}
}

func EngieVarTest() {
	myengie := tamber.NewEngie("UJVoOJrSoU4FfXpmM9R6")
	a, err := myengie.actor.Create(&tamber.ActorParams{
		Id: "68753A444D6F",
		Behaviors: &[]ActorBehavior{
			ActorBehavior{
				Behavior: "like",
				Item:     "9F45B8EK",
				Value:    0.23,
				Created:  1446417346,
			},
		},
	})
	if err != nil {
		fmt.Println(err)
	}
}
