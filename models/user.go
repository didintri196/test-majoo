package models

import (
	"crypto/md5"
	"encoding/hex"
	"majoo-test/forms"
	"majoo-test/request"
	"majoo-test/util"
	"strconv"
	"strings"
	"time"
)

type UserModel struct{}

func (M *UserModel) Create(body request.User) error {
	db := util.DBConnection.Session
	result := db.Table("tb_user").Create(&body)
	return result.Error
}

func (M *UserModel) FindById(id string) (data forms.User, err error) {
	db := util.DBConnection.Session
	result := db.Table("tb_user").Where("id_user = ?", id).First(&data)
	return data, result.Error
}

func (M *UserModel) ListWhereMercant(MercantID int, filter string, sort string, pageNo string, perPage string) (data []forms.User, err error, count int, errcount error) {

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
	result := db.Table("tb_user").Where("id_mercant=?", MercantID).Where("nama LIKE ?", "%"+filter+"%").Order(sorting).Find(&data)
	count = int(result.RowsAffected)
	db.Table("tb_user").Limit(perPageint).Offset(offset).Where("id_mercant=?", MercantID).Where("nama LIKE ?", "%"+filter+"%").Order(sorting).Find(&data)
	return data, err, count, errcount
}
func (M *UserModel) MD5(text string) string {
	algorithm := md5.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}
func (M *UserModel) Put(id string, body request.UserUpdate) (err error) {
	db := util.DBConnection.Session
	body.Password = M.MD5(body.Password)
	result := db.Table("tb_user").Model(&request.UserUpdate{}).Where("id_user = ?", id).Updates(body)
	return result.Error
}

func (M *UserModel) Delete(id string) (err error) {
	db := util.DBConnection.Session
	result := db.Table("tb_user").Where("id_user = ?", id).Delete(&forms.Product{})
	return result.Error
}

func (M *UserModel) FindByEmailPass(datalogin request.UserLogin) (data forms.User, err error) {
	db := util.DBConnection.Session
	result := db.Table("tb_user").Where("email = ? AND password = MD5(?)", datalogin.Email, datalogin.Password).First(&data)
	return data, result.Error
}

func (M *UserModel) UpdateLastLogin(id int) (err error) {
	db := util.DBConnection.Session
	body, _ := M.FindById(strconv.Itoa(id))
	body.LastLogin = int(time.Now().Unix())
	result := db.Table("tb_user").Model(&forms.User{}).Where("id_user = ?", id).Updates(body)
	return result.Error
}
