package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/ghodss/yaml"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Configure environment file",
	Run: func(cmd *cobra.Command, args []string) {
		tok := ""
		url := ""
		masked_tok := ""

		dotEnv, err := godotenv.Read("labers.env")
		if err == nil {
			tok = dotEnv["GITHUB_TOKEN"]
			if len(tok) > 3 {
				masked_tok = fmt.Sprint(strings.Repeat("*", utf8.RuneCountInString(tok)-4), string(tok[len(tok)-4:]))
			}
			url = dotEnv["GITHUB_URL"]
		}

		r := bufio.NewReader(os.Stdin)

		fmt.Printf("Github token [%v]: ", strings.TrimSpace(masked_tok))
		a, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Github URL [%v]: ", strings.TrimSpace(url))
		b, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if a != "\n" {
			tok = strings.TrimSpace(a)
		}
		if b != "\n" {
			url = strings.TrimSpace(b)
		}

		dotEnv = map[string]string{"GITHUB_TOKEN": tok, "GITHUB_URL": url}
		yDotEnv, err := yaml.Marshal(&dotEnv)
		err = ioutil.WriteFile("labers.env", yDotEnv, 0644)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
