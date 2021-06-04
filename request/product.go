package request

type Product struct {
	Sku        string `json:"sku" gorm:"column:sku"`
	Nama       string `json:"nama" gorm:"column:nama"`
	HargaBeli  int    `json:"harga_beli" gorm:"column:harga_beli"`
	HargaJual  int    `json:"harga_jual" gorm:"column:harga_jual"`
	Stock      int    `json:"stock" gorm:"column:stock"`
	SupplierID int    `json:"id_supplier" gorm:"column:id_supplier"`
	MercantID  int    `json:"id_mercant" gorm:"column:id_mercant"`
	Foto       string `json:"foto" gorm:"column:foto"`
	IsActive   string `json:"is_aktif" gorm:"column:is_aktif"`
}
