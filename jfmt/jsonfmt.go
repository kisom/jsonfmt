// Package jfmt contains JSON formatting functions split out from the
// main program to facilitate testing.
package jfmt

import (
	"bytes"
	"encoding/json"
)

// Pretty returns a pretty-printed version of the JSON data passed
// in.
func Pretty(jsonData []byte) ([]byte, error) {
	var buf = &bytes.Buffer{}
	err := json.Indent(buf, jsonData, "", "    ")
	if err != nil {
		return nil, err
	}
	return append(buf.Bytes(), 0xa), err
}

// Compact returns a compacted version of the JSON data passed in.
func Compact(jsonData []byte) ([]byte, error) {
	var buf = &bytes.Buffer{}
	err := json.Compact(buf, jsonData)
	return buf.Bytes(), err
}
