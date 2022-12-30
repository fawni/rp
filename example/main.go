package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fawni/rp"
	"github.com/fawni/rp/rpc"
)

func main() {
	c, err := rp.NewClient("DISCORD_APP_ID")
	if err != nil {
		log.Fatalln("Could not connect to discord rich presence client:", err)
	}

	now := time.Now()
	if err := c.SetActivity(&rpc.Activity{
		State:      "Hey!",
		Details:    "Running on rp.go!",
		LargeImage: "largeimageid",
		LargeText:  "This is the large image",
		SmallImage: "smallimageid",
		SmallText:  "And this is the small image",
		Party: &rpc.Party{
			ID:         "-1",
			Players:    15,
			MaxPlayers: 24,
		},
		Timestamps: &rpc.Timestamps{
			Start: &now,
		},
		Buttons: []*rpc.Button{
			{
				Label: "GitHub",
				Url:   "https://github.com/fawni/rp",
			},
		},
	}); err != nil {
		log.Fatalln("Could not set activity:", err)
	}

	// Discord will only show the presence if the app is running
	// Sleep for a few seconds to see the update
	fmt.Println("Sleeping...")
	time.Sleep(time.Second * 10)
	if err := c.ResetActivity(); err != nil {
		log.Fatalln("Could not reset activity:", err)
	}
	fmt.Println("Reset!")
	time.Sleep(time.Second * 10)
}
