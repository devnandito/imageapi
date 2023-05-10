package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/devnandito/imageapi/models"
	"github.com/gorilla/sessions"
)

var (
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func CheckAuth() models.Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			session, _ := store.Get(r, "cookie-name")
			if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
				http.Redirect(w, r, "/", http.StatusFound)
				return
			} else {
				f(w, r)
			}
		}
	}
}

func Logging() models.Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}

// func CheckAuth() models.Middleware {
// 	return func(f http.HandlerFunc) http.HandlerFunc {
// 		return func(w http.ResponseWriter, r *http.Request) {
// 			flag := false
// 			response, err := cls.GetOneClientGorm(1)

// 			if err != nil {
// 				panic(err)
// 			}

// 			if response.Ci == "3796986" {
// 				flag = true
// 				fmt.Println("User authenticated", response)
// 				fmt.Println("Checking authentication")
// 			}

// 			if flag {
// 				f(w, r)
// 			}else {
// 				fmt.Println("User not authenticated", response)
// 				http.Redirect(w, r, "/", http.StatusFound)
// 			}
// 		}
// 	}
// }