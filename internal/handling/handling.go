// Copyright 2026 Ivan Guerreschi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package handling provides higher-level business logic for managing
// links entries. It sits between the CLI layer and the storage layer,
// offering operations such as listing, creating, deleting, and searching
// link entries.
package handling

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/nullzeiger/linkliste/internal/storage"
	"github.com/nullzeiger/linkliste/internal/types"
)

// Link is an alias to types.Link for convenience within this package.
type Link = types.Link

// SearchResult represents a single search result containing the index and
// the corresponding Link entry.
type SearchResult struct {
	Index int
	Link  Link
}

// All retrieves all stored accounts and returns them formatted as strings,
// each containing index and field details. It is used primarily by the CLI
// when listing entries.
func All() ([]string, error) {
	links, err := storage.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read storage: %w", err)
	}

	entries := make([]string, 0, len(links))
	for i, link := range links {
		// Format each link as a readable CLI entry.
		entries = append(entries,
			fmt.Sprintf("[%d] %s", i+1, link.String()))
	}
	return entries, nil
}

// Create appends a new link entry to the storage file.
// It performs no validation—validation should be done at the CLI or higher layer.
func Create(link Link) error {
	if link.Description == "" || link.Name == "" || link.Link == "" {
		return errors.New("invalid link data: missing required fields")
	}
	return storage.Append(link)
}

// Create appends a new link entry to the storage file.
// It performs no validation—validation should be done at the CLI or higher layer.
func Delete(index int) (bool, error) {
	links, err := storage.Read()
	if err != nil {
		return false, err
	}

	// Validate index bounds
	if index < 0 || index >= len(links) {
		return false, fmt.Errorf("index %d out of range", index)
	}

	// Remove the entry using slice manipulation.
	links = slices.Delete(links, index, index+1)

	return true, storage.Write(links)
}

// Search scans all stored links and returns those matching the given
// keyword (case-insensitive). It compares the keyword with the name,
// description, and link fields.
//
// The result is a slice of structs containing both the index of the match
// and a copy of the corresponding account.
func Search(key string) ([]SearchResult, error) {
	links, err := storage.Read()
	if err != nil {
		return nil, err
	}

	key = strings.ToLower(key)
	results := make([]SearchResult, 0)

	// Match links based on any field containing the keyword.
	for i, link := range links {
		if strings.Contains(strings.ToLower(link.Description), key) ||
			strings.Contains(strings.ToLower(link.Name), key) ||
			strings.Contains(strings.ToLower(link.Link), key) {

			results = append(results, SearchResult{Index: i, Link: link})
		}
	}

	return results, nil
}
