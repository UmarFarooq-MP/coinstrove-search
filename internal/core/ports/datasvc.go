package ports

import (
	"coinstrove-search/internal/core/domain/model"
)

type DataSVC interface {
	GetCoinDetails()
	GetListOfCoins()
	GetListOfExchanges()
	UpdateDB(message model.Exchange)
}
