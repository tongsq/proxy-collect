package config

type TunnelConfig struct {
	Proto string             `yaml:"proto"`
	Host  string             `yaml:"host"`
	Port  string             `yaml:"port"`
	Users []TunnelUserConfig `yaml:"users"`
}

type TunnelUserConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Limiter  string `yaml:"limiter"`
}
