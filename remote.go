package main

import "fmt"
import "net/http"
import "io/ioutil"
import "net/url"
import "os"
import "encoding/xml"
import "strconv"
import "strings"

type Command uint8

const (
	Status Command = iota
	Playlist
	Browse
	VLMCommand
)

const requestString = "http://%v:%v/requests/%v"

var resources = map[Command]string{
	Status:   "status.xml",
	Playlist: "playlist.xml",
	Browse:   "browse.xml",
	VLMCommand: "vlm_cmd.xml",
}

type Remote struct {
	Host     string
	Port     string
	Username string
	Password string
}

func (remote *Remote) do(command Command, parameters map[string]string) (string, error) {
	urlString := fmt.Sprintf(requestString, remote.Host, remote.Port, resources[command])
	url, err := url.Parse(urlString)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if parameters != nil {
		urlQuery := url.Query()
		for key, value := range parameters {
			urlQuery.Set(key, value)
		}
		url.RawQuery = urlQuery.Encode()
	}
	
	client := new(http.Client)
	req, err := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(remote.Username, remote.Password)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(body), nil
}

func (remote *Remote) Status() (*StatusResult, error) {
	xmlString, err := remote.do(Status, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result := StatusResult{}
	err = xml.Unmarshal([]byte(xmlString), &result)

	return &result, nil
}

func (remote *Remote) PlayURI(uri string) (err error) {
	_, err = remote.do(Status, map[string]string{
		"command": "in_play",
		"input": uri,
	})
	return
}

func (remote *Remote) AddURI(uri string) (err error) {
	_, err = remote.do(Status, map[string]string{
		"command": "in_enqueue",
		"input": uri,
	})
	return
}

func (remote *Remote) DeleteItem(id int) (err error) {
	_, err = remote.do(Status, map[string]string{
		"command": "pl_delete",
		"id": string(id),
	})
	return
}

func (remote *Remote) ClearPlaylist() (err error) {
	_, err = remote.do(Status, map[string]string{
		"command": "pl_empty",
	})
	return
}

func (remote *Remote) Shuffle() (err error) {
	_, err = remote.do(Status, map[string]string{
		"command": "pl_random",
	})
	return
}

func (remote *Remote) Loop() (err error) {
	_, err = remote.do(Status, map[string]string{
		"command": "pl_loop",
	})
	return
}

func (remote *Remote) Repeat() (err error) {
	_, err = remote.do(Status, map[string]string{
		"command": "pl_repeat",
	})
	return
}

func (remote *Remote) Fullscreen() (err error) {
	_, err = remote.do(Status, map[string]string{
		"command": "fullscreen",
	})
	return
}

func (remote *Remote) IncreaseVolume(increment int) (err error) {
	_, err = remote.do(Status, map[string]string{
		"command": "volume",
		"val" : fmt.Sprintf("+%v", increment),
	})
	return
}

func (remote *Remote) DecreaseVolume(increment int) (err error) {
	_, err = remote.do(Status, map[string]string{
		"command": "volume",
		"val" : fmt.Sprintf("-%v", increment),
	})
	return
}

func (remote *Remote) Volume(level int) (err error) {
	_, err = remote.do(Status, map[string]string{
		"command": "volume",
		"val" : strconv.Itoa(level),
	})
	return
}

func (remote *Remote) Seek(seconds int) (err error) {
	sign := "+"
	if seconds < 0 {
		sign = "-"
	}
	
	_, err = remote.do(Status, map[string]string{
		"command": "seek",
		"val" : fmt.Sprintf("%v%v", sign, seconds),
	})
	return
}

func (remote *Remote) SeekTo(seconds int) (err error) {
	_, err = remote.do(Status, map[string]string{
		"command": "seek",
		"val" : strconv.Itoa(seconds),
	})
	return
}

func (remote *Remote) Play() (err error) {
	_, err = remote.do(Status, map[string]string{
		"command": "pl_pause",
	})
	return
}

func (remote *Remote) Stop() (err error) {
	_, err = remote.do(Status, map[string]string{
		"command": "pl_stop",
	})
	return
}

func (remote *Remote) Next() (err error) {
	_, err = remote.do(Status, map[string]string{
		"command": "pl_next",
	})
	return
}

func (remote *Remote) Previous() (err error) {
	_, err = remote.do(Status, map[string]string{
		"command": "pl_previous",
	})
	return
}

func (remote *Remote) Playlist() (*PlaylistResult, error) {
	xmlString, err := remote.do(Playlist, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	type Result struct {
		Name string `xml:"name,attr"`
		Items []PlaylistResult `xml:"node"`
	}

	result := Result{}
	err = xml.Unmarshal([]byte(xmlString), &result)

	for _, playlist := range result.Items {
		if playlist.Name == "Playlist" {
			return &playlist, nil
		}
	}

	return nil, nil
}

func (remote *Remote) Browse(directory string) (*[]DirectoryEntry, error) {
	xmlString, err := remote.do(Browse, map[string]string{
		"dir" : directory,
	})
	
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	type Result struct {
		Items []DirectoryEntry `xml:"element"`
	}

	result := Result{}
	err = xml.Unmarshal([]byte(xmlString), &result)

	return &result.Items, nil
}

func main() {
	remote := &Remote{
		Host:     "192.168.1.101",
		Port:     "8080",
		Username: "",
		Password: "vlcremote",
	}
}
