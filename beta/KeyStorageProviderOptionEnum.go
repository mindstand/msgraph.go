// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// KeyStorageProviderOption undocumented
type KeyStorageProviderOption string

const (
	// KeyStorageProviderOptionVUseTpmKspOtherwiseUseSoftwareKsp undocumented
	KeyStorageProviderOptionVUseTpmKspOtherwiseUseSoftwareKsp KeyStorageProviderOption = "UseTpmKspOtherwiseUseSoftwareKsp"
	// KeyStorageProviderOptionVUseTpmKspOtherwiseFail undocumented
	KeyStorageProviderOptionVUseTpmKspOtherwiseFail KeyStorageProviderOption = "UseTpmKspOtherwiseFail"
	// KeyStorageProviderOptionVUsePassportForWorkKspOtherwiseFail undocumented
	KeyStorageProviderOptionVUsePassportForWorkKspOtherwiseFail KeyStorageProviderOption = "UsePassportForWorkKspOtherwiseFail"
	// KeyStorageProviderOptionVUseSoftwareKsp undocumented
	KeyStorageProviderOptionVUseSoftwareKsp KeyStorageProviderOption = "UseSoftwareKsp"
)

// KeyStorageProviderOptionPUseTpmKspOtherwiseUseSoftwareKsp returns a pointer to KeyStorageProviderOptionVUseTpmKspOtherwiseUseSoftwareKsp
func KeyStorageProviderOptionPUseTpmKspOtherwiseUseSoftwareKsp() *KeyStorageProviderOption {
	v := KeyStorageProviderOptionVUseTpmKspOtherwiseUseSoftwareKsp
	return &v
}

// KeyStorageProviderOptionPUseTpmKspOtherwiseFail returns a pointer to KeyStorageProviderOptionVUseTpmKspOtherwiseFail
func KeyStorageProviderOptionPUseTpmKspOtherwiseFail() *KeyStorageProviderOption {
	v := KeyStorageProviderOptionVUseTpmKspOtherwiseFail
	return &v
}

// KeyStorageProviderOptionPUsePassportForWorkKspOtherwiseFail returns a pointer to KeyStorageProviderOptionVUsePassportForWorkKspOtherwiseFail
func KeyStorageProviderOptionPUsePassportForWorkKspOtherwiseFail() *KeyStorageProviderOption {
	v := KeyStorageProviderOptionVUsePassportForWorkKspOtherwiseFail
	return &v
}

// KeyStorageProviderOptionPUseSoftwareKsp returns a pointer to KeyStorageProviderOptionVUseSoftwareKsp
func KeyStorageProviderOptionPUseSoftwareKsp() *KeyStorageProviderOption {
	v := KeyStorageProviderOptionVUseSoftwareKsp
	return &v
}
