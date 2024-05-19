package configs

import "github.com/spf13/viper"

type DatabaseSettings struct {
	Root      string
	Password  string
	Host      string
	Port      int
	Dbname    string
	Charset   string
	ParseTime string
	Loc       string
}
type JWTSettings struct {
	Issuer    string
	Subject   string
	SecretKey string
}
type Setting struct {
	vp *viper.Viper
}

var (
	DbSettings  DatabaseSettings
	JwtSettings JWTSettings
)
