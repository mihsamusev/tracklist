package main

import (
	"fmt"
	"os"
	"strings"
)


type Args struct {
	filename string
	whitelist []string
}

func addToWhitelist(whitelist []string, content string) []string {
	separator := ","
    parts := strings.Split(content, separator)

    for _, part := range parts {
        whitelist = append(whitelist, part)
    }
	return whitelist
}

func ParseArgs() (Args, error) {
	if len(os.Args) < 2 {
        return Args{}, fmt.Errorf("Please provide a .cue filename to parse\n")
    }

    filename := os.Args[1]
    _, err := os.Stat(filename)
    if os.IsNotExist(err) {
        fmt.Println("Track file {} does not exist")
        return Args{}, fmt.Errorf("invalid input: track file '%s' does not exist\n", filename)
    }


	whitelist := make([]string, 0)
	if len(os.Args) == 3 {
		whitelist = addToWhitelist(whitelist, os.Args[2])
	}
	
	return Args{filename: filename, whitelist: whitelist}, nil
}
