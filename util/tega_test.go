package util

import (
	"log"
	"testing"
)

func TestOnInit(t *testing.T) {
	hosts := ListHosts()
	log.Printf("%v", hosts)
}
