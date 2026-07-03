package core_http_types

import (
	"encoding/json"
	"testing"
)

func TestNullableUnmarshalJSON(t *testing.T) {
	type request struct {
		Description Nullable[string] `json:"description"`
	}

	t.Run("field omitted", func(t *testing.T) {
		var req request
		if err := json.Unmarshal([]byte(`{}`), &req); err != nil {
			t.Fatalf("unmarshal request: %v", err)
		}

		if req.Description.Set {
			t.Fatal("expected omitted field to remain unset")
		}
		if req.Description.Value != nil {
			t.Fatal("expected omitted field value to remain nil")
		}
	})

	t.Run("null", func(t *testing.T) {
		var req request
		if err := json.Unmarshal([]byte(`{"description": null}`), &req); err != nil {
			t.Fatalf("unmarshal request: %v", err)
		}

		if !req.Description.Set {
			t.Fatal("expected null field to be set")
		}
		if req.Description.Value != nil {
			t.Fatal("expected null field value to be nil")
		}
	})

	t.Run("value", func(t *testing.T) {
		var req request
		if err := json.Unmarshal([]byte(`{"description": "text"}`), &req); err != nil {
			t.Fatalf("unmarshal request: %v", err)
		}

		if !req.Description.Set {
			t.Fatal("expected value field to be set")
		}
		if req.Description.Value == nil {
			t.Fatal("expected value field value to be non-nil")
		}
		if *req.Description.Value != "text" {
			t.Fatalf("expected value %q, got %q", "text", *req.Description.Value)
		}
	})
}
