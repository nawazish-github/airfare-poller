package database

import (
	"log"

	"github.com/nawazish-github/airfare-poller/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DatabaseClient struct {
	connection *mgo.Collection
}

func (client *DatabaseClient) DialDBAt(ipAddr string, database string, schema string) (*mgo.Collection, error) {
	if client == nil {
		log.Output(10, "db client is nilllllll")

	}
	session, dbConErr := mgo.Dial(ipAddr)
	if dbConErr != nil {
		return nil, dbConErr
	}

	client.connection = session.DB(database).C(schema)
	return client.connection, nil
}

func (client *DatabaseClient) GetDataFor(srcDest string) ([]models.Response, error) {
	var results []models.Response
	fetchError := client.connection.Find(bson.M{"srcdest": srcDest}).All(&results)
	return results, fetchError
}

func (client *DatabaseClient) UpdateDataFor(srcDest string, data *models.AirfareResponse) error {
	data.SrcDest = srcDest
	updateError := client.connection.Insert(data)
	return updateError
}
