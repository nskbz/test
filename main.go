package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func main() {
	f, _ := os.Open("vxlan.md")
	b, _ := io.ReadAll(f)
	i := base64.StdEncoding.EncodedLen(len(b))
	buf := make([]byte, i)
	base64.StdEncoding.Encode(buf, b)
	fmt.Println(string(buf))
}
