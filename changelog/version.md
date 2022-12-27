## Pre v2.1.4

Versions v2.1 to v2.1.4 is stored under ./changelog/old

## v2.2.0

Changes in Griffin backend changelog style. 
Schema changed to handle dashboard information. 

## Main feature
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

## Main feature
- Swagger endpoint added for 4 new points ( [84f2628]() )
  - `/payment/execute` PUT
  - `/payment/oneoff` POST
  - `/payment/employee` GET
  - `/payment/employer` GET
- Unused payment related endpoint docs deleted

## v2.2.2

Fix cycle import. 

## BUGFIX
- Fix import cycle. ( [a1d27d6]() )

## v2.2.3

Fix swagger document. ( [4604618]() ) 
- Delete `/payment/future`, `/payment/made` etc.

Fix Price API. ( [b825477]() )
- Add `REGION` to environment file. If region is US, use Binance US api ending in `.us` not `.com`.