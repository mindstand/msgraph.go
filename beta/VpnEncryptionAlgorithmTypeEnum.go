// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// VpnEncryptionAlgorithmType undocumented
type VpnEncryptionAlgorithmType string

const (
	// VpnEncryptionAlgorithmTypeVAes256 undocumented
	VpnEncryptionAlgorithmTypeVAes256 VpnEncryptionAlgorithmType = "Aes256"
	// VpnEncryptionAlgorithmTypeVDes undocumented
	VpnEncryptionAlgorithmTypeVDes VpnEncryptionAlgorithmType = "Des"
	// VpnEncryptionAlgorithmTypeVTripleDes undocumented
	VpnEncryptionAlgorithmTypeVTripleDes VpnEncryptionAlgorithmType = "TripleDes"
	// VpnEncryptionAlgorithmTypeVAes128 undocumented
	VpnEncryptionAlgorithmTypeVAes128 VpnEncryptionAlgorithmType = "Aes128"
	// VpnEncryptionAlgorithmTypeVAes128Gcm undocumented
	VpnEncryptionAlgorithmTypeVAes128Gcm VpnEncryptionAlgorithmType = "Aes128Gcm"
	// VpnEncryptionAlgorithmTypeVAes256Gcm undocumented
	VpnEncryptionAlgorithmTypeVAes256Gcm VpnEncryptionAlgorithmType = "Aes256Gcm"
)

// VpnEncryptionAlgorithmTypePAes256 returns a pointer to VpnEncryptionAlgorithmTypeVAes256
func VpnEncryptionAlgorithmTypePAes256() *VpnEncryptionAlgorithmType {
	v := VpnEncryptionAlgorithmTypeVAes256
	return &v
}

// VpnEncryptionAlgorithmTypePDes returns a pointer to VpnEncryptionAlgorithmTypeVDes
func VpnEncryptionAlgorithmTypePDes() *VpnEncryptionAlgorithmType {
	v := VpnEncryptionAlgorithmTypeVDes
	return &v
}

// VpnEncryptionAlgorithmTypePTripleDes returns a pointer to VpnEncryptionAlgorithmTypeVTripleDes
func VpnEncryptionAlgorithmTypePTripleDes() *VpnEncryptionAlgorithmType {
	v := VpnEncryptionAlgorithmTypeVTripleDes
	return &v
}

// VpnEncryptionAlgorithmTypePAes128 returns a pointer to VpnEncryptionAlgorithmTypeVAes128
func VpnEncryptionAlgorithmTypePAes128() *VpnEncryptionAlgorithmType {
	v := VpnEncryptionAlgorithmTypeVAes128
	return &v
}

// VpnEncryptionAlgorithmTypePAes128Gcm returns a pointer to VpnEncryptionAlgorithmTypeVAes128Gcm
func VpnEncryptionAlgorithmTypePAes128Gcm() *VpnEncryptionAlgorithmType {
	v := VpnEncryptionAlgorithmTypeVAes128Gcm
	return &v
}

// VpnEncryptionAlgorithmTypePAes256Gcm returns a pointer to VpnEncryptionAlgorithmTypeVAes256Gcm
func VpnEncryptionAlgorithmTypePAes256Gcm() *VpnEncryptionAlgorithmType {
	v := VpnEncryptionAlgorithmTypeVAes256Gcm
	return &v
}
