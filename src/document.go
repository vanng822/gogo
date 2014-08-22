package main

import "fmt"

type Document map[string]interface{}

// Has check if a key exist in document
func (d Document) Has(k string) bool {
	_, ok := d[k]
	return ok
}

// Get returns value of a key if key exists else panic
func (d Document) Get(k string) interface{} {
	if v, ok := d[k]; ok {
		return v
	}
	panic(fmt.Sprintf("Try to access field '%s' which does not exist", k))
}

// Set add a key/value to document
func (d Document) Set(k string, v interface{}) {
	d[k] = v
}

type Document2 Document


type SolrResult struct {
	docs []Document
}

type ResultParser interface {
	Parse(response string) (*SolrResult, error)
}

type ResultParser2 struct {}

func (parser *ResultParser2) Parse(response string) (*SolrResult, error) {
	result := &SolrResult{}
	docs := make([]Document, 2)
	docs[0] = Document2{"hell": "world"}
	docs[1] = Document2{"hello": "kitty"}
	result.docs = docs
	return result, nil
}

func main() {
	p := ResultParser2{}
	res, err := p.Parse("")
	
	fmt.Println(res.docs[0].Get("hello").(string))
}