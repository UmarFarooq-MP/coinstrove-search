package http

import (
	"coinstrove-search/internal/core/domain"
	"coinstrove-search/internal/core/domain/dto"
	"coinstrove-search/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type DataSVCHandler struct {
	dataService []ports.DataSVC
}

func NewDataSVCHandler(dataService []ports.DataSVC) *DataSVCHandler {
	return &DataSVCHandler{dataService: dataService}
}

func (handler *DataSVCHandler) GetCoinDetails(ctx *gin.Context) {

}

// UpdateCoinInfo will be invoked whenever there is new message via queue
func (handler *DataSVCHandler) UpdateCoinInfo(message dto.Message) {
	if message.ErrorMessage == "" {
		handler.dataService[0].UpdateDB(domain.ConvertUpdateDtoToModelObject(message))
	}
}
