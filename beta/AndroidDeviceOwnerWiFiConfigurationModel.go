// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidDeviceOwnerWiFiConfiguration By providing the configurations in this profile you can instruct the Android device to connect to desired Wi-Fi endpoint. By specifying the authentication method and security types expected by Wi-Fi endpoint you can make the Wi-Fi connection seamless for end user. This profile provides limited and simpler security types than Enterprise Wi-Fi profile.
type AndroidDeviceOwnerWiFiConfiguration struct {
	DeviceConfiguration
	// NetworkName Network Name
	NetworkName *string `json:"networkName,omitempty"`
	// Ssid This is the name of the Wi-Fi network that is broadcast to all devices.
	Ssid *string `json:"ssid,omitempty"`
	// ConnectAutomatically Connect automatically when this network is in range. Setting this to true will skip the user prompt and automatically connect the device to Wi-Fi network.
	ConnectAutomatically *bool `json:"connectAutomatically,omitempty"`
	// ConnectWhenNetworkNameIsHidden When set to true, this profile forces the device to connect to a network that doesn't broadcast its SSID to all devices.
	ConnectWhenNetworkNameIsHidden *bool `json:"connectWhenNetworkNameIsHidden,omitempty"`
	// WiFiSecurityType Indicates whether Wi-Fi endpoint uses an EAP based security type.
	WiFiSecurityType *AndroidDeviceOwnerWiFiSecurityType `json:"wiFiSecurityType,omitempty"`
	// PreSharedKey This is the pre-shared key for WPA Personal Wi-Fi network.
	PreSharedKey *string `json:"preSharedKey,omitempty"`
	// PreSharedKeyIsSet This is the pre-shared key for WPA Personal Wi-Fi network.
	PreSharedKeyIsSet *bool `json:"preSharedKeyIsSet,omitempty"`
}
