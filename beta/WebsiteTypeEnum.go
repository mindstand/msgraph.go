// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WebsiteType undocumented
type WebsiteType string

const (
	// WebsiteTypeVOther undocumented
	WebsiteTypeVOther WebsiteType = "Other"
	// WebsiteTypeVHome undocumented
	WebsiteTypeVHome WebsiteType = "Home"
	// WebsiteTypeVWork undocumented
	WebsiteTypeVWork WebsiteType = "Work"
	// WebsiteTypeVBlog undocumented
	WebsiteTypeVBlog WebsiteType = "Blog"
	// WebsiteTypeVProfile undocumented
	WebsiteTypeVProfile WebsiteType = "Profile"
)

// WebsiteTypePOther returns a pointer to WebsiteTypeVOther
func WebsiteTypePOther() *WebsiteType {
	v := WebsiteTypeVOther
	return &v
}

// WebsiteTypePHome returns a pointer to WebsiteTypeVHome
func WebsiteTypePHome() *WebsiteType {
	v := WebsiteTypeVHome
	return &v
}

// WebsiteTypePWork returns a pointer to WebsiteTypeVWork
func WebsiteTypePWork() *WebsiteType {
	v := WebsiteTypeVWork
	return &v
}

// WebsiteTypePBlog returns a pointer to WebsiteTypeVBlog
func WebsiteTypePBlog() *WebsiteType {
	v := WebsiteTypeVBlog
	return &v
}

// WebsiteTypePProfile returns a pointer to WebsiteTypeVProfile
func WebsiteTypePProfile() *WebsiteType {
	v := WebsiteTypeVProfile
	return &v
}
