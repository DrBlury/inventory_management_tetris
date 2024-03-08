package repo

type Config struct {
	Host         string
	Port         int `default:"5432"`
	DatabaseName string
	Username     string
	Password     string
	Level        string `default:"debug"`
}
