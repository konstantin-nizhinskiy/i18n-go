package i18n

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
func GetGin(key string,c *gin.Context, data ...map[string]interface{}) string {
	return Get(key,getLangGin(c),data...)
}


func getLangGin(c *gin.Context) string {
	cookie,err:=c.Request.Cookie(cookieName)
	if err == nil {
		return cookie.Value
	}
	cookieNew:=http.Cookie{
		Name:     cookieName,
		Expires: time.Now().Add(365 * 24 * time.Hour),
		Value:    langDefault,
		HttpOnly: false,
		Secure: false,
		Path:     "/"}
	http.SetCookie(c.Writer, &cookieNew)
	return cookieNew.Value
}

