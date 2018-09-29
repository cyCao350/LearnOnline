package main

import (
	"LearnOnline/infra/init"
	"LearnOnline/ui/api-server"
	_ "./docs"
)

// @title LearnOnline-2018 API
// @version 1.0
// @description This is a server for LearnOnline-2018.

// @contact.name API Support
// @contact.email dacheng@ultrachain.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:5005
// @BasePath /v1/api

// @securityDefinitions.apiKey Bearer
// @type apiKey
// @in header
// @name Authorization
func main() {
	// 1 create table by gorm auto migrate
	defer initiator.POSTGRES.Close()

	// 2 start http server
	api_server.New().Start()
}
