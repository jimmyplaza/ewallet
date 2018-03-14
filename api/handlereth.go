package main

import (
	"net/http"
)

// nodeInfo for enode info
func nodeInfo(w http.ResponseWriter, r *http.Request) {
	// var err error

	result := New()

	// lan := bone.GetValue(r, "lan")

	result.State = 1
	Render.JSON(w, http.StatusOK, result)
	return
}
