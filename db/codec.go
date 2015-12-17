package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	nlan "github.com/araobp/nlan/model/nlan"
	"github.com/golang/protobuf/proto"
	"log"
	"os"
)

const DBFILE = "./conf.db"

var PUT byte = 1
var DELETE byte = 2

// Record in TLV format
func record(ope byte, pb proto.Message) []byte {
	var buf bytes.Buffer
	pbdata, _ := proto.Marshal(pb)
	// T
	buf.Write([]byte{ope})
	// L
	b := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(b, int64(len(pbdata)))
	fmt.Printf("len: %v\n", b)
	buf.Write(b)
	// V
	buf.Write(pbdata)
	return buf.Bytes()
}

// Appends a record to the file
func Append(f *os.File, ope byte, pb proto.Message) {
	fmt.Printf("ope: %v, pb: %v\n", ope, pb)
	r := record(ope, pb)
	_, _ = f.Write(r)
	fmt.Printf("wire: %v\n", r)
}

// Loads the data from the file
func Load(f *os.File, pb interface{}) {
	_, err := f.Seek(0, 0)
	if err != nil {
		log.Fatal("File seek failure")
	}
	reader := bufio.NewReader(f)
	for {
		ope, err := reader.Peek(1)
		if err != nil {
			break
		}
		b, _ := reader.Peek(binary.MaxVarintLen64)
		l, _ := binary.Varint(b)
		data, _ := reader.Peek(int(l))

		proto.Unmarshal(data, pb.(proto.Message))
		fmt.Printf("ope: %v, len: %v, pb: %v\n", ope, l, pb)
	}
}

func main() {
	f, err := os.OpenFile(DBFILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Unable to open the file: %s", DBFILE)
	}
	defer f.Close()
	props := nlan.VhostProps{
		Network: "10.10.10.10/24",
		Vhosts:  2,
	}
	vhosts := nlan.Vhosts{
		VhostProps: []*nlan.VhostProps{&props},
	}
	Append(f, PUT, &vhosts)
	pb := new(nlan.Vhosts)
	Load(f, pb)
}
