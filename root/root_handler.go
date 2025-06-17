package root

import "net/http"

func RootHandle(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "root/root.html")
}
