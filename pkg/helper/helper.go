package helper

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"os"

	"github.com/fatih/color"
	"google.golang.org/grpc/credentials"
)

var bold = color.Bold

var Green = color.New(color.FgGreen, bold).SprintFunc()
var Blue = color.New(color.FgBlue, bold).SprintFunc()
var Red = color.New(color.FgRed, bold).SprintFunc()
var Yellow = color.New(color.FgYellow, bold).SprintFunc()

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
