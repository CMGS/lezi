package main

import (
	"fmt"
)

func Publishv2(input chan []*FileEvent, registrar chan []*FileEvent, config *NetworkConfig) {
	for events := range input {
		for _, event := range events {
			fmt.Println(*event.Text, *event.Fields)
		}
	}
}
