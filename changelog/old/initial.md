# CHANGE LOG

## Creating Backend V2

### Initial Commit ~ DAO Creation

`7082f0a`

- initial commit

`0131cad`

- add all the edges and entity, gen CRUD Op functions


### API Connection ~ Deployment

`a3dfcce`

- RESTful API Connection

`08f34144`

- add delete operation to gcrud

`9fb38dab`

- add employee api connection works + changes in parameter handling

`565f61af`

- add query methods to gcrud

`807909eb` , `21b5b31c`

- finish employee api connection

### Deployment ~ Testing

`768d642`

- query with contract month added

`9b013e2`

- ping endpoint will test database ping

`617e858`

- swaggo adding for api testing, add new crud ops

`46c93ec`

- edit CRUD ops. add more options variable

`06dccba`

- validity checker for the contract period.

`e2a546b` `1ce08aa`

- swaggo docs added

`d53e083`

- change database schema. employ_type lost contract start. employer has work_start

`9d3814f`

- Redis caching mechanism added.

`8ee8517` `a3cd3da`

- New CRUD ops.

`27b43d8` `29cb81b`

- Unique constraint to Griffin ID (GID) for employee and employers.

`6c453fb`

- Change mysql schema.

`932819c`, `b3df707`

- Entity query now covers recover. Add errors to return. 
  - Such as - creating entity did not return error. Now it returns errors

`60c91bc` `c8bb617`

- Swagger Document job, especially `c8bb617`
- Adjust code to embrace changes made in `932819c`
  - error managing added.

`f6075fa`

- change VARCHAR(k) to VARCHAR(200). VARCHAR(k) should be able to get UUID.

`a01214a`

- Minor error correction

`1e64783`

- Login Bug: Only syntax - cannot find ID if id(corp_email) is duplicate. Add Unique() to Username and Corp_email on employee

`72e8adb`

- Change port number


`422486d`

- JWT authentication introduced

`9fa39bf`

- auto generated code - result of schema name change

`79f906f`

- trlog add to schema

`a8af987`

- Autogen preload data

`6fd9fd2`

- JWT Token added

`b6aafb3`

- Fix bug: edges memory addr not found error panic
