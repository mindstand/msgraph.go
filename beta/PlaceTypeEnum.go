// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PlaceType undocumented
type PlaceType string

const (
	// PlaceTypeVUnknown undocumented
	PlaceTypeVUnknown PlaceType = "Unknown"
	// PlaceTypeVRoom undocumented
	PlaceTypeVRoom PlaceType = "Room"
	// PlaceTypeVRoomList undocumented
	PlaceTypeVRoomList PlaceType = "RoomList"
	// PlaceTypeVPublicPlace undocumented
	PlaceTypeVPublicPlace PlaceType = "PublicPlace"
	// PlaceTypeVPersonalPlace undocumented
	PlaceTypeVPersonalPlace PlaceType = "PersonalPlace"
)

// PlaceTypePUnknown returns a pointer to PlaceTypeVUnknown
func PlaceTypePUnknown() *PlaceType {
	v := PlaceTypeVUnknown
	return &v
}

// PlaceTypePRoom returns a pointer to PlaceTypeVRoom
func PlaceTypePRoom() *PlaceType {
	v := PlaceTypeVRoom
	return &v
}

// PlaceTypePRoomList returns a pointer to PlaceTypeVRoomList
func PlaceTypePRoomList() *PlaceType {
	v := PlaceTypeVRoomList
	return &v
}

// PlaceTypePPublicPlace returns a pointer to PlaceTypeVPublicPlace
func PlaceTypePPublicPlace() *PlaceType {
	v := PlaceTypeVPublicPlace
	return &v
}

// PlaceTypePPersonalPlace returns a pointer to PlaceTypeVPersonalPlace
func PlaceTypePPersonalPlace() *PlaceType {
	v := PlaceTypeVPersonalPlace
	return &v
}
