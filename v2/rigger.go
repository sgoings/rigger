package main

import (
  "os"
  "log"

  "github.com/spf13/viper"
  "github.com/hashicorp/logutils"
  "github.com/spf13/cobra"
)


func main() {
  filter := &logutils.LevelFilter{
    Levels: []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
    MinLevel: logutils.LogLevel("INFO"),
    Writer: os.Stderr,
  }
  log.SetOutput(filter)

  viper.SetConfigName("Rigfile")
  viper.AddConfigPath("$HOME/.rigger/")
  viper.AddConfigPath(".")

  err := viper.ReadInConfig()
  if err != nil {
    if _,ok := err.(viper.UnsupportedConfigError); ok {
      log.Printf("[ERROR] No Rigfile exists.")
      os.Exit(1)
    } else {
      log.Printf("[ERROR] %s", err)
    }
  }

  var cmdUp = &cobra.Command{
    Use:   "up",
    Short: "Create my infrastructure",
    Long: `Do lots of work`,
    Run: func(cmd *cobra.Command, args []string) {
      log.Printf("[INFO] Rigger lifting!")
    },
  }

  var rootCmd = &cobra.Command{Use: "rigger"}
  rootCmd.AddCommand(cmdUp)
  rootCmd.Execute()
}
