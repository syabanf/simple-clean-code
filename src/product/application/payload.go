package application

import "sagara-test/src/product/domain/entity"

type VMProduct struct {
	GUID    string           `json:"guid"`
	Name    string           `json:"name"`
	SKU     string           `json:"sku"`
	Price   float64          `json:"price"`
	HPP     float64          `json:"hpp"`
	Picture []VMProductMedia `json:"picture"`
}

type VMProductMedia struct {
	GUID string `json:"guid"`
	URL  string `json:"url"`
}

func (data *VMProduct) ToEntity() (value entity.ModelProduct) {
	value = entity.ModelProduct{
		GUID:  data.GUID,
		Name:  data.Name,
		SKU:   data.SKU,
		Price: data.Price,
		HPP:   data.HPP,
	}
	return
}

func ToPayload(r entity.ModelProduct) VMProduct {
	return VMProduct{
		GUID:  r.GUID,
		Name:  r.GUID,
		SKU:   r.SKU,
		Price: r.Price,
		HPP:   r.HPP,
	}
}
