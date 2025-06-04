package route

import (
	"net/http"

	"golang.org/x/net/webdav"
)

func InitDavRoute() http.Handler {
	fs := &webdav.Handler{
		Prefix:     "/",
		FileSystem: webdav.Dir("./"),
		LockSystem: webdav.NewMemLS(),
	}
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		username, password, ok := req.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, "WebDAV: need authorized!", http.StatusUnauthorized)
			return
		}
		//验证用户名 / 密码
		if username != "admin" || password != "123456" {
			http.Error(w, "WebDAV: need authorized!", http.StatusUnauthorized)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")

		w.Header().Set("Access-Control-Allow-Credentials", "true")

		fs.ServeHTTP(w, req)
	})
}
