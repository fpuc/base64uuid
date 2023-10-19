package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/google/uuid"
)

func main() {
	var id uuid.UUID

	if len(os.Args) > 1 {
		if os.Args[1] == "-decode" || os.Args[1] == "-d" {
			if len(os.Args) < 3 {
				fmt.Println("Missing argument")

				os.Exit(1)
			}

			decode(os.Args[2])

			return
		}

		id = uuid.MustParse(os.Args[1])
	} else {
		id = uuid.New()
	}

	b64id := base64.StdEncoding.EncodeToString(id[:])

	fmt.Println(b64id)
}

func decode(data string) {
	decoded, err := base64.StdEncoding.DecodeString(data)

	if err != nil {
		fmt.Println("Invalid base64 string")

		os.Exit(1)
	}

	id, err := uuid.FromBytes(decoded)

	if err != nil {
		fmt.Println("Invalid UUID")

		os.Exit(1)
	}

	fmt.Println(id.String())
}
