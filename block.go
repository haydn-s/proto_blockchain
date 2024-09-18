package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Name      string
	Email     string
	Cell      string
	Hash      string
	PrevHash  string
}

func SetHash(b Block) string {
	record := string(b.Index) + string(b.Timestamp) + b.Name + b.Email + b.Cell + string(b.PrevHash)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}

func CreateBlock(prevBlock Block, name string, email string, cell string) (Block, error) {
	var newBlock Block

	t := time.Now()

	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Name = name
	newBlock.Email = email
	newBlock.Cell = cell
	newBlock.PrevHash = prevBlock.Hash
	newBlock.Hash = SetHash(newBlock)

	return newBlock, nil
}

func BlockValidity(newBlock, prevBlock Block) bool {
	if prevBlock.Index+1 != newBlock.Index {
		return false
	}

	if prevBlock.Hash != newBlock.PrevHash {
		return false
	}

	if SetHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}
