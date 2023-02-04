## Pre v2.1.4

Versions v2.1 to v2.1.4 is stored under ./changelog/old

## v2.2.0

Changes in Griffin backend changelog style. 
Schema changed to handle dashboard information. 

### Main feature
- Add `payment` table(entity) and its edges. There are 3 sorts of payment ( [4b83278]() )
  - scheduled payments
  - executed payments.
  - one-off payments: payments made outside of the regular payment schedule.
    - They don't have a `payment_scheduled` value
- Add CRUD operation for 3 sorts of payment ( [ca3b7db]() )
  - `UpdatePaymentExecuted`
  - `CreatePaymentScheduled`
  - `CreatePaymentOneOff`
  - Create Payment Schedules whenever the employee is created.
    - `CreatePermanent` 
    - `CreateFreelance`
    - Both Use `CreatePaymentScheduled`
  
## v2.2.1

Change in Griffin Swagger document

### Main feature
- Swagger endpoint added for 4 new points ( [84f2628]() )
  - `/payment/execute` PUT
  - `/payment/oneoff` POST
  - `/payment/employee` GET
  - `/payment/employer` GET
- Unused payment related endpoint docs deleted

## v2.2.2

Fix cycle import. 

### BUGFIX
- Fix import cycle. ( [a1d27d6]() )

## v2.2.3

### Sub features

Fix swagger document. ( [4604618]() ) 
- Delete `/payment/future`, `/payment/made` etc.

Fix Price API. ( [b825477]() )
- Add `REGION` to environment file. If region is US, use Binance US api ending in `.us` not `.com`.

## v2.2.4

### Main features

Add future, past, missed payment for employee. ( [2aba1d8]() )
- `payment/future`, `payment/past` `payment/miss` respectively. 
- `func parseInterval` for parsing interval for `future` and `past`.

### Sub features
Add employee data to be displayed in payment function.- `QueryPaymentEmployer` ( [b3af2e8]() )

Swag document update. ( [be960ba]() )

## v2.2.5

### Main features

Add currency to API Payment return ( [6a36395]() )

Add payment frequency month ( [c819bc7]() )
- Add month validation. The date of the month should be under 25th - sine some month such as Feb does not have 30, and 31th. ( [c4cf6f9](), [e268f1f]() )

Delete USDT from currency ticker. ex) MATICUSDT -> MATIC


## v2.3.5

### Main features

- Adhere to <b>Common Backend Design</b> rules. Change api directories. Edit function result. ( [6d9dbf3](https://github.com/griffin-cryptopayroll/griffin-backendv2/commit/6d9dbf3538695cc26ed21ef3424b2c7737dff31d) )

### Bugs and fixes

- `GenerateEmployee` function in `api_v0` was not returning after giving error processing failed. Now return `api_base.CommonResponse` when failed. ( [6d9dbf3](https://github.com/griffin-cryptopayroll/griffin-backendv2/commit/6d9dbf3538695cc26ed21ef3424b2c7737dff31d) ) 