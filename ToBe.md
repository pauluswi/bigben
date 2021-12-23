Get eWallet Balance
curl -X GET \
  http://localhost:3000/v1/ewallet/balance \
  -H 'authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE1MDc3NDM4MTgsImlhdCI6MTUwNzc0MjkxOCwic3ViIjoiNTlkZTUzZDVhYzM5ZmQ1ODQ3MGRjODI4In0.mUry4SFaWRqRrBmNF1RBBnJMvcvJBYAktqczpMj8r2w' \
  -H 'cache-control: no-cache' \
  -H 'postman-token: 6df0eb80-e0fc-5f47-4b72-2f3f165eeaaf'

Make a Deposit to your eWallet
curl -X POST \
  http://localhost:3000/v1/ewallet/deposit \
  -H 'authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE1MDc3NDM4MTgsImlhdCI6MTUwNzc0MjkxOCwic3ViIjoiNTlkZTUzZDVhYzM5ZmQ1ODQ3MGRjODI4In0.mUry4SFaWRqRrBmNF1RBBnJMvcvJBYAktqczpMj8r2w' \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/x-www-form-urlencoded' \
  -H 'postman-token: 66218aae-19ee-3761-e0c0-53823d0d4820' \
  -d 'amount=10&card=4111111111111111'
Note: You can simulate a Payment Rejected by the PaymentGateway using this card 4242424242424242

Get eWallet Transactions
curl -X GET \
  http://localhost:3000/v1/ewallet/transactions \
  -H 'authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE1MDc3NDM4MTgsImlhdCI6MTUwNzc0MjkxOCwic3ViIjoiNTlkZTUzZDVhYzM5ZmQ1ODQ3MGRjODI4In0.mUry4SFaWRqRrBmNF1RBBnJMvcvJBYAktqczpMj8r2w' \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/x-www-form-urlencoded' \
  -H 'postman-token: ff68cdff-9fac-9647-4594-70315ab1f4cd'
  
Make a Transfer to another eWallet
curl -X POST \
  http://localhost:3000/v1/ewallet/transfer \
  -H 'authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE1MDc3NjYyOTgsImlhdCI6MTUwNzc2NTM5OCwic3ViIjoiNTlkZWE4ZDA2YzkyYmQ2ZTdkZjZiMzMwIn0.PGSdiEpPG43ihnJldKFY-MMqNzaGb4PwOylUbA05AVY' \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/x-www-form-urlencoded' \
  -H 'postman-token: 78116228-4061-257a-f47c-37033d474596' \
  -d 'amount=100&destinationAccountNumber=1001'
Note: Every Transaction will generate a fee that will be discounted from the eWallet Balance and will be credited to the Master Account according to this table.

Amount | Percent | Fixed rate |---|---|---|---|---| x <= 1,000 | 3.0% | $8.00 1,000 > x <= 5,000 | 2.5% | $6.00 5,000 > x <= 10,000. | 2.0% | $4.00 10,000 > x | 1.0% | $3.00

Triggers a Withdrawal from your eWallet
curl -X POST \
  http://localhost:3000/v1/ewallet/withdrawal \
  -H 'authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE1MDc3NzI5NjAsImlhdCI6MTUwNzc3MjA2MCwic3ViIjoiNTlkZWE4ZDA2YzkyYmQ2ZTdkZjZiMzMwIn0.SF8OdwKfT-fiWbkhUgnTKWfyeZCY_p3ek4j2dPVukuc' \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/x-www-form-urlencoded' \
  -H 'postman-token: d2292d62-cefd-e7b9-311a-12fe92795c79' \
  -d 'amount=1500&card=4111111111111111'