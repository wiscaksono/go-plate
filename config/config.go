package config

var (
	DB_USER     = GetENV("DB_USER")
	DB_PASSWORD = GetENV("DB_PASSWORD")
	DB_NAME     = GetENV("DB_NAME")

	PGADMIN_EMAIL    = GetENV("PGADMIN_EMAIL")
	PGADMIN_PASSWORD = GetENV("PGADMIN_PASSWORD")

	APP_PORT = GetENV("APP_PORT")
)
