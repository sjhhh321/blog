package conf

type Config struct {
	System System `yaml:"system"`
	Log    Log    `yaml:"log"`
	DB     DB     `yaml:"db"`  //默认是读库
	DB1    DB     `yaml:"db1"` // 默认是写库
}
