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

	c, err := NewClient(fake)
	if err != nil {
		t.Fatalf("NewClient failed: %s", err)
	}

	if err = c.User("uname"); err != nil {
		t.Fatal("User failed: %s", err)
	}

	if err = c.Pass("password1"); err == nil {
		t.Fatal("Pass succeeded inappropriately")
	}

	if err = c.Auth("uname", "password2"); err != nil {
		t.Fatal("Auth failed: %s", err)
	}

	if err = c.Noop(); err != nil {
		t.Fatal("Noop failed: %s", err)
	}

	bcmdbuf.Flush()
	if basicClient != cmdbuf.String() {
		t.Fatalf("Got:\n%s\nExpected:\n%s", cmdbuf.String(), basicClient)
	}
}

var basicServer = `+OK good morning
+OK send PASS
-ERR [AUTH] mismatched username and password
+OK send PASS
+OK welcome
+OK
`

var basicClient = `USER uname
PASS password1
USER uname
PASS password2
NOOP
`
