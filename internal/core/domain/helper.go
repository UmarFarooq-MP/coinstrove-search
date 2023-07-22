package domain

import (
	"coinstrove-search/internal/core/domain/dto"
	"coinstrove-search/internal/core/domain/model"
	"strings"
)

func ConvertUpdateDtoToModelObject(message dto.Message) model.Exchange {
	exchange := model.Exchange{}
	exchange.ExchangeName = strings.ToUpper(message.Data.ExchangeName)
	for _, value := range message.Data.Currencies {
		exchange.Currencies = append(exchange.Currencies, model.Currency{Name: value.Name, Price: value.Price})
	}
	return exchange
}
