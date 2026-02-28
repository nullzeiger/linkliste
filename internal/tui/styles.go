// Copyright 2026 Ivan Guerreschi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tui

import "github.com/charmbracelet/lipgloss"

// Style fot TUI.
var (
	TitleStyle   = lipgloss.NewStyle().Bold(true).Render
	LabelStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#7D7D7D")).Render
	DetailStyle  = lipgloss.NewStyle().Padding(1).Border(lipgloss.NormalBorder())
	HeaderStyle  = lipgloss.NewStyle().Bold(true).Render
	CursorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff5faf"))
	SpinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff5faf"))
)

// Row return style for row.
func Row(left, right string) string {
	return lipgloss.JoinHorizontal(lipgloss.Top, left, "    ", right)
}
