package httpstudy

import (
	"net/http"
)

// HTTPMain ...
func HTTPMain() {
	http.ListenAndServe(":1233", http.FileServer(http.Dir(".")))
}
