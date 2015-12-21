package radiot

import (
	"time"
)

const (
	RADIOT_FLAG   string = "radio-t"
	RADIOT_DOMAIN string = "http://www.radio-t.com"

	PIRATES_FLAG   string = "pirates"
	PIRATES_DOMAIN string = "http://pirates.radio-t.com"

	PODCAST_PAGE string = "/categories/podcast/"
	TIME_FORMAT  string = "2006-01-02T15:04:05-07:00"
)

type (
	Podcast struct {
		Name     string     `json:"-"`
		Domain   string     `json:"-"`
		Episodes []*Episode `json:"episodes"`
	}

	Episode struct {
		Title    string    `json:"title"`
		Link     string    `json:"link,omitempty"`
		Time     time.Time `json:"time"`
		Image    string    `json:"image,omitempty"`
		Download string    `json:"download,omitempty"`
		Pirates  *Pirates  `json:"pirates,omitempty"`
		Themes   []*Theme  `json:"themes,omitempty"`
		Sponsor  *Theme    `json:"sponsor,omitempty"`
	}

	Pirates struct {
		Link     string `json:"link,omitempty"`
		Download string `json:"download,omitempty"`
	}

	Theme struct {
		Title string `json:"title,omitempty"`
		Link  string `json:"link,omitempty"`
	}
)

func NewPodcast(isRadioT bool) *Podcast {
	if isRadioT {
		return &Podcast{
			Name:   RADIOT_FLAG,
			Domain: RADIOT_DOMAIN,
		}
	}
	return &Podcast{
		Name:   PIRATES_FLAG,
		Domain: PIRATES_DOMAIN,
	}
}
