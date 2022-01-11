package rest

import (
	"chsback/common"
	"chsback/db"
	"chsback/entity"
	"chsback/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func Login(rw http.ResponseWriter, r *http.Request) {

	reqbody := r.Body
	bodyBytes, err := ioutil.ReadAll(reqbody)

	if utils.HasError(err, common.READ_DATA_ERROR) {
		http.Error(rw, common.INTERNAL_SERVER_ERROR, http.StatusInternalServerError)
		return
	}

	var user entity.User
	err = json.Unmarshal(bodyBytes, &user)

	if utils.HasError(err, common.UNMARSHAL_ERROR) {
		http.Error(rw, common.INTERNAL_SERVER_ERROR, http.StatusInternalServerError)
		return
	}

	var userDB entity.User
	result := db.GetDB().Find(&userDB, "username=?", user.Username)

	if result.RecordNotFound() {
		http.Error(rw, "Username does not exist", http.StatusUnauthorized)
		return
	}

	if utils.Encode(user.Password) != userDB.Password {
		http.Error(rw, "Incorrect password", http.StatusUnauthorized)
		return
	}

	session, err := store.Get(r, "cookieMonster")

	if utils.HasError(err, "Session error") {
		http.Error(rw, "Session error", http.StatusInternalServerError)
		return
	}

	if session.IsNew {

		session.Values["sessionID"] = uuid.NewString()
		session.Options.MaxAge = 3600
		session.Options.HttpOnly = true
		sessionError := session.Save(r, rw)

		if utils.HasError(sessionError, "Session save error") {
			http.Error(rw, "Session save error", http.StatusInternalServerError)
			return
		}

		dbSession := entity.Session{
			ID:             (session.Values["sessionID"]).(string),
			UserID:         userDB.ID,
			ExpirationDate: time.Now().UTC().Add(time.Hour),
		}

		result = db.GetDB().Create(&dbSession)

		if result.Error != nil {
			http.Error(rw, result.Error.Error(), http.StatusInternalServerError)
		}

	}
	rw.Write([]byte(userDB.Username))
}
