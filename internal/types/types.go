// Copyright 2026 Ivan Guerreschi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"fmt"
	"time"
)

// Link represents a link with metadata
type Link struct {
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Name        string    `json:"name"`
}

// String implements fmt.Stringer interface
func (l Link) String() string {
	formattedDate := l.Date.Format("02-01-2006 15:04:05")
	return fmt.Sprintf("Date: %s, Description: %s, Link: %s, Name: %s",
		formattedDate, l.Description, l.Link, l.Name)
}
