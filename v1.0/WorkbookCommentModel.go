// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WorkbookComment undocumented
type WorkbookComment struct {
	Entity
	// Content undocumented
	Content *string `json:"content,omitempty"`
	// ContentType undocumented
	ContentType *string `json:"contentType,omitempty"`
	// Replies undocumented
	Replies []WorkbookCommentReply `json:"replies,omitempty"`
}
