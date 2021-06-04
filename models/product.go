package models

import (
	"majoo-test/forms"
	"majoo-test/request"
	"majoo-test/util"
	"strconv"
	"strings"
)

type ProductModel struct{}

func (M *ProductModel) Create(body request.Product) error {
	db := util.DBConnection.Session
	result := db.Table("tb_barang").Create(&body)
	return result.Error
}

func (M *ProductModel) ListByMercant(MercantID int, filter string, sort string, pageNo string, perPage string) (data []forms.Product, err error, count int, errcount error) {

	sorting := sort
	if strings.Contains(sort, "asc") {
		sorting = strings.Replace(sort, "|asc", "", -1)
		sorting = sorting + " asc"
	} else if strings.Contains(sort, "desc") {
		sorting = strings.Replace(sort, "|desc", "", -1)
		sorting = sorting + " desc"
	} else {
		sorting = ""
	}
	pageNoint, _ := strconv.Atoi(pageNo)
	perPageint, _ := strconv.Atoi(perPage)
	offset := (pageNoint - 1) * perPageint
	db := util.DBConnection.Session
	result := db.Table("tb_barang").Where("id_mercant=?", MercantID).Where("nama LIKE ?", "%"+filter+"%").Order(sorting).Find(&data)
	count = int(result.RowsAffected)
	db.Table("tb_barang").Limit(perPageint).Offset(offset).Where("nama LIKE ?", "%"+filter+"%").Order(sorting).Find(&data)
	return data, err, count, errcount
}

func (M *ProductModel) FindById(id string) (data forms.Product, err error) {
	db := util.DBConnection.Session
	result := db.Table("tb_barang").Where("id_barang = ?", id).First(&data)
	return data, result.Error
}

func (M *ProductModel) Put(id string, body request.Product) (err error) {
	db := util.DBConnection.Session
	result := db.Table("tb_barang").Model(&request.Product{}).Where("id_barang = ?", id).Updates(body)
	return result.Error
}

func (M *ProductModel) Delete(id string) (err error) {
	db := util.DBConnection.Session
	result := db.Table("tb_barang").Where("id_barang = ?", id).Delete(&forms.Product{})
	return result.Error
}
