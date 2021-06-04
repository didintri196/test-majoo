package controllers

import (
	"errors"
	"majoo-test/forms"
	"majoo-test/models"
	"majoo-test/request"
	"majoo-test/responses"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

var usermodel = new(models.UserModel)

func GetSessionID(c *gin.Context) (idaccount string, err error) {
	tokenString := c.Request.Header.Get("Authorization")
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		err = errors.New("Authentication failed " + err.Error())
		return "", err
	} else {
		idaccount = claims["id"].(string)
		return idaccount, nil
	}
}

// @Summary Create User
// @Description Create User
// @Accept  json
// @Produce  json
// @Param account body request.User true "Create User"
// @Success 201 {object} responses.UserDefault
// @Failure 400 {object} responses.UserDefault "Invalid Form"
// @Failure 401 {object} responses.AuthError "Unauthorized"
// @Failure 500 {object} responses.UserDefault "Something Wrong on Server"
// @Router /user [post]
// @Security ApiKeyAuth
// @Tags User
func (A *UserController) CreateHandler(c *gin.Context) {
	var body request.User
	// App level validation
	bindErr := c.BindJSON(&body)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, &responses.UserDefault{
			Status:  http.StatusBadRequest,
			Message: bindErr.Error(),
		})
		return
	}
	// Inserting data
	insertErr := usermodel.Create(body)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, &responses.UserDefault{
			Status:  http.StatusInternalServerError,
			Message: insertErr.Error(),
		})
		return
	} else {
		c.JSON(http.StatusCreated, &responses.UserDefault{
			Status:  http.StatusCreated,
			Message: "Success create data",
		})
		return
	}

}

// @Summary Get By Session User
// @Description Show Session User
// @Accept  json
// @Produce  json
// @Success 200 {object} responses.UserData
// @Failure 401 {object} responses.AuthError "Unauthorized"
// @Failure 404 {object} responses.UserDefault "Not Found"
// @Router /user/session [get]
// @Security ApiKeyAuth
// @Tags User
func (A *UserController) ShowSessionHandler(c *gin.Context) {
	id, login_err := GetSessionID(c)
	data, _ := usermodel.FindById(id)
	if login_err == nil {
		// Check if resource exist
		if data.UserID == 0 {
			c.JSON(http.StatusNotFound, &responses.UserDefault{
				Status:  http.StatusNotFound,
				Message: "Not found",
			})
			return
		} else {
			c.JSON(http.StatusOK, &responses.UserData{
				Status:  http.StatusOK,
				Message: "Berhasil menampilkan data",
				Data:    data,
			})
			return
		}
	} else {
		c.JSON(http.StatusNotFound, &responses.UserDefault{
			Status:  http.StatusNotFound,
			Message: login_err.Error(),
		})
		return
	}
}

