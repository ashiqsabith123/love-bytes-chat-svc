package main

import (
	"github.com/ashiqsabith123/chat-svc/pkg/config"
	"github.com/ashiqsabith123/chat-svc/pkg/db"
	logs "github.com/ashiqsabith123/love-bytes-proto/log"
)

func main() {

	err := logs.InitLogger("./pkg/logs/log.log")
	if err != nil {

		logs.ErrLog.Fatalln("Error while initilizing logger", err)
	}

	config, err := config.LoadConfig()

	if err != nil {
		logs.ErrLog.Fatalln("Error while loading config", err)
	}

	db.ConnectToDatabase(config)

}
