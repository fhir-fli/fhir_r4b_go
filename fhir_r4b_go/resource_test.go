package fhir_r4b_go

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// TestSerializationFromFiles tests serialization and deserialization of FHIR resources from JSON files.
func TestSerializationFromFiles(t *testing.T) {
	// Path to the directory containing test JSON files
	testDataDir := "testdata"

	// Iterate through all JSON files in the directory
	err := filepath.Walk(testDataDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip non-JSON files
		if info.IsDir() || filepath.Ext(path) != ".json" {
			return nil
		}

		t.Run(info.Name(), func(t *testing.T) {
			// Read the JSON file
			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("Failed to read file %s: %v", path, err)
			}

			// Detect the resource type
			var raw map[string]interface{}
			if err := json.Unmarshal(data, &raw); err != nil {
				t.Fatalf("Failed to parse JSON: %v", err)
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
				t.Fatalf("Failed to deserialize JSON into %s: %v", resourceType, err)
			}

			// Serialize the resource back to JSON
			serializedData, err := json.Marshal(resource)
			if err != nil {
				t.Fatalf("Failed to serialize %s: %v", resourceType, err)
			}

			// Compare the original and re-serialized JSON
			if !jsonEqual(data, serializedData) {
				t.Errorf("Mismatch between original and re-serialized JSON for %s:\nOriginal: %s\nRe-Serialized: %s", resourceType, data, serializedData)
			}
		})

		return nil
	})

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
