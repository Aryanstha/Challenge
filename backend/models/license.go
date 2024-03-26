// models/license.go
package models

type License struct {
	ID           int    `json:"id"`
	LicenseKey   string `json:"licenseKey"`
	CustomerName string `json:"customerName"`
	ExpiryDate   string `json:"expiryDate"`
	MaxUsers     int    `json:"maxUsers"`
}
