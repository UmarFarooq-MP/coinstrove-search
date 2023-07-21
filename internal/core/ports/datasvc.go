package ports

type DataSVC interface {
	GetCoinDetails()
	GetListOfCoins()
	GetListOfExchanges()
	UpdateDB()
}
