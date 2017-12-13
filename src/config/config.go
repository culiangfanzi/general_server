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
	PmpList []string	`yaml:"pmp_list"`
	S2sList []string	`yaml:"s2s_list"`
}

type DetailControlConfig struct {
	SuccRate float32		`yaml:"succ_rate"`
	AvgCost	 int			`yaml:"avg_cost"`
	DiffCost int			`yaml:"diff_cost"`
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