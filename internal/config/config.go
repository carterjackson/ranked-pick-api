package config

import (
	"database/sql"

	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/carterjackson/ranked-pick-api/internal/env"
	"github.com/go-chi/jwtauth/v5"
)

type AppConfig struct {
	Port             int
	Env              string
	Db               *sql.DB
	Queries          *db.Queries
	AccessTokenAuth  *jwtauth.JWTAuth
	RefreshTokenAuth *jwtauth.JWTAuth
}

var Config *AppConfig

func InitConfig() {
	Config = &AppConfig{
		AccessTokenAuth:  jwtauth.New("HS256", []byte(env.GetRequiredString("ACCESS_TOKEN_SECRET")), nil),
		RefreshTokenAuth: jwtauth.New("HS256", []byte(env.GetRequiredString("REFRESH_TOKEN_SECRET")), nil),
	}
	ParseFlags()
	PrepareDatabase()
}
