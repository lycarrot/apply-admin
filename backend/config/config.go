package config

type Server struct {
	AutoCode AutoCode        `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	Jwt      Jwt             `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap      Zap             `mapstructure:"zap" json:"zap" yaml:"zap"`
	System   System          `mapstructure:"system" json:"system" yaml:"system"`
	Mysql    Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	DBList   []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	Redis    Redis           `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email    Email           `mapstructure:"email" json:"email" yaml:"email"`
	Local    Local           `mapstructure:"local" json:"local" yaml:"local"`
	Cors     Cors            `mapstructure:"cors" json:"cors" yaml:"cors"`
	Captcha  Captcha         `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
}
