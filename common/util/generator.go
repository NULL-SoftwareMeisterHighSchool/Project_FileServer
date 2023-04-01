package util

import "github.com/gofrs/uuid"

var uuidGen *uuid.Gen = uuid.NewGen()

func CreateUUID() string {
	_uuid, err := uuidGen.NewV4()
	if err != nil {
		panic(err)
	}
	return _uuid.String()
}
