// Copyright 2026 Ivan Guerreschi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package storage provides low-level functions for creating, reading,
// writing, and appending link data to the JSON storage file.
// The file path and permissions are managed via the util package.
package storage

import (
	"encoding/json"
	"os"

	"github.com/nullzeiger/linkliste/internal/types"
	"github.com/nullzeiger/linkliste/internal/util"
)

// Create initializes the storage file if it does not already exist.
// It ensures that the file path returned by util.FilePath()
// exists and contains an empty JSON array ([]).
// If the file already exists, the function does nothing.
func Create() error {
	path := util.FilePath()

	// If the storage file already exists, nothing needs to be done.
	if util.FileExists(path) {
		return nil
	}

	// Create the empty file.
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Initialize the file with an empty JSON array.
	_, err = file.Write([]byte("[]"))
	return err
}

// Read loads all stored links from the JSON file into a slice.
// Returns an error if the file cannot be read or the JSON is malformed.
func Read() ([]types.Link, error) {
	path := util.FilePath()

	// Read raw file contents.
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Decode JSON into a slice of Link structs.
	var links []types.Link
	err = json.Unmarshal(data, &links)
	return links, err
}

// Write replaces the entire storage file with the provided slice of links.
// The file is overwritten using the permissions defined in util.Perm.
func Write(links []types.Link) error {
	path := util.FilePath()

	// Encode links as pretty-printed JSON.
	jsonData, err := json.MarshalIndent(links, "", "  ")
	if err != nil {
		return err
	}

	// Overwrite the storage file with new data.
	return os.WriteFile(path, jsonData, util.Perm)
}

// Append reads the existing links from storage, adds the new link,
// and writes all links back to the file.
// Returns an error if reading or writing fails.
func Append(link types.Link) error {
	links, err := Read()
	if err != nil {
		return err
	}

	// Add the new entry.
	links = append(links, link)

	// Add the new entry.
	return Write(links)
}
