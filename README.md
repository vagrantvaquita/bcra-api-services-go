Go Client for BCRA API Web services
==================================

[![GoDoc](https://godoc.org/googlemaps.github.io/maps?status.svg)](https://godoc.org/googlemaps.github.io/maps)
[![Go Report Card](https://goreportcard.com/badge/github.com/googlemaps/google-maps-services-go)](https://goreportcard.com/report/github.com/googlemaps/google-maps-services-go)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/googlemaps/google-maps-services-go)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

## Description

Use Go? This library brings many [BCRA Web Services APIs] to your Go application.

The Go Client for Google Maps Services is a Go Client library for the following Google Maps Platform
APIs:

- [Entidades API]
- [Cheque API]
- [Datos Variable API]
- [Principales Variables API]

> [!TIP]
> See the [Google Maps Platform Cloud Client Library for Go](https://github.com/googleapis/google-cloud-go/tree/main/maps) for our newer APIs
> including Address Validation API, Datasets API, Fleet Engine, new Places API, and Routes API.

## Requirements

- Go 1.7 or later.
- No [API key] needed currently.

## Installation

To install the Go Client for Google Maps Services, please execute the following `go get` command.

```bash
go get vagrantvaquita.github.io/go/bcra
```

## Documentation

View the [reference documentation](https://godoc.org/googlemaps.github.io/maps).

Additional documentation about the APIs is available at:

- [Entidades API]
- [Cheque API]
- [Datos Variable API]
- [Principales Variables API]

## Usage

Sample usage of the Directions API with an API key:

```go
package main

import (
	"log"

    "github.com/kr/pretty"
	"vagrantvaquita.github.io/go/bcra"
)

func main() {
	c, err := NewClient()
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	req := &ChequesRequest{
		CodigoEntidad: 11,
		NumeroCheque:  20377516,
	}
	res, err := c.Cheques(req)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	pretty.Println(route)
}
```