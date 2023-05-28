package converters

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var r = regexp.MustCompile(`[\/\-]([0-9a-f]+)\/?$`)

func ConvertId(c *fiber.Ctx) string {
	path := c.Path()
	postId := idFromParams(c)
	if postId == "" {
		return idFromPath(path)
	}
	return ""
}

func idFromPath(path string) string {
	if strings.HasPrefix(path, "/tag/") {
		return ""
	}

	matches := r.FindStringSubmatch(path)
	if len(matches) != 2 {
		return ""
	}

	return matches[1]
}

func idFromParams(c *fiber.Ctx) string {
	ru := c.Query("redirectUrl", "")
	if ru != "" {
		pu, err := url.Parse(ru)
		if err != nil {
			return ""
		}

		return idFromPath(pu.Path)
	}
	return ""
}
