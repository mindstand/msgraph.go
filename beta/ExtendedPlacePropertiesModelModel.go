// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ExtendedPlacePropertiesModel undocumented
type ExtendedPlacePropertiesModel struct {
	// Object is the base model of ExtendedPlacePropertiesModel
	Object
	// PriceRange undocumented
	PriceRange *string `json:"priceRange,omitempty"`
	// MenuURL undocumented
	MenuURL *string `json:"menuUrl,omitempty"`
	// BusinessURL undocumented
	BusinessURL *string `json:"businessUrl,omitempty"`
	// FormattedAddress undocumented
	FormattedAddress *string `json:"formattedAddress,omitempty"`
	// OpeningHoursSpecifications undocumented
	OpeningHoursSpecifications []OpeningHoursSpecification `json:"openingHoursSpecifications,omitempty"`
	// TimeZone undocumented
	TimeZone *string `json:"timeZone,omitempty"`
}
