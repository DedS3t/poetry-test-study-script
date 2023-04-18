package pkg

import (
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"strings"
	"time"

	"github.com/deds3t/poem-study/app/models"
)

const (
	POEM_FOLDER = "poems"
)

type Poem struct {
	Name    string
	Author  string
	Stanzas []string
}

func CreatePoem(dto models.PoemDto) *Poem {
	rand.Seed(time.Now().UnixNano())
	data, err := ioutil.ReadFile(filepath.Join(POEM_FOLDER, dto.Poem))
	if err != nil {
		panic(err)

	}

	return &Poem{
		Name:    dto.Name,
		Author:  dto.Author,
		Stanzas: strings.Split(string(data), "\n\n"),
	}
}

func extractFromStanza(stanza string, percent float32) string {
	lines := strings.Split(stanza, "\n")
	numLines := int(float32(len(lines)) * percent)
	idx := rand.Intn(len(lines) - numLines)
	var out string

	for i := idx; i < idx+numLines; i++ {
		out += lines[i] + "\n"
	}
	return out

}

func (p Poem) GetRandomPart() string {
	if len(p.Stanzas) == 0 {
		return ""
	}

	if len(p.Stanzas) == 1 {
		// single stanza poem would need to extract from stanza
		lines := len(strings.Split(p.Stanzas[0], "\n"))
		if lines >= 6 {
			return extractFromStanza(p.Stanzas[0], float32(4)/float32(lines))
		}
		return extractFromStanza(p.Stanzas[0], 0.5)
	}

	stanza := ""

	for len(stanza) < 10 {
		stanza = p.Stanzas[rand.Intn(len(p.Stanzas))]
	}

	lines := len(strings.Split(stanza, "\n"))

	if lines > 4 {
		return extractFromStanza(stanza, float32(4)/float32(lines))
	}

	return stanza
}
