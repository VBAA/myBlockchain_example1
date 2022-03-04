package core

//Blcokchain keeps a sequence of Blocks
type Blockchain struct {
	Blocks []*Block //数组
}

//创世纪块,data固定/默认，没有前一个块
//NewGenesisBlock creates and returns genesis Block
func NewGenesisBlcok() *Block {
	return NewBlock("Genesis BLock", []byte{})
}

//创建一个新的区块链
//NewBlockchain creates a new Blockchain with genesis Blcok
func NewBlockchain() *Blockchain {
	return &Blockchain{Blocks: []*Block{NewGenesisBlcok()}}
}

//AddBlock saves provided data as a block in the block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, preBlock.Hash) //保证hash的存在，确保数据没有被伪造，保存数据形成完整链条，不可修改
	bc.Blocks = append(bc.Blocks, newBlock)
}
