package rest

import (
	"chsback/common"
	"chsback/db"
	"chsback/entity"
	"chsback/utils"
	"encoding/json"
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
