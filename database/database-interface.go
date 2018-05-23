package database

import (
	"github.com/nawazish-github/airfare-poller/models"
	"gopkg.in/mgo.v2"
)

//refactor to interface type ... interface{}
type IDatabase interface {
	GetDataFor(srcDest string, database string, schema string) ([]models.Response, error)
	UpdateDataFor(srcDest string, data []*models.AirfareResponse) error
	DialDBAt(ipAddr string) (*mgo.Collection, error)
}
