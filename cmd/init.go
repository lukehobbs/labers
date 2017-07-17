// Copyright © 2017 Luke Hobbs <lukeehobbs@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

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
