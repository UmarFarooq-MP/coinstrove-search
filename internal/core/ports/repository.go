package ports

type DataSVCRepository interface {
	GetCoinByName()
	GetListOfExchanges()
	GetListOfCoins()
	UpdateDB()
}
