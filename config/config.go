package config

import "os"

type Config struct {
	AdminPassword string
	Secret        string
	DB_Loc        string
	SMTP_PASSWORD string
}

func NewConfig() *Config {
	ret := new(Config)

	ret.AdminPassword = os.Getenv("ADMIN_PASSWORD")
	ret.Secret = os.Getenv("SECRET")
	ret.DB_Loc = os.Getenv("DB_LOC")
	ret.SMTP_PASSWORD = os.Getenv("SMTP_PASSWORD")

	if ret.AdminPassword == "" || ret.Secret == "" || ret.DB_Loc == "" || ret.SMTP_PASSWORD == "" {
		panic("Some config is missing.  Not going to say what here though!")
	}

	return ret
}
