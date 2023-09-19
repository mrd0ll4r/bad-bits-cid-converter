package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/ipfs/go-cid"
	"github.com/minio/sha256-simd"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		line := reader.Text()
		encoded, err := processLine(line)
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "%s: %v\n", line, err)
			if err != nil {
				panic(err)
			}
			continue
		}

		fmt.Println(*encoded)
	}
}

func processLine(line string) (*string, error) {
	// Format the line into an ipfs:// URI
	if strings.HasPrefix(line, "/ipfs/") {
		line = line[len("/ipfs/"):]
	}

	if !strings.HasPrefix(line, "ipfs://") {
		line = "ipfs://" + line
	}

	u, err := url.Parse(line)
	if err != nil {
		return nil, fmt.Errorf("unable to build IPFS URI: %w", err)
	}

	// Extract necessary parts of the URI.
	c := u.Host
	path := u.EscapedPath()

	cc, err := cid.Parse(c)
	if err != nil {
		return nil, fmt.Errorf("invalid CID: %w", err)
	}

	// Construct a v1 CID.
	v1Cid := cid.NewCidV1(cc.Type(), cc.Hash())

	// Append / or /<path>, if present.
	var out string
	if len(path) > 0 {
		out = fmt.Sprintf("%s/%s", v1Cid, path)
	} else {
		out = fmt.Sprintf("%s/", v1Cid)
	}

	// Hash and base16 encode the result.
	hasher := sha256.New()
	hasher.Write([]byte(out))
	encoded := hex.EncodeToString(hasher.Sum(nil))

	return &encoded, nil
}
