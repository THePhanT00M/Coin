package main

import (
	"github.com/THePhanT00M/Coin/explorer"
	"github.com/THePhanT00M/Coin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
