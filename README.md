# Go Scryfall API Client

This package is auto-geneated from the [Scryfall API definition](https://github.com/jdharmon/scryfallapi).

## Getting Started

1. Get the scryfall package

```
$ go get github.com/jdharmon/scryfallapi-go
```

2. Edit main.go

```csharp
package main

import (
    "fmt"
    "context"
    "github.com/jdharmon/scryfallapi-go"
)

func main() {
    client := scryfall.NewCardsClient()
    ctx := context.TODO()
    card, err := client.GetRandom(ctx)
    if err != nil {
        panic(err)
    }
    fmt.Println("Got card:", *card.Name)
}
```

3. Run

```
$ go run main.go
```