package main

import (
  "os"
  "log"

  "github.com/hashicorp/logutils"

  "github.com/deis/rigger/v2/cmd"
)

func main() {
  filter := &logutils.LevelFilter{
    Levels: []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
    MinLevel: logutils.LogLevel("INFO"),
    Writer: os.Stderr,
  }
  log.SetOutput(filter)

  cmd.Execute()
}
