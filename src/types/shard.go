package types

type Shard struct {
	Id int `yaml:"id"`
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`
}
