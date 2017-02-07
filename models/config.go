package models

type HTTPConfig struct {
	Serve bool `yaml:"serve_api"`
	Port  int  `yaml:"port"`
}

type Config struct {
	Datastore  string `yaml:"datastore"`
	ServerMode bool   `yaml:"server_mode"`
	LogLevel   string `yaml:"log_level"`
	WebHook    string `yaml:"webhook"`
	LogToFile  bool   `yaml:"file_log"`
	HTTPConfig `yaml:"http"`
}
