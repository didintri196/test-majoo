package controllers

import (
	"bytes"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"majoo-test/config"
	"majoo-test/forms"
	"majoo-test/models"
	"majoo-test/request"
	"majoo-test/responses"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductController struct{}

var productmodel = new(models.ProductModel)

func (A *ProductController) Base64Uploader(file, folder string) (path string, err error) {
	dir := config.GetDir()
	fileNameBase := uuid.New().String()
	idx := strings.Index(file, ";base64,")
	if idx < 0 {
		return "", errors.New("Invalid Base64 Image")
	} else {
		reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(file[idx+8:]))
		buff := bytes.Buffer{}
		_, err := buff.ReadFrom(reader)
		if err != nil {
			return "", err
		} else {
			fileName := fileNameBase + ".png"
			errcreate := ioutil.WriteFile(dir+"/../assets/images/"+folder+"/"+fileName, buff.Bytes(), 0644)
			if errcreate != nil {
				return "", errcreate
			} else {
				return "/public/images/" + folder + "/" + fileName, errcreate
			}
		}
	}
}

// @Summary Create Product
// @Description Create Product
// @Accept  json
// @Produce  json
// @Param account body request.Product true "Create Product"
// @Success 201 {object} responses.ProductDefault
// @Failure 400 {object} responses.ProductDefault "Invalid Form"
// @Failure 401 {object} responses.AuthError "Unauthorized"
// @Failure 500 {object} responses.ProductDefault "Something Wrong on Server"
// @Router /product [post]
// @Security ApiKeyAuth
// @Tags Product
func (A *ProductController) CreateHandler(c *gin.Context) {
	var body request.Product

	// App level validation
	bindErr := c.BindJSON(&body)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, &responses.ProductDefault{
			Status:  http.StatusBadRequest,
			Message: bindErr.Error(),
		})
		return
	}
	// Inserting data
	if body.Foto != "" {
		path, err_upload := A.Base64Uploader(body.Foto, "product")
		if err_upload == nil {
			body.Foto = path
			insertErr := productmodel.Create(body)
			if insertErr != nil {
				c.JSON(http.StatusInternalServerError, &responses.ProductDefault{
					Status:  http.StatusInternalServerError,
					Message: insertErr.Error(),
				})
				return
			} else {
				c.JSON(http.StatusCreated, &responses.ProductDefault{
					Status:  http.StatusCreated,
					Message: "Success create data",
				})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, &responses.ProductDefault{
				Status:  http.StatusBadRequest,
				Message: err_upload.Error(),
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, &responses.ProductDefault{
			Status:  http.StatusBadRequest,
			Message: bindErr.Error(),
		})
		return
	}

}

// @Summary Get list data Product In Mercant
// @Description Show all data Product Mercant with pagination
// @Accept  json
// @Produce  json
// @Param   filter     query    string     false        "Some Fiter"
// @Param   sort     query    string     true        "name|asc"
// @Param   page     query    int     true        "1"
// @Param   per_page     query    int     true        "5"
// @Success 200 {object} responses.ListProduct
// @Failure 401 {object} responses.AuthError "Unauthorized"
// @Failure 500 {object} responses.ProductDefault "Something Wrong on Server"
// @Router /product [get]
// @Security ApiKeyAuth
// @Tags Product
func (A *ProductController) ListHandler(c *gin.Context) {
	filter := c.Query("filter")
	sort := c.Query("sort")
	pageNo := c.Query("page")
	perPage := c.Query("per_page")
	iduser, _ := GetSessionID(c)
	data_user, _ := usermodel.FindById(iduser)
	if sort == "" {
		sort = "id"
	}
	if pageNo == "" {
		pageNo = "1"
	}
	if perPage == "" {
		perPage = "5"
	}

	pp, _ := strconv.Atoi(perPage)
	pn, _ := strconv.Atoi(pageNo)
	data, err, count, errcount := productmodel.ListByMercant(data_user.MercantID, filter, sort, pageNo, perPage)
	lastPage := float64(count) / float64(pp)
	if pp != 0 {
		if count%pp == 0 {
			lastPage = lastPage
		} else {
			lastPage = lastPage + 1
		}
	} else {
		lastPage = float64(count) / float64(5)
	}
	if err != nil && errcount != nil {
		c.JSON(http.StatusInternalServerError, &responses.ProductDefault{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	} else {
		if len(data) == 0 {
			c.JSON(http.StatusOK, &responses.ListProduct{
				Total:       count,
				PerPage:     pp,
				CurrentPage: pn,
				LastPage:    int(lastPage),
				NextPage:    0,
				PrevPage:    0,
				Data:        []forms.Product{},
				Status:      http.StatusOK,
			})
			return
		} else {
			nextpage := pn
			if pn < int(lastPage) {
				nextpage = pn + 1
			}
			prevpage := 1
			if pn > 1 {
				prevpage = pn - 1
			}
			c.JSON(200, &responses.ListProduct{
				Total:       count,
				PerPage:     pp,
				CurrentPage: pn,
				LastPage:    int(lastPage),
				NextPage:    nextpage,
				PrevPage:    prevpage,
				Data:        data,
				Status:      http.StatusOK,
			})
			return
		}

	}
}

// @Summary Get By Id Product
// @Description Show by id Product
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "Some ID"
// @Success 200 {object} responses.ProductData
// @Failure 401 {object} responses.AuthError "Unauthorized"
// @Failure 404 {object} responses.ProductDefault "Not Found"
// @Router /product/view/{id} [get]
// @Security ApiKeyAuth
// @Tags Product
func (A *ProductController) ShowHandler(c *gin.Context) {
	id := c.Param("id")
	data, _ := productmodel.FindById(id)
	// Check if resource exist
	if data.ProductID == 0 {
		c.JSON(http.StatusNotFound, &responses.ProductDefault{
			Status:  http.StatusNotFound,
			Message: "Not found",
		})
		return
	} else {
		c.JSON(http.StatusOK, &responses.ProductData{
			Status:  http.StatusOK,
			Message: "Berhasil menampilkan data",
			Data:    data,
		})
		return
	}
}

// @Summary Update Product by ID
// @Description Update Product by ID
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "Some ID"
// @Param data body request.Product true "Update Product"
// @Success 201 {object} responses.ProductDefault
// @Failure 406 {object} responses.ProductDefault "Invalid Form"
// @Failure 401 {object} responses.AuthError "Error Not Logged In"
// @Failure 404 {object} responses.ProductDefault "ID not found"
// @Router /product/{id} [put]
// @Security ApiKeyAuth
// @Tags Product
func (A *ProductController) PutHandler(c *gin.Context) {
	id := c.Param("id")

	var data request.Product

	// App level validation
	bindErr := c.BindJSON(&data)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, &responses.ProductDefault{
			Status:  http.StatusBadRequest,
			Message: bindErr.Error(),
		})
		return
	}

	// Check if resource exist
	cekdata, _ := productmodel.FindById(id)
	if cekdata.ProductID == 0 {
		c.JSON(http.StatusNotFound, &responses.ProductDefault{
			Status:  http.StatusNotFound,
			Message: "Not found",
		})
		return
	}

	if data.Foto != "" {
		path, err_upload := A.Base64Uploader(data.Foto, "product")
		if err_upload != nil {
			data.Foto = cekdata.Foto
		} else {
			data.Foto = path
		}
	} else {
		data.Foto = cekdata.Foto
	}

	// Updating data
	err := productmodel.Put(id, data)

	if err != nil {
		c.JSON(http.StatusBadRequest, &responses.ProductDefault{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusCreated, &responses.ProductDefault{
			Status:  http.StatusCreated,
			Message: "Berhasil mengubah data",
		})
		return
	}
}

// @Summary Delete Product by ID
// @Description Delete Product by ID
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "Some ID"
// @Success 200 {object} responses.ProductDefault
// @Failure 406 {object} responses.ProductDefault "Invalid Form"
// @Failure 401 {object} responses.AuthError "Error Not Logged In"
// @Failure 404 {object} responses.ProductDefault "ID not found"
// @Router /product/{id} [delete]
// @Security ApiKeyAuth
// @Tags Product
func (A *ProductController) DeleteHandler(c *gin.Context) {
	id := c.Param("id")

	// Check if resource exist
	cekdata, _ := productmodel.FindById(id)
	if cekdata.ProductID == 0 {
		c.JSON(http.StatusNotFound, &responses.ProductDefault{
			Status:  http.StatusNotFound,
			Message: "Not found",
		})
		return
	}

	// Deleting data
	err := productmodel.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &responses.ProductDefault{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, &responses.ProductDefault{
			Status:  http.StatusOK,
			Message: "Berhasil menghapus data",
		})
		return
	}
}
