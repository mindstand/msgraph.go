// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// GiphyRatingType undocumented
type GiphyRatingType string

const (
	// GiphyRatingTypeVModerate undocumented
	GiphyRatingTypeVModerate GiphyRatingType = "Moderate"
	// GiphyRatingTypeVStrict undocumented
	GiphyRatingTypeVStrict GiphyRatingType = "Strict"
	// GiphyRatingTypeVUnknownFutureValue undocumented
	GiphyRatingTypeVUnknownFutureValue GiphyRatingType = "UnknownFutureValue"
)

// GiphyRatingTypePModerate returns a pointer to GiphyRatingTypeVModerate
func GiphyRatingTypePModerate() *GiphyRatingType {
	v := GiphyRatingTypeVModerate
	return &v
}

// GiphyRatingTypePStrict returns a pointer to GiphyRatingTypeVStrict
func GiphyRatingTypePStrict() *GiphyRatingType {
	v := GiphyRatingTypeVStrict
	return &v
}

// GiphyRatingTypePUnknownFutureValue returns a pointer to GiphyRatingTypeVUnknownFutureValue
func GiphyRatingTypePUnknownFutureValue() *GiphyRatingType {
	v := GiphyRatingTypeVUnknownFutureValue
	return &v
}
