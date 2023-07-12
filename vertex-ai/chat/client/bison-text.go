package client

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2/google"
)

const (
	// API_ENDPOINT is the API endpoint for the Vertex AI Prediction API.
	API_ENDPOINT = "us-central1-aiplatform.googleapis.com"
	// PROJECT_ID is the GCP project ID.
	PROJECT_ID = "cloud-large-language-models"
	// ENDPOINT_ID is the Vertex AI Endpoint ID.
	ENDPOINT_ID = "4511608470067216384"
	// PREDICT_URL is the URL for the Vertex AI Prediction API.
	PREDICT_URL = "https://" + API_ENDPOINT + "/v1/projects/" + PROJECT_ID + "/locations/us-central1/endpoints/" + ENDPOINT_ID + ":predict"
)

// BisonPayload is the payload for the Vertex AI Prediction API.
type BisonPayload struct {
	Instances []struct {
		Content string `json:"content"`
	} `json:"instances"`
	Parameters struct {
		Temperature   float64 `json:"temperature"`
		MaxDecodeStep int     `json:"maxDecodeSteps"`
		TopP          float64 `json:"topP"`
		TopK          int     `json:"topK"`
	} `json:"parameters"`
}

// GetToken returns the access token.
func GetToken() string {
	ctx := context.Background()
	scopes := []string{"https://www.googleapis.com/auth/cloud-platform"}
	// Option1: Retrieve the default credentials from the environment.
	credentials, err := google.FindDefaultCredentials(ctx, scopes...)
	// Option2: Retrieve the default credentials from the key.json
	// credentials, err := google.CredentialsFromJSON(ctx, []byte(key), scopes...)
	if err != nil {
		fmt.Printf("GetToken: Unable to find default credentials: %v", err)
		return ""
	}
	// Create a new access token from the retrieved credentials.
	token, err := credentials.TokenSource.Token()
	if err != nil {
		fmt.Printf("GetToken: Unable to retrieve token from token source: %v", err)
		return ""
	}
	return token.AccessToken
}

// SetBisonPayload returns a BisonPayload.
func SetBisonPayload(content string) *BisonPayload {
	bp := BisonPayload{}
	bp.Instances = append(bp.Instances, struct {
		Content string `json:"content"`
	}{Content: content})
	bp.Parameters.Temperature = 0.2
	bp.Parameters.MaxDecodeStep = 256
	bp.Parameters.TopP = 0.95
	bp.Parameters.TopK = 40
	return &bp
}

// GetPrediction returns the prediction.
func GetPrediction(payload *BisonPayload) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("GetPrediction: Unable to marshal payload: %v", err)
		return
	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", PREDICT_URL, body)
	if err != nil {
		fmt.Printf("GetPrediction: Unable to create request: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+GetToken())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("GetPrediction: Unable to execute request: %v", err)
	}
	defer resp.Body.Close()
	// Read body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("GetPrediction: Unable to read response body: %v", err)
		return
	}
	fmt.Printf("GetPrediction: Status: %v", resp.Status)
	fmt.Printf("GetPrediction: Body: %v", string(bodyBytes))
}
