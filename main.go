package main


import (
	"bufio"
//	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"math/big"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"golang.org/x/text/unicode/norm"
	"github.com/mmaker/otot/encodings/arab"
	"github.com/mmaker/otot/twio"
	"github.com/mmaker/otot/twutils"
	"github.com/mmaker/otot/proto/dh"
	"github.com/mmaker/otot/proto/ot"
)


var (
	register = flag.Bool("register", false, "Get user credentials.")
	credentials = flag.String("cred", "credentials", "Credentials file.")
	start = flag.Bool("start", false, "Is initiating the protocol?")
	partner = flag.String("with", "", "Partner")
	proto_dh = flag.Bool("dh", false, "Perform a DH key exchange.")
	proto_ot = flag.Bool("ot", false, "Perform Oblivious Transfer")

)

func GetTokens() {
	url, credentials, _ := anaconda.AuthorizationURL("oob")
	var pin string
	fmt.Printf("Please visit <" + url + "> and enter the pin: ")
	fmt.Scanf("%s", &pin)

	credentials, _, _ = anaconda.GetCredentials(credentials, pin)

	fmt.Println(credentials.Token)
	fmt.Println(credentials.Secret)
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func startDH(api *anaconda.TwitterApi, isInitiating bool, partner string) {
	var r io.Reader = twio.NewTwitterReader(api)
	r = arab.NewDecoder(r)
	r = norm.NFC.Reader(r)

	//r = base64.NewDecoder(base64.StdEncoding, r)

	var w io.Writer = twio.NewTwitterWriter(api, partner)
 	w = arab.NewEncoder(w)
	w = norm.NFC.Writer(w)
	//w = base64.NewEncoder(base64.StdEncoding, w)

	var key *big.Int
	if isInitiating {
		key = dh.StartServer(r, w)
	} else {
		key = dh.StartClient(r, w)
	}
	fmt.Println("Receieved ", key)

}

func startOT(api *anaconda.TwitterApi, isSender bool, partner string) {
	var r io.Reader = twio.NewTwitterReader(api)
	r = arab.NewDecoder(r)

	var w io.Writer = twio.NewTwitterWriter(api, partner)
	w = arab.NewEncoder(w)

	if isSender {
		r := bufio.NewReader(os.Stdin)
		fst, _ := r.ReadString('\n')
		snd, _ := r.ReadString('\n')
		choices := []string{fst, snd}
		ot.StartSender(r, w, choices)
	} else {
		var choice int
		var msg string
		fmt.Scanf("%d", choice)
		msg = ot.StartReceiver(r, w, choice)
		fmt.Println("Got: ", msg)
	}
}

func main() {
	flag.Parse()

	if *register {
		GetTokens()
		os.Exit(0)
	}
	api, _ := twutils.GetApi(*credentials)

	if *proto_dh {
		startDH(api, *start, *partner)
	} else if *proto_ot {
		startOT(api, *start, *partner)
	}
}
