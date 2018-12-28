package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:9099", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	svc_port := r.URL.Query().Get("svc_port")
	svc_ip := r.URL.Query().Get("svc_ip")
	fmt.Sprintf("svc url is http://%s:%s", svc_ip, svc_port)
	if len(svc_ip) == 0 {
		return
	}
	_, err := http.Get(fmt.Sprintf("http://%s:%s", svc_ip, svc_port))

	if err != nil {
		fmt.Fprint(w, "unReady")
	}else {
		fmt.Fprint(w, "Ready")
	}

}
