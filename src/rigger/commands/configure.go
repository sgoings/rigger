package commands

import (
  "log"
  "fmt"
  // "os"

  "github.com/jacobstr/confer"
  "github.com/docopt/docopt-go"
)

func Configure(argv []string, logger *log.Logger) (err error) {
  usage := `usage: rigger configure
options:
  -h, --help
`

  args, _ := docopt.Parse(usage, nil, true, "", false)
  logger.Println(args)

  // pwd, err := os.Getwd()

  config := confer.NewConfig()


  config.ReadPaths("/home/sgoings/.rigger/Rigfile.yaml", "Rigfile.yaml")

  fmt.Printf("Platform version: %s\n", config.Get("platform.version"))
  fmt.Printf("Clients version: %s\n", config.Get("clients.git.sha1"))

  return
}