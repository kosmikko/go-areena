package responses

import (
	"strconv"
	"strings"
	"time"
)

// Program details
type Program struct {
	Apiversion string `json:"apiVersion"`
	Data       struct {
		Alternativeid []string `json:"alternativeId"`
		Audio         []struct {
			Format []struct {
				Inscheme string `json:"inScheme"`
				Key      string `json:"key"`
				Type     string `json:"type"`
			} `json:"format"`
			Language []string `json:"language"`
			Type     string   `json:"type"`
		} `json:"audio"`
		Collection    string `json:"collection"`
		Contentrating struct {
			Agerestriction int    `json:"ageRestriction"`
			Ratingsystem   string `json:"ratingSystem"`
			Reason         []struct {
				Key   string `json:"key"`
				Title struct {
					En string `json:"en"`
					Fi string `json:"fi"`
					Sv string `json:"sv"`
				} `json:"title"`
				Type string `json:"type"`
			} `json:"reason"`
			Title struct {
				En string `json:"en"`
				Fi string `json:"fi"`
				Sv string `json:"sv"`
			} `json:"title"`
			Type string `json:"type"`
		} `json:"contentRating"`
		Countryoforigin []string `json:"countryOfOrigin"`
		Creator         []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"creator"`
		Description struct {
			Fi string `json:"fi"`
			Sv string `json:"sv"`
		} `json:"description"`
		Duration      string `json:"duration"`
		Episodenumber int    `json:"episodeNumber"`
		ID            string `json:"id"`
		Image         struct {
			Available bool   `json:"available"`
			ID        string `json:"id"`
			Type      string `json:"type"`
			Version   int    `json:"version"`
		} `json:"image"`
		Indexdatamodified time.Time `json:"indexDataModified"`
		Itemtitle         struct {
			Fi string `json:"fi"`
		} `json:"itemTitle"`
		Originaltitle struct {
			Unknown string `json:"unknown"`
		} `json:"originalTitle"`
		Partofseries struct {
			Countryoforigin []string `json:"countryOfOrigin"`
			Coverimage      struct {
				Available bool   `json:"available"`
				ID        string `json:"id"`
				Type      string `json:"type"`
				Version   int    `json:"version"`
			} `json:"coverImage"`
			Creator []struct {
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"creator"`
			Description struct {
				Fi string `json:"fi"`
			} `json:"description"`
			ID    string `json:"id"`
			Image struct {
				Available bool   `json:"available"`
				ID        string `json:"id"`
				Type      string `json:"type"`
				Version   int    `json:"version"`
			} `json:"image"`
			Indexdatamodified time.Time `json:"indexDataModified"`
			Interactions      []struct {
				Title struct {
					Fi string `json:"fi"`
				} `json:"title"`
				Type string `json:"type"`
				URL  string `json:"url"`
			} `json:"interactions"`
			Subject []struct {
				Broader struct {
					ID string `json:"id"`
				} `json:"broader"`
				ID       string `json:"id"`
				Inscheme string `json:"inScheme"`
				Key      string `json:"key"`
				Title    struct {
					Fi string `json:"fi"`
					Sv string `json:"sv"`
				} `json:"title"`
				Type string `json:"type"`
			} `json:"subject"`
			Title struct {
				Fi string `json:"fi"`
				Sv string `json:"sv"`
			} `json:"title"`
			Type string `json:"type"`
		} `json:"partOfSeries"`
		Productionid     string `json:"productionId"`
		Publicationevent []struct {
			Duration string    `json:"duration"`
			Endtime  time.Time `json:"endTime"`
			ID       string    `json:"id"`
			Media    struct {
				Available bool   `json:"available"`
				ID        string `json:"id"`
				Duration  string `json:"duration"`
			} `json:"media"`
			Region  string `json:"region"`
			Service struct {
				ID string `json:"id"`
			} `json:"service"`
			Starttime      time.Time `json:"startTime"`
			Temporalstatus string    `json:"temporalStatus"`
			Type           string    `json:"type"`
		} `json:"publicationEvent"`
		Subject []struct {
			Broader struct {
				ID string `json:"id"`
			} `json:"broader"`
			ID       string `json:"id"`
			Inscheme string `json:"inScheme"`
			Key      string `json:"key"`
			Title    struct {
				Fi string `json:"fi"`
				Sv string `json:"sv"`
			} `json:"title"`
			Type string `json:"type"`
		} `json:"subject"`
		Subtitling []struct {
			Language []string `json:"language"`
			Type     string   `json:"type"`
		} `json:"subtitling"`
		Title struct {
			Fi string `json:"fi"`
			Sv string `json:"sv"`
		} `json:"title"`
		Type         string `json:"type"`
		Typecreative string `json:"typeCreative"`
		Typemedia    string `json:"typeMedia"`
		Video        struct {
			Format []struct {
				Inscheme string `json:"inScheme"`
				Key      string `json:"key"`
				Type     string `json:"type"`
			} `json:"format"`
			Language []interface{} `json:"language"`
			Type     string        `json:"type"`
		} `json:"video"`
	} `json:"data"`
	Meta struct {
		ID string `json:"id"`
	} `json:"meta"`
}

// MediaID return media id
func (p *Program) MediaID() (id string) {
	for _, evt := range p.Data.Publicationevent {
		if evt.Temporalstatus == "currently" && evt.Type == "OnDemandPublication" && evt.Media.Available {
			return evt.Media.ID
		}
	}
	return
}

// Title return title for program
func (p *Program) Title() (title string) {
	titleParts := []string{p.Data.Title.Fi}
	if p.Data.Episodenumber > 0 {
		titleParts = append(titleParts, strconv.Itoa(p.Data.Episodenumber))
	}
	if p.Data.Itemtitle.Fi != "" {
		titleParts = append(titleParts, p.Data.Itemtitle.Fi)
	}

	return strings.Join(titleParts, " - ")
}
