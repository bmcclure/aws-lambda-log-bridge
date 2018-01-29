package config

type LogSourceConfig struct {
	Id string `json:"id"`
}

type LogDestinationConfig struct {
	Id string `json:"id"`
}

type LogFileConfig struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"`
}

type AppConfig struct {
	Name string `json:"name"`
	Source LogSourceConfig `json:"source"`
	Destination LogDestinationConfig `json:"destination"`
	LogFiles []LogFileConfig `json:"logFiles"`
}
