package config

type Getter struct {
	Method string   `yaml:"method"`
	Agent  string   `yaml:"agent"`
	Regexp string   `yaml:"regexp"`
	Proto  string   `yaml:"proto"`
	Urls   []string `yaml:"urls"`
	Proxy  bool     `yaml:"proxy"`
}
