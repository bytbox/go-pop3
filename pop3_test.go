package pop3

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"testing"
)

type faker struct {
	io.ReadWriter
}

func (f faker) Close() error {
	return nil
}

func TestBasic (t *testing.T) {
	basicServer := strings.Join(strings.Split(basicServer, "\n"), "\r\n")
	basicClient := strings.Join(strings.Split(basicClient, "\n"), "\r\n")

	var cmdbuf bytes.Buffer
	bcmdbuf := bufio.NewWriter(&cmdbuf)
	var fake faker
	fake.ReadWriter = bufio.NewReadWriter(bufio.NewReader(strings.NewReader(basicServer)), bcmdbuf)
	bcmdbuf.Flush()
	if basicClient != cmdbuf.String() {
		t.Fatalf("Got:\n%s\nExpected:\n%s", cmdbuf.String(), basicClient)
	}
}

var basicServer = ``

var basicClient = ``
