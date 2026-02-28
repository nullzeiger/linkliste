// Copyright 2026 Ivan Guerreschi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

// Update the state of TUI.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		leftWidth := int(float64(m.Width) * 0.4)
		if leftWidth < 20 {
			leftWidth = m.Width / 2
		}
		m.L.SetSize(leftWidth, m.Height-2)

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
