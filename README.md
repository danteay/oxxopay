# OxxoPay Go

## Installing

```bash
go get https://github.com/danteay/oxxopay
```

## Importing

```go
import op "github.com/danteay/oxxopay"
```

## Usage

### Configuring

```go
privateApiKey := "key_asdRfdasASfadf"

client := new(op.OpClient)
client.Init(privateApiKey)
```

### CreateOrder

```go
data := op.RequestData{
  "line_items": []op.RequestData{
    op.RequestData{
      "name":       "La divina comedia",
      "unit_price": 12345,
      "quantity":   1,
    },
  },
  "currency": "MXN",
  "customer_info": op.RequestData{
    "name":  "Dante Aligeri",
    "email": "dante@inferno.com",
    "phone": "+525561463627",
    "charges": []op.RequestData{
      op.RequestData{
        "payment_method": op.RequestData{
          "type": "oxxo_cash",
        },
      },
    },
  },
}

response, errReq := client.CreateOrder(data)

fmt.Println(err)
fmt.Println(string(response))
```

### Get Order Info

```go
orderId := "ord_bjhbkjntgfouvasd"

response, errReq := client.GetOrder(orderId)

fmt.Println(err)
fmt.Println(string(response))
```

### Get List Orders

```go
response, errReq := client.GetListOrders(orderId)

fmt.Println(err)
fmt.Println(string(response))
```