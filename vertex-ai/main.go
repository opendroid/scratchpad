package main

import (
	_ "embed"
	"fmt"

	"github.com/opendroid/scratchpad/vertex-ai/chat/client"
)

//go:embed chat/client/chat.txt
var prompt string

func main() {
	fmt.Println("Getting prediction from post...")
	data := client.SetBisonPayloadJSON(prompt)
	client.GetPrediction(data)
	fmt.Println("Getting prediction from client...")
	client.GetPredictionFromClient(prompt)
}
