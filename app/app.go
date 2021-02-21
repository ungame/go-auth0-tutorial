package app

import (
	"encoding/gob"

	"github.com/gorilla/sessions"
)

var SecretKey = []byte("APP_SECRET_KEY")

var Store *sessions.FilesystemStore

func Init() error {
	Store = sessions.NewFilesystemStore("", SecretKey)
	gob.Register(map[string]interface{}{})
	return nil
}
