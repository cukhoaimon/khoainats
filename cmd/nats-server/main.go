package main

import (
	"flag"
	"log"

	"github.com/nats-io/nats-server/v2/server"
)

func main() {
	var (
		host           string
		port           int
		certFile       string
		keyFile        string
		clientCertFile string
		clientKeyFile  string
		caFile         string
	)

	flag.StringVar(&host, "host", "localhost", "Client connection host/IP.")
	flag.IntVar(&port, "port", 4222, "Client connection port.")
	flag.StringVar(&certFile, "tls.cert.server", "cert.pem", "TLS cert file.")
	flag.StringVar(&keyFile, "tls.key.server", "key.pem", "TLS key file.")
	flag.StringVar(&caFile, "tls.ca", "ca.pem", "TLS CA file.")
	flag.StringVar(&clientCertFile, "tls.cert.client", "client-cert.pem", "TLS cert file.")
	flag.StringVar(&clientKeyFile, "tls.key.client", "client-key.pem", "TLS key file.")

	flag.Parse()

	serverTlsConfig, err := server.GenTLSConfig(&server.TLSConfigOpts{
		CertFile: certFile,
		KeyFile:  keyFile,
		CaFile:   caFile,
		Verify:   true,
		Timeout:  2,
	})
	if err != nil {
		log.Fatalf("tls config: %v", err)
	}

	opts := server.Options{
		Host:      host,
		Port:      port,
		TLSConfig: serverTlsConfig,
	}

	ns, err := server.NewServer(&opts)
	if err != nil {
		log.Fatalf("server init: %v", err)
	}

	go ns.Start()
	defer ns.Shutdown()
}
