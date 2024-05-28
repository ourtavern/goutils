package goutils

import (
	"bufio"
	"embed"
	"io"
	"os"
	"strings"
)

var Env = make(map[string]string)

func EnvLoader(embeds ...embed.FS) {
	var reader io.Reader
	var err error

	file, err := os.Open(".env")
	if err != nil {
		LogWarn("Warning: .env file not found. Will read the embedded go.env file.")
		if len(embeds) > 0 {
			f, err := embeds[0].Open("go.env")
			if err != nil {
				LogError("Error opening embedded go.env file: " + err.Error())
				return
			}
			defer f.Close()
			reader = f
		} else {
			LogError("No embedded file system provided and .env file not found")
			return
		}
	} else {
		defer file.Close()
		reader = file
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "=") {
			pair := strings.SplitN(line, "=", 2)
			if len(pair) == 2 {
				Env[pair[0]] = pair[1]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		LogError("Error reading .env file: " + err.Error())
	}

	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if len(pair) == 2 {
			Env[pair[0]] = pair[1]
		}
	}
}
