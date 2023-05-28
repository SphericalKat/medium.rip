package converters

import (
	"encoding/json"
	"testing"

	"github.com/medium.rip/pkg/entities"
)

func TestRanges(t *testing.T) {
	ranges := ranges("strong and emphasized only", []entities.Markup{
		{
			Type:  "STRONG",
			Start: 0,
			End:   10,
		},
		{
			Type:  "EM",
			Start: 7,
			End:   21,
		},
	})

	if len(ranges) != 4 {
		t.Errorf("Expected 4 ranges, got %d", len(ranges))
	}

	if ranges[0].Range[0] != 0 || ranges[0].Range[1] != 7 {
		t.Errorf("Expected range to be [0, 7], got %v", ranges[0].Range)
	}

	if ranges[0].Markups[0].Type != "STRONG" {
		t.Errorf("Expected markup to be STRONG, got %s", ranges[0].Markups[0].Type)
	}

	if ranges[1].Range[0] != 7 || ranges[1].Range[1] != 10 {
		t.Errorf("Expected range to be [7, 10], got %v", ranges[1].Range)
	}

	if ranges[1].Markups[0].Type != "STRONG" {
		t.Errorf("Expected markup to be STRONG, got %s", ranges[1].Markups[0].Type)
	}

	if ranges[2].Range[0] != 10 || ranges[2].Range[1] != 21 {
		t.Errorf("Expected range to be [10, 21], got %v", ranges[2].Range)
	}

	if ranges[2].Markups[0].Type != "EM" {
		t.Errorf("Expected markup to be EM, got %s", ranges[2].Markups[0].Type)
	}

	if ranges[3].Range[0] != 21 || ranges[3].Range[1] != 26 {
		t.Errorf("Expected range to be [21, 26], got %v", ranges[3].Range)
	}

	if len(ranges[3].Markups) != 0 {
		t.Errorf("Expected markup to be empty, got %v", ranges[3].Markups)
	}
}

func TestConvert(t *testing.T) {
	jsonData := `{
		"name": "254a",
		"text": "Early Flush prevents subsequent changes to the headers (e.g to redirect or change the status code). In the React + NodeJS world, it’s common to delegate redirects and error throwing to a React app rendered after the data has been fetched. This won’t work if you’ve already sent an early <head> tag and a 200 OK status.",
		"type": "P",
		"href": null,
		"layout": null,
		"markups": [
			{
				"title": null,
				"type": "CODE",
				"href": null,
				"userId": null,
				"start": 287,
				"end": 293,
				"anchorType": null
			}
		],
		"iframe": null,
		"metadata": null
	}`
	p := new(entities.Paragraph)
	_ = json.Unmarshal([]byte(jsonData), p)

	ConvertMarkup(p.Text, p.Markups)

	// if markup != "<strong>strong </strong><em><strong>and</strong></em><em> emphasized</em> only" {
	// 	t.Errorf("Expected markup to be <strong>strong </strong><em><strong>and</strong></em><em> emphasized</em> only, got %s", markup)
	// }
}
