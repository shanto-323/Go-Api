package middlewere

import (
	"errors"
	"net/http"

	"github.com/shanto-323/Go-Api/api"
	"github.com/shanto-323/Go-Api/internal/tools"
  log "github.com/sirupsen/logrus"
)


var UnAuthError = errors.New("Invalid username or Token")

func Auth(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    var username string = r.URL.Query().Get("username")
    var token = r.Header.Get("Auth")
    var err error

    //Local Var
    if username == "" || token == "" {
      log.Error(UnAuthError)
      api.RequestErrorHandler(w,UnAuthError)
      return
    }

    var database *tools.DatabaseInterface
    database , err = tools.NewDatabase()

    if err != nil {
      api.InternalErrorHandler(w)
      return
    }

    var login *tools.LoginDetail
    login = (*database).GetUserLoginDetails(username)

    if(login == nil || (token != (*login).AuthToken)){
      log.Error(UnAuthError)
      api.RequestErrorHandler(w, UnAuthError)
      return
    }

    next.ServeHTTP(w,r)
  })
}
