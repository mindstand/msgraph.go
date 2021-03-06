// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MediaDirection undocumented
type MediaDirection string

const (
	// MediaDirectionVInactive undocumented
	MediaDirectionVInactive MediaDirection = "Inactive"
	// MediaDirectionVSendOnly undocumented
	MediaDirectionVSendOnly MediaDirection = "SendOnly"
	// MediaDirectionVReceiveOnly undocumented
	MediaDirectionVReceiveOnly MediaDirection = "ReceiveOnly"
	// MediaDirectionVSendReceive undocumented
	MediaDirectionVSendReceive MediaDirection = "SendReceive"
)

// MediaDirectionPInactive returns a pointer to MediaDirectionVInactive
func MediaDirectionPInactive() *MediaDirection {
	v := MediaDirectionVInactive
	return &v
}

// MediaDirectionPSendOnly returns a pointer to MediaDirectionVSendOnly
func MediaDirectionPSendOnly() *MediaDirection {
	v := MediaDirectionVSendOnly
	return &v
}

// MediaDirectionPReceiveOnly returns a pointer to MediaDirectionVReceiveOnly
func MediaDirectionPReceiveOnly() *MediaDirection {
	v := MediaDirectionVReceiveOnly
	return &v
}

// MediaDirectionPSendReceive returns a pointer to MediaDirectionVSendReceive
func MediaDirectionPSendReceive() *MediaDirection {
	v := MediaDirectionVSendReceive
	return &v
}
