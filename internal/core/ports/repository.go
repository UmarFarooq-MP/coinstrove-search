package ports

import (
	"coinstrove-search/internal/core/domain/model"
)

type DataSVCRepository interface {
	GetCoinByName()
	GetListOfExchanges()
	GetListOfCoins()
	UpdateDB(message model.Exchange)
}
