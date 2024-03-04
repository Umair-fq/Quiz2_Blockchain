package main

import (
	"fmt"
	"/blockchain"
)

func main() {
	originBlock := blockchain.GenerateNewBlock("Origin Block", "")

	var chain []*blockchain.Block
	chain = append(chain, originBlock)

	// Append some blocks to the blockchain
	chain = append(chain, blockchain.GenerateNewBlock("Block One", chain[len(chain)-1].ThisHash))
	chain = append(chain, blockchain.GenerateNewBlock("Block Two", chain[len(chain)-1].ThisHash))

	// Display blockchain before changes
	fmt.Println("Blockchain prior to updates:")
	blockchain.ShowBlocks(chain)

	// Updating a block's content
	fmt.Println("\nUpdating Block One's content...")
	blockchain.UpdateBlock(chain[1], "Updated Block One")

	// Display blockchain after changes
	fmt.Println("\nBlockchain following updates:")
	blockchain.ShowBlocks(chain)
}
