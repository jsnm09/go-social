package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Index        int
	Timestamp    string
	Data         string
	PreviousHash string
	Hash         string
}

var Blockchain []Block

// function to generate a SHA-256 hash for a block
func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PreviousHash
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

// function to create new block based on previous block's hash
func generateBlock(oldBlock Block, data string) Block {
	newBlock := Block{
		Index:        oldBlock.Index + 1,
		Timestamp:    time.Now().String(),
		Data:         data,
		PreviousHash: oldBlock.Hash,
	}
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

// bool function to check if new block is valid based on previous blocks
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PreviousHash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func main() {

	genesisBlock := Block{0, time.Now().String(), "Genesis Block", "", ""}
	genesisBlock.Hash = calculateHash(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)

	newBlock := generateBlock(genesisBlock, "First block after Genesis")
	if isBlockValid(newBlock, genesisBlock) {
		Blockchain = append(Blockchain, newBlock)
	}
	newBlock2 := generateBlock(newBlock, "Second block")
	if isBlockValid(newBlock2, newBlock) {
		Blockchain = append(Blockchain, newBlock2)
	}

	for _, block := range Blockchain {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("PreviousHash: %s,\n", block.PreviousHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println()
	}
}
