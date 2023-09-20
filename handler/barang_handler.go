package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/muhammadsaefulr/inventori-barang/models"
	"github.com/muhammadsaefulr/inventori-barang/usecase"
)

type BarangHandler struct {
	BarangUseCase *usecase.BarangUseCase
}

func NewBarangHandler(r *gin.RouterGroup, barang_u *usecase.BarangUseCase) {
	handler := &BarangHandler{
		BarangUseCase: barang_u,
	}

	r.POST("/barang", handler.CreateBarang)
	r.GET("/barang/:id", handler.GetBarangId)
	r.GET("barang/kategori/:kategori", handler.GetCategory)
	r.GET("/barang", handler.GetAllListBarang)
	r.GET("/barang/date", handler.GetAllDataFromYear)
	r.GET("/barang/status/:status", handler.GetStatus)
	r.PUT("/barang/:id", handler.UpdateBarang)
	r.DELETE("/barang/:id", handler.DeleteBarangFromId)
}

func (barang_h *BarangHandler) GetAllListBarang(c *gin.Context) {
	barang, err := barang_h.BarangUseCase.GetAllListBarang()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, barang)
}

func (barang_h *BarangHandler) GetBarangId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	barang, err := barang_h.BarangUseCase.GetBarangId(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, barang)
}

func (barang_h *BarangHandler) GetStatus(c *gin.Context) {
	status := c.Param("status")

	barang, err := barang_h.BarangUseCase.GetStatus(status)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	c.JSON(http.StatusOK, barang)
}

func (barang_h *BarangHandler) GetCategory(c *gin.Context) {
	kategori := c.Param("kategori")

	barang, err := barang_h.BarangUseCase.GetCategory(kategori)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Kategori Barang Tidak Ditemukan"})
		return
	}
	c.JSON(http.StatusOK, barang)
}

func (barang_h *BarangHandler) CreateBarang(c *gin.Context) {
	var barang models.Barang

	if err := c.ShouldBindJSON(&barang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := barang_h.BarangUseCase.CreateBarang(&barang); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Barang Berhasil Dibuat"})
}

func (barang_h *BarangHandler) UpdateBarang(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	barang, err := barang_h.BarangUseCase.GetBarangId(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Id Object Not Found"})
	}

	barang.CreatedAt = time.Now()

	if err := c.ShouldBindJSON(barang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := barang_h.BarangUseCase.UpdateBarang(barang); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Barang Berhasil Di Update"})

}

func (barang_h *BarangHandler) DeleteBarangFromId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Object Not Found"})
	}

	barang, err := barang_h.BarangUseCase.GetBarangId(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Barang Tidak Ditemukan"})
		return
	}

	if err := barang_h.BarangUseCase.DeleteBarangFromId(barang); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil Di hapus"})
}

func (barang_h *BarangHandler) GetAllDataFromYear(c *gin.Context) {
	StartDateStr := c.Query("start_date")
	EndDateStr := c.Query("end_date")

	barang, err := barang_h.BarangUseCase.GetAllDataFromYear(StartDateStr, EndDateStr)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": barang})
}
