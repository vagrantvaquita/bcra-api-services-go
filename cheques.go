package bcra

import (
	"errors"
	"log"
	"net/url"
)

var chequesAPI = &APIConfig{
	host: "https://api.bcra.gob.ar",
	path: "/cheques/v1.0/denunciados/%d/%d",
}

// Cheques issues the GET/cheques request and retrieves the Response
func (api *Client) Cheques(req *ChequesRequest) (*Cheque, error) {
	if req.CodigoEntidad == 0 {
		return nil, errors.New("bcra: CodigoEntidad is missing")
	}

	if req.NumeroCheque == 0 {
		return nil, errors.New("bcra: NumeroCheque is missing")
	}

	var res Response[Cheque]

	chequesAPI.AddPathValues(req.CodigoEntidad, req.NumeroCheque)

	if err := api.GetJson(chequesAPI, req, &res); err != nil {
		log.Fatal(err)
	}

	if err := res.StatusError(); err != nil {
		return nil, err
	}

	return &res.Results, nil
}

// ChequesRequest is the functional options struct for GET/cheques
type ChequesRequest struct {
	// CodigoEntidad is the code of the entity that issued the check.
	CodigoEntidad int32
	// NumeroCheque is the ID number on the check.
	NumeroCheque int32
}

func (v *ChequesRequest) params() url.Values {
	q := make(url.Values)
	return q
}
