package main

import (
  "github.com/docopt/docopt-go"
  "bytes"
  "log"

  "rigger/commands"
)

func main() {
  usage := `usage: rigger [--version] [--help] <command> [<args>...]
options:
   -h, --help

The most commonly used rigger commands are:
   configure      prepare a Deis provision/deployment

See 'rigger help <command>' for more information on a specific command.
`
  args, _ := docopt.Parse(usage, nil, true, "rigger version 0.1.0", true)

  var buf bytes.Buffer
  logger := log.New(&buf, "logger: ", log.Lshortfile)

  logger.Println("global arguments:")
  logger.Println(args)

  logger.Println("command arguments:")
  cmd := args["<command>"].(string)
  cmdArgs := args["<args>"].([]string)

  err := runCommand(cmd, cmdArgs, logger)
  if err != nil {
    logger.Fatalf("%s is not a rigger command. See 'rigger help'", cmd)
  }
}

func runCommand(cmd string, args []string, logger *log.Logger) (err error) {
  argv := make([]string, 1)
  argv[0] = cmd
  argv = append(argv, args...)
  switch cmd {
  case "configure":
    return commands.Configure(argv, logger)
  }

  return nil
}
