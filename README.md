### Create an inventory item:
`POST` to `/items/create`

_Sample Payload_: 
```
{
    "name": "Beef",
    "description": "lean",
    "amount": 60
}
```

### Read inventory items:
`GET` to `/items/read`

### Update inventory items:
`PATCH` to `/items/edit`

_Sample Payload_: 
```
{
    "itemId": "1f709c78-7960-11ec-af55-a70c809bdb5e",
    "amount": 5
}
```

### Delete inventory items:
`DELETE` to `/items/delete/:id`

_Sample Request_: `/items/delete/0c022abe-7960-11ec-af55-a70c809bdb5e`

## Extra Requirement: Shipments
---

### Create a shipment
`POST` to `/shipments/create`

_Sample Payload_: 
```
{
    "address": "Downtown Montreal",
    "contents": [
        {
            "itemId": "2ae314a7-7960-11ec-af55-a70c809bdb5e",
            "amount" : 10
        },
        {
            "itemId": "338bba99-7960-11ec-af55-a70c809bdb5e",
            "amount" : 5
        }
    ]
}
```

### Read shipments:
`GET` to `/shipments/read`
