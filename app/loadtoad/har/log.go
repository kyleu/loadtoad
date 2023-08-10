package har

type Creator struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Comment string `json:"comment,omitempty"`
}

type Browser struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Comment string `json:"comment"`
}

type Log struct {
	Version string   `json:"version"`
	Creator *Creator `json:"creator"`
	Browser *Browser `json:"browser"`
	Pages   Pages    `json:"pages,omitempty"`
	Entries Entries  `json:"entries"`
	Comment string   `json:"comment"`
}
