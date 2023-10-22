package main

import (
	"io"
	"log"
	"net/http"
	"os"
)


var GOOGLE_METADATA_URL = "http://metadata.google.internal/computeMetadata/v1/instance"


func parseMetadataResponse(w http.ResponseWriter, resp *http.Response, err error){

	if err != nil {
		//Handle Error
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`"gcloud error": ` + err.Error()))
	} else {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			//Handle Error
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`"error while parsing gcloud response": ` + err.Error()))
		} else {
			// Set the return Content-Type as JSON like before
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(string(data)))
		}
	}
}


func getInstanceExternalIp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			gcp_client := &http.Client{}
			req, _ := http.NewRequest("GET", GOOGLE_METADATA_URL +
				 "/network-interfaces/0/access-configs/0/external-ip", nil)
			req.Header.Set("Metadata-Flavor", "Google")
			resp, err := gcp_client.Do(req)
			parseMetadataResponse(w, resp, err)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Invalid method requested"))
	}
}


func getInstanceName(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			gcp_client := &http.Client{}
			req, _ := http.NewRequest("GET", GOOGLE_METADATA_URL +
				"/name", nil)
			req.Header.Set("Metadata-Flavor", "Google")
			resp, err := gcp_client.Do(req)
			parseMetadataResponse(w, resp, err)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Invalid method requested"))
	}
}


// Implement health interface
func getHealth (w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Invalid method requested"))
	}
}


func main() {

	var bind_port = "1000"

	arguments := os.Args
	if len(arguments) > 1 {
		bind_port = arguments[1]
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", getHealth)
	mux.HandleFunc("/instance/name", getInstanceName)
	mux.HandleFunc("/instance/external-ip", getInstanceExternalIp)

	log.Println("Service is listening on localhost:" + bind_port)
	http.ListenAndServe(":" + bind_port, mux)

}
