package utils

import (
	"crypto/rsa"
	"fmt"
	"net"

	"github.com/18F/cf-cdn-service-broker/lego/acme"
)

func preCheckDNS(fqdn, value string) (bool, error) {
	record, err := net.LookupTXT(fqdn)
	if err != nil {
		return false, err
	}
	if len(record) == 1 && record[0] == value {
		return true, nil
	}
	return false, fmt.Errorf("DNS precheck failed on name %s, value %s", fqdn, value)
}

func init() {
	acme.PreCheckDNS = preCheckDNS
}

type User struct {
	Email        string
	Registration *acme.RegistrationResource
	OrderURL     string
	key          rsa.PrivateKey
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRegistration() *acme.RegistrationResource {
	return u.Registration
}

func (u *User) GetPrivateKey() *rsa.PrivateKey {
	return &u.key
}

func (u *User) SetPrivateKey(key rsa.PrivateKey) {
	u.key = key
}
