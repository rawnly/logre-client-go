package configuration

type Configuration struct {
	BoxId string `json:"box_id"`
	Debug bool   `json:"debug"`
}

var Config = Configuration{
	Debug: false,
}
