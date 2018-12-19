package controller

import "net/http"

func handleLetter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello there"))
}

func PermissionLetterHandler() {
	http.HandleFunc("/", handleLetter)
}
