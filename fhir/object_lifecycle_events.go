// Code generated by FHIR Generator. DO NOT EDIT.
// This example FHIR value set is comprised of lifecycle event codes. The FHIR Actor value set is based on DICOM Audit Message, ParticipantObjectDataLifeCycle; ISO Standard, TS 21089-2017;
package fhir_r4b_go

// ObjectLifecycleEvents represents the FHIR ValueSet as an enumeration
type ObjectLifecycleEvents int

const (
	// Origination / Creation: 
	ObjectLifecycleEvents_Value1 ObjectLifecycleEvents = iota
	// Import / Copy: 
	ObjectLifecycleEvents_Value2 ObjectLifecycleEvents = iota
	// Amendment: 
	ObjectLifecycleEvents_Value3 ObjectLifecycleEvents = iota
	// Verification: 
	ObjectLifecycleEvents_Value4 ObjectLifecycleEvents = iota
	// Translation: 
	ObjectLifecycleEvents_Value5 ObjectLifecycleEvents = iota
	// Access / Use: 
	ObjectLifecycleEvents_Value6 ObjectLifecycleEvents = iota
	// De-identification: 
	ObjectLifecycleEvents_Value7 ObjectLifecycleEvents = iota
	// Aggregation / summarization / derivation: 
	ObjectLifecycleEvents_Value8 ObjectLifecycleEvents = iota
	// Report: 
	ObjectLifecycleEvents_Value9 ObjectLifecycleEvents = iota
	// Export: 
	ObjectLifecycleEvents_Value10 ObjectLifecycleEvents = iota
	// Disclosure: 
	ObjectLifecycleEvents_Value11 ObjectLifecycleEvents = iota
	// Receipt of disclosure: 
	ObjectLifecycleEvents_Value12 ObjectLifecycleEvents = iota
	// Archiving: 
	ObjectLifecycleEvents_Value13 ObjectLifecycleEvents = iota
	// Logical deletion: 
	ObjectLifecycleEvents_Value14 ObjectLifecycleEvents = iota
	// Permanent erasure / Physical destruction: 
	ObjectLifecycleEvents_Value15 ObjectLifecycleEvents = iota
	// Access/View Record Lifecycle Event: Occurs when an agent causes the system to obtain and open a record entry for inspection or review.
	ObjectLifecycleEvents_Access ObjectLifecycleEvents = iota
	// Add Legal Hold Record Lifecycle Event: Occurs when an agent causes the system to tag or otherwise indicate special access management and suspension of record entry deletion/destruction, if deemed relevant to a lawsuit or which are reasonably anticipated to be relevant or to fulfill organizational policy under the legal doctrine of “duty to preserve”.
	ObjectLifecycleEvents_Hold ObjectLifecycleEvents = iota
	// Amend (Update) Record Lifecycle Event: Occurs when an agent makes any change to record entry content currently residing in storage considered permanent (persistent).
	ObjectLifecycleEvents_Amend ObjectLifecycleEvents = iota
	// Archive Record Lifecycle Event: Occurs when an agent causes the system to create and move archive artifacts containing record entry content, typically to long-term offline storage.
	ObjectLifecycleEvents_Archive ObjectLifecycleEvents = iota
	// Attest Record Lifecycle Event: Occurs when an agent causes the system to capture the agent’s digital signature (or equivalent indication) during formal validation of record entry content.
	ObjectLifecycleEvents_Attest ObjectLifecycleEvents = iota
	// Decrypt Record Lifecycle Event: Occurs when an agent causes the system to decode record entry content from a cipher.
	ObjectLifecycleEvents_Decrypt ObjectLifecycleEvents = iota
	// De-Identify (Anononymize) Record Lifecycle Event: Occurs when an agent causes the system to scrub record entry content to reduce the association between a set of identifying data and the data subject in a way that might or might not be reversible.
	ObjectLifecycleEvents_Deidentify ObjectLifecycleEvents = iota
	// Deprecate Record Lifecycle Event: Occurs when an agent causes the system to tag record entry(ies) as obsolete, erroneous or untrustworthy, to warn against its future use.
	ObjectLifecycleEvents_Deprecate ObjectLifecycleEvents = iota
	// Destroy/Delete Record Lifecycle Event: Occurs when an agent causes the system to permanently erase record entry content from the system.
	ObjectLifecycleEvents_Destroy ObjectLifecycleEvents = iota
	// Disclose Record Lifecycle Event: Occurs when an agent causes the system to release, transfer, provision access to, or otherwise divulge record entry content.
	ObjectLifecycleEvents_Disclose ObjectLifecycleEvents = iota
	// Encrypt Record Lifecycle Event: Occurs when an agent causes the system to encode record entry content in a cipher.
	ObjectLifecycleEvents_Encrypt ObjectLifecycleEvents = iota
	// Extract Record Lifecycle Event: Occurs when an agent causes the system to selectively pull out a subset of record entry content, based on explicit criteria.
	ObjectLifecycleEvents_Extract ObjectLifecycleEvents = iota
	// Link Record Lifecycle Event: Occurs when an agent causes the system to connect related record entries.
	ObjectLifecycleEvents_Link ObjectLifecycleEvents = iota
	// Merge Record Lifecycle Event: Occurs when an agent causes the system to combine or join content from two or more record entries, resulting in a single logical record entry.
	ObjectLifecycleEvents_Merge ObjectLifecycleEvents = iota
	// Originate/Retain Record Lifecycle Event: Occurs when an agent causes the system to: a) initiate capture of potential record content, and b) incorporate that content into the storage considered a permanent part of the health record.
	ObjectLifecycleEvents_Originate ObjectLifecycleEvents = iota
	// Pseudonymize Record Lifecycle Event: Occurs when an agent causes the system to remove record entry content to reduce the association between a set of identifying data and the data subject in a way that may be reversible.
	ObjectLifecycleEvents_Pseudonymize ObjectLifecycleEvents = iota
	// Re-activate Record Lifecycle Event: Occurs when an agent causes the system to recreate or restore full status to record entries previously deleted or deprecated.
	ObjectLifecycleEvents_Reactivate ObjectLifecycleEvents = iota
	// Receive/Retain Record Lifecycle Event: Occurs when an agent causes the system to a) initiate capture of data content from elsewhere, and b) incorporate that content into the storage considered a permanent part of the health record.
	ObjectLifecycleEvents_Receive ObjectLifecycleEvents = iota
	// Re-identify Record Lifecycle Event: Occurs when an agent causes the system to restore information to data that allows identification of information source and/or information subject.
	ObjectLifecycleEvents_Reidentify ObjectLifecycleEvents = iota
	// Remove Legal Hold Record Lifecycle Event: Occurs when an agent causes the system to remove a tag or other cues for special access management had required to fulfill organizational policy under the legal doctrine of “duty to preserve”.
	ObjectLifecycleEvents_Unhold ObjectLifecycleEvents = iota
	// Report (Output) Record Lifecycle Event: Occurs when an agent causes the system to produce and deliver record entry content in a particular form and manner.
	ObjectLifecycleEvents_Report ObjectLifecycleEvents = iota
	// Restore Record Lifecycle Event: Occurs when an agent causes the system to recreate record entries and their content from a previous created archive artefact.
	ObjectLifecycleEvents_Restore ObjectLifecycleEvents = iota
	// Transform/Translate Record Lifecycle Event: Occurs when an agent causes the system to change the form, language or code system used to represent record entry content.
	ObjectLifecycleEvents_Transform ObjectLifecycleEvents = iota
	// Transmit Record Lifecycle Event: Occurs when an agent causes the system to send record entry content from one (EHR/PHR/other) system to another.
	ObjectLifecycleEvents_Transmit ObjectLifecycleEvents = iota
	// Unlink Record Lifecycle Event: Occurs when an agent causes the system to disconnect two or more record entries previously connected, rendering them separate (disconnected) again.
	ObjectLifecycleEvents_Unlink ObjectLifecycleEvents = iota
	// Unmerge Record Lifecycle Event: Occurs when an agent causes the system to reverse a previous record entry merge operation, rendering them separate again.
	ObjectLifecycleEvents_Unmerge ObjectLifecycleEvents = iota
	// Verify Record Lifecycle Event: Occurs when an agent causes the system to confirm compliance of data or data objects with regulations, requirements, specifications, or other imposed conditions based on organizational policy.
	ObjectLifecycleEvents_Verify ObjectLifecycleEvents = iota
)

