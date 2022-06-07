package health_controller

import (
	"fmt"
	"log"
	"net/http"
)

// HealthController displays stats about the api
type HealthController struct{}

func (c HealthController) Health(w http.ResponseWriter, _ *http.Request) {
	if _, err := fmt.Fprintf(w, "Welcome to the Fill Up Station\nCurrent Status: %s\n", "Healthy"); err != nil {
		log.Println(err)
	}
}
