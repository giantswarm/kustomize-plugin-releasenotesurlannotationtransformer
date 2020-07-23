package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"
)

func handleDocument(provider, annotationKey string, document []byte) {
	var releaseObject map[string]interface{}
	err := yaml.UnmarshalStrict(document, &releaseObject)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	release := unstructured.Unstructured{Object: releaseObject}

	if release.GetKind() == "Release" {
		annotations := release.GetAnnotations()
		if annotations == nil {
			annotations = map[string]string{}
		}
		name := release.GetName()

		annotations[annotationKey] = fmt.Sprintf("https://github.com/giantswarm/releases/tree/master/%s/%s", provider, name)

		release.SetAnnotations(annotations)

		r, err := yaml.Marshal(release.Object)
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
