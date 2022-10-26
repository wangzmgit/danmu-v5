package dto

type DatabasegDto struct {
	DbName    string
	DbHost    string
	DbPort    int
	DbUser    string
	DbPwd     string
	CacheHost string
	CachePort int
	CachePwd  string
}

type EmailConfigDto struct {
	Open     bool
	Name     string
	Host     string
	Port     int
	Address  string
	Password string
}

type StorageConfigDto struct {
	Https           bool
	Oss             bool
	MaxImgSize      uint
	MaxVideoSize    uint
	Bucket          string
	Endpoint        string
	AccesskeyId     string
	AccesskeySecret string
	Domain          string
}

type OtherConfigDto struct {
	Prefix   string
	MaxRes   int
	FePath   string
	Email    string
	Password string
}
