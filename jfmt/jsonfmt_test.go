package jfmt

import (
	"bytes"
	"testing"
)

var (
	original = []byte(`{
  "some_key": "hello",
  "another_key": {
    "widgets": [1, 2, 3, 4],
    "thing": "amajigger"
  }
}`)
	compacted = []byte(`{"some_key":"hello","another_key":{"widgets":[1,2,3,4],"thing":"amajigger"}}`)
	pretty    = []byte(`{
    "some_key": "hello",
    "another_key": {
        "widgets": [
            1,
            2,
            3,
            4
        ],
        "thing": "amajigger"
    }
}`)
)

func TestPretty(t *testing.T) {
	out, err := Pretty(original)
	if err != nil {
		t.Fatalf("%v", err)
	}

	if !bytes.Equal(out, pretty) {
		t.Fatalf("expected\n%s\n\nbut have\n%s", string(pretty), string(out))
	}
}

func TestCompact(t *testing.T) {
	out, err := Compact(original)
	if err != nil {
		t.Fatalf("%v", err)
	}

	if !bytes.Equal(out, compacted) {
		t.Fatalf("expected\n%s\n\nbut have\n%s", string(compacted), string(out))
	}
}
