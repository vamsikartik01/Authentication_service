package models

type Config struct {
	Service string `json:"service"`
	Jwt     Jwt    `json:"jwt"`
	Mysql   Mysql  `json:"mysql"`
}

type Jwt struct {
	Secret string `json:"secret"`
}

type Mysql struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}
