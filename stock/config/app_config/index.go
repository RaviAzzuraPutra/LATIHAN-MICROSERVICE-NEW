package app_config

import "os"

var APP string

func App_Config() {
	APP = os.Getenv("STOCK_APP_PORT")
}
