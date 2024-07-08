package bcra

import (
	"fmt"
)

// commonResponse contains the common response fields to most API calls inside
// the BCRA APIs. This is used internally.
type commonResponse struct {
	// Status contains the status of the request.
	Status int `json:"status"`

	// ErrorMessages is the explanatory field added when Status is an error.
	ErrorMessages []string `json:"errorMessages"`
}

// ChequeDetalle represents details of the check.
type ChequeDetalle struct {
	DenominacionEntidad string `json:"denominacionEntidad"`
	Sucursal            int    `json:"sucursal"`
	NumeroCuenta        int    `json:"numeroCuenta"`
	Causal              string `json:"causal"`
}

// Cheque represents the status of the check.
type Cheque struct {
	Denunciado         bool          `json:"denunciado"`
	FechaProcesamiento string        `json:"fechaProcesamiento"`
	Detalle            ChequeDetalle `json:"detalle"`
}

// Entidad represents details for a single entity.
type Entidad struct {
	CodigoEntidad int    `json:"codigoEntidad"`
	Denominacion  string `json:"denominacion"`
}

// PrincipalesVariables represents details for a single variable.
type PrincipalesVariables struct {
	IdVariable  int     `json:"idVariable"`
	CdSerie     int     `json:"cdSerie"`
	Descripcion string  `json:"descripcion"`
	Fecha       string  `json:"fecha"`
	Valor       float32 `json:"valor"`
}

// DatosVariable represents details for a single variable.
type DatosVariable struct {
	IdVariable int     `json:"idVariable"`
	Fecha      string  `json:"fecha"`
	Valor      float32 `json:"valor"`
}

// Response contains all field from an API response.
type Response[T any] struct {
	Results T `json:"results"`
	commonResponse
}

// StatusError returns an error if this object has a Status different
// from 200 or ErrorMessages is not nil.
func (c *commonResponse) StatusError() error {
	if c.Status != 200 && c.ErrorMessages != nil {
		return fmt.Errorf("bcra: %d - %s", c.Status, c.ErrorMessages)
	}
	return nil
}
