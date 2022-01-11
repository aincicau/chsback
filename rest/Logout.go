package rest

import (
	"chsback/common"
	"chsback/db"
	"chsback/entity"
	"chsback/utils"
	"net/http"
)

func Logout(rw http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "cookieMonster")
	if utils.HasError(err, "Session error") {
		http.Error(rw, "Session error", http.StatusInternalServerError)
		return
	}

	if !session.IsNew {

		var sessionDB entity.Session
		result := db.GetDB().Find(&sessionDB, "id=?", (session.Values["sessionID"]).(string))

		if result.Error != nil {
			http.Error(rw, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		result = db.GetDB().Delete(&sessionDB)

		if result.Error != nil {
			http.Error(rw, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		session.Options.MaxAge = -1
		sessionError := session.Save(r, rw)
		if utils.HasError(sessionError, "Session save error") {
			http.Error(rw, "Session save error", http.StatusInternalServerError)
			return
		}

	} else {
		http.Error(rw, common.AUTHENTICATION_NEEDED+common.PLEASE_LOGIN, http.StatusUnauthorized)
		return
	}
}
