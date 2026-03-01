// Copyright 2026 Ivan Guerreschi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tui

import (
	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
)

const (
	minListWidth      = 20
	widthPercentage   = 0.4
	heightPadding     = 2
	minTerminalHeight = 5
)

func calculateListDimensions(width, height int) (int, int) {
	listWidth := int(float64(width) * widthPercentage)
	if listWidth < minListWidth {
		listWidth = width / 2
	}

	listHeight := max(height-heightPadding, minTerminalHeight)

	return listWidth, listHeight
}

// Update the state of TUI.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.Width, m.Height = msg.Width, msg.Height
		listW, listH := calculateListDimensions(m.Width, m.Height)
		m.L.SetSize(listW, listH)

	case ErrMsg:
		m.Err = msg
		m.Loading = false
		return m, nil

	case LinksLoadedMsg:
		m.Loading = false
		m.Links = msg
		items := make([]list.Item, len(msg))
		for i, link := range msg {
			items[i] = ListItem{Link: link, Index: i}
		}
		cmd = m.L.SetItems(items)
		cmds = append(cmds, cmd)

	case spinner.TickMsg:
		if m.Loading {
			var spinnerCmd tea.Cmd
			m.Spinner, spinnerCmd = m.Spinner.Update(msg)
			cmds = append(cmds, spinnerCmd)
		}
	}

	m.L, cmd = m.L.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
