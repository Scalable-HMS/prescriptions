package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wryonik/appointment/models"
)

type Prescriptions struct {
	DoctorId          uint
	PatientId         uint
	HospitalId        uint
	PrescriptionFiles string
	Date              time.Time
	CombinedPdfUrl    string
}

type CreatePrescriptionInput struct {
	DoctorId          uint      `json:"doctor_id" binding:"required"`
	PatientId         uint      `json:"patient_id" binding:"required"`
	HospitalId        uint      `json:"hospital_id" binding:"required"`
	PrescriptionFiles string    `json:"prescription_files" binding:"required"`
	Date              time.Time `json:"date_time" binding:"required"`
}

type UpdatePrescriptionInput struct {
	DoctorId          uint      `json:"doctor_id" binding:"required"`
	PatientId         uint      `json:"patient_id" binding:"required"`
	HospitalId        uint      `json:"hospital_id" binding:"required"`
	PrescriptionFiles string    `json:"prescription_files" binding:"required"`
	Date              time.Time `json:"date_time" binding:"required"`
}

// GET /prescriptions
// Find all prescriptions
func FindPrescriptions(c *gin.Context) {
	var prescriptions []Prescriptions
	models.DB.Find(&prescriptions)

	c.JSON(http.StatusOK, prescriptions)
}

// GET /prescriptions/:id
// Find a prescription
func FindPrescription(c *gin.Context) {
	// Get model if exist
	var prescription Prescriptions
	if err := models.DB.Where("id = ?", c.Param("id")).First(&prescription).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, prescription)
}

// POST /prescriptions
// Create new prescription
func CreatePrescription(c *gin.Context) {
	// Validate input
	var input CreatePrescriptionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create prescription
	prescription := Prescriptions{DoctorId: input.DoctorId, PatientId: input.PatientId, HospitalId: input.HospitalId, Date: input.Date, PrescriptionFiles: input.PrescriptionFiles}
	models.DB.Create(&prescription)

	c.JSON(http.StatusOK, gin.H{"data": prescription})
}

// PATCH /prescriptions/:id
// Update a prescription
func UpdatePrescription(c *gin.Context) {
	// Get model if exist
	var prescription Prescriptions
	if err := models.DB.Where("id = ?", c.Param("id")).First(&prescription).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdatePrescriptionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&prescription).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": prescription})
}

// DELETE /prescriptions/:id
// Delete a prescription
func DeletePrescription(c *gin.Context) {
	// Get model if exist
	var prescription Prescriptions
	if err := models.DB.Where("id = ?", c.Param("id")).First(&prescription).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&prescription)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
