package assignment01bca

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

var Blockchain []Block

// CalculateHash calculates the hash of a given string using SHA256
func CalculateHash(input string) string {
	hash := sha256.Sum256([]byte(input))
	return fmt.Sprintf("%x", hash)
}

// NewBlock creates a new block and appends it to the blockchain
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{Transaction: transaction, Nonce: nonce, PreviousHash: previousHash}
	block.Hash = CalculateHash(transaction + string(nonce) + previousHash)
	Blockchain = append(Blockchain, *block)
	return block
}

// DisplayBlocks displays all blocks in the blockchain
func DisplayBlocks() {
	for _, block := range Blockchain {
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Hash: %s\n\n", block.Hash)
	}
}

// ChangeBlock changes the transaction of a block at a given index
func ChangeBlock(index int, newTransaction string) {
	if index < 0 || index >= len(Blockchain) {
		fmt.Println("Invalid block reference.")
		return
	}
	Blockchain[index].Transaction = newTransaction
	Blockchain[index].Hash = CalculateHash(newTransaction + string(Blockchain[index].Nonce) + Blockchain[index].PreviousHash)
}

// VerifyChain verifies the integrity of the blockchain
func VerifyChain() bool {
	for i := 1; i < len(Blockchain); i++ {
		currentBlock := Blockchain[i]
		previousBlock := Blockchain[i-1]

		if currentBlock.PreviousHash != previousBlock.Hash {
			fmt.Println("Previous hash does not match.")
			return false
		}

		if currentBlock.Hash != CalculateHash(currentBlock.Transaction+string(currentBlock.Nonce)+currentBlock.PreviousHash) {
			fmt.Println("Current hash does not match calculated hash.")
			return false
		}
	}
	return true
}
