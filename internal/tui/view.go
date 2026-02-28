// Copyright 2026 Ivan Guerreschi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tui

import (
	"fmt"
)

// View the TUI.
func (m Model) View() string {
	if m.Err != nil {
		return fmt.Sprintf("\n  Error: %v\n\n  Press q to exit.\n", m.Err)
	}

	if m.Loading {
		return fmt.Sprintf("\n %s Loading links...\n", m.Spinner.View())
	}

	leftWidth := int(float64(m.Width) * 0.4)
	if leftWidth < 20 {
		leftWidth = m.Width / 2
	}
	rightWidth := m.Width - leftWidth - 6

	left := m.L.View()

	var right string
	selectedItem := m.L.SelectedItem()

	if selectedItem != nil {
		if i, ok := selectedItem.(ListItem); ok {
			link := i.Link
			title := TitleStyle(link.Name)

			url := link.Link
			if url == "" {
				url = "<no url>"
			}

			rightBody := fmt.Sprintf("%s\n\n%s %d / %d\n\n%s %s\n\n%s %s\n\n%s %s\n\n%s %s",
				title,
				LabelStyle("Index:"), i.Index+1, len(m.Links),
				LabelStyle("Name:"), link.Name,
				LabelStyle("URL:"), url,
				LabelStyle("Description:"), link.Description,
				LabelStyle("Date:"), link.Date.Format("02-01-2006 15:04:05"),
			)
			right = DetailStyle.Width(rightWidth).Render(rightBody)
		}
	} else {
		right = DetailStyle.Width(rightWidth).Render("No links found..")
	}

	row := Row(left, right)
	header := HeaderStyle("Linkliste TUI - Press q to exit.")

	return header + "\n\n" + row
}
