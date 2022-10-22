package routes

import "net/http"

func CarHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Youness!\n"))
}
