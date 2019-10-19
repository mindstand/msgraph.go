// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "encoding/json"

// WorkbookFilterCriteria undocumented
type WorkbookFilterCriteria struct {
	// Color undocumented
	Color *string `json:"color,omitempty"`
	// Criterion1 undocumented
	Criterion1 *string `json:"criterion1,omitempty"`
	// Criterion2 undocumented
	Criterion2 *string `json:"criterion2,omitempty"`
	// DynamicCriteria undocumented
	DynamicCriteria *string `json:"dynamicCriteria,omitempty"`
	// FilterOn undocumented
	FilterOn *string `json:"filterOn,omitempty"`
	// Icon undocumented
	Icon *WorkbookIcon `json:"icon,omitempty"`
	// Operator undocumented
	Operator *string `json:"operator,omitempty"`
	// Values undocumented
	Values json.RawMessage `json:"values,omitempty"`
}
