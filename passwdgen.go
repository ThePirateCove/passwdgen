package main

import (
  "fmt"
  "log"
  "os"
  "math/rand"
  "github.com/urfave/cli/v2"
  "time"
  //"strings"
  "strconv"
)

func generatePassword() string {
  r := rand.New(rand.NewSource(time.Now().UnixNano()))

  charList := []string{
    "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
    "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
    "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
    "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
    "0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
  }

  optionalChars := []string{
    "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "=", "+",
    "[", "]", "{", "}", "|", ";", ":", "'", "\"", ",", ".", "<", ">", "?", "~",
  }

  password := ""
  for i := 0; i < 12; i++ {
    // randomly choose whether to use charList or optionalChars
    var chosenChars []string
    if r.Intn(2) == 0 {
      chosenChars = charList
    } else {
      chosenChars = optionalChars
    }

    // randomly select a character from chosenChars
    randomChar := chosenChars[r.Intn(len(chosenChars))]
    password += randomChar
  }

  return password
}

func main() {
  app := &cli.App{
    EnableBashCompletion: true,
    Commands: []*cli.Command{
      {
        Name:    "genpasswd",
        Aliases: []string{"gp"},
        Usage:   "Generates passwords a specified amount of times | <1-10>",
        Flags: []cli.Flag{
          &cli.IntFlag{
            Name:  "times, t",
            Value: 0,
            Usage: "Number of passwords to generate",
          },
        },
        Action: func(cCtx *cli.Context) error {
          timesStr := cCtx.String("times")
          times, err := strconv.Atoi(timesStr)
          if err != nil {
            return err
          }

          for i := 0; i < times; i++ {
            fmt.Println(generatePassword())
          }
          return nil
        },
      },
      {
        Name:    "template",
        Aliases: []string{"t"},
        Usage:   "options for task templates",
        Subcommands: []*cli.Command{
          {
            Name:  "add",
            Usage: "add a new template",
            Action: func(cCtx *cli.Context) error {
              fmt.Println("new task template: ", cCtx.Args().First())
              return nil
            },
          },
          {
            Name:  "remove",
            Usage: "remove an existing template",
            Action: func(cCtx *cli.Context) error {
              fmt.Println("removed task template: ", cCtx.Args().First())
              return nil
            },
          },
        },
      },
    },
  }

  if err := app.Run(os.Args); err != nil {
    log.Fatal(err)
  }
}
