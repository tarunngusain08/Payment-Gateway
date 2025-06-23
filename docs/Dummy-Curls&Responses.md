# Dummy CURL Requests & Responses

---

## Deposit

**Request**
```sh
curl --location 'http://localhost:8000/deposit' \
  --header 'Content-Type: application/json' \
  --data '{"account_id": "user123", "amount": 100.0}'
```

**Response**
```json
{
  "success": true
}
```

---

## Withdrawal

**Request**
```sh
curl --location 'http://localhost:8000/withdrawal' \
  --header 'Content-Type: application/json' \
  --data '{"account_id": "user123", "amount": 50.0}'
```

**Response**
```json
{
  "success": true
}
```

---

## Gateway A Callback (JSON)

**Request**
```sh
curl --location 'http://localhost:8000/callback/gateway-a' \
  --header 'Content-Type: application/json' \
  --data '{
    "transaction_id": "txn123",
    "status": "success",
    "metadata": {},
    "gateway_ref": "gwref123",
    "amount": 100.0,
    "currency": "USD",
    "timestamp": "2024-06-01T12:00:00Z"
  }'
```

**Response**
```json
{
  "status": "success",
  "message": "Successfully processed callback for transaction: txn123"
}
```

---

## Gateway B Callback (XML)

**Request**
```sh
curl --location 'http://localhost:8000/callback/gateway-b' \
  --header 'Content-Type: application/xml' \
  --data '<HandleCallbackRequest>
    <TransactionID>txn456</TransactionID>
    <Status>failed</Status>
    <GatewayRef>gwref456</GatewayRef>
    <Amount>50.0</Amount>
    <Currency>USD</Currency>
    <Timestamp>2024-06-01T12:05:00Z</Timestamp>
  </HandleCallbackRequest>'
```

**Response**
```xml
<HandleCallbackResponse>
  <Status>success</Status>
  <Message>Gateway B callback processed for transaction: txn456, ref: gwref456</Message>
</HandleCallbackResponse>
```

