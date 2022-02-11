### Create an inventory item:
`POST` to `https://shopifychallenge-338719.ue.r.appspot.com/items/create`

_Sample Payload_: 
```
{
    "name": "Beef",
    "description": "lean",
    "amount": 60
}
```

### Read inventory items:
`GET` to `https://shopifychallenge-338719.ue.r.appspot.com/items/read`

### Update inventory items:
`PATCH` to `https://shopifychallenge-338719.ue.r.appspot.com/items/edit`

_Sample Payload_: 
```
{
    "itemId": "1f709c78-7960-11ec-af55-a70c809bdb5e",
    "amount": 5
}
```

### Delete inventory items:
`DELETE` to `https://shopifychallenge-338719.ue.r.appspot.com/items/delete/:id`

_Sample Request_: `https://shopifychallenge-338719.ue.r.appspot.com/items/delete/0c022abe-7960-11ec-af55-a70c809bdb5e`

## Extra Requirement: Shipments
---

### Create a shipment
`POST` to `https://shopifychallenge-338719.ue.r.appspot.com/shipments/create`

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
`GET` to `https://shopifychallenge-338719.ue.r.appspot.com/shipments/read`
