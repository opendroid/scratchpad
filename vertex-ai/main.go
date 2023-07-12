package main

import (
	_ "embed"
	"fmt"

	"github.com/opendroid/scratchpad/vertex-ai/chat/client"
)

//go:embed chat/client/chat.txt
var prompt string

func main() {
	data := client.SetBisonPayload(prompt)
	client.GetPrediction(data)
	fmt.Println("Done")
}
