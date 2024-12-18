package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/antchfx/xmlquery"
)

// FhirXhtml represents the FHIR primitive type "xhtml".
type FhirXhtml struct {
	Value   string   `json:"value,omitempty"`
	Element *Element `json:"_value,omitempty"`
}

// NewFhirXhtml creates a new FhirXhtml instance after validation.
func NewFhirXhtml(input string, element *Element) (*FhirXhtml, error) {
	if err := validateXhtml(input); err != nil {
		return nil, err
	}
	return &FhirXhtml{Value: input, Element: element}, nil
}

// validateXhtml validates the XHTML string structure.
func validateXhtml(xhtml string) error {
	doc, err := xmlquery.Parse(strings.NewReader(xhtml))
	if err != nil {
		return errors.New("invalid XHTML: " + err.Error())
	}

	// Validate the root element.
	root := doc.SelectElement("*")
	if root == nil || !isAllowedElement(root.Data) {
		return errors.New("invalid XHTML: root element not allowed")
	}

	// Validate attributes and child elements recursively.
	if err := validateElement(root, true); err != nil {
		return err
	}

	return nil
}

// isAllowedElement checks if the element is allowed in XHTML.
func isAllowedElement(name string) bool {
	allowedElements := map[string]bool{
		"div": true, "p": true, "b": true, "i": true, "em": true, "strong": true,
		"ul": true, "ol": true, "li": true, "span": true, "br": true, "a": true,
		"img": true, "h1": true, "h2": true, "h3": true, "h4": true, "h5": true,
		"h6": true, "table": true, "thead": true, "tbody": true, "tr": true,
		"th": true, "td": true, "pre": true, "code": true, "blockquote": true,
		"hr": true, "caption": true, "colgroup": true, "col": true, "dl": true,
		"dt": true, "dd": true, "big": true, "small": true, "tt": true,
	}
	return allowedElements[name]
}

// validateElement recursively validates an element's attributes and children.
func validateElement(element *xmlquery.Node, isRoot bool) error {
	for _, attr := range element.Attr {
		if !isAllowedAttribute(attr.Name.Local, element.Data, isRoot) {
			return errors.New("invalid attribute in element " + element.Data)
		}
	}

	// Recursively validate child elements.
	for child := element.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == xmlquery.ElementNode {
			if !isAllowedElement(child.Data) {
				return errors.New("invalid child element: " + child.Data)
			}
			if err := validateElement(child, false); err != nil {
				return err
			}
		}
	}

	return nil
}

// isAllowedAttribute validates attributes for a given element.
func isAllowedAttribute(attrName, elementName string, isRoot bool) bool {
	allowedAttributes := map[string]bool{
		"style": true, "class": true, "src": true, "href": true, "name": true,
		"alt": true, "title": true, "colspan": true, "rowspan": true, "width": true,
		"height": true, "align": true, "valign": true, "border": true, "span": true,
	}

	// Allow xmlns attribute for the root element.
	if isRoot && attrName == "xmlns" {
		return true
	}

	return allowedAttributes[attrName]
}

// MarshalJSON serializes FhirXhtml into JSON.
func (f *FhirXhtml) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{}
	if f.Value != "" {
		data["value"] = f.Value
	}
	if f.Element != nil {
		data["_value"] = f.Element
	}
	return json.Marshal(data)
}

// UnmarshalJSON deserializes JSON into FhirXhtml.
func (f *FhirXhtml) UnmarshalJSON(data []byte) error {
	temp := struct {
		Value   string   `json:"value"`
		Element *Element `json:"_value"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	if err := validateXhtml(temp.Value); err != nil {
		return err
	}

	f.Value = temp.Value
	f.Element = temp.Element
	return nil
}

// String returns the XHTML content as a string.
func (f *FhirXhtml) String() string {
	return f.Value
}

// Clone creates a deep copy of FhirXhtml.
func (f *FhirXhtml) Clone() *FhirXhtml {
	var elementCopy *Element
	if f.Element != nil {
		elementCopy = f.Element.Copy()
	}
	return &FhirXhtml{
		Value:   f.Value,
		Element: elementCopy,
	}
}

// Equal checks equality between two FhirXhtml instances.
func (f *FhirXhtml) Equal(other *FhirXhtml) bool {
	return f.Value == other.Value && compareElements(f.Element, other.Element)
}
