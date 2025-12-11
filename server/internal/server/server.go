package server

import (
	"fmt"
	"net/http"

	"github.com/chomba-mumba/proxy-server-ad-block/internal/configs"
)

// start server on defined port
func Run() error {
	//load config
	config, err := configs.NewConfiguration()
	if err != nil {
		return fmt.Errorf("could not laod configurations: %v", err)
	}

	// create new router
	router := http.NewServeMux()

	//health check endpoint
	router.HandleFunc("/health", health)

	router.HandleFunc("/proxy", ProxyHandler)

	//running proxy server
	fmt.Printf("running proxy server on port %s\n", config.Server.ListenPort)

	err = http.ListenAndServe(config.Server.Host+":"+config.Server.ListenPort, router)
	if err != nil {
		return fmt.Errorf("could not start the server: %v", err)
	}
	return nil
}
