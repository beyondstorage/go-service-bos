[![Build Status](https://github.com/beyondstorage/go-service-bos/workflows/Unit%20Test/badge.svg?branch=master)](https://github.com/beyondstorage/go-service-bos/actions?query=workflow%3A%22Unit+Test%22)
[![Integration Tests](https://teamcity.beyondstorage.io/app/rest/builds/buildType:(id:Services_Bos_IntegrationTests)/statusIcon)](https://teamcity.beyondstorage.io/buildConfiguration/Services_Bos_IntegrationTests)
[![License](https://img.shields.io/badge/license-apache%20v2-blue.svg)](https://github.com/Xuanwo/storage/blob/master/LICENSE)
[![](https://img.shields.io/matrix/beyondstorage@go-service-bos:matrix.org.svg?logo=matrix)](https://matrix.to/#/#beyondstorage@go-service-bos:matrix.org)

# go-services-bos

BOS(Baidu Object Storage) service support for [go-storage](https://github.com/beyondstorage/go-storage).

## Install

```go
go get github.com/beyondstorage/go-service-bos
```

## Usage

```go
import (
	"log"

	_ "github.com/beyondstorage/go-service-bos/"
	"github.com/beyondstorage/go-storage/v4/services"
)

func main() {
	store, err := services.NewStoragerFromString("bos://bucket_name/path/to/workdir")
	if err != nil {
		log.Fatal(err)
	}

	// Write data from io.Reader into hello.txt
	n, err := store.Write("hello.txt", r, length)
}
```

- See more examples in [go-storage-example](https://github.com/beyondstorage/go-storage-example).
- Read [more docs](https://beyondstorage.io/docs/go-storage/services/bos) about go-service-bos.
