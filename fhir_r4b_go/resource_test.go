package fhir_r4b_go

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp" // For detailed diff output
)

// TestSerializationFromFiles tests serialization and deserialization of FHIR resources from JSON files.
func TestSerializationFromFiles(t *testing.T) {
	// Path to the directory containing test JSON files
	testDataDir := "testdata"

	err := filepath.Walk(testDataDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip non-JSON files
		if info.IsDir() || filepath.Ext(path) != ".json" {
			return nil
		}

		// Read the JSON file
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("Failed to read file %s: %v", path, err)
		}

		// Detect the resource type
		var raw map[string]interface{}
		if err := json.Unmarshal(data, &raw); err != nil {
			t.Fatalf("Failed to parse JSON in file %s: %v", path, err)
		}

		resourceType, ok := raw["resourceType"].(string)
		if !ok {
			t.Fatalf("Missing or invalid 'resourceType' in file %s", path)
		}

		// Create a new resource instance
		resource := newResourceFromType(resourceType)
		if resource == nil {
			t.Fatalf("Unsupported resource type %s in file %s", resourceType, path)
		}

		// Deserialize the JSON into the resource
		if err := json.Unmarshal(data, resource); err != nil {
			t.Fatalf("Failed to deserialize JSON into %s for file %s: %v", resourceType, path, err)
		}

		// Serialize the resource back to JSON
		serializedData, err := json.Marshal(resource)
		if err != nil {
			t.Fatalf("Failed to serialize %s for file %s: %v", resourceType, path, err)
		}

		// Compare the original and re-serialized JSON
		if !jsonEqual(data, serializedData) {
			diff := cmp.Diff(prettyJSON(data), prettyJSON(serializedData))
			t.Fatalf("Serialization mismatch for file %s (%s):\n%s", path, resourceType, diff)
		}

		// Return an error to stop further file iteration after the first failure
		return nil
	})

	// Handle errors from walking the directory
	if err != nil {
		t.Fatalf("Error walking test data directory: %v", err)
	}
}

// jsonEqual compares two JSON byte slices for structural equality.
func jsonEqual(a, b []byte) bool {
	var objA, objB interface{}

	if err := json.Unmarshal(a, &objA); err != nil {
		return false
	}
	if err := json.Unmarshal(b, &objB); err != nil {
		return false
	}

	return objA == objB
}

// prettyJSON formats JSON byte slices into indented strings for better readability.
func prettyJSON(data []byte) string {
	var out bytes.Buffer
	if err := json.Indent(&out, data, "", "  "); err != nil {
		return string(data)
	}
	return out.String()
}
