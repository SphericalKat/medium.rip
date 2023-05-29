package converters

import (
	"fmt"
	"html"
	"sort"
	"strings"
	"unicode/utf16"

	"github.com/medium.rip/pkg/entities"
)

type RangeWithMarkup struct {
	Range   []int
	Markups []entities.Markup
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func ranges(text string, markups []entities.Markup) []RangeWithMarkup {
	ranges := make([]RangeWithMarkup, 0)

	// first, get all the borders of the markups
	markupBoundaries := make([]int, 0)
	for _, m := range markups {
		markupBoundaries = append(markupBoundaries, []int{int(m.Start), int(m.End)}...)
	}

	// include the start and end indexes of the text
	markupBoundaries = append([]int{0}, markupBoundaries...)
	markupBoundaries = append(markupBoundaries, len(utf16.Encode([]rune(text))))

	// remove duplicates
	markupBoundaries = unique(markupBoundaries)

	// sort slice
	sort.Slice(markupBoundaries, func(i, j int) bool {
		return markupBoundaries[i] < markupBoundaries[j]
	})

	// attach markup to every range
	for i := 0; i < len(markupBoundaries)-1; i++ {
		start := markupBoundaries[i]
		end := markupBoundaries[i+1]

		// check if this markup is covered by the range
		coveredMarkups := make([]entities.Markup, 0)
		for _, m := range markups {
			if (int(m.Start) >= start && int(m.Start) < end) || (int(m.End-1) >= start && int(m.End-1) < end) {
				coveredMarkups = append(coveredMarkups, m)
			}
		}

		// append the range
		ranges = append(ranges, RangeWithMarkup{
			Range:   []int{start, end},
			Markups: coveredMarkups,
		})
	}

	return ranges
}

func ConvertMarkup(text string, markups []entities.Markup) string {
	if len(markups) == 0 {
		return html.EscapeString(text)
	}

	var markedUp strings.Builder
	for _, r := range ranges(text, markups) {
		// handle utf-16
		utf16Text := utf16.Encode([]rune(text))
		ranged := utf16Text[r.Range[0]:r.Range[1]]
		textToWrap := string(utf16.Decode(ranged))
		markedUp.WriteString(wrapInMarkups(textToWrap, r.Markups, false))
	}

	return markedUp.String()
}

func wrapInMarkups(child string, markups []entities.Markup, childIsMarkup bool) string {
	if len(markups) == 0 {
		return child
	}
	if !childIsMarkup {
		child = html.EscapeString(child)
	}
	markedUp := markupNodeInContainer(child, markups[0])
	return wrapInMarkups(markedUp, markups[1:], true)
}

func markupNodeInContainer(child string, markup entities.Markup) string {
	switch markup.Type {
	case "A":
		if markup.Href != nil {
			return fmt.Sprintf(`<a href="%s">%s</a>`, *markup.Href, child)
		} else if markup.UserID != nil {
			return fmt.Sprintf(`<a href="https://medium.com/u/%s">%s</a>`, markup.UserID, child)
		}
	case "CODE":
		return fmt.Sprintf(`<code>%s</code>`, child)
	case "EM":
		return fmt.Sprintf(`<em>%s</em>`, child)
	case "STRONG":
		return fmt.Sprintf(`<strong>%s</strong>`, child)
	default:
		return fmt.Sprintf(`<code>%s</code>`, child)
	}
	return child
}
