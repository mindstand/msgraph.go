// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AddContentFooterAction undocumented
type AddContentFooterAction struct {
	InformationProtectionAction
	// UIElementName undocumented
	UIElementName *string `json:"uiElementName,omitempty"`
	// Text undocumented
	Text *string `json:"text,omitempty"`
	// FontName undocumented
	FontName *string `json:"fontName,omitempty"`
	// FontSize undocumented
	FontSize *int `json:"fontSize,omitempty"`
	// FontColor undocumented
	FontColor *string `json:"fontColor,omitempty"`
	// Alignment undocumented
	Alignment *ContentAlignment `json:"alignment,omitempty"`
	// Margin undocumented
	Margin *int `json:"margin,omitempty"`
}
