package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"sigs.k8s.io/yaml"
)

func handleDocument(provider, annotationKey string, document []byte) {
	var release map[string]interface{}
	err := yaml.Unmarshal(document, &release)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	kind := release["kind"]

	if kind == "Release" {
		annotations := release["metadata"].(map[string]interface{})
		name := annotations["name"]

		annotations[annotationKey] = fmt.Sprintf("https://github.com/giantswarm/releases/tree/master/%s/%s", provider, name)

		r, err := yaml.Marshal(release)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%s\n---\n", r)
	}
}

func main() {
	provider := os.Args[2]
	annotationKey := os.Args[3]

	var buf bytes.Buffer
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				buf.WriteString(line)
				handleDocument(provider, annotationKey, buf.Bytes())
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		if strings.TrimSpace(line) == "---" {
			handleDocument(provider, annotationKey, buf.Bytes())
			buf.Reset()
		} else {
			buf.WriteString(line)
		}
	}
}
