package main

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"io"
	"os"
	"strings"
	"time"
)

var (
	GCPBucketID    = "ivr-audio-storage"
	GCPOutBucketID = "gs://ivr-audio-storage/WCR/USCanadaEarnerSupport/GOOG-USM"
	TwentyFiveMb   = 25 * 1024 * 1024
)

// GetGCSFileNames returns a list of all files in the GCP storage bucket
func GetGCSFileNames() ([]string, error) {
	// Read all files from GCP storage
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()
	bucket, err := storage.NewClient(ctx)
	if err != nil {
		return []string{}, fmt.Errorf("ReadGCSFiles: %s", err.Error())
	}
	query := &storage.Query{Prefix: "WCR/USCanadaEarnerSupport"}
	it := bucket.Bucket(GCPBucketID).Objects(ctx, query)
	var files []string
	for it != nil {
		obj, err := it.Next()
		if err == iterator.Done {
			return files, nil
		}
		if err != nil {
			return files, fmt.Errorf("ReadGCSFiles: %s", err.Error())
		}
		if strings.Contains(obj.Name, ".WAV") && obj.Size < int64(TwentyFiveMb) {
			files = append(files, obj.Name)
		}
	}
	return files, nil
}

// ReadGCSObject returns the contents of a file in the GCP storage bucket
func ReadGCSObject(fileName string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()
	bucket, err := storage.NewClient(ctx)
	if err != nil {
		return []byte{}, fmt.Errorf("ReadGCSObject: %s", err.Error())
	}
	rc, err := bucket.Bucket(GCPBucketID).Object(fileName).NewReader(ctx)
	if err != nil {
		return []byte{}, fmt.Errorf("ReadGCSObject: %s", err.Error())
	}
	defer func() { _ = rc.Close() }()
	data, err := io.ReadAll(rc)
	if err != nil {
		return []byte{}, fmt.Errorf("ReadGCSObject: %s", err.Error())
	}
	return data, nil
}

// TranscribeGCSFileUsingWhisper transcribes a file in the GCP storage bucket
func TranscribeGCSFileUsingWhisper(fileName string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	bucket, err := storage.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("ReadGCSObject: %s", err.Error())
	}
	rc, err := bucket.Bucket(GCPBucketID).Object(fileName).NewReader(ctx)
	if err != nil {
		return "", fmt.Errorf("ReadGCSObject: %s", err.Error())
	}
	defer func() { _ = rc.Close() }()
	return TranscribeUsingWhisper(rc)
}

// CreateTxtFileName creates a txt file name from a wav file name
func CreateTxtFileName(fileName, dir string) string {
	values := strings.FieldsFunc(fileName, func(r rune) bool { return r == '/' || r == '.' })
	if len(values) < 2 {
		return "temp.txt"
	}
	name := values[len(values)-2]
	if dir[len(dir)-1] != '/' {
		return dir + "/" + name + ".txt"
	}
	return dir + name + ".txt"
}

// SaveTranscription saves the transcription to a txt file in local storage
func SaveTranscription(fileName string, transcription string) error {
	// create a file
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("SaveTranscription: create file: %s", err.Error())
	}
	defer func() { _ = file.Close() }()
	// write to file
	_, err = file.WriteString(transcription)
	if err != nil {
		return fmt.Errorf("SaveTranscription: Save error: %s", err.Error())
	}
	if err := file.Close(); err != nil {
		return fmt.Errorf("SaveTranscription: Close error: %s", err.Error())
	}

	return nil
}
