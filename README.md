# go-services-bos

BOS(Baidu Object Storage) service support for [go-storage](https://github.com/beyondstorage/go-storage).

## Notes

**This package has been moved to [go-storage](https://github.com/beyondstorage/go-storage/tree/master/services/bos).**

```shell
go get go.beyondstorage.io/services/bos/v2
```

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
