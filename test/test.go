package test

import (
	"fmt"
	tamber "github.com/tamber/tamber-go"
	"github.com/tamber/tamber-go/actor"
	// "github.com/tamber/tamber-go/behavior"
	"github.com/tamber/tamber-go/engine"
	"github.com/tamber/tamber-go/item"
	"github.com/tamber/tamber-go/property"
)

func Test() {
	tamber.DefaultKey = "UhurbcIiFDIt6yuzDOAO"
	a, err := actor.Remove(&tamber.ActorParams{
		Id: "2197054086",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v", a)
	}

	// a, err = actor.Retrieve(&tamber.ActorParams{
	// 	Id: "2197054086",
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%v", a)
	// }

	p, err := property.Create(&tamber.PropertyParams{
		Name: "length",
		Type: "float",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nProperty: %v\n", p)
	}

	i, err := item.AddProperties(&tamber.ItemParams{
		Id: "HZNP",
		Properties: &map[string]interface{}{
			"length": 4,
		},
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nItem: %v\n", i)
	}

	// b, err := behavior.Create(&tamber.BehaviorParams{
	// 	Name:         "friend",
	// 	Desirability: 0.5,
	// 	Type:         "decay",
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%v", b)
	// }

}

func EngieVarTest() {
	e := &engine.API{}
	e.Init("UhurbcIiFDIt6yuzDOAO", nil)
	a, err := e.Actors.AddBehaviors(&tamber.ActorParams{
		Id: "2197054086",
		Behaviors: &[]tamber.ActorBehavior{
			tamber.ActorBehavior{
				Behavior: "like",
				Item:     "HZNP",
				Value:    1.0,
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
