package main

import (
	"fmt"

	"github.com/joyrexus/slack"
)

func main() {
	client := slack.NewClient("xoxb-14316137920-HnYQrHOS7pOtkLKEDRh1i30w")
	for _, m := range client.GetUsers().Members {
		fmt.Println(m.Name)
	}
}
