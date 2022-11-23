package main

import (
	"fmt"
	"grpcOpenApi/api"
	"grpcOpenApi/pkg/server"

	"github.com/gin-gonic/gin"
	"github.com/kube-tarian/kad/agent/pkg/logging"
)

func main() {
	log := logging.NewLogger()
	log.Infof("Starting server")

	s := server.NewServer(log)

	r := gin.Default()
	_ = api.RegisterHandlers(r, s)

	if err := r.Run(fmt.Sprintf("%s:%d", "0.0.0.0", 52002)); err != nil {
		log.Errorf("server run returned, reason: %s", err.Error())
	}

	log.Infof("Stopped server")
}
