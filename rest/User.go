package rest

import (
	"chsback/common"
	"chsback/db"
	"chsback/entity"
	"chsback/utils"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	reqBody := r.Body
	bodyBytes, err := ioutil.ReadAll(reqBody)
	if utils.HasError(err, common.READ_DATA_ERROR) {
		http.Error(rw, common.INTERNAL_SERVER_ERROR, http.StatusInternalServerError)
		return
	}

	var user entity.User
	err = json.Unmarshal(bodyBytes, &user)
	if utils.HasError(err, common.UNMARSHAL_ERROR) {
		http.Error(rw, common.DATA_INTEGRITY_ERROR, http.StatusBadRequest)
		return
	}

	user.Password = base64.StdEncoding.EncodeToString([]byte(user.Password))

	result := db.GetDB().Find(&entity.User{}).Where("email=?", user.Email)

	if result.RecordNotFound() {
		result = db.GetDB().Create(&user)
		if result.Error != nil {
			logrus.WithError(result.Error).Error(common.USER_CREATION_ERROR)
			http.Error(rw, common.USER_CREATION_ERROR, http.StatusBadRequest)
			return
		}
	} else {
		logrus.WithError(result.Error).Error(common.USER_EXISTS)
		http.Error(rw, common.USER_EXISTS, http.StatusBadRequest)
	}
}
