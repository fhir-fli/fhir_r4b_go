// Code generated by FHIR Generator. DO NOT EDIT.
// The type of process where the audit event originated from.
package fhir_r4b_go

// AuditEventSourceType represents the FHIR ValueSet as an enumeration
type AuditEventSourceType int

const (
	// User Device: End-user display device, diagnostic device.
	AuditEventSourceType_Value1 AuditEventSourceType = iota
	// Data Interface: Data acquisition device or instrument.
	AuditEventSourceType_Value2 AuditEventSourceType = iota
	// Web Server: Web Server process or thread.
	AuditEventSourceType_Value3 AuditEventSourceType = iota
	// Application Server: Application Server process or thread.
	AuditEventSourceType_Value4 AuditEventSourceType = iota
	// Database Server: Database Server process or thread.
	AuditEventSourceType_Value5 AuditEventSourceType = iota
	// Security Server: Security server, e.g. a domain controller.
	AuditEventSourceType_Value6 AuditEventSourceType = iota
	// Network Device: ISO level 1-3 network component.
	AuditEventSourceType_Value7 AuditEventSourceType = iota
	// Network Router: ISO level 4-6 operating software.
	AuditEventSourceType_Value8 AuditEventSourceType = iota
	// Other: Other kind of device (defined by DICOM, but some other code/system can be used).
	AuditEventSourceType_Value9 AuditEventSourceType = iota
)

// String converts the enum to its string representation
func (e AuditEventSourceType) String() string {
	switch e {
	case AuditEventSourceType_Value1: return "User Device"
	case AuditEventSourceType_Value2: return "Data Interface"
	case AuditEventSourceType_Value3: return "Web Server"
	case AuditEventSourceType_Value4: return "Application Server"
	case AuditEventSourceType_Value5: return "Database Server"
	case AuditEventSourceType_Value6: return "Security Server"
	case AuditEventSourceType_Value7: return "Network Device"
	case AuditEventSourceType_Value8: return "Network Router"
	case AuditEventSourceType_Value9: return "Other"
	default: return "Unknown"
	}
}
