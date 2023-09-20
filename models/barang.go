package models

import "time"

type Barang struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	NamaBarang string    `json:"namaBarang" gorm:"type:varchar(50)"`
	Keterangan string    `json:"keterangan" `
	StokBarang uint      `json:"stokBarang"`
	Category   string    `json:"kategori"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt" gorm:"type:datetime;AutoCreateTime"`
}
