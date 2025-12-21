package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/chomba-mumba/proxy-server-ad-block/internal/configs"
)

// start server on defined port
func Run() error {
	//load config
	config, err := configs.NewConfiguration()
	if err != nil {
		return fmt.Errorf("[PROXY SERVER] could not load configurations: %v", err)
	}

	server := &http.Server{
		Addr:         fmt.Sprintf("localhost:%s", config.Server.ListenPort),
		Handler:      http.HandlerFunc(handleRequest),
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		IdleTimeout:  2 * time.Second,
	}

	//running proxy server
	fmt.Printf("[PROXY SERVER] running proxy server on port %s\n", config.Server.ListenPort)

	err = server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("[PROXY SERVER] could not start the server: %v", err)
	}
	return nil
}
