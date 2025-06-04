package v1

import (
	"errors"
	"net/http"

	"github.com/IceWhaleTech/CasaOS-Common/utils/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

var (
	errUnsupportedMethod = errors.New("webdav: unsupported method")
	errPrefixMismatch    = errors.New("webdav: prefix mismatch")
)

const (
	StatusMulti               = 207
	StatusUnprocessableEntity = 422
	StatusLocked              = 423
	StatusFailedDependency    = 424
	StatusInsufficientStorage = 507
)

func StatusText(code int) string {
	switch code {
	case StatusMulti:
		return "Multi-Status"
	case StatusUnprocessableEntity:
		return "Unprocessable Entity"
	case StatusLocked:
		return "Locked"
	case StatusFailedDependency:
		return "Failed Dependency"
	case StatusInsufficientStorage:
		return "Insufficient Storage"
	}
	return http.StatusText(code)
}

func ServeWebDAV(ctx echo.Context) error {
	status, err := http.StatusBadRequest, errUnsupportedMethod
	w := ctx.Response().Writer
	r := ctx.Request()
	switch r.Method {
	case "OPTIONS":
		status, err = handleOptions(w, r)
	}
	if status != 0 {
		w.WriteHeader(status)
		if status != http.StatusNoContent {
			w.Write([]byte(StatusText(status)))
		}
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return nil
}

func handleOptions(w http.ResponseWriter, r *http.Request) (status int, err error) {

	logger.Error("handleOptions", zap.String("saaaa", "sasdasdas"))
	// reqPath := r.URL.Path
	// if err != nil {
	// 	return status, err
	// }
	// ctx := r.Context()
	// user := ctx.Value("user").(*model.User)
	// reqPath, err = user.JoinPath(reqPath)
	// if err != nil {
	// 	return 403, err
	// }
	allow := "OPTIONS, LOCK, PUT, MKCOL"
	// if fi, err := fs.Get(ctx, reqPath, &fs.GetArgs{}); err == nil {
	// 	if fi.IsDir() {
	// 		allow = "OPTIONS, LOCK, DELETE, PROPPATCH, COPY, MOVE, UNLOCK, PROPFIND"
	// 	} else {
	// 		allow = "OPTIONS, LOCK, GET, HEAD, POST, DELETE, PROPPATCH, COPY, MOVE, UNLOCK, PROPFIND, PUT"
	// 	}
	// }
	w.Header().Set("Allow", allow)
	// http://www.webdav.org/specs/rfc4918.html#dav.compliance.classes
	w.Header().Set("DAV", "1, 2")
	// http://msdn.microsoft.com/en-au/library/cc250217.aspx
	w.Header().Set("MS-Author-Via", "DAV")
	return 0, nil
}
