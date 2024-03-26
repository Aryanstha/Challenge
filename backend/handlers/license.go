package handlers

import (
	"encoding/json"
	"github.com/aryanstha/saas-license/backend/models"
	"net/http"
)

// GetLicenseHandler returns the license details
func GetLicenseHandler(w http.ResponseWriter, r *http.Request) {
	license := models.GetLicense()
	json.NewEncoder(w).Encode(license)
}

// UpdateLicenseHandler updates the license details
func UpdateLicenseHandler(w http.ResponseWriter, r *http.Request) {
	var license models.License
	json.NewDecoder(r.Body).Decode(&license)
	models.UpdateLicense(license)
	json.NewEncoder(w).Encode(license)
}

// DeleteLicenseHandler deletes the license details
func DeleteLicenseHandler(w http.ResponseWriter, r *http.Request) {
	models.DeleteLicense()
	w.WriteHeader(http.StatusNoContent)
}

// CreateLicenseHandler creates the license details
func CreateLicenseHandler(w http.ResponseWriter, r *http.Request) {
	var license models.License
	json.NewDecoder(r.Body).Decode(&license)
	models.CreateLicense(license)
	json.NewEncoder(w).Encode(license)
}

// ValidateLicenseHandler validates the license details
func ValidateLicenseHandler(w http.ResponseWriter, r *http.Request) {
	var license models.License
	json.NewDecoder(r.Body).Decode(&license)
	valid := models.ValidateLicense(license)
	json.NewEncoder(w).Encode(valid)
}
