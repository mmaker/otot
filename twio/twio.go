package twio

import (
	"fmt"
	"io"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/ChimeraCoder/anaconda"
)


const PREFIX = "@%s "
const TIMEOUT = 20
const RPOLL = 5 * time.Second
const WPOLL = 10 * time.Second

type reader struct {
	t *anaconda.TwitterApi
	prefix string
	sleep int
	sinceId string
	buf []byte
	attempts int
}

type writer struct {
	t *anaconda.TwitterApi
	u string
}


func (r *reader) SeekEnd() error {
	v := url.Values{}
	v.Set("count", "1")
	mentions, err := r.t.GetMentionsTimeline(v)

	if len(mentions) > 0 {
		r.sinceId = mentions[0].IdStr
	}
	return err
}

func (r *reader) Read(p []byte) (n int, err error) {
	log.Println("Now Trying to read…")

	if len(r.buf) > 0 {
		n = copy(p, r.buf)
		if n < len(r.buf) {
			r.buf = r.buf[n:]
			return n, err
		}
	}

	if r.attempts == TIMEOUT {
		return 0, io.EOF
	} else {
		r.attempts++
		time.Sleep(RPOLL)
	}

	v := url.Values{}
	v.Set("count", "30")
	v.Set("since_id", r.sinceId)
	mentions, err := r.t.GetMentionsTimeline(v)
	if err != nil {
		return n, err
	}

	if len(mentions) > 0 {
		r.attempts = 0
	}

	var written int
	var msg []byte
	for i := len(mentions) - 1; i >= 0; i-- {
		text := mentions[i].Text
		log.Println(text)

		if !strings.HasPrefix(text, r.prefix) {
			continue
		}

		msg = []byte(text[len(r.prefix):])
		written = copy(p, msg)
		n += written
		p = p[written:]

		if written < len(msg) {
			r.buf = msg[written:]
			break
		}
		r.sinceId = mentions[i].IdStr
	}
	return n, err
}


func (w *writer) Write(p []byte) (i int, err error) {
	v := url.Values{}
	log.Println("Now Writing…")

	msg := string(p)
	prefix := fmt.Sprintf(PREFIX, w.u)
	tweetLen := 140 - len(prefix)
	for i = 0; i < len(msg); time.Sleep(WPOLL)  {
		if i + tweetLen  >= len(msg) {
			_, err = w.t.PostTweet(prefix + msg[i:], v)
			if err != nil {
				break
			}
			i = len(msg)
		} else {
			_, err = w.t.PostTweet(prefix + msg[i:i+tweetLen], v)
			if err != nil {
				break
			}
			i += tweetLen
		}
	}
	return
}


func NewTwitterReader(api *anaconda.TwitterApi) io.Reader {
	v := url.Values{}
	self, err := api.GetSelf(v)
	if err != nil {
		log.Fatal(err)
	}

	r := reader{
		t: api,
		prefix: fmt.Sprintf(PREFIX, self.ScreenName),
		sleep: 15,
		buf: []byte(""),
	}
	r.SeekEnd()
	return &r

}

func NewTwitterWriter(api *anaconda.TwitterApi, username string) io.Writer {
	return &writer{
		t: api,
		u: username,
	}
}
