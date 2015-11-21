package twio

import (
	"crypto/rand"
	"bytes"
	"net/url"
	"testing"
	"flag"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/mmaker/otot/twutils"
)

var sender, receiver *anaconda.TwitterApi

func TestWhoAmI(t *testing.T) {
	v := url.Values{}
	su, err := sender.GetSelf(v)
	if err != nil || su.ScreenName != "otsender" {
		t.Errorf("error: %s (username '%s')", err, su.ScreenName)
	}

	ru, err := receiver.GetSelf(v)
	if ru.ScreenName != "otreceiver" {
		t.Errorf("error: %s(username '%s')", err, ru.ScreenName)
	}
}


func TestWriterAndReader(t *testing.T) {
	w := NewTwitterWriter(sender, "otreceiver")
	r := NewTwitterReader(receiver)

	z, _ := rand.Prime(rand.Reader, 1024)
	msg := []byte(z.String())
	n, err := w.Write(msg)
	if n != len(msg) || err != nil {
		t.Errorf("Error writing: %s (wrote %s)", err, n)
	}

	got := make([]byte, len(msg))
	n, err = r.Read(got)
	if n != len(msg)  || err != nil {
		t.Errorf("Error reading: %s (read %d)", err, n)
	}

	if bytes.Compare(msg, got) != 0 {
		t.Errorf("'%s' != '%s'", msg, got)
	}


}


func TestMain(m *testing.M) {
	flag.Parse()
	sender, _ = twutils.GetApi("/home/maker/dev/go/src/github.com/mmaker/otot/sender")
	receiver, _ = twutils.GetApi("/home/maker/dev/go/src/github.com/mmaker/otot/receiver")
	os.Exit(m.Run())
}
