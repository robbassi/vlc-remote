package main

type VideoEffects struct {
	Gamma float64 `xml:"gamma"`
	Contrast float64 `xml:"contrast"`
	Saturation float64 `xml:"saturation"`
	Brightness float64 `xml:"brightness"`
	Hue float64 `xml:"hue"`
}

type Info struct {
	Name string `xml:"name,attr"`
	Value string `xml:",innerxml"`
}

type Category struct {
	Name string `xml:"name,attr"`
	Info []Info `xml:"info"`
}

type Stats struct {
	AverageInputBitRate float64 `xml:"averageinputbitrate"`
	PlayedABuffers int `xml:"playedabuffers"`
	LostPictures int `xml:"lostpictures"`
	DemuxDiscontinuity int `xml:"demuxdiscontinuity"`
	LostBuffers int `xml:"lostbuffers"`
	SentBytes int `xml:"sentbytes"`
	InputBitRate float64 `xml:"inputbitrate"`
	DecodedAudio int `xml:"decodedaudio"`
	SentPackets int `xml:"sentpackets"`
	SendBitRate float64 `xml:"sendbitrate"`
	DisplayedPictures int `xml:"displayedpictures"`
	DemuxBitRate float64 `xml:"demuxbitrate"`
	DemuxCorrupted int `xml:"demuxcorrupted"`
	DemuxReadPackets int `xml:dmeuxreadpackets"`
	AverageDemuxBitRate float64 `xml:"averagedemuxbitrate"`
	DemuxReadBytes int `xml:"demuxreadbytes"`
	DecodedVideo int `xml:"decodedvideo"`
	ReadBytes int `xml:"readbytes"`
	ReadPackets int `xml:"readpackets"`
}

type StatusResult struct {
	Random bool `xml:"random"`
	APIVersion int8 `xml:"apiversion"`
	Volume int32 `xml:"volume"`
	Fullscreen bool `xml:"fullscreen"`
	State string `xml:"state"`
	//AudioFilters []AudioFilter `xml:"audiofilters"`
	Position float64 `xml:"position"`
	Loop bool `xml:"loop"`
	Equalizer string `xml:"equalizer"`
	Length int `xml:"length"`
	Repeat bool `xml:"repeat"`
	VideoEffects VideoEffects `xml:"videoeffects"`
	Time int `xml:"time"`
	AudioDelay int `xml:"audiodelay"`
	Version string `xml:"version"`
	SubtitleDelay int `xml:"subtitledelay"`
	CurrentItemID int `xml:"currentplid"`
	Rate int `xml:"rate"`
	Information []Category `xml:"information>category"`
	Stats Stats `xml:"stats"`
}
