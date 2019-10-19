// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeliveryOptimizationMaxCacheSizePercentage undocumented
type DeliveryOptimizationMaxCacheSizePercentage struct {
	DeliveryOptimizationMaxCacheSize
	// MaximumCacheSizePercentage Specifies the maximum cache size that Delivery Optimization can utilize, as a percentage of disk size (1-100). Valid values 1 to 100
	MaximumCacheSizePercentage *int `json:"maximumCacheSizePercentage,omitempty"`
}
