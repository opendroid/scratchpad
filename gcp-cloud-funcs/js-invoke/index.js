"use strict";
const { GoogleAuth } = require('google-auth-library');
const auth = new GoogleAuth({ keyFilename: "/workspaces/scratchpad/gcp-cloud-funcs-invoke/key.json" });
// urls are secured test functions deployed to GCP
const urls = {
  charge: "https://charge-api-cvqrtz3jva-uc.a.run.app",
  atV1: "https://us-central1-sau2020.cloudfunctions.net/log-test-v1",
  atV2: "https://log-test-cvqrtz3jva-uc.a.run.app",
};

const testUrl = urls.charge;
const data = {
  "message": {
    "id": "1",
    "value": "Hi, where is my order?",
    "type": "MESSAGE_TYPE_CUSTOMER"
  },
  "thread_id": "0000-0000-0000",
  "prediction_type": "PREDICTION_TYPE_SUGGEST_RESPONSE"
};

// Created a request to the secured function
async function request(url, data) {
  const targetAudience = url;
  const client = await auth.getIdTokenClient(targetAudience);

  const res = await client.request({method: 'POST', url, data});
  return res.data;
}

request(testUrl, data).then(console.log).catch(err => {
    console.error(err.message);
    process.exitCode = 1;
});