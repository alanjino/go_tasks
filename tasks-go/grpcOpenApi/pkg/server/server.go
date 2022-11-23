package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kube-tarian/kad/agent/pkg/logging"
)

type Server struct {
	log logging.Logger
}

func NewServer(log logging.Logger) *Server {
	return &Server{
		log: log,
	}
}

func (s *Server) GetStatus(c *gin.Context) {
	s.log.Infof("Get Status is invoked, %+v", c)
	c.IndentedJSON(http.StatusOK, "OK")
}

func (s *Server) PostHello(c *gin.Context) {
	s.log.Infof("Post Hello is invoked, %+v", c.Request.Body)
	payload, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		s.log.Errorf("payload read failed, %v", err)
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.String(http.StatusOK, fmt.Sprintf("Hello %v", string(payload)))
}

func (s *Server) PostBye(c *gin.Context) {
	s.log.Infof("Post Hello is invoked, %+v", c.Request.Body)
	payload, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		s.log.Errorf("payload read failed, %v", err)
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.String(http.StatusOK, fmt.Sprintf("bye %v", string(payload)))
}

func (s *Server) GetBye(c *gin.Context) {
	s.log.Infof("Get Status is invoked, %+v", c)
	c.IndentedJSON(http.StatusOK, "OK")
}
