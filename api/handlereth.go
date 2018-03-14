package main

import (
	"encoding/json"
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
		result.State = 0
		result.Content["error"] = err
		Render.JSON(w, http.StatusOK, result)
		return
	}

	Render.JSON(w, http.StatusOK, nodeFormat(node))
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
		result.Content["error"] = "No block struct found"
		Render.JSON(w, http.StatusOK, result)
		return
	}
	if err != nil {
		log.Println(err)
		result.State = 0
		result.Content["error"] = err
		Render.JSON(w, http.StatusOK, result)
		return
	}

	Render.JSON(w, http.StatusOK, blockFormat(block))
	return
}

// transationInfo
func transationInfo(w http.ResponseWriter, r *http.Request) {
	result := New()
	transationhash := bone.GetValue(r, "transation_hash")

	t, err := ethClient.EthGetTransactionByHash(transationhash)

	if err != nil {
		log.Println(err)
		result.State = 0
		result.Content["error"] = err
		Render.JSON(w, http.StatusOK, result)
		return
	}

	Render.JSON(w, http.StatusOK, transFormat(t))
	return
}

// startMiner
func startMiner(w http.ResponseWriter, r *http.Request) {
	result := New()

	ret, err := ethClient.Call("miner_start")

	if err != nil {
		log.Println(err)
		result.State = 0
		result.Content["error"] = err
		Render.JSON(w, http.StatusOK, result)
		return
	}

	Render.JSON(w, http.StatusOK, ret)
	return
}

// stopMiner
func stopMiner(w http.ResponseWriter, r *http.Request) {
	result := New()

	ret, err := ethClient.Call("miner_stop")

	if err != nil {
		log.Println(err)
		result.State = 0
		result.Content["error"] = err
		Render.JSON(w, http.StatusOK, result)
		return
	}

	Render.JSON(w, http.StatusOK, ret)
	return
}

// sendTrans
func sendTrans(w http.ResponseWriter, r *http.Request) {
	result := New()
	type transhash struct {
		Transation string `json:"transation"`
	}

	var trans transhash
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&trans); err != nil {
		log.Println(err)
		result.State = 0
		result.Content["error"] = err.Error()
		Render.JSON(w, http.StatusOK, result)
		return
	}

	t, err := ethClient.EthSendRawTransaction(trans.Transation)

	if err != nil {
		log.Println(err)
		result.State = 0
		result.Content["error"] = err
		Render.JSON(w, http.StatusOK, result)
		return
	}

	Render.JSON(w, http.StatusOK, t)
	return
}
