package main

import "fmt"
import "net/http"
import "io/ioutil"
import "net/url"
import "os"
import "encoding/xml"

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

func main() {
	remote := &Remote{
		Host:     "192.168.1.101",
		Port:     "8080",
		Username: "",
		Password: "vlcremote",
	}

	status, err := remote.Status()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
		
	switch os.Args[1] {
	case "play":
		remote.Play()
	case "stop":
		remote.Stop()
	case "next":
		remote.Next()
	case "prev":
		remote.Previous()
	}

	fmt.Printf("%+#v\n", status)
}
