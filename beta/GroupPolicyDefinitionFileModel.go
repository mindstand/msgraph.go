// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// GroupPolicyDefinitionFile The entity represents an ADMX (Administrative Template) XML file. The ADMX file contains a collection of group policy definitions and their locations by category path. The group policy definition file also contains the languages supported as determined by the language dependent ADML (Administrative Template) language files.
type GroupPolicyDefinitionFile struct {
	Entity
	// DisplayName The localized friendly name of the ADMX file.
	DisplayName *string `json:"displayName,omitempty"`
	// Description The localized description of the policy settings in the ADMX file. The default value is empty.
	Description *string `json:"description,omitempty"`
	// LanguageCodes The supported language codes for the ADMX file.
	LanguageCodes []string `json:"languageCodes,omitempty"`
	// TargetPrefix Specifies the logical name that refers to the namespace within the ADMX file.
	TargetPrefix *string `json:"targetPrefix,omitempty"`
	// TargetNamespace Specifies the URI used to identify the namespace within the ADMX file.
	TargetNamespace *string `json:"targetNamespace,omitempty"`
	// PolicyType Specifies the type of group policy.
	PolicyType *GroupPolicyType `json:"policyType,omitempty"`
	// Revision The revision version associated with the file.
	Revision *string `json:"revision,omitempty"`
	// LastModifiedDateTime The date and time the entity was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Definitions undocumented
	Definitions []GroupPolicyDefinition `json:"definitions,omitempty"`
}
