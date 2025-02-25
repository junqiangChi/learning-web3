package geth

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestGeth(t *testing.T) {
	log.Println("geth GethDemo...")
	GethDemo()
	log.Println("geth GethDemo1...")
	GethDemo1()
}
