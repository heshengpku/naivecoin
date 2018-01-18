package block

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"time"

	"github.com/op/go-logging"

	"naivecoin/utils"
)

var log = logging.MustGetLogger("block")

// Block - Block struct
type Block struct {
	Index     int    `json:"index"`
	PreHash   string `json:"preHash"`
	Nonce     int    `json:"nonce"`
	Timestamp int64  `json:"timestamp"`
	Data      string `json:"data"`
	Hash      string `json:"hash"`
}

// BlockChain - a Chain of Block
type BlockChain []*Block

// LocalBlockChain - the local blockchain
var LocalBlockChain BlockChain

// BlockFile - the blocks data file
const BlockFile = "data/blocks.json"

func init() {
	buf, err := ioutil.ReadFile(BlockFile)
	if err != nil {
		log.Error("Open the block file failed!")
	}
	var blockchain BlockChain
	err = json.Unmarshal(buf, &blockchain)
	if err != nil {
		LocalBlockChain = NewBlockChain()
	} else {
		LocalBlockChain = blockchain
	}

}

// NewBlock - Create a new block
func NewBlock(index int, preHash string, timestamp time.Time, data string) *Block {
	return &Block{Index: index, PreHash: preHash, Timestamp: timestamp.Unix(), Data: data}
}

// NewBlockChain - Initial a new blockchain
func NewBlockChain() BlockChain {
	var blockchain BlockChain
	blockchain = append(blockchain, NewGenesisBlock())
	return blockchain
}

func (block *Block) calculateBlockHash() string {
	return utils.CalculateHash(fmt.Sprintf("%d%s%d%d%v", block.Index, block.PreHash, block.Nonce, block.Timestamp, block.Data))
}

// NewGenesisBlock - generate the genesis block
func NewGenesisBlock() *Block {
	genesisTime, _ := time.Parse(time.RFC3339, "2018-01-01T08:00:00Z") // Happy new year, 2018 Beijing time
	block := NewBlock(0, "0", genesisTime, "Genesis Block")
	// block.powRun()
	block.Nonce = 6952812
	block.Hash = "000000998cc79a8e302cbc78abb87599e6cc8671553834abef72d206c78a71c6"
	return block
}

func (blockchain BlockChain) getLatestBlock() *Block {
	return blockchain[len(blockchain)-1]
}

func (blockchain BlockChain) generateNextBlock(data string) *Block {
	preBlock := blockchain.getLatestBlock()
	nextBlock := NewBlock(preBlock.Index+1, preBlock.Hash, time.Now(), data)
	nextBlock.powRun()
	// nextBlock.Hash = nextBlock.calculateBlockHash()
	// log.Println(nextBlock)
	return nextBlock
}

func (blockchain *BlockChain) addBlock(newBlock *Block) error {
	// log.SetPrefix("Add a block - ")
	err := blockchain.isValidNewBlock(newBlock)
	if err != nil {
		return err
	}
	// log.Println(newBlock)
	*blockchain = append(*blockchain, newBlock)
	// log.Println("Success")
	return nil
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
	if !newBlock.powCheck() {
		return errors.New("invalid proof of work")
	}
	return nil
}

// GetAllBlocks - return the whole blockchain
func GetAllBlocks() BlockChain {
	return LocalBlockChain
}

// GetBlockByIndex - get the block by index, return nil when non-exist
func GetBlockByIndex(index int) *Block {
	if index < 0 || index >= len(LocalBlockChain) {
		return nil
	}
	return LocalBlockChain[index]
}

// GetBlockByHash - get the block by hash, return nil when non-exist
func GetBlockByHash(hash string) *Block {
	for _, block := range LocalBlockChain {
		if block.Hash == hash {
			return block
		}
	}
	return nil
}

// GetLatestBlock - get the latest block
func GetLatestBlock() *Block {
	return LocalBlockChain.getLatestBlock()
}

// MineBlock - mine a new block
func MineBlock(data string) *Block {
	block := LocalBlockChain.generateNextBlock(data)
	LocalBlockChain.addBlock(block)
	return block
}

// AddBlock - add a block to the blockchain, return true/false when succeeded/failed
func AddBlock(block *Block) error {
	return LocalBlockChain.addBlock(block)
}

const targetBits = 24

func (block *Block) powRun() {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	var hashInt big.Int
	maxNonce := math.MaxInt64

	start := time.Now()
	for block.Nonce = 0; block.Nonce < maxNonce; block.Nonce++ {
		hash := block.calculateBlockHash()
		hashInt.SetString(hash, 16)

		if hashInt.Cmp(target) == -1 {
			log.Debugf("\r%s", hash)
			block.Hash = hash
			break
		}
	}
	log.Debug(block.Index, " POW using ", time.Since(start))
}

func (block *Block) powCheck() bool {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	var hashInt big.Int
	hashInt.SetString(block.Hash, 16)
	return hashInt.Cmp(target) == -1
}
