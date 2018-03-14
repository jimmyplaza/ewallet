package main

import (
	"encoding/json"
	"log"
	"math/big"

	"github.com/onrik/ethrpc"
)

func nodeFormat(node json.RawMessage) interface{} {
	type nodeAlias struct {
		Enode string `json:"enode"`
		Name  string `json:"name"`
	}

	byteNode, err := json.Marshal(node)
	if err != nil {
		log.Println(err)
	}
	var n nodeAlias
	err = json.Unmarshal(byteNode, &n)
	if err != nil {
		log.Println(err)
	}
	return n
}

func blockFormat(block *ethrpc.Block) interface{} {
	type blockAlias struct {
		Number           int                  `json:"-"`
		Hash             string               `json:"hash"`
		ParentHash       string               `json:"parentHash"`
		Nonce            string               `json:"-"`
		Sha3Uncles       string               `json:"-"`
		LogsBloom        string               `json:"-"`
		TransactionsRoot string               `json:"-"`
		StateRoot        string               `json:"-"`
		Miner            string               `json:"miner"`
		Difficulty       big.Int              `json:"-"`
		TotalDifficulty  big.Int              `json:"-"`
		ExtraData        string               `json:"-"`
		Size             int                  `json:"-"`
		GasLimit         int                  `json:"gasLimit"`
		GasUsed          int                  `json:"gasUsed"`
		Timestamp        int                  `json:"-"`
		Uncles           []string             `json:"-"`
		Transactions     []ethrpc.Transaction `json:"-"`
		Difficulty_      int64                `json:"difficulty"`
		TotalDifficulty_ int64                `json:"totalDifficulty"`
	}

	byteBlock, err := json.Marshal(block)
	if err != nil {
		log.Println(err)
	}
	var b blockAlias
	err = json.Unmarshal(byteBlock, &b)
	if err != nil {
		log.Println(err)
	}
	return b
}

func transFormat(trans *ethrpc.Transaction) interface{} {
	type transAlias struct {
		Hash             string  `json:"hash"`
		Nonce            int     `json:"nonce"`
		BlockHash        string  `json:"blockHash"`
		BlockNumber      *int    `json:"blockNumber"`
		TransactionIndex *int    `json:"-"`
		From             string  `json:"from"`
		To               string  `json:"to"`
		Value            big.Int `json:"-"`
		Gas              int     `json:"gas"`
		GasPrice         big.Int `json:"-"`
		Input            string  `json:"-"`

		GasPrice_ int64 `json:"gasPrice"`
		Value_    int64 `json:"value"`
	}

	byteTrans, err := json.Marshal(trans)
	if err != nil {
		log.Println(err)
	}
	var t transAlias
	err = json.Unmarshal(byteTrans, &t)
	if err != nil {
		log.Println(err)
	}
	return t
}
