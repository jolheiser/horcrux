//go:generate go run horcrux.example.go
//go:generate go run router/templates/templates.go
//go:generate go fmt ./...

package main

import (
	"fmt"
	"net/http"

	"go.jolheiser.com/horcrux/config"
	"go.jolheiser.com/horcrux/router"

	"go.jolheiser.com/beaver"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		beaver.Fatal(err)
	}
	port := fmt.Sprintf(":%s", cfg.Port)
	beaver.Infof("horcrux is listening at http://localhost%s", port)
	if err := http.ListenAndServe(port, router.New(cfg)); err != nil {
		beaver.Fatal(err)
	}
}
