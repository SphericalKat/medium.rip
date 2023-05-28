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
	return fmt.Sprintf("%s/%d/%d/%s", IMAGE_HOST, i.width(), i.height(), i.ID)
}

func (i *Image) height() int64 {
	if i.OriginalWidth > MAX_WIDTH {
		return i.OriginalHeight * int64(i.ratio())
	} else {
		return i.OriginalHeight
	}
}

func (i *Image) ratio() float32 {
	return float32(MAX_WIDTH) / float32(i.OriginalWidth)
}

func ConvertParagraphs(paragraphs []entities.Paragraph) string {
	if len(paragraphs) == 0 {
		return ""
	}

	var ps strings.Builder

	for i, p := range paragraphs {
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
			listItems := convertOli(paragraphs[i:])
			ps.WriteString(fmt.Sprintf("<ol>%s</ol>", listItems))
		case "ULI":
			listItems := convertUli(paragraphs[i:])
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
		img := Image{ID : p.Metadata.ID}
		img.Initialize(&p.Metadata.OriginalWidth, &p.Metadata.OriginalHeight)
		return fmt.Sprintf("<figure><img src=\"%s\" width=\"%d\" /><figcaption>%s</figcaption></figure>", img.src(), img.width(), captionMarkup)
	} else {
		return ""
	}
}

func convertOli(ps []entities.Paragraph) string {
	if len(ps) != 0 && ps[0].Type == "OLI" {
		p := ps[0]
		children := ConvertMarkup(p.Text, p.Markups)
		return fmt.Sprintf("<li>%s</li>", children) + convertOli(ps[1:])
	} else {
		return ""
	}
}

func convertUli(ps []entities.Paragraph) string {
	if len(ps) != 0 && ps[0].Type == "ULI" {
		p := ps[0]
		children := ConvertMarkup(p.Text, p.Markups)
		return fmt.Sprintf("<li>%s</li>", children) + convertUli(ps[1:])
	} else {
		return ""
	}
}
