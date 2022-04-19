package cmd

import (
	"vending/app/adapter/http"
	"vending/app/types/constants"
)

func run(mode string) {
	http.NewHttp(constants.DebugMode)
}
