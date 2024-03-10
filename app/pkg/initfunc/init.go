package initfunc

import (
	"OCsemmerApp/pkg/model"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func InitUpdatePassword() {
	udata := map[string]string{
		"0000": "guest",
		"0001": "Microsoft",
		"0010": "Bitcoin",
		"0011": "personalcomputer",
		"0100": "iPhone",
		"0101": "enigma",
		"0110": "linux",
		"0111": "ruby",
		"1000": "ryzen",
		"1001": "shimomura",
		"1010": "spiderman",
	}
	for id, password := range udata {
		hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
		if err != nil {
			log.Panicln(err)
		}

		if err := model.UpdatePass(id, string(hashed)); err != nil {
			log.Panicln(err)
		}
	}
}
