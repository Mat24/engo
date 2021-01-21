package game

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/EngoEngine/engo"
)

var audioFiles = []string{}

// openFile is the desktop-specific way of opening a file
func openFile(url string) (io.ReadCloser, error) {
	return os.Open(url)
}

func parse() (Map, error) {
	var m Map
	f, err := openFile("../assets/example.tmx")
	if err != nil {
		return m, fmt.Errorf("unable to open resource: %s", err)
	}
	defer f.Close()

	d, err := ioutil.ReadAll(f)
	if err != nil {
		return m, err
	}
	err = xml.Unmarshal(d, &m)
	return m, err
}

// Map is the root element of a TMX map
type Map struct {
	Sounds []Sound `xml:"sounds"`
}
type Sound struct {
	Audios []Audio `xml:"audio"`
}
type Audio struct {
	Source string `xml:"source,attr"`
}

func loadAudioFilesFromTMX() error {
	var m Map
	m, err := parse()
	for _, s := range m.Sounds {
		for _, v := range s.Audios {
			audioFiles = append(audioFiles, v.Source)
		}
	}

	for _, v := range audioFiles {
		err = engo.Files.Load(v)
		if err != nil {
			log.Println(fmt.Sprintf("unable to load %v err = %v", v, err))
			return err
		}
	}
	return nil
}
