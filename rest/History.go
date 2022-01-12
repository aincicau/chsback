package rest

import (
	"chsback/common"
	"chsback/db"
	"chsback/entity"
	"chsback/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetHistory(rw http.ResponseWriter, r *http.Request) {
	sessionDetails, err := store.Get(r, "cookieMonster")

	if utils.HasError(err, common.INTERNAL_SERVER_ERROR) {
		http.Error(rw, "Cookie error", http.StatusInternalServerError)
		return
	}

	if sessionDetails.IsNew {
		http.Error(rw, common.AUTHENTICATION_NEEDED+common.PLEASE_LOGIN, http.StatusUnauthorized)
		return
	} else {
		sessionID := (sessionDetails.Values["sessionID"]).(string)

		var sessionDB entity.Session
		result := db.GetDB().Find(&sessionDB, "id=?", sessionID)
		if result.RecordNotFound() || result.RowsAffected == 0 {
			http.Error(rw, common.AUTHENTICATION_NEEDED+common.PLEASE_LOGIN, http.StatusUnauthorized)
			return
		}

		var histories []entity.History
		result = db.GetDB().Find(&histories, "user_id=?", sessionDB.UserID)

		if result.RecordNotFound() {
			http.Error(rw, common.NO_RECORD_FOUND, http.StatusInternalServerError)
			return
		}
		if result.Error != nil {
			http.Error(rw, result.Error.Error(), http.StatusInternalServerError)
		}

		dataToBeSent, err := json.Marshal(histories)
		if utils.HasError(err, common.MARSHAL_ERROR) {
			http.Error(rw, common.INTERNAL_SERVER_ERROR, http.StatusInternalServerError)
			return
		}

		rw.Write(dataToBeSent)
	}
}

func PostHistory(rw http.ResponseWriter, r *http.Request) {
	sessionDetails, err := store.Get(r, "cookieMonster")

	if utils.HasError(err, common.INTERNAL_SERVER_ERROR) {
		http.Error(rw, "Cookie error", http.StatusInternalServerError)
		return
	}

	if sessionDetails.IsNew {
		http.Error(rw, common.AUTHENTICATION_NEEDED+common.PLEASE_LOGIN, http.StatusUnauthorized)
		return
	} else {
		sessionID := (sessionDetails.Values["sessionID"]).(string)

		var sessionDB entity.Session
		result := db.GetDB().Find(&sessionDB, "id=?", sessionID)
		if result.RecordNotFound() || result.RowsAffected == 0 {
			http.Error(rw, common.AUTHENTICATION_NEEDED+common.PLEASE_LOGIN, http.StatusUnauthorized)
			return
		}

		reqBody := r.Body
		bodyBytes, err := ioutil.ReadAll(reqBody)
		if utils.HasError(err, common.READ_DATA_ERROR) {
			http.Error(rw, common.INTERNAL_SERVER_ERROR, http.StatusInternalServerError)
			return
		}

		var history entity.History
		err = json.Unmarshal(bodyBytes, &history)
		if utils.HasError(err, common.UNMARSHAL_ERROR) {
			http.Error(rw, common.DATA_INTEGRITY_ERROR, http.StatusBadRequest)
			return
		}
		history.UserID = sessionDB.UserID

		result = db.GetDB().Create(&history)
		if result.Error != nil {
			http.Error(rw, "Unable to add history", http.StatusInternalServerError)
			return
		}
	}
}
