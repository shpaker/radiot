package radiot

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

const (
	JSON_FILENAME string = "radiot.json"
)

type JsonFile struct {
	podcast *Podcast
	file    string
}

func NewJsonFile(pathToFolder string) (rt *JsonFile, err error) {

	rt = &JsonFile{
		file:    pathToFolder + JSON_FILENAME,
		podcast: NewPodcast(true),
	}

	// create json file if !exist
	if _, err := os.Stat(rt.file); os.IsNotExist(err) {

		file, err := os.Create(rt.file)
		defer file.Close()
		if err != nil {
			return nil, err
		}
		rt.writeToJsonFile()
	}

	rt.readJsonFile()
	return rt, nil
}

func (rt *JsonFile) UpdateJsonFile() (err error) {
	var newpodcast *Podcast

	newpodcast = NewPodcast(true)
	if err = newpodcast.Update(); err != nil {
		return err
	}

	for _, episode := range newpodcast.Episodes {
		isNewEpisode := true
		for _, oldEpisode := range rt.podcast.Episodes {
			if episode.Time.Equal(oldEpisode.Time) {
				isNewEpisode = false
			}
		}
		if isNewEpisode {
			rt.podcast.Episodes = append(rt.podcast.Episodes, episode)
		}
	}

	rt.writeToJsonFile()

	return nil
}

func (rt *JsonFile) GetLatestEpisodeId() int {
	return len(rt.podcast.Episodes) - 1
}

func (rt *JsonFile) GetEpisodeMessage(id int) (str string, err error) {
	episode, err := rt.podcast.Show(id)
	if err != nil {
		return
	}

	var themesText string

	for _, th := range episode.Themes {
		themesText += "  * " + th.Title + "\n  "
		if th.Link != "" {
			themesText += th.Link + "\n"
		}
	}

	str = `Выпуск ` + episode.Title + `
` + episode.Image + `

ТЕМЫ ВЫПУСКА
` + themesText

	if episode.Sponsor != nil {
		str += "\nСПОНСОРЫ ЭТОГО ВЫПУСКА\n  " + episode.Sponsor.Title + " " + episode.Sponsor.Link + "\n"
	}

	str += "\n" + strings.ToUpper(episode.Title) + " " + episode.Download + "\n"

	if episode.Pirates != nil {
		str += "\nПИРАТЫ РТ  " + episode.Pirates.Download
	}

	rt.writeToJsonFile()
	return
}

func (rt *JsonFile) readJsonFile() (err error) {
	file, err := ioutil.ReadFile(rt.file)
	if err != nil {
		return err
	}
	json.Unmarshal(file, &rt.podcast)

	return nil
}

func (rt *JsonFile) writeToJsonFile() (err error) {
	sort.Sort(rt.podcast)

	jsondata, err := json.Marshal(rt.podcast)
	if err != nil {
		return err
	}

	jsondata, err = json.MarshalIndent(rt.podcast, "", "\t")
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(rt.file, jsondata, 0644); err != nil {
		return err
	}

	return nil
}
