"use strict";

// Example of JSON:
const jsonStr1 = `{"name": "JSON String 1", "description": "This is a valid JSON 1"}`;
const jsonStr2 = `{"name": "JSON String 2", description: "This is a invalid JSON 2"}`; // Error: needs "description" to be a in quotes
const json1 = {"name": "JSON 1", "description": "This is a valid JSON 1"};
const json2 = {"name": "JSON 2", description: "This is a valid JSON 2"};

// getJSON:
function getJSONFromString(jsonString) {
  try {
    return JSON.parse(jsonString);
  } catch (e) {
    return {error: "Invalid JSON string", e: e.message}
  }
}

// getJSON2:
function getJSONFromJSON(json) {
  try {
    return JSON.parse(JSON.stringify(json));
  } catch (e) {
    return {error: "Invalid JSON", e: e.message}
  }
}

// getStringifyJSON: prints a JSON object as string, in one line
function getStringifyJSON(json) {
  try {
    return JSON.stringify(json);
  } catch (e) {
    console.log(e);
    return {error: "Invalid JSON", e: e}
  }
}

// test getJSONFromString:
console.log(`JSON string1:` + getStringifyJSON(getJSONFromString(jsonStr1)));
console.log(`JSON string2: ${getStringifyJSON(getJSONFromString(jsonStr2))}`);
// test getJSONFromJSON:
console.log(`JSON 1: ${getStringifyJSON(getJSONFromJSON(json1))}`);
console.log(`JSON 2: ${getStringifyJSON(getJSONFromJSON(json2))}`);

var t = ''