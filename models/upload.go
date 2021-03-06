package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/uploadexpress/app/constants"
)

type Upload struct {
	Id             string   `json:"id" bson:"_id,omitempty"`
	Name           string   `json:"name" bson:"name"`
	Files          []*File  `json:"files" bson:"files"`
	Backgrounds    []*Image `json:"backgrounds" bson:"backgrounds"`
	DownloadCount  int      `json:"download_count" bson:"download_count"`
	Public         bool     `json:"public" bson:"public"`
	Ready          bool     `json:"-" bson:"ready"`
	Gallery        bool     `json:"gallery" bson:"gallery"`
	RequestId      string   `json:"request_id" bson:"request_id"`
	Date           int64    `json:"date" bson:"date"`
	ExpirationDate *int64   `json:"expiration_date,omitempty" bson:"expiration_date,omitempty"`
}

func (upload *Upload) BeforeCreate() {
	// Assigns an ID to each file
	for _, file := range upload.Files {
		if file.Id == "" {
			file.Id = bson.NewObjectId().Hex()
		}
	}

	upload.Id = bson.NewObjectId().Hex()
	upload.Ready = false
	upload.Date = time.Now().Unix()
}

func (upload *Upload) Size() constants.ByteSize {
	var size int64 = 0
	for _, file := range upload.Files {
		size += file.Size
	}
	return constants.ByteSize(size)
}

const UploadsCollection = "uploads"
