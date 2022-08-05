package main

import (
	"go-mod/blockchain"
	"log"
	"net/http"
	"text/template"
)

const (
	port        string = ":3000"
	templateDir string = "templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	// tmpl := template.Must(template.ParseFiles("templates/pages/home.gohtml"))
	data := homeData{"Home", blockchain.Getblockchain().Allblocks()}
	templates.ExecuteTemplate(rw, "home", data)

}

func main() {

	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(port, nil))
}

// chain := blockchain.Getblockchain()
// chain.Addblock("Second block")
// chain.Addblock("Third block")
// chain.Addblock("Fourth block")
// for _, block := range chain.Allblocks() {
// 	fmt.Printf("Data: %s\n", block.Data)
// 	fmt.Printf("Hash: %s\n", block.Hash)
// 	fmt.Printf("PrevHash: %s\n", block.PrevHash)
// }
