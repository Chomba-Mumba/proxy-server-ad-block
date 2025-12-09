package server

import (
	"fmt"
	"net/http"
	"net/url"

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
	mux := http.NewServeMux()

	//health check endpoint
	mux.HandleFunc("/health", health)
	//fmt.Printf("erererer %v", config)
	// regsiter config resource and register them into router
	for _, resource := range config.Resources {
		url, _ := url.Parse(resource.DestinationURL)
		proxy := NewProxy(url)
		mux.HandleFunc(resource.Endpoint, ProxyRequestHandler(proxy, url, resource.Endpoint))
	}

	fmt.Printf("running proxy server on port %s", config.Server.ListenPort)
	//running proxy server
	err = http.ListenAndServe(config.Server.Host+":"+config.Server.ListenPort, mux)
	if err != nil {
		return fmt.Errorf("could not start the server: %v", err)
	}
	return nil
}
