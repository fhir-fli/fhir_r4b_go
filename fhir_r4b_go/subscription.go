// Code generated by FHIR Generator. DO NOT EDIT.

package fhir_r4b_go

import (
	"encoding/json")

// Subscription
// The subscription resource is used to define a push-based subscription from a server to another system. Once a subscription is registered with the server, the server checks every resource that is created or updated, and if the resource matches the given criteria, it sends a message on the defined "channel" so that another system can take an appropriate action.
type Subscription struct {
	DomainResource
	Id *FhirString `json:"id,omitempty"`
	Meta *FhirMeta `json:"meta,omitempty"`
	ImplicitRules *FhirUri `json:"implicitrules,omitempty"`
	Language *CommonLanguages `json:"language,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Contained []*Resource `json:"contained,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Status *SubscriptionStatusCodes `json:"status,omitempty"`
	Contact []*ContactPoint `json:"contact,omitempty"`
	End *FhirInstant `json:"end,omitempty"`
	Reason *FhirString `json:"reason,omitempty"`
	Criteria *FhirString `json:"criteria,omitempty"`
	Error *FhirString `json:"error,omitempty"`
	Channel *SubscriptionChannel `json:"channel,omitempty"`
}

// NewSubscription creates a new Subscription instance
func NewSubscription() *Subscription {
	return &Subscription{}
}

// FromJSON populates Subscription from JSON data
func (m *Subscription) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts Subscription to JSON data
func (m *Subscription) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of Subscription
func (m *Subscription) Clone() *Subscription {
	if m == nil { return nil }
	return &Subscription{
		Id: m.Id.Clone(),
		Meta: m.Meta.Clone(),
		ImplicitRules: m.ImplicitRules.Clone(),
		Language: m.Language.Clone(),
		Text: m.Text.Clone(),
		Contained: cloneSlices(m.Contained),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Status: m.Status.Clone(),
		Contact: cloneSlices(m.Contact),
		End: m.End.Clone(),
		Reason: m.Reason.Clone(),
		Criteria: m.Criteria.Clone(),
		Error: m.Error.Clone(),
		Channel: m.Channel.Clone(),
	}
}

// Equals checks for equality with another Subscription instance
func (m *Subscription) Equals(other *Subscription) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !m.Meta.Equals(other.Meta) { return false }
	if !m.ImplicitRules.Equals(other.ImplicitRules) { return false }
	if !m.Language.Equals(other.Language) { return false }
	if !m.Text.Equals(other.Text) { return false }
	if !compareSlices(m.Contained, other.Contained) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Status.Equals(other.Status) { return false }
	if !compareSlices(m.Contact, other.Contact) { return false }
	if !m.End.Equals(other.End) { return false }
	if !m.Reason.Equals(other.Reason) { return false }
	if !m.Criteria.Equals(other.Criteria) { return false }
	if !m.Error.Equals(other.Error) { return false }
	if !m.Channel.Equals(other.Channel) { return false }
	return true
}

// SubscriptionChannel
// Details where to send notifications when resources are received that meet the criteria.
type SubscriptionChannel struct {
	BackboneElement
	Id *FhirString `json:"id,omitempty"`
	Extension_ []*FhirExtension `json:"extension,omitempty"`
	ModifierExtension []*FhirExtension `json:"modifierextension,omitempty"`
	Type *SubscriptionChannelType `json:"type,omitempty"`
	Endpoint *FhirUrl `json:"endpoint,omitempty"`
	Payload *FhirCode `json:"payload,omitempty"`
	Header []*FhirString `json:"header,omitempty"`
}

// NewSubscriptionChannel creates a new SubscriptionChannel instance
func NewSubscriptionChannel() *SubscriptionChannel {
	return &SubscriptionChannel{}
}

// FromJSON populates SubscriptionChannel from JSON data
func (m *SubscriptionChannel) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToJSON converts SubscriptionChannel to JSON data
func (m *SubscriptionChannel) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// Clone creates a deep copy of SubscriptionChannel
func (m *SubscriptionChannel) Clone() *SubscriptionChannel {
	if m == nil { return nil }
	return &SubscriptionChannel{
		Id: m.Id.Clone(),
		Extension_: cloneSlices(m.Extension_),
		ModifierExtension: cloneSlices(m.ModifierExtension),
		Type: m.Type.Clone(),
		Endpoint: m.Endpoint.Clone(),
		Payload: m.Payload.Clone(),
		Header: cloneSlices(m.Header),
	}
}

// Equals checks for equality with another SubscriptionChannel instance
func (m *SubscriptionChannel) Equals(other *SubscriptionChannel) bool {
	if m == nil && other == nil { return true }
	if m == nil || other == nil { return false }
	if !m.Id.Equals(other.Id) { return false }
	if !compareSlices(m.Extension_, other.Extension_) { return false }
	if !compareSlices(m.ModifierExtension, other.ModifierExtension) { return false }
	if !m.Type.Equals(other.Type) { return false }
	if !m.Endpoint.Equals(other.Endpoint) { return false }
	if !m.Payload.Equals(other.Payload) { return false }
	if !compareSlices(m.Header, other.Header) { return false }
	return true
}
