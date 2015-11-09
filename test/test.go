package test

import (
	"fmt"
	"tamber-go-master/actor"
)

func Test() {
	tamber.DefaultKey = "UJVoOJrSoU4FfXpmM9R6"
	a, err := actor.Create(&ActorParams{
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

func EngieVarTest() {
	myengie := tamber.NewEngie("UJVoOJrSoU4FfXpmM9R6")
	a, err := myengie.actor.Create(&ActorParams{
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
