package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/medium.rip/internal/config"
	"github.com/medium.rip/pkg/entities"
)

func PostData(postId string) (*entities.MediumResponse, error) {
	if config.Conf.Env == "dev" {
		file, err := os.ReadFile("response.json")
		if err != nil {
			return nil, err
		}

		mr, err := entities.UnmarshalMediumResponse(file)
		if err != nil {
			log.Printf("Error unmarshalling body from response %v\n", err)
			return nil, err
		}

		return &mr, nil
	}

	// http client to post data
	url := "https://medium.com/_/graphql"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`query {
        post(id: "%s") {
          title
          createdAt
          creator {
            id
            name
          }
          content {
            bodyModel {
              paragraphs {
                name
                text
                type
                href
                layout
                markups {
                  title
                  type
                  href
                  userId
                  start
                  end
                  anchorType
                }
                iframe {
                  mediaResource {
                    href
                    iframeSrc
                    iframeWidth
                    iframeHeight
                  }
                }
                metadata {
                  id
                  originalWidth
                  originalHeight
                }
              }
            }
          }
        }
      }`, postId))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Printf("Error constructing request %v\n", err)
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		log.Printf("Error making request to Medium API %v\n", err)
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading body from response %v\n", err)
		return nil, err
	}

	mr, err := entities.UnmarshalMediumResponse(body)
	if err != nil {
		log.Printf("Error unmarshalling body from response %v\n", err)
		return nil, err
	}

	return &mr, nil
}
