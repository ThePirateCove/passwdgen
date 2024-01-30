package main

import (
  "fmt"
  "log"
  "os"
  "math/rand"
  "github.com/urfave/cli/v2"
  "time"
  "strings"
)

func contains(slice []string, s string) bool {
  for _, value := range slice {
    if value == s {
      return true
    }
  }
  return false
}

func removeChars(slice []string, removeChars string) []string {
  var result []string
  for _, str := range slice {
    for _, char := range removeChars {
      str = strings.ReplaceAll(str, string(char), "")
    }
    result = append(result, str)
  }
  return result
}

func generatePassword(length int, opt bool, bl string) string {
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

	for _, char := range bl {
    if opt == false && contains(charList, string(char)) {
			charList = removeChars(charList, string(char))
    } else if opt == true && contains(optionalChars, string(char)) {
			optionalChars = removeChars(optionalChars, string(char))
    } else {
			fmt.Println("\n=+PasswdGen Runtime Error: BOOLEAN ERR...+=\n")
    }
	}

	password := ""
	for i := 0; i < length; i++ {
    var chosenChars []string
    if opt == false {
      chosenChars = charList
    } else {
      chosenChars = append(charList, optionalChars...)
		}
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
        Usage:   "Generate password(s) a specified amount of times with options | <1-10>",
        Flags: []cli.Flag{
          &cli.IntFlag{
            Name:  "length, l",
            Value: 12,
            Usage: "Length of the password",
          },
          &cli.IntFlag{
            Name:  "times, t",
            Value: 1,
            Usage: "Number of passwords to generate",
          },
          &cli.BoolFlag{
            Name:  "optional, opt",
            Value: true,
            Usage: "Enable/Disable optional characters",
          },
          &cli.StringFlag{
            Name: "blacklist, bl",
            Value: "",
            Usage: "Characters you want blacklisted | <Takes in a String>",
          },
        },
        Action: func(cCtx *cli.Context) error {
          length := cCtx.Int("length")
          times := cCtx.Int("times")
          opt := cCtx.Bool("optional")
          blacklist := strings.TrimSuffix(cCtx.String("blacklist"), "\n")

          for i := 0; i < times; i++ {
            fmt.Println(generatePassword(length, opt, blacklist))
          }
          return nil
        },
      },
    },
  }

  if err := app.Run(os.Args); err != nil {
    log.Fatal(err)
  }
}
