package client

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2/google"
	"io"
	"net/http"
)

const (
	// APIEndpoint is the API endpoint for the Vertex AI Prediction API.
	APIEndpoint = "us-central1-aiplatform.googleapis.com"
	// ProjectID is the GCP project ID.
	ProjectID = "conversational-ai-digital-bots"
	// ModelID is the Vertex AI Endpoint ID.
	ModelID = "text-bison@001"
	// PredictURL is the URL for the Vertex AI Prediction API.
	PredictURL = "https://" + APIEndpoint + "/v1/projects/" + ProjectID + "/locations/us-central1/publishers/google/models/" + ModelID + ":predict"
)

// BisonPayloadJSON is the payload for the Vertex AI Prediction API.
type BisonPayloadJSON struct {
	Instances []struct {
		Content string `json:"content"`
	} `json:"instances"`
	Parameters struct {
		Temperature     float64 `json:"temperature"`
		MaxOutputTokens int     `json:"maxOutputTokens"`
		CandidateCount  int     `json:"candidateCount"`
		TopP            float64 `json:"topP"`
		TopK            int     `json:"topK"`
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

// SetBisonPayloadJSON returns a BisonPayloadJSON.
func SetBisonPayloadJSON(content string) *BisonPayloadJSON {
	bp := BisonPayloadJSON{}
	bp.Instances = append(bp.Instances, struct {
		Content string `json:"content"`
	}{Content: content})
	bp.Parameters.Temperature = 0.2
	bp.Parameters.MaxOutputTokens = 1024
	bp.Parameters.CandidateCount = 1
	bp.Parameters.TopP = 0.95
	bp.Parameters.TopK = 40
	return &bp
}

// GetPrediction returns the prediction.
func GetPrediction(payload *BisonPayloadJSON) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("GetPrediction: Unable to marshal payload: %v", err)
		return
	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", PredictURL, body)
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
	defer func() { _ = resp.Body.Close() }()
	// Read body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("GetPrediction: Unable to read response body: %v", err)
		return
	}
	fmt.Printf("GetPrediction: Status: %v\n", resp.Status)
	fmt.Printf("GetPrediction: Body: %v\n", string(bodyBytes))
}
