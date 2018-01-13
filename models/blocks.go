package models

import (
	"errors"
	"fmt"
	"log"
	"time"

	"naivecoin/utils"
)

// Block - Block struct
type Block struct {
	Index     int    `json:"index"`
	PreHash   string `json:"preHash"`
	Timestamp int64  `json:"timestamp"`
	Data      string `json:"data"`
	Hash      string `json:"hash"`
}

// BlockChain - a Chain of Block
type BlockChain []*Block

var LocalBlockChain BlockChain

func init() {
	LocalBlockChain = NewBlockChain()
}

// NewBlock - Create a new block
func NewBlock(index int, preHash string, timestamp time.Time, data string) *Block {
	// log.SetPrefix("New Block - ")
	var b Block
	b.Index = index
	b.PreHash = preHash
	b.Timestamp = timestamp.Unix()
	b.Data = data
	// log.Printf("Success index: %d", b.Index)
	return &b
}

// NewBlockChain - Initial a new blockchain
func NewBlockChain() BlockChain {
	// log.SetPrefix("New BlockChain - ")
	var blockchain BlockChain
	blockchain = append(blockchain, getGenesisBlock())
	// log.Println("Success")
	return blockchain
}

func (block *Block) calculateBlockHash() string {
	return utils.CalculateHash(fmt.Sprintf("%d%s%d%v", block.Index, block.PreHash, block.Timestamp, block.Data))
}

func getGenesisBlock() *Block {
	genesisTime, _ := time.Parse(time.RFC3339, "2018-01-01T00:00:00Z")
	block := NewBlock(0, "0", genesisTime, "Genesis Block")
	block.Hash = block.calculateBlockHash()
	// log.Println(block)
	return block
}

func (blockchain BlockChain) getLatestBlock() *Block {
	return blockchain[len(blockchain)-1]
}

func (blockchain BlockChain) generateNextBlock(data string) *Block {
	preBlock := blockchain.getLatestBlock()
	nextBlock := NewBlock(preBlock.Index+1, preBlock.Hash, time.Now(), data)
	nextBlock.Hash = nextBlock.calculateBlockHash()
	// log.Println(nextBlock)
	return nextBlock
}

func (blockchain *BlockChain) addBlock(newBlock *Block) bool {
	// log.SetPrefix("Add a block - ")
	err := blockchain.isValidNewBlock(newBlock)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	// log.Println(newBlock)
	*blockchain = append(*blockchain, newBlock)
	// log.Println("Success")
	return true
}

func (blockchain BlockChain) isValidNewBlock(newBlock *Block) error {
	preBlock := blockchain.getLatestBlock()
	if preBlock.Index+1 != newBlock.Index {
		return errors.New("invalid index")
	}
	if preBlock.Hash != newBlock.PreHash {
		return errors.New("invalid previous hash")
	}
	if newBlock.calculateBlockHash() != newBlock.Hash {
		return errors.New("invalid hash")
	}
	return nil
}

func GetAllBlocks() BlockChain {
	return LocalBlockChain
}

func GetBlockByIndex(index int) *Block {
	if index < 0 || index >= len(LocalBlockChain) {
		return nil
	}
	return LocalBlockChain[index]
}

func GetBlockByHash(hash string) *Block {
	for _, block := range LocalBlockChain {
		if block.Hash == hash {
			return block
		}
	}
	return nil
}

func GetLatestBlock() *Block {
	return LocalBlockChain.getLatestBlock()
}

func MineBlock(data string) *Block {
	block := LocalBlockChain.generateNextBlock(data)
	LocalBlockChain.addBlock(block)
	return block
}

func AddBlock(block *Block) bool {
	return LocalBlockChain.addBlock(block)
}
