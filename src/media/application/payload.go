package application

import "sagara-test/src/media/domain/entity"

type VMMedia struct {
	GUID string `json:"guid"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func (data *VMMedia) ToEntity() (value entity.ModelMedia) {
	value = entity.ModelMedia{
		GUID: data.GUID,
		Name: data.Name,
		Type: data.Type,
	}
	return
}

func ToPayload(r entity.ModelMedia) VMMedia {
	return VMMedia{
		GUID: r.GUID,
		Name: r.GUID,
		Type: r.Type,
	}
}
