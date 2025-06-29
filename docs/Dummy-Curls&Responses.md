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
<img width="898" alt="Screenshot 2025-06-22 at 12 01 21 PM" src="https://github.com/user-attachments/assets/f6aa81dc-fa71-4f19-b48b-822707ef6bf6" />

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
<img width="902" alt="Screenshot 2025-06-22 at 12 01 26 PM" src="https://github.com/user-attachments/assets/18591439-2dbc-4d41-bd53-a386e264debe" />

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
<img width="901" alt="Screenshot 2025-06-22 at 12 01 32 PM" src="https://github.com/user-attachments/assets/ef550926-1b44-4e71-96ef-52c21ccf2eb4" />

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
<img width="900" alt="Screenshot 2025-06-22 at 12 01 38 PM" src="https://github.com/user-attachments/assets/cf731604-d5a8-46a2-8a28-52b0f36e290e" />


