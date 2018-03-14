package main

import (
	"context"
	"log"
	"net/http"
	"reflect"
	"strconv"

	// "github.com/getamis/eth-client/client"
	"github.com/getamis/eth-client/client"
	"github.com/go-zoo/bone"
	"github.com/onrik/ethrpc"
)

// nodeInfo for enode info
func nodeInfo(w http.ResponseWriter, r *http.Request) {
	result := New()

	client := ethrpc.New("http://127.0.0.1:8545")
	node, err := client.Call("admin_nodeInfo")

	if err != nil {
		log.Println(err)
	}
	log.Println(reflect.TypeOf(node))

	result.State = 1
	result.Content["info"] = nodeFormat(node)
	Render.JSON(w, http.StatusOK, result)
	return
}

// nodeInfo for enode info
func nodeInfo_old(w http.ResponseWriter, r *http.Request) {
	result := New()

	url := "http://127.0.0.1:8545"
	c, err := client.Dial(url)
	if err != nil {
		log.Println("Failed to dial, url: ", url, ", err: ", err)
		return
	}
	info, err := c.NodeInfo(context.Background())
	log.Printf("%#v", info)

	result.State = 1
	result.Content["info"] = info
	Render.JSON(w, http.StatusOK, result)
	return
}

// blockInfo
func blockInfo(w http.ResponseWriter, r *http.Request) {
	result := New()
	blocknum := bone.GetValue(r, "block_number")
	bn, _ := strconv.Atoi(blocknum)

	client := ethrpc.New("http://127.0.0.1:8545")
	block, err := client.EthGetBlockByNumber(bn, false)
	if block == nil {
		result.State = 0
		result.Content["error"] = "no struct "
		Render.JSON(w, http.StatusOK, result)
		return
	}
	if err != nil {
		log.Println(err)
	}
	log.Printf("%#v", block)

	result.State = 1
	result.Content["block"] = blockFormat(block)
	Render.JSON(w, http.StatusOK, result)
	return
}

// transationInfo
func transationInfo(w http.ResponseWriter, r *http.Request) {
	result := New()
	transationhash := bone.GetValue(r, "transation_hash")

	client := ethrpc.New("http://127.0.0.1:8545")
	t, err := client.EthGetTransactionByHash(transationhash)

	if err != nil {
		log.Println(err)
	}
	log.Printf("%#v", t)

	result.State = 1
	result.Content["transation"] = transFormat(t)
	Render.JSON(w, http.StatusOK, result)
	return
}
