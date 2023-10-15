package client

import (
	aiplatform "cloud.google.com/go/aiplatform/apiv1"
	"cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/structpb"
	"time"
)

const ( // EndPointURL for the AI
	APIPort     = "443"
	EndPointURL = "projects/" + ProjectID + "/locations/us-central1/publishers/google/models/" + ModelID
)

// createPredictionRequest creates a prediction request from prompt
func createPredictionRequest(prompt string) *aiplatformpb.PredictRequest {
	// Set up prompt in instance
	instance, err := structpb.NewValue(map[string]interface{}{
		"prompt": prompt,
	})
	if err != nil {
		fmt.Printf("CreatePredictionRequest: instance: %v", err)
		return nil
	}

	// Set up configuration PB
	params, err := structpb.NewValue(map[string]interface{}{
		"temperature":     0.2,
		"maxOutputTokens": 1024,
		"topK":            40,
		"topP":            0.95,
	})
	if err != nil {
		fmt.Printf("CreatePredictionRequest: params: %v", err)
		return nil
	}

	// Create prediction request
	req := &aiplatformpb.PredictRequest{
		Endpoint:   EndPointURL,
		Instances:  []*structpb.Value{instance},
		Parameters: params,
	}

	return req
}

func GetPredictionFromClient(prompt string) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(60*time.Second))
	defer cancel()

	// Create AI platform prediction client
	client, err := aiplatform.NewPredictionClient(ctx, option.WithEndpoint(APIEndpoint+":"+APIPort))
	if err != nil {
		fmt.Printf("GetPrediction2: Unable to create client: %v", err)
		return
	}
	defer func() { _ = client.Close() }()

	if req := createPredictionRequest(prompt); req != nil {
		response, err := client.Predict(ctx, req)
		if err != nil {
			fmt.Printf("GetPredictionFromClient: Unable to get prediction: %v", err)
			return
		} else {
			respMap := response.Predictions[0].GetStructValue().AsMap()
			fmt.Printf("Response: %v\n", respMap["content"])
		}
	}
}
