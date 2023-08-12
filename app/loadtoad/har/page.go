package har

type Page struct {
	StartedDateTime string      `json:"startedDateTime"`
	ID              string      `json:"id"`
	Title           string      `json:"title"`
	PageTiming      *PageTiming `json:"pageTiming"`
	Comment         string      `json:"comment,omitempty"`
}

type Pages []*Page

type PageTiming struct {
	OnContentLoad int    `json:"onContentLoad"`
	OnLoad        int    `json:"onLoad"`
	Comment       string `json:"comment"`
}

type PageTimings struct {
	Total   int    `json:"total,omitempty"`
	Blocked int    `json:"blocked,omitempty"`
	DNS     int    `json:"dns,omitempty"`
	Connect int    `json:"connect,omitempty"`
	Send    int    `json:"send,omitempty"`
	Wait    int    `json:"wait,omitempty"`
	Receive int    `json:"receive,omitempty"`
	Ssl     int    `json:"ssl,omitempty"`
	Comment string `json:"comment,omitempty"`
}
