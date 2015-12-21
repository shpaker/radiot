package radiot

import (
	"errors"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

/*
Возвращает подкаст номер episodeId
*/
func (podcast *Podcast) Show(episodeId int) (*Episode, error) {
	var strEpisode string = strconv.Itoa(episodeId)

	for _, episode := range podcast.Episodes {
		if strings.Split(episode.Title, " ")[1] == strEpisode {

			if episode.Themes == nil {

				if episode.Link != "" {

					// http request
					doc, err := goquery.NewDocument(episode.Link)
					if err != nil {
						return nil, err
					}

					link, _ := doc.Find(".entry-content audio").Attr("src")
					episode.Download = link

					link, _ = doc.Find("img").Attr("src")
					if link[0] == '/' {
						link = podcast.Domain + link
					}
					episode.Image = link

					link, _ = doc.Find(".entry-content em a").Attr("href")
					episode.Sponsor = &Theme{
						Title: doc.Find(".entry-content em a").Text(),
						Link:  link,
					}

					doc.Find(".entry-content ul li").Each(func(i int, docTheme *goquery.Selection) {
						link, _ := docTheme.Find("a").Attr("href")

						episode.Themes = append(episode.Themes, &Theme{
							Title: docTheme.Text(),
							Link:  link,
						})

					})
				}
			}

			if episode.Pirates != nil {
				doc, err := goquery.NewDocument(episode.Pirates.Link)
				if err != nil {
					return nil, err
				}
				link, _ := doc.Find(".entry-content audio").Attr("src")
				episode.Pirates.Download = link
			}

			return episode, nil
		}
	}

	return nil, errors.New("Searching episodes is empty: " + strconv.Itoa(episodeId))
}
