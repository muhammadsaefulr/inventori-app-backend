package usecase

import (
	"errors"

	"github.com/muhammadsaefulr/inventori-barang/models"
	"gorm.io/gorm"
)

type BarangUseCase struct {
	DB *gorm.DB
}

func NewBarangUseCase(db *gorm.DB) *BarangUseCase {
	return &BarangUseCase{DB: db}
}

func (uc *BarangUseCase) GetAllListBarang() ([]models.Barang, error) {
	var barang []models.Barang

	if err := uc.DB.Find(&barang).Error; err != nil {
		return nil, err
	}
	return barang, nil
}

func (uc *BarangUseCase) GetCategory(kategori string) ([]models.Barang, error) {
	var barang []models.Barang

	if err := uc.DB.Where("category = ?", kategori).Find(&barang).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Catgeory Not Found")
		}
		return nil, err

	}
	return barang, nil
}

func (uc *BarangUseCase) CreateBarang(barang *models.Barang) error {
	if err := uc.DB.Create(barang).Error; err != nil {
		return nil
	}
	return nil
}

func (uc *BarangUseCase) GetBarangId(id uint) (*models.Barang, error) {
	var barang models.Barang

	if err := uc.DB.First(&barang, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Barang tidak ditemukan")
		}
		return nil, err
	}
	return &barang, nil
}

func (uc *BarangUseCase) GetStatus(Status string) ([]models.Barang, error) {
	var barang []models.Barang

	if err := uc.DB.Where("status = ?", Status).Find(&barang).Error; err != nil {
		return nil, err
	}
	return barang, nil
}

func (uc *BarangUseCase) UpdateBarang(barang *models.Barang) error {
	if err := uc.DB.Save(barang).Error; err != nil {
		return nil
	}
	return nil
}

func (uc *BarangUseCase) DeleteBarangFromId(barang *models.Barang) error {
	if err := uc.DB.Delete(barang).Error; err != nil {
		return nil
	}
	return nil
}

func (uc *BarangUseCase) GetAllDataFromYear(StartDateStr string, EndDataStr string) ([]models.Barang, error) {
	var barang []models.Barang

	if err := uc.DB.Where("created_at BETWEEN ? AND ?", StartDateStr, EndDataStr).Find(&barang).Error; err != nil {
		return nil, err
	}

	return barang, nil
}
