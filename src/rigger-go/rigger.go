package main

import (
  "fmt"

  docopt "github.com/docopt/docopt-go"
)

func main() {
  usage := `rigger is a tool to deploy Deis on a variety of cloud providers.

Usage: 
  rigger <command>

Subcommands:

  configure         configure a Deis environment
`

  args, _ := docopt.Parse(usage, nil, false, "0.0.1", false)
  command := args["<command>"]

  switch command {
  case "configure":
    fmt.Println("You're now configuring!")
  case "checkout":
  default:
    fmt.Println(args)
  }
}
