// Copyright 2026 Ivan Guerreschi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nullzeiger/linkliste/internal/storage"
	"github.com/nullzeiger/linkliste/internal/types"
)

type ListItem struct {
	Link  types.Link
	Index int
}

func (i ListItem) Title() string {
	return fmt.Sprintf("%2d. %s", i.Index+1, i.Link.Name)
}

func (i ListItem) Description() string {
	return i.Link.Link
}

func (i ListItem) FilterValue() string {
	return i.Link.Name
}

type LinksLoadedMsg []types.Link
type ErrMsg error

type Model struct {
	L       list.Model
	Links   []types.Link
	Spinner spinner.Model
	Width   int
	Height  int
	Loading bool
	Err     error
}

func InitialModel() Model {
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.FilterInput.Prompt = " Search: "
	l.FilterInput.Cursor.Style = CursorStyle
	l.SetShowTitle(false)

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = SpinnerStyle

	return Model{
		L:       l,
		Spinner: s,
		Loading: true,
	}
}

func FetchLinksCmd() tea.Msg {
	links, err := storage.Read()
	if err != nil {
		return ErrMsg(err)
	}
	return LinksLoadedMsg(links)
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.Spinner.Tick, FetchLinksCmd)
}
