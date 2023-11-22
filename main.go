package main

import (
	"github.com/THePhanT00M/Coin/cli"
	"github.com/THePhanT00M/Coin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
