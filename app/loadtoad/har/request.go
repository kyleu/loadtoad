package har

import "net/url"

type PostParam struct {
	Name        string `json:"name"`
	Value       string `json:"value,omitempty"`
	FileName    string `json:"fileName,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	Comment     string `json:"comment,omitempty"`
}

type PostParams []*PostParam

type PostData struct {
	MimeType string     `json:"mimeType"`
	Params   PostParams `json:"params"`
	Text     string     `json:"text"`
	Comment  string     `json:"comment,omitempty"`
}

type Request struct {
	Method      string    `json:"method"`
	URL         string    `json:"url"`
	HTTPVersion string    `json:"httpVersion"`
	Cookies     Cookies   `json:"cookies"`
	Headers     NVPs      `json:"headers"`
	QueryString NVPs      `json:"queryString"`
	PostData    *PostData `json:"postData"`
	HeaderSize  int       `json:"headerSize"`
	BodySize    int       `json:"bodySize"`
	Comment     string    `json:"comment"`
}

func (r *Request) Size() int {
	if r.PostData == nil {
		return 0
	}
	return len(r.PostData.Text)
}

func (r *Request) GetURL() *url.URL {
	u, err := url.Parse(r.URL)
	if err != nil {
		return nil
	}
	return u
}
