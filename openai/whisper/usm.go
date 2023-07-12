package main

import (
	"context"
	"fmt"
	"time"

	speech "cloud.google.com/go/speech/apiv2"
	"cloud.google.com/go/speech/apiv2/speechpb"
)

// TranscribeGCSBucketUsingUSM transcribes a file in the GCP storage bucket using USM.
// See [Google Speech-to-text v2]
// [Google Speech-to-text v2]: https://cloud.google.com/speech-to-text/v2/docs/batch-recognize
func TranscribeGCSBucketUsingUSM() error {
	ctx := context.Background()
	client, err := speech.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("TranscribeGCSBucketUsingUSM: NewClient: %s", err.Error())
	}
	defer func() { _ = client.Close() }()
	req := &speechpb.BatchRecognizeRequest{
		RecognitionOutputConfig: &speechpb.RecognitionOutputConfig{
			Output: &speechpb.RecognitionOutputConfig_GcsOutputConfig{
				GcsOutputConfig: &speechpb.GcsOutputConfig{
					Uri: GCPOutBucketID,
				},
			},
		},
		Files: []*speechpb.BatchRecognizeFileMetadata{
			{
				AudioSource: &speechpb.BatchRecognizeFileMetadata_Uri{
					Uri: "gs://ivr-audio-storage/WCR/USCanadaEarnerSupport/00011ea1-519c-4aee-a0b5-1045fffe6d10.WAV",
				},
			},
		},
	}
	op, err := client.BatchRecognize(ctx, req)
	if err != nil {
		return fmt.Errorf("TranscribeGCSBucketUsingUSM: BatchRecognize: %s", err.Error())
	}
	resp, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("TranscribeGCSBucketUsingUSM: Wait: %s", err.Error())
	}
	// TODO: Use resp.
	fmt.Printf("TranscribeGCSBucketUsingUSM: %+v\n", resp)
	return nil
}

/**
curl -X POST \
    -H "Authorization: Bearer $(gcloud auth application-default print-access-token)" \
    -H "Content-Type: application/json; charset=utf-8" \
    --data "{\"languageCodes\": \"en-US\", \"model\": \"latest_long\"}" \
    https://speech.googleapis.com/v2/projects/ajayt-cx-experiments/locations/global/recognizers?recognizer_id=RECOGNIZER_ID

*/

func TranscribeGCSBuckets() {
	// Read all files from GCP storage
	if err := TranscribeGCSBucketUsingUSM(); err != nil {
		fmt.Printf("main Error: %s\n", err.Error())
		return
	}
	objects, err := GetGCSFileNames()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	fmt.Printf("Found %d files\n", len(objects))

	for i, obj := range objects {
		out := CreateTxtFileName(obj, "data/out")
		if IsAlreadyTranscribed(out) {
			fmt.Printf("Skipping: exists: %d:%s\n", i+1, out)
			continue
		}
		fmt.Printf("Processing: %d:%s\n", i+1, obj)
		start := time.Now()
		if text, err := TranscribeGCSFileUsingWhisper(obj); err != nil {
			fmt.Printf("Error: %d:%s:%s\n", i+1, obj, err.Error())
		} else {
			duration := time.Since(start)
			if err := SaveTranscription(out, text); err != nil {
				fmt.Printf("Error: %d:%s:%s\n", i+1, out, err.Error())
			} else {
				fmt.Printf("Saved: %d:%s:%s\n", i+1, out, duration)
			}
		}
		time.Sleep(5 * time.Second)
	}
}
