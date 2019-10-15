# go-cisco-webex-teams #
go-cisco-webex-teams is a Go client library for the [Cisco Webex Teams API](https://developer.webex.com/index.html).

## Usage ##

```go
import	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"
```

## Documentation

https://godoc.org/github.com/jbogarin/go-cisco-webex-teams/sdk

## Changes

* 2019-08-12: **Tag v0.2.0**: _Breaking change_, moved from resty v1 to resty v2. Include paginate option in List query params
* 2019-09-10: **Tag v0.3.0**: _Breaking change_, removed complexity from client, resty is a dependency for the library but it is not longer necessary to import it in the code using the SDK.
* 2019-10-14: **Tag v0.4.0**: _Breaking change_, added Go modules functionality

## Authorization Token

Authorization token can be defined in environment variable as WEBEX_TEAMS_ACCESS_TOKEN or within the code:

```go
Client = webexteams.NewClient()
Client.SetAuthToken("<WEBEX TEAMS TOKEN>")
```

## TODO ##

1. Documentation
  * In the code files
  * In the README
2. Examples
3. Testing

## License ##

This library is distributed under the MIT license found in the [LICENSE](./LICENSE) file.

## Inspiration ##

This library is inspired by the following ones:

- [godo](https://github.com/digitalocean/godo)
- [go-github](https://github.com/google/go-github)

