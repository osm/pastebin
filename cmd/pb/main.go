package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/osm/pastebin"
)

func main() {
	apiKey := os.Getenv("PB_API_KEY")
	if apiKey == "" {
		fmt.Fprintf(os.Stderr, "error: set PB_API_KEY in your environment\n")
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	stdin, _ := reader.ReadString('\n')
	code := strings.TrimSuffix(stdin, "\n")

	pb := pastebin.New(apiKey)
	if url, err := pb.NewPaste(code, "", pastebin.Unlisted, pastebin.TenMinutes); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	} else {
		fmt.Println(url)
	}
}
