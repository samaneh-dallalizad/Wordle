package app

import (
	"fmt"
	"net/http"

	//"strings"
	"time"

	"wordleGame/internal/infrastructure"

	"github.com/gin-gonic/gin"
)

type ApplicationState struct {
	HTTPServer *http.Server
	Handler    *gin.Engine
	Config     *infrastructure.ApplicationEnvironment
}

type ApplicationServer struct {
	State ApplicationState
}

func (s *ApplicationServer) registerHandlers() {

	s.State.Handler.Static("/scripts", fmt.Sprintf("%sweb/scripts", s.State.Config.WebAssetsFolder))
	s.State.Handler.Static("/styles", fmt.Sprintf("%sweb/styles", s.State.Config.WebAssetsFolder))
	s.State.Handler.Static("/assets", fmt.Sprintf("%sweb/assets", s.State.Config.WebAssetsFolder))

	s.State.Handler.Handle(http.MethodGet, "/", s.HomePageHandler())

	s.State.Handler.NoRoute(s.err404PageHandler())

}

func NewApplicationServer(userOptions *ApplicationState) *ApplicationServer {
	state := userOptions
	if state == nil {
		state = &ApplicationState{}
	}

	if state.Config == nil {
		config := infrastructure.GetConfig()
		state.Config = &config
	}

	if state.Handler == nil {
		gin.SetMode(state.Config.ApplicationMode)

		state.Handler = gin.Default()
	}

	if state.HTTPServer == nil {
		state.HTTPServer = &http.Server{
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout:  100 * time.Second,
			Addr:         state.Config.BindAddr,
			Handler:      state.Handler,
		}
	}

	srv := ApplicationServer{
		State: ApplicationState{
			HTTPServer: state.HTTPServer,
			Handler:    state.Handler,
			Config:     state.Config,
		},
	}

	srv.registerHandlers()

	return &srv
}