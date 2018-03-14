package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-zoo/bone"
)

// nodeInfo for enode info
func nodeInfo(w http.ResponseWriter, r *http.Request) {
	result := New()

	node, err := ethClient.Call("admin_nodeInfo")

	if err != nil {
		log.Println(err)
	}

	result.State = 1
	result.Content["info"] = nodeFormat(node)
	Render.JSON(w, http.StatusOK, result)
	return
}

// blockInfo
func blockInfo(w http.ResponseWriter, r *http.Request) {
	result := New()
	blocknum := bone.GetValue(r, "block_number")
	bn, _ := strconv.Atoi(blocknum)

	block, err := ethClient.EthGetBlockByNumber(bn, false)
	if block == nil {
		result.State = 0
		result.Content["error"] = "no struct "
		Render.JSON(w, http.StatusOK, result)
		return
	}
	if err != nil {
		log.Println(err)
	}

	result.State = 1
	result.Content["block"] = blockFormat(block)
	Render.JSON(w, http.StatusOK, result)
	return
}

// transationInfo
func transationInfo(w http.ResponseWriter, r *http.Request) {
	result := New()
	transationhash := bone.GetValue(r, "transation_hash")

	t, err := ethClient.EthGetTransactionByHash(transationhash)

	if err != nil {
		log.Println(err)
	}

	result.State = 1
	result.Content["transation"] = transFormat(t)
	Render.JSON(w, http.StatusOK, result)
	return
}
