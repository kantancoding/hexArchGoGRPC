package rpc

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"hex-arch-go-grpc/internal/ports"
)

// Adapter is an rpc adapter the is compatible with
// the
type Adapter struct {
	api ports.APIPort
}

// NewAdapter creates a grpc entry adapter that is
// compatible with gRPCEntryPort
func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

// Run is the HTTP entrypoint for the application
func (httpa Adapter) Run() {
	http.HandleFunc("/add", httpa.addition)
	http.HandleFunc("/subtract", httpa.addition)
	http.HandleFunc("/multiply", httpa.addition)
	http.HandleFunc("/divide", httpa.addition)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", "8090"), nil); err != nil {
		log.Fatalf("Failed to serve HTTP server over port 8090: %v", err)
	}
}

type success struct {
	Answer int32 `json:"answer"`
}

func (httpa Adapter) addition(w http.ResponseWriter, req *http.Request) {
	a := req.FormValue("a")
	if a == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	b := req.FormValue("b")
	if b == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// only non negative numbers
	aInt64, err := strconv.ParseUint(a, 10, 32)
	bInt64, err := strconv.ParseUint(b, 10, 32)
	aInt32 := int32(aInt64)
	bInt32 := int32(bInt64)

	answer, err := httpa.api.GetAddition(aInt32, bInt32)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	status := success{Answer: answer}

	if err := json.NewEncoder(w).Encode(status); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}
