package app_config

import "os"

var PORT string

func APP_CONFIG() {
	PORT = os.Getenv("ORDER_APP_PORT")
}
