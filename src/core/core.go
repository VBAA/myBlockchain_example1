package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

//type Blockchain struct {
//	Block []*Block
//}

//Block keeps block headers
type Block struct {
	Timestamp     int64  //区块创建时间戳
	Data          []byte //区块包含的数据
	PrevBlockHash []byte //前一个区块的哈希值
	Hash          []byte //区块自身的哈希值，用于校验区块数据有效
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{Timestamp: time.Now().Unix(), Data: []byte(data), PrevBlockHash: prevBlockHash, Hash: []byte{}}
	block.SetHash()
	return block
}

//SetHash calculates and sets block hash
//sha安全散列算法，位数越高安全程度越高
//sha256计算散列值
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}
