package common

import (
	"bytes"
	"testing"
)

func TestWriteLog(t *testing.T) {

	var b bytes.Buffer
	b.WriteString("TTT")
	WriteLog("test", &b)
}
