package config

var(
	_appConfig *AppConfig = nil
)

type AppConfig struct {
	BindingConf BingdingConfig 	`yaml:"binding_config"`
	VastConf VastConfig		`yaml:"vast_config"`
	ControlConf  ControlConfig	`yaml:"control_config"`
}

type BingdingConfig struct {
	Addr string `yaml:"addr"`
}

type VastConfig struct {
	VideoList []string	`yaml:"video_list"`
	NoVideoList []string	`yaml:"no_video_list"`
}

type DetailControlConfig struct {
	SuccRate float32		`yaml:"succ_rate"`
	AvgCost	 int			`yaml:"avg_cost"`
}

type AdxControl struct {

}

type ControlConfig struct {
	S2cControl DetailControlConfig `yaml:"s2s_control"`
	AdxControl DetailControlConfig `yaml:"adx_control"`
}

func GetConfig() *AppConfig {
	return _appConfig
}

func SetConfig(conf *AppConfig)  {
	_appConfig = conf
}