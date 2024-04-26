package database

type Config struct {
	Driver     string `yaml:"driver"`
	Connection string `yaml:"connection"`
}
