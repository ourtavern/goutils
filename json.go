package goutils

import (
	"embed"
	"encoding/json"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var Json = make(map[string]map[string]interface{})

func JsonLoader(folder string, embeds ...embed.FS) {
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
			filename := strings.TrimSuffix(info.Name(), ".json")
			if err := loadJsonFile(path, filename); err != nil {
				LogError("Error loading JSON file: " + path + " - " + err.Error())
			}
		}
		return nil
	})
	if err != nil {
		LogWarn("Warning: Error walking the path " + folder + ": " + err.Error())
	}

	if len(embeds) > 0 {
		embedFs := embeds[0]
		err := fs.WalkDir(embedFs, folder, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() && strings.HasSuffix(d.Name(), ".json") {
				filename := strings.TrimSuffix(filepath.Base(path), ".json")
				f, err := embedFs.Open(path)
				if err != nil {
					LogError("Error opening embedded file: " + path + " - " + err.Error())
					return nil
				}
				defer f.Close()
				if err := loadJsonReader(f, filename); err != nil {
					LogError("Error loading embedded JSON file: " + path + " - " + err.Error())
				}
			}
			return nil
		})
		if err != nil {
			LogError("Error walking embedded filesystem: " + err.Error())
		}
	}
}

func loadJsonFile(path string, filename string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return loadJsonReader(file, filename)
}

func loadJsonReader(reader io.Reader, filename string) error {
	decoder := json.NewDecoder(reader)
	var jsonData map[string]interface{}
	if err := decoder.Decode(&jsonData); err != nil {
		return err
	}
	if Json[filename] == nil {
		Json[filename] = make(map[string]interface{})
	}
	for k, v := range jsonData {
		Json[filename][k] = v
	}
	return nil
}
