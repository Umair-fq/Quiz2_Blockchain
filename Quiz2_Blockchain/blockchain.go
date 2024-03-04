package blockchain

// GenerateNewBlock creates a new block using the provided data and previous hash
func GenerateNewBlock(content string, prevHash string) *Block {
	newBlock := &Block{
		CreatedAt: time.Now().Unix(),
		Content:   content,
		PrevHash:  prevHash,
	}
	newBlock.ThisHash = hashBlock(newBlock)
	return newBlock
}

// ShowBlocks displays all blocks in the blockchain, printing key information
func ShowBlocks(chain []*Block) {
	for _, b := range chain {
		fmt.Printf("CreatedAt: %d, Content: %s, PrevHash: %s, Hash: %s\n",
			b.CreatedAt, b.Content, b.PrevHash, b.ThisHash)
	}
}

// UpdateBlock modifies the content of a given block and updates its hash to reflect the change
func UpdateBlock(b *Block, newContent string) {
	b.Content = newContent
	b.ThisHash = hashBlock(b)
}
