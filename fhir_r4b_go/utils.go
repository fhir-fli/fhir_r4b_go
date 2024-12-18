package fhir_r4b_go

import "time"

// Helper functions for creating pointers.
func intPtr(v int) *int { return &v }
func intPtrIfNotNil(v *int) *int {
	if v == nil {
		return nil
	}
	return intPtr(*v)
}

func int64Ptr(v int64) *int64 { return &v }
func int64PtrIfNotNil(v *int64) *int64 {
	if v == nil {
		return nil
	}
	return int64Ptr(*v)
}

func strPtrIfNotNil(v *string) *string {
	if v == nil {
		return nil
	}
	return strPtr(*v)
}

func strPtr(v string) *string { return &v }
func timePtr(t time.Time) *time.Time {
	return &t
}
