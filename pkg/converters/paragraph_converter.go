package converters

import (
	"fmt"
	"strings"

	"github.com/medium.rip/pkg/entities"
)

const IMAGE_HOST = "https://cdn-images-1.medium.com/fit/c"
const MAX_WIDTH = 800
const FALLBACK_HEIGHT = 600

type Image struct {
	ID             string
	OriginalHeight int64
	OriginalWidth  int64
}

func (i *Image) Initialize(originalWidth *int64, originalHeight *int64) {
	if originalWidth != nil {
		i.OriginalWidth = *originalWidth
	} else {
		i.OriginalWidth = MAX_WIDTH
	}

	if originalHeight != nil {
		i.OriginalHeight = *originalHeight
	} else {
		i.OriginalHeight = FALLBACK_HEIGHT
	}
}

func (i *Image) width() int64 {
	if i.OriginalWidth > MAX_WIDTH {
		return MAX_WIDTH
	} else {
		return i.OriginalWidth
	}
}

func (i *Image) src() string {
	return fmt.Sprintf("https://miro.medium.com/v2/resize:fit:1200/%s", i.ID)
}

func ConvertParagraphs(paragraphs []entities.Paragraph) string {
	if len(paragraphs) == 0 {
		return ""
	}

	var ps strings.Builder

	skipCount := 0
	for i, p := range paragraphs {
		if skipCount > 0 {
			skipCount--
			continue
		}

		switch p.Type {
		case "BQ", "MIXTAPE_EMBED", "PQ":
			children := ConvertMarkup(p.Text, p.Markups)
			ps.WriteString(fmt.Sprintf("<blockquote><p>%s</p></blockquote>", children))
		case "H2":
			children := ConvertMarkup(p.Text, p.Markups)
			if p.Name != "" {
				ps.WriteString(fmt.Sprintf("<h2 id=\"%s\">%s</h2>", p.Name, children))
			} else {
				ps.WriteString(fmt.Sprintf("<h2>%s</h2>", children))
			}
		case "H3":
			children := ConvertMarkup(p.Text, p.Markups)
			if p.Name != "" {
				ps.WriteString(fmt.Sprintf("<h3 id=\"%s\">%s</h3>", p.Name, children))
			} else {
				ps.WriteString(fmt.Sprintf("<h3>%s</h3>", children))
			}
		case "H4":
			children := ConvertMarkup(p.Text, p.Markups)
			if p.Name != "" {
				ps.WriteString(fmt.Sprintf("<h4 id=\"%s\">%s</h4>", p.Name, children))
			} else {
				ps.WriteString(fmt.Sprintf("<h4>%s</h4>", children))
			}
		// TODO: handle IFRAME
		case "IMG":
			ps.WriteString(convertImg(p))
		case "OLI":
			listItems, skip := convertOli(paragraphs[i:])
			skipCount = skip
			ps.WriteString(fmt.Sprintf("<ol>%s</ol>", listItems))
		case "ULI":
			listItems, skip := convertUli(paragraphs[i:])
			skipCount = skip
			ps.WriteString(fmt.Sprintf("<ul>%s</ul>", listItems))
		case "P":
			children := ConvertMarkup(p.Text, p.Markups)
			ps.WriteString(fmt.Sprintf("<p>%s</p>", children))
		case "PRE":
			children := ConvertMarkup(p.Text, p.Markups)
			ps.WriteString(fmt.Sprintf("<pre>%s</pre>", children))
		case "SECTION_CAPTION":
			// unused
		default:
		}
	}

	return ps.String()
}

func convertImg(p entities.Paragraph) string {
	if p.Metadata != nil {
		captionMarkup := ConvertMarkup(p.Text, p.Markups)
		img := Image{ID: p.Metadata.ID}
		img.Initialize(&p.Metadata.OriginalWidth, &p.Metadata.OriginalHeight)
		return fmt.Sprintf("<figure><img src=\"%s\" width=\"%d\" /><figcaption>%s</figcaption></figure>", img.src(), img.width(), captionMarkup)
	} else {
		return ""
	}
}

func convertOli(ps []entities.Paragraph) (string, int) {
	var sb strings.Builder
	count := 0

	for _, p := range ps {
		if p.Type == "OLI" {
			children := ConvertMarkup(p.Text, p.Markups)
			sb.WriteString(fmt.Sprintf("<li>%s</li>", children))
			count++
		} else {
			break
		}
	}
	
	return sb.String(), count
}

func convertUli(ps []entities.Paragraph) (string, int) {
	var sb strings.Builder
	count := 0

	for _, p := range ps {
		if p.Type == "ULI" {
			children := ConvertMarkup(p.Text, p.Markups)
			sb.WriteString(fmt.Sprintf("<li>%s</li>", children))
			count++
		} else {
			break
		}
	}
	
	return sb.String(), count
}
