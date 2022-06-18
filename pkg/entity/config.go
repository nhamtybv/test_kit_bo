package entity

type Config struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ConfigList struct {
	Count   int               `json:"count"`
	Configs map[string]string `json:"configs"`
}
