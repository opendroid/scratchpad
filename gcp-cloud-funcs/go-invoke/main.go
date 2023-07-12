package main

import (
	"github.com/opendroid/scratchpad/gcp-cloud-funcs/go-invoke/cf"
)

// Test a GCP cloud function call
var (
	// test URLS
	urls = map[string]string{
		"charge": "https://charge-api-cvqrtz3jva-uc.a.run.app",
		"atV1":   "https://us-central1-sau2020.cloudfunctions.net/log-test-v1",
		"atV2":   "https://log-test-cvqrtz3jva-uc.a.run.app",
	}
)

func main() {
	// Keep compiler happy
	_ = cf.AuthToken
}
