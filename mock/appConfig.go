package mock

import "github.com/larb26656/assessment-tax/config"

func NewMockAppConfig() config.AppConfig {
	return config.AppConfig{
		Port:          "8080",
		DatabaseUrl:   "postgres://postgres:postgres@127.0.0.1:5432/ktaxes?sslmode=disable",
		AdminUsername: "adminTax",
		AdminPassword: "admin!",
	}
}
