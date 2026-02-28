// Copyright 2026 Ivan Guerreschi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage_test

import (
	"os"
	"testing"
	"time"

	"github.com/nullzeiger/linkliste/internal/storage"
	"github.com/nullzeiger/linkliste/internal/types"
	"github.com/nullzeiger/linkliste/internal/util"
)

func setupTempStorage(t *testing.T) string {
	t.Helper()

	tempHome := t.TempDir()
	t.Setenv("HOME", tempHome)
	t.Setenv("USERPROFILE", tempHome)

	path := util.FilePath()

	if err := storage.Create(); err != nil {
		t.Fatalf("storage.Create() failed: %v", err)
	}

	return path
}

func TestCreate(t *testing.T) {
	path := setupTempStorage(t)

	if !util.FileExists(path) {
		t.Fatalf("File %s should exist after Create()", path)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if string(data) != "[]" {
		t.Fatalf("File content = %s; want []", string(data))
	}
}

func TestWriteAndRead(t *testing.T) {
	setupTempStorage(t)

	links := []types.Link{
		{Date: time.Now(), Description: "Search Engine", Link: "google.com", Name: "Google"},
		{Date: time.Now(), Description: "Go language", Link: "go.dev", Name: "Go"},
	}

	if err := storage.Write(links); err != nil {
		t.Fatalf("Write() failed: %v", err)
	}

	readlinks, err := storage.Read()
	if err != nil {
		t.Fatalf("Read() failed: %v", err)
	}

	if len(readlinks) != 2 {
		t.Fatalf("Read returned %d links; want 2", len(readlinks))
	}

	if readlinks[0].Name != "Google" {
		t.Fatalf("Link Name = %s; want Googlw", readlinks[0].Name)
	}
}

func TestAppend(t *testing.T) {
	setupTempStorage(t)

	link1 := types.Link{Date: time.Now(), Description: "Go language", Link: "go.dev", Name: "Go"}

	link2 := types.Link{Date: time.Now(), Description: "Search Engine", Link: "google.com", Name: "Google"}

	if err := storage.Append(link1); err != nil {
		t.Fatalf("Append() failed: %v", err)
	}

	if err := storage.Append(link2); err != nil {
		t.Fatalf("Append() failed: %v", err)
	}

	links, err := storage.Read()
	if err != nil {
		t.Fatalf("Read() failed: %v", err)
	}

	if len(links) != 2 {
		t.Fatalf("Read returned %d links; want 2", len(links))
	}

	if links[0].Name != "Go" || links[1].Name != "Google" {
		t.Fatalf("links data mismatch: %v", links)
	}
}

func TestAppendWithoutCreate(t *testing.T) {
	t.Helper()

	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	link := types.Link{}

	err := storage.Append(link)
	if err == nil {
		t.Fatalf("Append() should fail if storage file does not exist")
	}
}
