// Copyright 2013 Daniel Jo. All rights reserved.
// Use of this source code is governed by a BSD-sytle
// license that can be found in the LICENSE file.

// Package format defines an API for formatting strings for printing in a fixed
// width font. The functionality is similar to that of the Unix command fmt.
package format

const _pad = "                                             "

// Split separates a string to fit within a text area with a width defined by
// the parameter w. Words do not fit on a line will be deferred to the next. The
// parameter tw defines the width of tab characters.
// TODO: Split hyphenated words and handle words longer than line itself.
func Split(w, tw int, s string) []string {
	var (
		lines []string
		line  = make([]rune, 0, w)
		word  = make([]rune, 0, w)
	)

	for _, c := range s {
		switch c {
		case ' ', '\n', '\t':
			if len(word) > 0 {
				if len(line)+len(word) > w {
					lines = append(lines, string(line))
					line = line[:0]
				}
				line = append(line, word...)
				word = word[:0]
			}
			switch c {
			case '\n':
				lines = append(lines, string(line))
					line = line[:0]
			case '\t':
				var spaces = tw - len(line)%tw
				if len(line)+spaces < w {
					line = append(line, []rune(_pad)[:spaces]...)
				} else {
					line = append(line, []rune(_pad)[:w-len(line)]...)
				}
			default:
				if len(line) < w {
					line = append(line, c)
				}
			}
		default:
			word = append(word, c)
		}
	}

	if len(word) > 0 {
		if len(line)+len(word) > w {
			lines = append(lines, string(line))
			line = line[:0]
		}
		line = append(line, word...)
		lines = append(lines, string(line))
	}

	return lines
}