// String converts the enum to its string representation
func (e ObjectLifecycleEvents) String() string {
	switch e {
	case ObjectLifecycleEvents_Value1: return "Origination / Creation"
	case ObjectLifecycleEvents_Value2: return "Import / Copy"
	case ObjectLifecycleEvents_Value3: return "Amendment"
	case ObjectLifecycleEvents_Value4: return "Verification"
	case ObjectLifecycleEvents_Value5: return "Translation"
	case ObjectLifecycleEvents_Value6: return "Access / Use"
	case ObjectLifecycleEvents_Value7: return "De-identification"
	case ObjectLifecycleEvents_Value8: return "Aggregation / summarization / derivation"
	case ObjectLifecycleEvents_Value9: return "Report"
	case ObjectLifecycleEvents_Value10: return "Export"
	case ObjectLifecycleEvents_Value11: return "Disclosure"
	case ObjectLifecycleEvents_Value12: return "Receipt of disclosure"
	case ObjectLifecycleEvents_Value13: return "Archiving"
	case ObjectLifecycleEvents_Value14: return "Logical deletion"
	case ObjectLifecycleEvents_Value15: return "Permanent erasure / Physical destruction"
	case ObjectLifecycleEvents_Access: return "Access/View Record Lifecycle Event"
	case ObjectLifecycleEvents_Hold: return "Add Legal Hold Record Lifecycle Event"
	case ObjectLifecycleEvents_Amend: return "Amend (Update) Record Lifecycle Event"
	case ObjectLifecycleEvents_Archive: return "Archive Record Lifecycle Event"
	case ObjectLifecycleEvents_Attest: return "Attest Record Lifecycle Event"
	case ObjectLifecycleEvents_Decrypt: return "Decrypt Record Lifecycle Event"
	case ObjectLifecycleEvents_Deidentify: return "De-Identify (Anononymize) Record Lifecycle Event"
	case ObjectLifecycleEvents_Deprecate: return "Deprecate Record Lifecycle Event"
	case ObjectLifecycleEvents_Destroy: return "Destroy/Delete Record Lifecycle Event"
	case ObjectLifecycleEvents_Disclose: return "Disclose Record Lifecycle Event"
	case ObjectLifecycleEvents_Encrypt: return "Encrypt Record Lifecycle Event"
	case ObjectLifecycleEvents_Extract: return "Extract Record Lifecycle Event"
	case ObjectLifecycleEvents_Link: return "Link Record Lifecycle Event"
	case ObjectLifecycleEvents_Merge: return "Merge Record Lifecycle Event"
	case ObjectLifecycleEvents_Originate: return "Originate/Retain Record Lifecycle Event"
	case ObjectLifecycleEvents_Pseudonymize: return "Pseudonymize Record Lifecycle Event"
	case ObjectLifecycleEvents_Reactivate: return "Re-activate Record Lifecycle Event"
	case ObjectLifecycleEvents_Receive: return "Receive/Retain Record Lifecycle Event"
	case ObjectLifecycleEvents_Reidentify: return "Re-identify Record Lifecycle Event"
	case ObjectLifecycleEvents_Unhold: return "Remove Legal Hold Record Lifecycle Event"
	case ObjectLifecycleEvents_Report: return "Report (Output) Record Lifecycle Event"
	case ObjectLifecycleEvents_Restore: return "Restore Record Lifecycle Event"
	case ObjectLifecycleEvents_Transform: return "Transform/Translate Record Lifecycle Event"
	case ObjectLifecycleEvents_Transmit: return "Transmit Record Lifecycle Event"
	case ObjectLifecycleEvents_Unlink: return "Unlink Record Lifecycle Event"
	case ObjectLifecycleEvents_Unmerge: return "Unmerge Record Lifecycle Event"
	case ObjectLifecycleEvents_Verify: return "Verify Record Lifecycle Event"
	default: return "Unknown"
	}
}
