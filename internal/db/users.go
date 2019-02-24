package db

import (
	"fmt"

	"github.com/sauerbraten/maitred/pkg/auth"
)

type User struct {
	Name      string         `json:"name"`
	PublicKey auth.PublicKey `json:"public_key"`
}

func (db *Database) AddUser(name, pubkey string) error {
	_, err := db.Exec("insert into `users` (`name`, `pubkey`) values (?, ?)", name, pubkey)
	if err != nil {
		return fmt.Errorf("db: error inserting '%s' (%s) into database: %v", name, pubkey, err)
	}
	return nil
}

func (db *Database) GetPublicKey(name string) (pubkey auth.PublicKey, err error) {
	var _pubkey string
	err = db.
		QueryRow("select `pubkey` from `users` where `name` = ?", name).
		Scan(&_pubkey)
	if err != nil {
		err = fmt.Errorf("db: error retrieving public key of '%s' from database: %v", name, err)
		return
	}

	pubkey, err = auth.ParsePublicKey(_pubkey)
	if err != nil {
		err = fmt.Errorf("db: error parsing public key '%s' from database: %v", _pubkey, err)
	}
	return
}
