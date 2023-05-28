package converters

import (
	"testing"

	"github.com/medium.rip/pkg/entities"
)

func TestConvertOli(t *testing.T) {
	oli := entities.Paragraph{
		Name:    "1-1",
		Type:    "OLI",
		Text:    "This is an ordered list item.",
		Markups: []entities.Markup{},
	}
	olis := []entities.Paragraph{oli}
	oliHTML := ConvertParagraphs(olis)
	expected := "<ol><li>This is an ordered list item.</li></ol>"
	if oliHTML != expected {
		t.Errorf("ConvertParagraphs(olis) = %s; want %s", oliHTML, expected)
	}
}