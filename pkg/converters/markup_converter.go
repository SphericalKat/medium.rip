package converters

import (
	"sort"

	"github.com/medium.rip/pkg/entities"
)

type RangeWithMarkup struct {
	Range  []int
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
	markupBoundaries = append(markupBoundaries, len(text))

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
			if (int(m.Start) >= start && int(m.Start) < end) || (int(m.End - 1) >= start && int(m.End - 1) < end) {
				coveredMarkups = append(coveredMarkups, m)
			}
		}

		// append the range
		ranges = append(ranges, RangeWithMarkup{
			Range: []int{start, end},
			Markups: coveredMarkups,
		})
	}

	return ranges
}

func Convert(text string, markups []entities.Markup) {
	// for _, m := range markups {
	// 	switch m.Type {
	// 	case "A":
	// 		if m.Href != nil {

	// 		} else if {
	// 			m.UserID != nil {

	// 			}
	// 		}
	// 	case "CODE":
	// 	case "EM":
	// 	case "STRONG":
	// 	}
	// }
}