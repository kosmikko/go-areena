package responses

type Playout struct {
	Meta struct {
		ID string `json:"id"`
	} `json:"meta"`
	Data []struct {
		Subtitles []struct {
			Lang string `json:"lang"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"subtitles"`
		Protocol       string `json:"protocol"`
		Multibitrate   bool   `json:"multibitrate"`
		Formatof       string `json:"formatOf"`
		Width          int    `json:"width"`
		Type           string `json:"type"`
		URL            string `json:"url"`
		Live           bool   `json:"live"`
		Protectiontype string `json:"protectionType"`
		Height         int    `json:"height"`
	} `json:"data"`
	Notifications struct {
	} `json:"notifications"`
}

// EncodedURL get the url for playout
func (playout *Playout) EncodedURL() string {
	for _, p := range playout.Data {
		if p.URL != "" {
			return p.URL
		}
	}
	return ""
}

func (playout *Playout) SubtitleURL(lang string) string {
	for _, p := range playout.Data {
		subs := p.Subtitles
		for _, s := range subs {
			if s.Lang == lang {
				return s.URI
			}
		}
	}
	return ""
}
