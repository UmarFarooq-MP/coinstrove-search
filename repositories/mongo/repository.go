package mongo

import (
	"coinstrove-search/consts"
	"coinstrove-search/internal/core/domain/model"
	"coinstrove-search/internal/core/ports"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type mongoRepository struct {
	Client   *mongo.Client
	Database string
	Timeout  time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewMongoRepository(mongoURL, mongoDB string, mongoTimeout int) (ports.DataSVCRepository, error) {
	repo := &mongoRepository{
		Timeout:  time.Duration(mongoTimeout) * time.Second,
		Database: mongoDB,
	}
	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		log.Fatalf("Failed to connect DB with message %v", err)
	}
	repo.Client = client
	return repo, nil
}

func (mongorepo *mongoRepository) GetCoinByName()      {}
func (mongorepo *mongoRepository) GetListOfExchanges() {}
func (mongorepo *mongoRepository) GetListOfCoins()     {}
func (mongorepo *mongoRepository) UpdateDB(message model.Exchange) {
	//update the latest first before inserting
	updateTheLates(message.ExchangeName, mongorepo.Client)

	switch message.ExchangeName {
	case consts.BITFINEX:
		updateDBHelper(consts.BITFINEX, mongorepo.Client, message)
	case consts.GATEIO:
		updateDBHelper(consts.GATEIO, mongorepo.Client, message)
	case consts.KRAKEN:
		updateDBHelper(consts.KRAKEN, mongorepo.Client, message)
	case consts.COINBASE:
		updateDBHelper(consts.COINBASE, mongorepo.Client, message)
	case consts.BITPAY:
		updateDBHelper(consts.BITPAY, mongorepo.Client, message)
	case consts.BINANCE:
		updateDBHelper(consts.BINANCE, mongorepo.Client, message)
	case consts.HUOBI:
		updateDBHelper(consts.HUOBI, mongorepo.Client, message)
	case consts.BITSTAMP:
		updateDBHelper(consts.BITSTAMP, mongorepo.Client, message)
	case consts.KUCOIN:
		updateDBHelper(consts.KUCOIN, mongorepo.Client, message)
	}
}
