package main

import (
	"fmt"

	"github.com/JunkieJI/travel-senpai/cmd"
	"github.com/JunkieJI/travel-senpai/config"
	"github.com/JunkieJI/travel-senpai/store"
)

func main() {
	cmd.Execute()

	config.Init()
	s := store.NewStore(config.GetDSN())
	fmt.Println(s.Ping())
}
