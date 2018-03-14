package main

import (
	"encoding/json"
	"log"
	"math/big"

	"github.com/onrik/ethrpc"
)

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
