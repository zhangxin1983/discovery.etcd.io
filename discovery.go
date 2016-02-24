package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"
	"github.com/coreos/go-systemd/activation"

	_ "github.com/coreos/discovery.etcd.io/http"
)

var addr = flag.String("addr", "", "web service address")

func main() {
	var ETCD_CONN = os.Getenv("ETCD_CONN")
	var BASE_URL = os.Getenv("BASE_URL")
	if( ETCD_CONN == "" || BASE_URL == "" ){
		panic("Need envronment ETCD_CONN and BASE_URL")
	}
	
	if( !strings.Contains(ETCD_CONN,"http://") ){
		os.Setenv("ETCD_CONN", "http://" + ETCD_CONN)
	}
	
	if( !strings.Contains(BASE_URL,"http://") ){
		os.Setenv("BASE_URL", "http://" + BASE_URL)
	}
	
	log.SetFlags(0)
	flag.Parse()

	if *addr != "" {
		http.ListenAndServe(*addr, nil)
	}

	listeners, err := activation.Listeners(true)
	if err != nil {
		panic(err)
	}

	if len(listeners) != 1 {
		panic("Unexpected number of socket activation fds")
	}

	http.Serve(listeners[0], nil)
}
