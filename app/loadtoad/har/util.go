package har

type HarWrapper struct {
	Log *Log `json:"log"`
}

type NVP struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Comment string `json:"comment,omitempty"`
}

type NVPs []*NVP

func (p NVPs) GetValue(n string) string {
	for _, x := range p {
		if x.Name == n {
			return x.Value
		}
	}
	return ""
}

type Content struct {
	Size        int    `json:"size"`
	Compression int    `json:"compression,omitempty"`
	MimeType    string `json:"mimeType"`
	Text        string `json:"text,omitempty"`
	Encoding    string `json:"encoding,omitempty"`
	Comment     string `json:"comment,omitempty"`
	File        string `json:"_file,omitempty"`
}
