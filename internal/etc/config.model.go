package etc

type Configuration struct {
	Name     string   `toml:"name"`
	Version  string   `toml:"version"`
	Db       db       `toml:"db"`
	Web      web      `toml:"web"`
	Influx   influx   `toml:"influx"`
	DownLink downLink `toml:"down-link"`
}

var Config *Configuration

type db struct {
	Host     string `toml:"host"`
	Database string `toml:"database"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

type web struct {
	Listen string `toml:"listen"`
}

type influx struct {
	Host     string `toml:"host"`
	Database string `toml:"database"`
	Enable   int    `toml:"enable"`
}

type downLink struct {
	IdleTime  int64 `toml:"idle-time"`
	FetchSize int   `toml:"fetch-size"`
	LogEnable bool  `toml:"log-enable"`
}
