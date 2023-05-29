package converters

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/medium.rip/pkg/entities"
	log "github.com/sirupsen/logrus"
)

func ConvertEmbedded(media entities.MediaResource) string {
	if media.IframeSrc == "" {
		return customEmbed(media)
	} else {
		return fmt.Sprintf("<iframe src=\"%s\" width=\"%d\" height=\"%d\" frameborder=\"0\" allowfullscreen=\"true\" ></iframe>", media.IframeSrc, media.IframeWidth, media.IframeHeight)
	}
}

func customEmbed(media entities.MediaResource) string {
	if strings.HasPrefix(media.Href, "https://gist.github.com") {
		return fmt.Sprintf("<script src=\"%s.js\"></script>", media.Href)
	} else {
		url, err := url.Parse(media.Href)
		var caption string
		if err != nil {
			log.Warnf("Error parsing url %s", media.Href)
			caption = media.Href
		} else {
			caption = fmt.Sprintf("Embedded content at %s", url.Host)
		}
		return fmt.Sprintf("<figure><a href=\"%s\">%s</a></figure>", media.Href, caption)
	}
}