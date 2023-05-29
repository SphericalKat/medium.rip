package entities

import "encoding/json"

func UnmarshalMediumResponse(data []byte) (MediumResponse, error) {
	var r MediumResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MediumResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MediumResponse struct {
	Data Data `json:"data"`
}

type Data struct {
	Post Post `json:"post"`
}

type Post struct {
	Title     string  `json:"title"`
	CreatedAt int64   `json:"createdAt"`
	Creator   Creator `json:"creator"`
	Content   Content `json:"content"`
}

type Content struct {
	BodyModel BodyModel `json:"bodyModel"`
}

type BodyModel struct {
	Paragraphs []Paragraph `json:"paragraphs"`
}

type MediaResource struct {
	Href         string `json:"href"`
	IframeSrc    string `json:"iframeSrc"`
	IframeWidth  int64  `json:"iframeWidth"`
	IframeHeight int64  `json:"iframeHeight"`
}

type Iframe struct {
	MediaResource MediaResource `json:"mediaResource"`
}

type Paragraph struct {
	Name     string      `json:"name"`
	Text     string      `json:"text"`
	Type     Type        `json:"type"`
	Href     interface{} `json:"href"`
	Layout   *string     `json:"layout"`
	Markups  []Markup    `json:"markups"`
	Iframe   *Iframe     `json:"iframe"`
	Metadata *Metadata   `json:"metadata"`
}

type Markup struct {
	Title      *string     `json:"title"`
	Type       string      `json:"type"`
	Href       *string     `json:"href"`
	UserID     interface{} `json:"userId"`
	Start      int64       `json:"start"`
	End        int64       `json:"end"`
	AnchorType *string     `json:"anchorType"`
}

type Metadata struct {
	ID             string `json:"id"`
	OriginalWidth  int64  `json:"originalWidth"`
	OriginalHeight int64  `json:"originalHeight"`
}

type Creator struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Type string

const (
	H3  Type = "H3"
	H4  Type = "H4"
	Img Type = "IMG"
	P   Type = "P"
	Pre Type = "PRE"
)
