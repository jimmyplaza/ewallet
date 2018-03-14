package main

import "net/http"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := make(map[string]string)
	Render.HTML(w, http.StatusOK, "index", ctx)
}
