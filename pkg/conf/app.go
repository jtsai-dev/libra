package conf

type Config struct {
	App      App      `yaml:"App"`
	Wechat   Wechat   `yaml:"Wechat"`
	Server   Server   `yaml:"Server"`
	Database Database `yaml:"Database"`
	Redis    Redis    `yaml:"Redis"`
	Log      Log      `yaml:"Log"`
}

type App struct {
	TokenExpiredSeconds int      `yaml:"TokenExpiredSeconds"`
	PageSize            int      `yaml:"PageSize"`
	JwtSecret           string   `yaml:"JwtSecret"`
	ImageSavePath       string   `yaml:"ImageSavePath"`
	ImageMaxSize        int      `yaml:"ImageMaxSize"`
	ImageAllowExts      []string `yaml:"ImageAllowExts"`
	ExportPath          string   `yaml:"ExportPath"`
}

type Wechat struct {
	AppId         string `yaml:"AppId"`
	AppSecret     string `yaml:"AppSecret"`
	DefaultAvatar string `yaml:"DefaultAvatar"`
}

type Server struct {
	RunMode string `yaml:"RunMode"`
	Port    int    `yaml:"Port"`
}

type Database struct {
	Type     string `yaml:"Type"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Host     string `yaml:"Host"`
	Name     string `yaml:"Name"`
	Charset  string `yaml:"Charset"`
}

type Redis struct {
	Address  string `yaml:"Address"`
	Password string `yaml:"Password"`
	DBIndex  int    `yaml:"DBIndex"`
}

type Log struct {
	Path            string `yaml:"Path"`
	TimestampFormat string `yaml:"TimestampFormat"`
}
