package radiot

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

/*
Загружает список подкатов
*/
func (podcast *Podcast) Update() (err error) {

	podcast.Episodes, err = parseArchiveHtml(podcast.Domain)
	if err != nil {
		return err
	}

	epTime, _ := time.Parse("2006-01-02T15:04", "2007-08-26T06:00")
	podcast.Episodes = append(podcast.Episodes, &Episode{
		Title:    "Радио-Т 49",
		Time:     epTime,
		Download: "http://archive.rucast.net/radio-t/media/rt_podcast49.mp3",
	})
	epTime, _ = time.Parse("2006-01-02T15:04", "2007-12-09T06:00")
	podcast.Episodes = append(podcast.Episodes, &Episode{
		Title:    "Радио-Т 64",
		Time:     epTime,
		Download: "http://archive.rucast.net/radio-t/media/rt_podcast64.mp3",
	})
	epTime, _ = time.Parse("2006-01-02T15:04", "2009-08-23T06:00")
	podcast.Episodes = append(podcast.Episodes, &Episode{
		Title:    "Радио-Т 150",
		Time:     epTime,
		Download: "http://archive.rucast.net/radio-t/media/rt_podcast150.mp3",
	})
	epTime, _ = time.Parse("2006-01-02T15:04", "2009-10-16T06:00")
	podcast.Episodes = append(podcast.Episodes, &Episode{
		Title:    "Радио-Т 259",
		Time:     epTime,
		Download: "http://archive.rucast.net/radio-t/media/rt_podcast259.mp3",
	})

	piratesEpisodes, err := parseArchiveHtml(PIRATES_DOMAIN)
	if err != nil {
		return err
	}

	for _, episode := range podcast.Episodes {
		episodeId := strings.Split(episode.Title, " ")[1]

		// todo convert to int
		for _, pirate := range piratesEpisodes {

			if len(strings.Split(pirate.Title, " ")) > 2 {
				pId := strings.Split(pirate.Title, " ")[2]
				if pId == episodeId {
					episode.Pirates = &Pirates{
						Link: pirate.Link,
					}
				}
			}
		}
	}
	return
}

func parseArchiveHtml(domain string) (episodes []*Episode, err error) {

	html, err := goquery.NewDocument(domain + PODCAST_PAGE)
	if err != nil {
		return nil, err
	}

	html.Find("article[role=article]").Each(func(i int, articleWrapper *goquery.Selection) {
		articleWrapper.Find("article").Each(func(i int, articleHtml *goquery.Selection) {
			article := &Episode{}
			article.Title = articleHtml.Find("h1 a").Text()
			article.Link, _ = articleHtml.Find("h1 a").Attr("href")
			article.Link = domain + article.Link
			timeString, _ := articleHtml.Find("time").Attr("datetime")
			article.Time, _ = time.Parse(TIME_FORMAT, timeString)
			episodes = append(episodes, article)
		})
	})

	return
}
