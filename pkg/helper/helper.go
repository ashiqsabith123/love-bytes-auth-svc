package helper

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"os"

	"google.golang.org/grpc/credentials"
)

func GetCertificates(ca_cert, server_cert, server_key string) (credentials.TransportCredentials, error) {
	caPem, err := os.ReadFile(ca_cert)
	if err != nil {
		return nil, errors.Join(errors.New("error while reading ca_cert file"), err)
	}

	// create cert pool and append ca's cert
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caPem) {
		return nil, errors.Join(errors.New("error while creating pool"), err)
	}

	// read server cert & key
	serverCert, err := tls.LoadX509KeyPair(server_cert, server_key)
	if err != nil {
		return nil, errors.Join(errors.New("error while loading 509 key pair"), err)
	}

	// configuration of the certificate what we want to
	conf := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	//create tls certificate
	tlsCredentials := credentials.NewTLS(conf)

	return tlsCredentials, nil
}
