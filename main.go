package main


import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/big"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/mmaker/otot/twio"
	"github.com/mmaker/otot/proto/dh"
	"github.com/mmaker/otot/proto/ot"
	"github.com/mmaker/otot/encodings"
)


var (
	register = flag.Bool("register", false, "Get user credentials.")
	credentials = flag.String("cred", "credentials", "Credentials file.")
	start = flag.Bool("start", false, "Is initiating the protocol?")
	partner = flag.String("with", "", "Partner")
	proto_dh = flag.Bool("dh", false, "Perform a DH key exchange.")
	proto_ot = flag.Bool("ot", false, "Perform Oblivious Transfer")
	enc = flag.String("enc", "none", "Select Encoding")
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

func startDH(c *encodings.TConn, isInitiating bool) {
	var key *big.Int
	if isInitiating {
		key = dh.StartServer(c)
	} else {
		key = dh.StartClient(c)
	}
	fmt.Println("Receieved ", key)

}

func startOT(c *encodings.TConn, isSender bool) {
	if isSender {
		stdin := bufio.NewReader(os.Stdin)
		fst, _ := stdin.ReadString('\n')
		snd, _ := stdin.ReadString('\n')
		choices := []string{fst, snd}
		ot.StartSender(c, choices)
	} else {
		var choice int
		var msg string
		fmt.Scanf("%d", choice)
		msg = ot.StartReceiver(c, choice)
		fmt.Println("Got: ", msg)
	}
}

func main() {
	flag.Parse()

	if *register {
		GetTokens()
		os.Exit(0)
	}
	api, _ := twio.GetApi(*credentials)

	var r io.Reader = twio.NewTwitterReader(api)
//	r = arab.NewDecoder(r)
	var w io.Writer = twio.NewTwitterWriter(api, *partner)
//	w = arab.NewEncoder(w)
	c := encodings.NewTConn(r, w)

	if *proto_dh {
		startDH(c, *start)
	} else if *proto_ot {
		startOT(c, *start)
	}
}
