package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
)

// This utility generates a random ED25519 libp2p identity, prints its corresponding peer ID and
// stores its marshalled private key at the path specified via `-out` flag.
func main() {
	out := flag.String("out", "", "The path to which to write the marshalled libp2p identity.")
	flag.Parse()

	if *out == "" {
		panic("-out must be specified")
	}

	k, _, err := crypto.GenerateEd25519Key(nil)
	if err != nil {
		panic(err)
	}
	mk, err := crypto.MarshalPrivateKey(k)
	if err != nil {
		panic(err)
	}
	pid, err := peer.IDFromPrivateKey(k)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Genered identity with peer ID: %s\n", pid.String())
	p, err := filepath.Abs(*out)
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(p, mk, 0666); err != nil {
		panic(err)
	}
	fmt.Printf("Stored marshalled identity at: %s\n", p)
}
