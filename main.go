package main

import (
	"fmt"
	"os"
)

func main() {
  // このコードは引数と標準出力を用いたサンプルコードです。
  // このコードは好きなように編集・削除してもらって構いません。
  // ---
  // This is a sample code to use arguments and outputs.
  // Edit and remove this code as you like.

	run(os.Args[1:])
}

func run(args []string) {
	for i, v := range args {
		fmt.Printf("argv[%d]: %s\n", i, v)
	}
}
