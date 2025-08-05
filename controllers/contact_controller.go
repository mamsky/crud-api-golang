package controllers

import (
	"crud/database"
	"crud/models"
	"crud/utils"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateContact godoc
// @Summary Buat kontak baru
// @Description Buat data kontak baru dan simpan ke database
// @Tags Contacts
// @Accept json
// @Produce json
// @Param contact body models.CreateContactRequest true "Contact"
// @Success 201 {object} models.CreateContactRequest
// @Failure 400
// @Router /contacts [post]
func CreateContact(c *fiber.Ctx) error {
	contact := new(models.Contact)

	// Parse request body
	if err := c.BodyParser(contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Input tidak valid",
		})
	}

	if !utils.IsValidEmail(contact.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Format email tidak valid",
		})
	}

	if !utils.IsValidPhone(contact.Phone) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Format nomor telepon tidak valid",
		})
	}
	if contact.Gender != "" && contact.Gender != "male" && contact.Gender != "female" {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Gender hanya boleh 'male' atau 'female'",
	})
}

	// Validasi sederhana (manual)
	if contact.Name == ""|| contact.Gender == "" || contact.Email == "" || contact.Phone == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name, Email, dan Phone wajib diisi",
		})
	}

	// Cek apakah email sudah ada
	var existingEmail models.Contact
	if err := database.DB.Where("email = ?", contact.Email).First(&existingEmail).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email sudah digunakan",
		})
	}

	// Cek apakah phone sudah ada
	var existingPhone models.Contact
	if err := database.DB.Where("phone = ?", contact.Phone).First(&existingPhone).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Nomor telepon sudah digunakan",
		})
	}

	// Generate UUID dan timestamp
	contact.ID = uuid.New().String()
	contact.CreatedAt = time.Now()
	contact.UpdatedAt = time.Now()

	// Simpan ke DB
	if err := database.DB.Create(&contact).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal menyimpan ke database",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(contact)
}

// GetContacts godoc
// @Summary Ambil semua kontak
// @Description Mengambil daftar semua data kontak dengan opsi pencarian, filter gender, dan pagination
// @Tags Contacts
// @Accept json
// @Produce json
// @Param page query int false "Nomor halaman" minimum(1)
// @Param limit query int false "Jumlah item per halaman (tidak boleh 0)" minimum(1)
// @Param gender query string false "Filter berdasarkan gender ('male' atau 'female')"
// @Success 200 
// @Failure 400 "Bad Request jika parameter tidak valid"
// @Failure 500 "Internal Server Error jika query gagal"
// @Router /contacts [get]
func GetContacts(c *fiber.Ctx) error {
	var contacts []models.Contact
	query := database.DB

	
	// Optional: filter by gender
	// Validasi gender
	if gender := c.Query("gender"); gender != "" {
		if gender != "male" && gender != "female" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Gender harus 'male' atau 'female'",
			})
		}
		query = query.Where("gender = ?", gender)
	}

	// Validasi limit
	limit, offset := utils.Paginate(c)
	if limit <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Limit Tidak boleh 0",
		})
	}

	// Query data
	result := query.Limit(limit).Offset(offset).Find(&contacts)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengambil data",
		})
	}

	return c.JSON(fiber.Map{
		"data":   contacts,
		"limit":  limit,
		"offset": offset,
		"count":  result.RowsAffected,
	})
}

// GetContactByID godoc
// @Summary Ambil kontak berdasarkan ID
// @Description Ambil detail kontak berdasarkan ID
// @Tags Contacts
// @Produce json
// @Param id path string true "Contact ID"
// @Success 200 {object} models.Contact
// @Failure 404
// @Router /contacts/{id} [get]
func GetContactByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var contact models.Contact

	err := database.DB.First(&contact, "id = ?", id).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Kontak tidak ditemukan",
		})
	}

	return c.JSON(contact)
}

// UpdateContact godoc
// @Summary Update kontak
// @Description Update data kontak berdasarkan ID
// @Tags Contacts
// @Accept json
// @Produce json
// @Param id path string true "Contact ID"
// @Param contact body models.CreateContactRequest true "Contact"
// @Success 200 {object} models.CreateContactRequest
// @Failure 400 
// @Router /contacts/{id} [put]
func UpdateContact(c *fiber.Ctx) error {
	id := c.Params("id")
	var contact models.Contact

	// Cek apakah ada kontak dengan ID itu
	if err := database.DB.First(&contact, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "id tidak ditemukan",
		})
	}

	// Bind data baru dari body
	if err := c.BodyParser(&contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Input tidak valid",
		})
	}

	if contact.Name == ""|| contact.Gender == "" || contact.Email == "" || contact.Phone == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name, Gender, Email, dan Phone wajib diisi",
		})
	}

	if !utils.IsValidEmail(contact.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Format email tidak valid",
		})
	}

	if !utils.IsValidPhone(contact.Phone) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Format nomor telepon tidak valid",
		})
	}
	
	// Cek apakah email sudah digunakan oleh kontak lain
	var count int64
	database.DB.Model(&models.Contact{}).
		Where("email = ? AND id != ?", contact.Email, id).
		Count(&count)
	if count > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email sudah digunakan oleh kontak lain",
		})
	}
	

	// Cek apakah phone sudah digunakan oleh kontak lain
	database.DB.Model(&models.Contact{}).
		Where("phone = ? AND id != ?", contact.Phone, id).
		Count(&count)
	if count > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Phone sudah digunakan oleh kontak lain",
		})
	}


	contact.UpdatedAt = time.Now()

	// Simpan ke DB
	if err := database.DB.Save(&contact).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal update kontak",
		})
	}

	return c.JSON(contact)
}

// DeleteContact godoc
// @Summary Hapus kontak
// @Description Hapus data kontak berdasarkan ID
// @Tags Contacts
// @Param id path string true "Contact ID"
// @Success 200 
// @Failure 404
// @Router /contacts/{id} [delete]
func DeleteContact(c *fiber.Ctx) error {
	id := c.Params("id")

	// Cek apakah data dengan ID tersebut ada
	var contact models.Contact
	if err := database.DB.First(&contact, "id = ?", id).Error; err != nil {
		// Jika tidak ditemukan
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data Tidak Ditemukan",
			})
		}
		// Jika error lain (contoh: koneksi DB)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Terjadi kesalahan saat mencari data",
		})
	}

	// Jika ditemukan, hapus data
	if err := database.DB.Delete(&contact).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal hapus kontak",
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Berhasil menghapus data dengan ID %s", id),
	})
}

