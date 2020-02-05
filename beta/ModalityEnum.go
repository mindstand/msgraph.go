// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Modality undocumented
type Modality string

const (
	// ModalityVUnknown undocumented
	ModalityVUnknown Modality = "Unknown"
	// ModalityVAudio undocumented
	ModalityVAudio Modality = "Audio"
	// ModalityVVideo undocumented
	ModalityVVideo Modality = "Video"
	// ModalityVVideoBasedScreenSharing undocumented
	ModalityVVideoBasedScreenSharing Modality = "VideoBasedScreenSharing"
	// ModalityVData undocumented
	ModalityVData Modality = "Data"
	// ModalityVUnknownFutureValue undocumented
	ModalityVUnknownFutureValue Modality = "UnknownFutureValue"
)

// ModalityPUnknown returns a pointer to ModalityVUnknown
func ModalityPUnknown() *Modality {
	v := ModalityVUnknown
	return &v
}

// ModalityPAudio returns a pointer to ModalityVAudio
func ModalityPAudio() *Modality {
	v := ModalityVAudio
	return &v
}

// ModalityPVideo returns a pointer to ModalityVVideo
func ModalityPVideo() *Modality {
	v := ModalityVVideo
	return &v
}

// ModalityPVideoBasedScreenSharing returns a pointer to ModalityVVideoBasedScreenSharing
func ModalityPVideoBasedScreenSharing() *Modality {
	v := ModalityVVideoBasedScreenSharing
	return &v
}

// ModalityPData returns a pointer to ModalityVData
func ModalityPData() *Modality {
	v := ModalityVData
	return &v
}

// ModalityPUnknownFutureValue returns a pointer to ModalityVUnknownFutureValue
func ModalityPUnknownFutureValue() *Modality {
	v := ModalityVUnknownFutureValue
	return &v
}
