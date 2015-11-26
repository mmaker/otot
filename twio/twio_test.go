package twio

import (
	"crypto/rand"
	"net/url"
	"testing"
	"flag"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/stretchr/testify/assert"

	"github.com/mmaker/otot/twutils"
)

var sender, receiver *anaconda.TwitterApi

func TestWhoAmI(t *testing.T) {
	v := url.Values{}
	su, err := sender.GetSelf(v)
	assert.Nil(t, err)
	assert.Equal(t, su.ScreenName, "otsender")
	ru, err := receiver.GetSelf(v)
	assert.Nil(t, err)
	assert.Equal(t, ru.ScreenName, "otreceiver")
}


func TestWriterAndReader(t *testing.T) {
	w := NewTwitterWriter(sender, "otreceiver")
	r := NewTwitterReader(receiver)

	z, _ := rand.Prime(rand.Reader, 10)
	msg := []byte("Hello! This is a random prime: " + z.String())
	_, err := w.Write(msg)
	assert.Nil(t, err)
	// assert.Len(t, msg, n)

	got := make([]byte, len(msg))
	_, err = r.Read(got)
	assert.Nil(t, err)
	assert.NotNil(t, got)
	// assert.Len(t, msg, n)
	assert.Equal(t, msg, got)
}


func TestMain(m *testing.M) {
	flag.Parse()
	sender, _ = twutils.GetApi("/home/maker/dev/go/src/github.com/mmaker/otot/sender")
	receiver, _ = twutils.GetApi("/home/maker/dev/go/src/github.com/mmaker/otot/receiver")
	os.Exit(m.Run())
}
