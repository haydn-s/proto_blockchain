package main

import (
	// Go Imports
	"bytes"
	"crypto/sha256"
	"strconv"
	// GitHub Imports
	/* "github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	*/)

type Block struct {
	Index     int
	Timestamp string
	FILL_DATA int
	Hash      string
	PrevHash  string
}

var Blockchain []Block

func (b *Block) SetHash() {
	index := b.Index
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevHash, b.FILL_DATA, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}
