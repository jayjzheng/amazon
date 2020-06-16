# Authentication Service
verifies user credentials and provides shortlived JWT.

## Technology
- JWT
- datastore of choice
- language of choice

## Domain Object
- User { login, password }

## Caller
- API Gateway

## Actions
- Create User
- Change Password
- Create Token
- Refresh Token

## Publishes
- userCreated
- passwordChanged
- tokenCreated
- tokenRefreshed
- JWTkeyChanged

## Subscribes
- userUpdated

## Open Discussions
### datastore
- both NoSQL and RDBMS works fine.

### programming language
depends on the team members, experience, preferences, etc.
