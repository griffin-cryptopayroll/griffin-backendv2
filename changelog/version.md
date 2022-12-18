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