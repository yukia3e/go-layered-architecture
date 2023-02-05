package config

import (
	"crypto/tls"
	"os"
	"strings"
)

func GetServerCert() string {
	return strings.Replace(os.Getenv("SERVER_CRT"), `\n`, "\n", -1)
}

func GetServerKey() string {
	return strings.Replace(os.Getenv("SERVER_KEY"), `\n`, "\n", -1)
}

func GetServerTLSCredential() (*tls.Certificate, error) {
	if GetServerCert() == "" || GetServerKey() == "" {
		return nil, nil
	}
	cert, err := tls.X509KeyPair([]byte(GetServerCert()), []byte(GetServerKey()))
	if err != nil {
		return nil, err
	}
	return &cert, nil
}
