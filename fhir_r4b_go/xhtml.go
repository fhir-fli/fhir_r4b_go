package fhir_r4b_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
)

// FhirXhtml represents the FHIR primitive type "xhtml".
type FhirXhtml struct {
	Value   *string  `json:"value,omitempty"`
	Element *Element `json:"_value,omitempty"`
}

// NewFhirXhtml creates a new FhirXhtml instance after validation.
func NewFhirXhtml(input string, element *Element) (*FhirXhtml, error) {
	if err := validateXhtml(input); err != nil {
		return nil, err
	}
	return &FhirXhtml{
		Value:   &input,
		Element: element,
	}, nil
}

// NewFhirXhtmlFromMap creates a FhirXhtml instance from a map.
func NewFhirXhtmlFromMap(data map[string]interface{}) (*FhirXhtml, error) {
	valueStr, ok := data["value"].(string)
	if !ok {
		return nil, errors.New("invalid or missing value for FhirXhtml")
	}

	if err := validateXhtml(valueStr); err != nil {
		return nil, fmt.Errorf("invalid XHTML value for FhirXhtml: %v", err)
	}

	xhtmlInstance := &FhirXhtml{Value: &valueStr}

	if elementData, ok := data["_value"].(map[string]interface{}); ok {
		xhtmlInstance.Element = &Element{}
		if err := mapToStruct(elementData, xhtmlInstance.Element); err != nil {
			return nil, fmt.Errorf("failed to parse _value for FhirXhtml: %v", err)
		}
	}

	return xhtmlInstance, nil
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
	return validateElement(root, true)
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

// MarshalJSON serializes FhirXhtml into JSON.
func (f *FhirXhtml) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{}
	if *f.Value != "" {
		data["value"] = f.Value
	}
	if f.Element != nil {
		data["_value"] = f.Element
	}
	return json.Marshal(data)
}

func (f *FhirXhtml) UnmarshalJSON(data []byte) error {
	temp := struct {
		Value   string   `json:"value"`
		Element *Element `json:"_value"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Validate XHTML content
	if err := validateXhtml(temp.Value); err != nil {
		return err
	}

	f.Value = &temp.Value
	f.Element = temp.Element
	return nil
}

// Clone creates a deep copy of FhirXhtml.
func (f *FhirXhtml) Clone() *FhirXhtml {
	if f == nil {
		return nil
	}
	return &FhirXhtml{
		Value:   f.Value,
		Element: f.Element.Clone(),
	}
}

// Equals checks equality between two FhirXhtml instances.
func (f *FhirXhtml) Equals(other *FhirXhtml) bool {
	if f == nil && other == nil {
		return true
	}
	if f == nil || other == nil {
		return false
	}
	return f.Value == other.Value && f.Element.Equals(other.Element)
}

// validateElement recursively validates an element's attributes and children.
func validateElement(element *xmlquery.Node, isRoot bool) error {
	for _, attr := range element.Attr {
		if !isAllowedAttribute(attr.Name.Local, isRoot) {
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
func isAllowedAttribute(attrName string, isRoot bool) bool {
	allowedAttributes := map[string]bool{
		"style": true, "class": true, "src": true, "href": true, "name": true,
		"alt": true, "title": true, "colspan": true, "rowspan": true, "width": true,
		"height": true, "align": true, "valign": true, "border": true, "span": true,
	}

	// Allow xmlns attribute for the root element.
	return (isRoot && attrName == "xmlns") || allowedAttributes[attrName]
}
