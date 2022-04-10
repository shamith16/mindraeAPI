package tunefindmodel

type ShowSearch struct {
	Show          Show     `json:"show"`
	Seasons       []Season `json:"seasons"`
	LatestEpisode Episode  `json:"latest_episode"`
}

type MusicSupervisorsComposers struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	NameStub string `json:"name_stub"`
	Image    Image  `json:"image"`
}

type Show struct {
	NameStub string `json:"name_stub"`
	Name     string `json:"name"`
	Image    Image  `json:"image"`
}
