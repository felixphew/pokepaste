package main

import (
	"bytes"
	"html/template"
	"net/http"
	"regexp"
	"strings"
)

type paste struct {
	Paste  []set
	Title  string
	Author string
	Notes  string
}

type set struct {
	Pokemon uint
	Form    uint
	Item    uint
	Text    template.HTML
}

var (
	reHead = regexp.MustCompile(`^(?:(.* \()([A-Z][a-z0-9:']+\.?(?:[- ][A-Za-z][a-z0-9:']*\.?)*)(\))|([A-Z][a-z0-9:']+\.?(?:[- ][A-Za-z][a-z0-9:']*\.?)*))(?:( \()([MF])(\)))?(?:( @ )([A-Z][a-z0-9:']+(?:[- ][A-Z][a-z0-9:']*)*))?( *)$`)
	reMove = regexp.MustCompile(`^(-)( *([A-Z][a-z\']*(?:[- ][A-Za-z][a-z\']*)*)(?: */ *[A-Z][a-z\']*(?:[- ][A-Za-z][a-z\']*)*)* *)$`)
	reStat = regexp.MustCompile(`^(\d+ HP)?( / )?(\d+ Atk)?( / )?(\d+ Def)?( / )?(\d+ SpA)?( / )?(\d+ SpD)?( / )?(\d+ Spe)?( *)$`)

	tmpl = template.Must(template.ParseFiles("paste.tmpl"))
)

func renderPaste(w http.ResponseWriter, text, title, author, notes []byte) {
	sets := bytes.Split(text, []byte("\r\n\r\n"))
	fpaste := paste{
		Paste:  make([]set, 0, len(sets)),
		Title:  string(title),
		Author: string(author),
		Notes:  string(notes),
	}

	for _, bset := range sets {
		if len(bset) == 0 {
			continue
		}

		var fset set
		var b strings.Builder

		lines := bytes.Split(bset, []byte("\r\n"))

		m := reHead.FindSubmatch(lines[0])
		if m == nil {
			template.HTMLEscape(&b, bset)
			fset.Text = template.HTML(b.String())
			fpaste.Paste = append(fpaste.Paste, fset)
			continue
		}

		if len(m[2]) != 0 {
			if p, ok := pokemonData[string(m[2])]; ok {
				fset.Pokemon = p["id"].(uint)
				fset.Form = p["form"].(uint)
				template.HTMLEscape(&b, m[1])
				b.WriteString(`<span class="type-`)
				b.WriteString(p["type"].(string))
				b.WriteString(`">`)
				template.HTMLEscape(&b, m[2])
				b.WriteString(`</span>`)
				template.HTMLEscape(&b, m[3])
			} else {
				template.HTMLEscape(&b, m[1])
				template.HTMLEscape(&b, m[2])
				template.HTMLEscape(&b, m[3])
			}
		} else if len(m[4]) != 0 {
			if p, ok := pokemonData[string(m[4])]; ok {
				fset.Pokemon = p["id"].(uint)
				fset.Form = p["form"].(uint)
				b.WriteString(`<span class="type-`)
				b.WriteString(p["type"].(string))
				b.WriteString(`">`)
				template.HTMLEscape(&b, m[4])
				b.WriteString(`</span>`)
			} else {
				template.HTMLEscape(&b, m[4])
			}
		}

		if len(m[6]) != 0 {
			template.HTMLEscape(&b, m[5])
			if m[6][0] == 'M' {
				b.WriteString(`<span class="gender-m">`)
			} else if m[6][0] == 'F' {
				b.WriteString(`<span class="gender-f">`)
			} else {
				b.WriteString(`<span>`)
			}
			template.HTMLEscape(&b, m[6])
			b.WriteString(`</span>`)
			template.HTMLEscape(&b, m[7])
		}

		if len(m[9]) != 0 {
			if i, ok := itemData[string(m[9])]; ok {
				fset.Item = i["id"].(uint)
				template.HTMLEscape(&b, m[8])
				if t, ok := i["type"]; ok {
					b.WriteString(`<span class="type-`)
					b.WriteString(t.(string))
					b.WriteString(`">`)
					template.HTMLEscape(&b, m[9])
					b.WriteString(`</span>`)
				} else {
					template.HTMLEscape(&b, m[9])
				}
			} else {
				template.HTMLEscape(&b, m[8])
				template.HTMLEscape(&b, m[9])
			}
		}

		b.Write(m[10])
		b.WriteByte('\n')

		for _, line := range lines[1:] {
			if m := reMove.FindSubmatch(line); m != nil {
				if mv, ok := moveData[string(m[3])]; ok {
					b.WriteString(`<span class="type-`)
					b.WriteString(mv["type"].(string))
					b.WriteString(`">`)
					template.HTMLEscape(&b, m[1])
					b.WriteString(`</span>`)
				} else {
					template.HTMLEscape(&b, m[1])
				}
				template.HTMLEscape(&b, m[2])
			} else if m := bytes.SplitAfterN(line, []byte(": "), 2); len(m) == 2 {
				b.WriteString(`<span class="attr">`)
				template.HTMLEscape(&b, m[0])
				b.WriteString(`</span>`)
				if len(m[0]) == 5 && m[0][1] == 'V' && m[0][2] == 's' {
					attr := m[1]
					if m := reStat.FindSubmatch(attr); m != nil {
						for i, stat := range [...]string{"hp", "atk", "def", "spa", "spd", "spe"} {
							if len(m[i * 2 + 1]) > 0 {
								b.WriteString(`<span class="stat-`)
								b.WriteString(stat)
								b.WriteString(`">`)
								template.HTMLEscape(&b, m[i * 2 + 1])
								b.WriteString(`</span>`)
							}
							template.HTMLEscape(&b, m[i * 2 + 2])
						}
					} else {
						template.HTMLEscape(&b, attr)
					}
				} else {
					template.HTMLEscape(&b, m[1])
				}
			} else {
				template.HTMLEscape(&b, line)
			}
			b.WriteByte('\n')
		}

		fset.Text = template.HTML(b.String())

		fpaste.Paste = append(fpaste.Paste, fset)
	}

	err := tmpl.Execute(w, fpaste)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
