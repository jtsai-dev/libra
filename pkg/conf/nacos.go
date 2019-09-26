package conf

type Nacos struct {
	ServerAddr     string `yaml:"ServerAddr"`
	ServerPort     uint64 `yaml:"ServerPort"`
	DataId         string `yaml:"DataId"`
	Group          string `yaml:"Group"`
	Tenant         string `yaml:"Tenant"`
	Extension      string `yaml:"Extension"`
	LogDir         string `yaml:"LogDir"`
	TimeoutMs      uint64 `yaml:"TimeoutMs"`
	ListenInterval uint64 `yaml:"ListenInterval"`
}
