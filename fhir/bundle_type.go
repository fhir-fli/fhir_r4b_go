// Code generated by FHIR Generator. DO NOT EDIT.
// Indicates the purpose of a bundle - how it is intended to be used.
package fhir_r4b_go

// BundleType represents the FHIR ValueSet as an enumeration
type BundleType int

const (
	// Document: The bundle is a document. The first resource is a Composition.
	BundleType_Document BundleType = iota
	// Message: The bundle is a message. The first resource is a MessageHeader.
	BundleType_Message BundleType = iota
	// Transaction: The bundle is a transaction - intended to be processed by a server as an atomic commit.
	BundleType_Transaction BundleType = iota
	// Transaction Response: The bundle is a transaction response. Because the response is a transaction response, the transaction has succeeded, and all responses are error free.
	BundleType_Transaction_response BundleType = iota
	// Batch: The bundle is a set of actions - intended to be processed by a server as a group of independent actions.
	BundleType_Batch BundleType = iota
	// Batch Response: The bundle is a batch response. Note that as a batch, some responses may indicate failure and others success.
	BundleType_Batch_response BundleType = iota
	// History List: The bundle is a list of resources from a history interaction on a server.
	BundleType_History BundleType = iota
	// Search Results: The bundle is a list of resources returned as a result of a search/query interaction, operation, or message.
	BundleType_Searchset BundleType = iota
	// Collection: The bundle is a set of resources collected into a single package for ease of distribution that imposes no processing obligations or behavioral rules beyond persistence.
	BundleType_Collection BundleType = iota
)

// String converts the enum to its string representation
func (e BundleType) String() string {
	switch e {
	case BundleType_Document: return "Document"
	case BundleType_Message: return "Message"
	case BundleType_Transaction: return "Transaction"
	case BundleType_Transaction_response: return "Transaction Response"
	case BundleType_Batch: return "Batch"
	case BundleType_Batch_response: return "Batch Response"
	case BundleType_History: return "History List"
	case BundleType_Searchset: return "Search Results"
	case BundleType_Collection: return "Collection"
	default: return "Unknown"
	}
}
