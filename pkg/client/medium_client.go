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
	if config.Conf.Env == "devd" {
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

	payload := strings.NewReader(fmt.Sprintf("{\"query\":\"query {\\n        post(id: \\\"%s\\\") {\\n          title\\n          createdAt\\n          creator {\\n            id\\n            name\\n          }\\n          content {\\n            bodyModel {\\n              paragraphs {\\n                name\\n                text\\n                type\\n                href\\n                layout\\n                markups {\\n                  title\\n                  type\\n                  href\\n                  userId\\n                  start\\n                  end\\n                  anchorType\\n                }\\n                iframe {\\n                  mediaResource {\\n                    href\\n                    iframeSrc\\n                    iframeWidth\\n                    iframeHeight\\n                  }\\n                }\\n                metadata {\\n                  id\\n                  originalWidth\\n                  originalHeight\\n                }\\n              }\\n            }\\n          }\\n        }\\n      }\",\"variables\":{}}", postId))

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
