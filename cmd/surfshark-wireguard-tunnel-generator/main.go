package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
)

type Server struct {
	Name string `json:"connectionName"`
	Key  string `json:"pubKey"`
}

func main() {
	privateKey_arg := flag.String("privateKey", "", "the privateKey to use")
	presharedKey_arg := flag.String("presharedKey", "", "the presharedKey to use")
	dnsservers_arg := flag.String("DNS", "162.252.172.57, 149.154.159.92", "a commad-delimited list of dns servers to use")
	output_arg := flag.String("output", "wgs", "Output directory")

	flag.Parse()

	var privateKey string
	if len(*privateKey_arg) == 0 {
		prompt := promptui.Prompt{
			Label: "Private key",
			Validate: func(input string) (err error) {
				if len(input) == 0 {
					return errors.New("no private key provided")
				}

				_, err = base64.StdEncoding.DecodeString(input)
				return
			},
			Mask:        '*',
			HideEntered: true,
		}

		var err error
		privateKey, err = prompt.Run()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		privateKey = *privateKey_arg
	}

	var presharedKey string
	if len(*presharedKey_arg) != 0 {
		presharedKey = "\nPresharedKey = " + *presharedKey_arg
	} else {
		presharedKey = ""
	}

	template :=
		`[Interface]
Address = 10.14.0.2/16
PrivateKey = %s
DNS = %s
[Peer]
PublicKey = %s%s
AllowedIPs = 0.0.0.0/0
Endpoint = %s:51820`

	var outputDirectory string
	if len(*output_arg) == 0 {
		prompt := promptui.Prompt{
			Label: "Output directory",
			Default: func() (dir string) {
				dir, _ = os.Getwd()
				return
			}(),
			AllowEdit: true,
			Validate: func(input string) (err error) {
				_, err = os.Stat(input)
				return
			},
		}

		var err error
		outputDirectory, err = prompt.Run()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		outputDirectory = *output_arg
	}

	if _, err := os.Stat(outputDirectory); os.IsNotExist(err) {
		err = os.MkdirAll(outputDirectory, 0755)
		if err != nil {
			log.Fatalf("Failed to create output directory: %v", err)
		}
	}

	for _, path := range [...]string{"generic", "double", "static", "obfuscated"} {
		resp, err := http.Get("https://api.surfshark.com/v4/server/clusters/" + path)
		if err != nil {
			log.Fatal(err)
		}
		defer func(resp *http.Response) {
			resp.Body.Close()
		}(resp)

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		connections := []Server{}
		err = json.Unmarshal(body, &connections)
		if err != nil {
			log.Fatal(err)
		}

		for _, connection := range connections {
			if len(connection.Key) == 0 {
				continue
			}

			filePrefix, _ := strings.CutSuffix(connection.Name, ".prod.surfshark.com")
			file, err := os.Create(filepath.Join(outputDirectory, filePrefix+".conf"))
			if err != nil {
				log.Fatal(err)
			}
			defer func(file *os.File) {
				file.Close()
			}(file)

			_, err = file.WriteString(fmt.Sprintf(template, privateKey, *dnsservers_arg, connection.Key, presharedKey, connection.Name))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
