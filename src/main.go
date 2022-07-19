package main

import (
	"fmt"
	"go-mod/blockchain"
	"html/template"
	"log"
	"net/http"
)

const port string = ":3000"

type homedata struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func main() {
	chain := blockchain.Getblockchain()
	chain.Addblock("Second block")
	chain.Addblock("Third block")
	chain.Addblock("Fourth block")
	for _, block := range chain.Allblocks() {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/home.gohtml"))
		data := homedata{"Home", blockchain.Getblockchain().Allblocks()}
		tmpl.Execute(w, data)
	})
	log.Fatal(http.ListenAndServe(port, nil))
}
