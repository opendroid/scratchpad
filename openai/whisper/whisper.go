package main

import (
	"context"
	"fmt"
	"github.com/rakyll/openai-go/audio"
	"io"
	"os"
	"time"

	openai "github.com/rakyll/openai-go"
)

// TranscribeUsingWhisper creates a transcription from the given location-data using OpenAI whisper API
func TranscribeUsingWhisper(filename io.Reader) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	openApiKey := os.Getenv("OPENAI_API_KEY")
	fmt.Printf("OPENAI_API_KEY: %s\n", openApiKey)
	session := openai.NewSession(openApiKey)
	client := audio.NewClient(session, "")
	resp, err := client.CreateTranscription(ctx, &audio.CreateTranscriptionParams{
		Language:    "en",
		Audio:       filename,
		AudioFormat: "mp3",
		Prompt:      `This transcript is between a Uber customer care Agent and Uber driver partner. It has two speakers, identify the speakers as Agent and Driver Partner. Split the output one line per speaker.\n\n`,
	})
	if err != nil {
		return "", fmt.Errorf("TranscribeUsingWhisper: error transcribing file: %v", err)
	}
	return resp.Text, nil
}

func IsAlreadyTranscribed(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

// TranscribeSafetyFiles transcribes all files in the local storage
func TranscribeSafetyFiles() {
	entry, err := os.ReadDir("location-data/safety")
	if err != nil {
		fmt.Printf("TranscribeSafetyFiles: %s", err.Error())
		return
	}
	for _, file := range entry {
		name := file.Name()
		if f, err := os.Open("location-data/safety/" + name); err != nil {
			fmt.Printf("TranscribeSafetyFiles: %s", err.Error())
			continue
		} else {
			outName := CreateTxtFileName(name, "location-data/safety")
			if IsAlreadyTranscribed(outName) {
				continue
			}
			if transcription, err := TranscribeUsingWhisper(f); err != nil {
				fmt.Printf("TranscribeSafetyFiles: %s", err.Error())
			} else {
				if err := SaveTranscription(outName, transcription); err != nil {
					fmt.Printf("TranscribeSafetyFiles: %s", err.Error())
				}
			}
		}
	}
}
