# Peerless Animate API
[![Travis Status](https://api.travis-ci.org/ubogdan/peerless.svg?branch=master)](https://travis-ci.org/ubogdan/peerless)
[![Go Report Card](https://goreportcard.com/badge/github.com/ubogdan/peerless)](https://goreportcard.com/report/github.com/ubogdan/peerless)
[![Go Doc](https://godoc.org/github.com/ubogdan/peerless?status.svg)](https://godoc.org/github.com/ubogdan/peerless)

## Library status
Work in progress ... 

## Examples

```go
api := New(StagingEndpoint, "Customer", "XXXXXX", "user@domain.com")
```

New number order
----------------
```go


```
Get order status
```go
res, err := api.GetOrderStatus("110001")if err != nil {
	return error
}
```

Update number
----------------
```go

```