// @Summary Get list data User Mercant
// @Description Show all data User Mercant with pagination
// @Accept  json
// @Produce  json
// @Param   filter     query    string     false        "Some Fiter"
// @Param   sort     query    string     true        "name|asc"
// @Param   page     query    int     true        "1"
// @Param   per_page     query    int     true        "5"
// @Success 200 {object} responses.ListUser
// @Failure 401 {object} responses.AuthError "Unauthorized"
// @Failure 500 {object} responses.UserDefault "Something Wrong on Server"
// @Router /user [get]
// @Security ApiKeyAuth
// @Tags User
func (A *UserController) ListHandlerMercant(c *gin.Context) {
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
	data, err, count, errcount := usermodel.ListWhereMercant(data_user.MercantID, filter, sort, pageNo, perPage)
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
		c.JSON(http.StatusInternalServerError, &responses.UserDefault{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	} else {
		if len(data) == 0 {
			c.JSON(http.StatusOK, &responses.ListUser{
				Total:       count,
				PerPage:     pp,
				CurrentPage: pn,
				LastPage:    int(lastPage),
				NextPage:    0,
				PrevPage:    0,
				Data:        []forms.User{},
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
			c.JSON(200, &responses.ListUser{
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

// @Summary Get By Id User
// @Description Show by id User
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "Some ID"
// @Success 200 {object} responses.UserData
// @Failure 401 {object} responses.AuthError "Unauthorized"
// @Failure 404 {object} responses.UserDefault "Not Found"
// @Router /user/view/{id} [get]
// @Security ApiKeyAuth
// @Tags User
func (A *UserController) ShowHandler(c *gin.Context) {
	id := c.Param("id")
	data, _ := usermodel.FindById(id)
	// Check if resource exist
	if data.UserID == 0 {
		c.JSON(http.StatusNotFound, &responses.UserDefault{
			Status:  http.StatusNotFound,
			Message: "Not found",
		})
		return
	} else {
		c.JSON(http.StatusOK, &responses.UserData{
			Status:  http.StatusOK,
			Message: "Berhasil menampilkan data",
			Data:    data,
		})
		return
	}
}

// @Summary Login User
// @Description Login User
// @Accept  json
// @Produce  json
// @Param user body request.UserLogin true "Login User"
// @Success 200 {object} responses.AuthOK
// @Failure 400 {object} responses.AuthError "Invalid Form"
// @Failure 401 {object} responses.AuthError "Unauthorized"
// @Router /user/login [post]
// @Tags User
func (A *UserController) UserLogin(c *gin.Context) {
	var data request.UserLogin
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.AuthError{
			Status:  http.StatusBadRequest,
			Message: "Invalid Form",
		})
	} else {
		dataakun, er := usermodel.FindByEmailPass(data)
		if er == nil {
			usermodel.UpdateLastLogin(dataakun.UserID)
			mySigningKey := []byte("secret")
			// Create the Claims
			claims := forms.Payload{
				strconv.Itoa(dataakun.UserID),
				time.Now().Add(time.Hour * 99999).Unix(),
				jwt.StandardClaims{},
			}
			convert := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			token, err := convert.SignedString(mySigningKey)

			if err != nil {
				c.JSON(http.StatusNotAcceptable, responses.AuthError{
					Message: err.Error(),
					Status:  http.StatusNotAcceptable,
				})
				c.Abort()
			} else {
				c.JSON(http.StatusOK, responses.AuthOK{
					Message: "Sukses berhasil login",
					Status:  http.StatusOK,
					Expired: time.Now().Add(time.Minute * 99999).Unix(),
					Token:   token,
					User:    dataakun,
				})
				c.Abort()

			}

		} else {
			c.JSON(http.StatusUnauthorized, responses.AuthError{
				Message: "Email / Password salah",
				Status:  http.StatusUnauthorized,
			})
			c.Abort()
		}
	}
}

// @Summary Update User by ID
// @Description Update User by ID
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "Some ID"
// @Param data body request.UserUpdate true "Update User"
// @Success 201 {object} responses.UserDefault
// @Failure 406 {object} responses.UserDefault "Invalid Form"
// @Failure 401 {object} responses.AuthError "Error Not Logged In"
// @Failure 404 {object} responses.UserDefault "ID not found"
// @Router /user/{id} [put]
// @Security ApiKeyAuth
// @Tags User
func (A *UserController) PutHandler(c *gin.Context) {
	id := c.Param("id")

	var data request.UserUpdate

	// App level validation
	bindErr := c.BindJSON(&data)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, &responses.UserDefault{
			Status:  http.StatusBadRequest,
			Message: bindErr.Error(),
		})
		return
	}

	// Check if resource exist
	cekdata, _ := usermodel.FindById(id)
	if cekdata.UserID == 0 {
		c.JSON(http.StatusNotFound, &responses.UserDefault{
			Status:  http.StatusNotFound,
			Message: "Not found",
		})
		return
	}

	// Updating data
	err := usermodel.Put(id, data)

	if err != nil {
		c.JSON(http.StatusBadRequest, &responses.UserDefault{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusCreated, &responses.UserDefault{
			Status:  http.StatusCreated,
			Message: "Berhasil mengubah data",
		})
		return
	}
}

// @Summary Delete User by ID
// @Description Delete User by ID
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "Some ID"
// @Success 200 {object} responses.UserDefault
// @Failure 406 {object} responses.UserDefault "Invalid Form"
// @Failure 401 {object} responses.AuthError "Error Not Logged In"
// @Failure 404 {object} responses.UserDefault "ID not found"
// @Router /user/{id} [delete]
// @Security ApiKeyAuth
// @Tags User
func (A *UserController) DeleteHandler(c *gin.Context) {
	id := c.Param("id")

	// Check if resource exist
	cekdata, _ := usermodel.FindById(id)
	if cekdata.UserID == 0 {
		c.JSON(http.StatusNotFound, &responses.UserDefault{
			Status:  http.StatusNotFound,
			Message: "Not found",
		})
		return
	}

	// Deleting data
	err := usermodel.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &responses.UserDefault{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, &responses.UserDefault{
			Status:  http.StatusOK,
			Message: "Berhasil menghapus data",
		})
		return
	}
}
