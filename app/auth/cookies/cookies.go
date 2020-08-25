package cookies

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Toringol/nonlinearity/app/auth/sessionManager"
	"github.com/Toringol/nonlinearity/app/model"
	"github.com/gorilla/securecookie"
	"github.com/labstack/echo"
)

var (
	SessManager *sessionManager.SessionManager
)

// cookieHandler - secure cookie with two concatinated parts
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32),
)

// SetSession - call to sessManager to create record in redisDB and set cookie in ctx
func SetSession(ctx echo.Context, userData *model.User) (*http.Cookie, error) {
	sessID, err := SessManager.Create(&model.Session{
		Username:  userData.Username,
		Useragent: ctx.Request().UserAgent(),
	})
	if err != nil {
		return nil, err
	}

	value := map[string]string{
		"session_id": sessID.ID,
		"user_id":    strconv.FormatInt(userData.ID, 10),
	}

	if encoded, err := cookieHandler.Encode("session_token", value); err == nil {
		expiration := time.Now().Add(24 * time.Hour)
		cookie := &http.Cookie{
			Name:    "session_token",
			Value:   encoded,
			Path:    "/",
			Expires: expiration,
		}
		ctx.SetCookie(cookie)
		return cookie, nil
	}

	return nil, nil
}

// ClearSession - call to sessManager to delete record in redisDB and clear cookie in ctx
func ClearSession(ctx echo.Context) error {
	err := SessManager.Delete(ReadSessionID(ctx))
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
	}
	ctx.SetCookie(cookie)

	return nil
}

// СheckSession - call to sessManager to check session and return sess(Username and UserAgent)
func СheckSession(ctx echo.Context) (*model.Session, error) {
	cookieSessionID := ReadSessionID(ctx)
	if cookieSessionID == nil {
		return nil, nil
	}

	sess, err := SessManager.Check(cookieSessionID)
	if err != nil {
		return nil, err
	}

	return sess, nil
}

// ReadSessionID - decrypt cookie and return sessionID
func ReadSessionID(ctx echo.Context) *model.SessionID {
	if cookie, err := ctx.Request().Cookie("session_token"); err == nil {
		value := make(map[string]string)
		if err = cookieHandler.Decode("session_token", cookie.Value, &value); err == nil {
			return &model.SessionID{ID: value["session_id"]}
		}
	}

	return nil
}

// ReadUserID - decrypt cookie and return userID
func ReadUserID(ctx echo.Context) *model.SessionID {
	if cookie, err := ctx.Request().Cookie("session_token"); err == nil {
		value := make(map[string]string)
		if err = cookieHandler.Decode("session_token", cookie.Value, &value); err == nil {
			return &model.SessionID{ID: value["user_id"]}
		}
	}

	return nil
}
