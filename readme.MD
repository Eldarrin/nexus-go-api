# Sonatype Nexus Go API

Based on their swagger 2.0 doc for the Nexus Go API. Version: 3.40.1

## Installation

```bash
go get -u github.com/Eldarrin/nexus-go-api
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/Eldarrin/nexus-go-api/client"
    "github.com/Eldarrin/nexus-go-api/models"
)

func main() {

    resp, err := client.Default().GetRepositories()
    if err != nil {
        fmt.Println(err)
    }

}
```

To generate your own client for internal use; use the following `create.sh`


## Examples


## License


## Authors



To