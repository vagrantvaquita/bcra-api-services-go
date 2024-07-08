package bcra

import (
	"log"
)

var entidadesAPI = &APIConfig{
	host: "https://api.bcra.gob.ar",
	path: "/cheques/v1.0/entidades",
}

// Entidades issues the GET/Entidades request and retrieves the Response
func (api *Client) Entidades() ([]Entidad, error) {

	var res Response[[]Entidad]

	if err := api.GetJson(entidadesAPI, nil, &res); err != nil {
		log.Fatal(err)
	}

	if err := res.StatusError(); err != nil {
		return nil, err
	}

	return res.Results, nil
}
