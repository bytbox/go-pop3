package pop3

import (
	"bufio"
	"bytes"
	"io"
	"net"
	"strings"
	"testing"
	"time"
)

type fakeAddr struct {}
func (fakeAddr) Network() string { return "" }
func (fakeAddr) String() string { return "" }

type faker struct {
	io.ReadWriter
}

func (f faker) Close() error {
	return nil
}

func (f faker) LocalAddr() net.Addr {
	return fakeAddr{}
}

func (f faker) RemoteAddr() net.Addr {
	return fakeAddr{}
}

func (f faker) SetDeadline(t time.Time) error {
	return nil
}

func (f faker) SetReadDeadline(t time.Time) error {
	return nil
}

func (f faker) SetWriteDeadline(t time.Time) error {
	return nil
}

func TestBasic (t *testing.T) {
	basicServer := strings.Join(strings.Split(basicServer, "\n"), "\r\n")
	basicClient := strings.Join(strings.Split(basicClient, "\n"), "\r\n")

	var cmdbuf bytes.Buffer
	bcmdbuf := bufio.NewWriter(&cmdbuf)
	var fake faker
	fake.ReadWriter = bufio.NewReadWriter(bufio.NewReader(strings.NewReader(basicServer)), bcmdbuf)

	_, err := NewClient(fake)
	if err != nil {
		t.Fatalf("NewClient failed: %s", err)
	}

	bcmdbuf.Flush()
	if basicClient != cmdbuf.String() {
		t.Fatalf("Got:\n%s\nExpected:\n%s", cmdbuf.String(), basicClient)
	}
}

var basicServer = `+OK good morning
`

var basicClient = ``
