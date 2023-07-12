package client

import (
	//lint:ignore SA1019
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"text/template"
)

// ModelType is the type of model to use for completion.
type ModelType string

const (
	// Davinci is most capable GPT-3 model. Can do any task the other models can do,
	// often with higher quality, longer output and better instruction-following.
	// Also supports inserting completions within text.
	Davinci ModelType = "text-davinci-003"
	// Curie is a smaller model that is faster and cheaper than Davinci.
	Curie ModelType = "text-curie-001"
	// Babbage is capable of straightforward tasks, very fast, and lower cost.
	Babbage ModelType = "text-babbage-001"
	// Ada is capable of very simple tasks, usually the fastest model in the GPT-3 series, and lowest cost.
	Ada ModelType = "text-ada-001"
)

// SampleGPTModel is the model to use for completion.
type SampleGPTModel struct {
	Type             ModelType `json:"model"`
	MaxTokens        int       `json:"max_tokens"`
	Description      string    `json:"description"`
	TrainingDataDate string    `json:"date"`
}

var (
	//go:embed text/davinci.txt
	davinviDescription string

	// DavinciModel is the model to use for completion.
	DavinciModel = SampleGPTModel{
		Type:             Davinci,
		MaxTokens:        4_000,
		Description:      davinviDescription,
		TrainingDataDate: "2021-07-01",
	}
	_ = DavinciModel

	//go:embed text/curie.txt
	curieDescription string

	// CurieModel is the  GPT3 model is very capable, faster and lower cost than Davinci.
	CurieModel = SampleGPTModel{
		Type:             Curie,
		MaxTokens:        2_048,
		Description:      curieDescription,
		TrainingDataDate: "2019-11-01",
	}
	_ = CurieModel

	//go:embed text/babbage.txt
	babbageDescription string

	// BabbageModel is GPT3 models that capable of straightforward tasks, very fast, and lower cost.
	BabbageModel = SampleGPTModel{
		Type:             Babbage,
		MaxTokens:        2_048,
		Description:      babbageDescription,
		TrainingDataDate: "2019-11-01",
	}
	_ = BabbageModel

	//go:embed text/ada.txt
	adaDescription string

	// AdaModel is GPT3 models that capable of very simple tasks, usually the fastest model in the GPT-3 series, and lowest cost.
	AdaModel = SampleGPTModel{
		Type:             Ada,
		MaxTokens:        2_048,
		Description:      adaDescription,
		TrainingDataDate: "2019-11-01",
	}

	// ModelsDisplayTemplate path to template, values are
	// /workspaces/scratchpad/openai/chat/client/text/models.tmpl
	ModelsDisplayTemplate = "/Users/ajayt/gocode/src/github.com/opendroid/scratchpad/openai/chat/client/text/models.tmpl"
)

// Permission is the permission to use for completion.
type Permission struct {
	ID                 string `json:"id,omitempty"`
	Object             string `json:"object,omitempty"`
	Created            uint64 `json:"created,omitempty"`
	AllowCreateEngine  bool   `json:"allow_create_engine,omitempty"`
	AllowSampling      bool   `json:"allow_sampling,omitempty"`
	AllowLogprobs      bool   `json:"allow_logprobs,omitempty"`
	AllowSearchIndices bool   `json:"allow_search_indices,omitempty"`
	AllowFineTuning    bool   `json:"allow_fine_tuning,omitempty"`
	Organization       string `json:"organization,omitempty"`
	Groups             string `json:"groups,omitempty"`
	IsBlocking         bool   `json:"is_blocking,omitempty"`
}

// GPTModel is the model to use for completion.
type GPTModel struct {
	ID      string `json:"id,omitempty"`
	Object  string `json:"object,omitempty"`
	Created uint64 `json:"created,omitempty"`
	OwnedBy string `json:"owned_by,omitempty"`
	Root    string `json:"root,omitempty"`
	Parent  string `json:"parent,omitempty"`
}

// Models response of: curl https://api.openai.com/v1/models -H 'Authorization: Bearer '
type Models struct {
	Object string     `json:"object,omitempty"`
	Model  []GPTModel `json:"data,omitempty"`
}

//go:embed text/models.json
var data string

// ListModels in OpenAI
func ListModels() {

	var GPTReturnedModels Models
	if err := json.Unmarshal([]byte(data), &GPTReturnedModels); err != nil {
		fmt.Printf("ModelFunc: error: %s", err.Error())
		return
	}

	tmpl := template.Must(template.ParseFiles(ModelsDisplayTemplate))
	if err := tmpl.Execute(os.Stdout, GPTReturnedModels.Model); err != nil {
		fmt.Printf("ModelFunc: error: %s", err.Error())
		return
	}
}
