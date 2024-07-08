package bcra

import (
	"testing"
)

func TestEntidades(t *testing.T) {

	client, err := NewClient()
	client.Debug = true

	if err != nil {
		t.Error(err)
	}

	res, err := client.Entidades()

	t.Logf("[RESPONSE] %+v", res)

	if err != nil {
		t.Errorf("r.Get returned non nil error, was %+v", err)
	}
}

func TestCheques(t *testing.T) {

	client, err := NewClient()
	client.Debug = true

	if err != nil {
		t.Error(err)
	}

	req := &ChequesRequest{
		CodigoEntidad: 11,
		NumeroCheque:  20377516,
	}

	res, err := client.Cheques(req)

	t.Logf("[RESPONSE] %+v", res)

	if err != nil {
		t.Errorf("r.Get returned non nil error, was %+v", err)
	}
}
