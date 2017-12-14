package database

import "gopkg.in/mgo.v2"

func connect() (*Session, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return err
	}

	return session
}
