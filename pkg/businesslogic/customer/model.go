package customer

import (
	"GameLoaders/pkg/businesslogic/task"
)

type Model struct {
	Id          int           `json:"id"`
	Money       float32       `json:"money"`
	Tasks       []*task.Task  `json:"tasks"`
	LoaderModel []interface{} `json:"loaders"`
}

//ToModel me - показать свои характеристики (деньги, зарегистрировавшиеся грузчики)
func (c *Customer) ToModel() interface{} {
	var loaderModel []interface{}
	for _, v := range c.loaders {
		loaderModel = append(loaderModel, v.ToModel())
	}
	return &Model{
		Id:          c.Id(),
		Money:       c.Wallet.GetInfo(),
		Tasks:       c.Tasks(),
		LoaderModel: loaderModel,
	}
}
