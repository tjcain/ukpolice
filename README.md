# ukpolice

[![Build Status](https://travis-ci.org/tjcain/ukpolice.svg?branch=master)](https://travis-ci.org/tjcain/ukpolice)

ukpolice is a Go client library for accessing the [data.police.uk api](https://data.police.uk/docs/)

## Usage

```go
import "github.com/tjcain/ukpolice"
```

Construct a new ukpolice client, then use the various services on the client to
access different parts of the data.police.uk API. It is recommended to pass in a http.Client with a longer timeout than default as some responses from the API can take over 60 seconds.

For example:

```go
customClient := http.Client{Timeout: time.Second * 120}
	client := ukpolice.NewClient(&customClient))

// list all available data sets.
avaliable, _, err := client.Avaliability.GetAvaliabilityInfo(context.Background())
```

Some API methods have optional parameters that can be passed. For example:

```go
searches, _, err := client.StopAndSearch.GetStopAndSearchesByForce(context.Background(),
        ukpolice.WithDate("2018-01"), ukpolice.WithForce("west-midlands"))
```

## Rate Limiting

The data.police.uk api sets a [rate limit of 15 requests per second](https://data.police.uk/docs/api-call-limits/). This limit is adhered to automatically by the package.

## Contributing

Contributions are always welcome.
