package main

import (
	"log"

	"github.com/versent/saml2aws/v2/cmd/saml2aws/commands"
	"github.com/versent/saml2aws/v2/pkg/flags"
)

var (
	// Version app version
	Version = "1.0.0"
	errtpl  = "%v\n"
)

func main() {
	err := saml2aws_do_login()
	if err != nil {
		log.Println("There was an error authenticating to AWS")
		log.Printf(errtpl, err)
	}
}

func saml2aws_do_login() (err error) {
	loginFlags := new(flags.LoginExecFlags)
	commonFlags := new(flags.CommonFlags)
	commonFlags.IdpAccount = "default"
	loginFlags.CommonFlags = commonFlags
	err = commands.Login(loginFlags)
	return
}
