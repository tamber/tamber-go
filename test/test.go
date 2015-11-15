package test

import (
	"fmt"
	tamber "github.com/tamber/tamber-go"
	"github.com/tamber/tamber-go/actor"
	// "github.com/tamber/tamber-go/discover"
	"github.com/tamber/tamber-go/engine"
	"github.com/tamber/tamber-go/item"
	"github.com/tamber/tamber-go/property"
)

func Test() {
	tamber.DefaultKey = "sBW1WHQ4bP4Ryfz3AQOo"
	fmt.Printf("\n\nResults\n------\n\n")
	// a, err := actor.Remove(&tamber.ActorParams{
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
		fmt.Printf("Property: %v", *p)
	}
	p, err = property.Create(&tamber.PropertyParams{
		Name: "color",
		Type: "string",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Property: %v", *p)
	}

	a, err := actor.Retrieve(&tamber.ActorParams{
		Id: "2197054086",
		GetRecs: &tamber.DiscoverParams{
			Filter: &map[string]interface{}{
				"or": []interface{}{
					map[string]interface{}{
						"lt": []interface{}{
							map[string]interface{}{
								"property": "length",
							},
							5.1,
						},
					},
					map[string]interface{}{
						"eq": []interface{}{
							map[string]interface{}{
								"property": "color",
							},
							"blue",
						},
					},
				},
			},
		},
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v", a)
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
		fmt.Printf("Item: %v", *i)
	}

	// d, err := discover.AddProperties(&tamber.ItemParams{
	// 	Id: "HZNP",
	// 	Properties: &map[string]interface{}{
	// 		"length": 4,
	// 	},
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("Item: %v", *i)
	// }
	fmt.Printf("\n\n")
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
	e.Init("80r2oX10Uw4XfZSxfh4O", nil)
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
