# shopping system

## key design principles
- microservice architecture
- event driven design
- prefer async communications between services

## key domains concepts
- users
- products
- shopping carts

## possible services
- user
- auth
- product
- search
- shopping
- inventory
- shipping

### user
- keep track of things user related
- accept requests such as user creation, address change, etc.
- corresponding events: userCreated, addressChanged, etc.

### auth
- all things authentication and authorization related
- accept requests to generate/refresh tokens
- corresponding events: tokenCreated, tokenExpired, etc.

### product
- manage all things product related
- accpet request such as product added, product spec change, etc.
- corresponding events: productAdded, specChanged, etc.

### search
- read only service allowing users to search domain objects
- accept requests: search product, search orders, etc.

### shopping

### inventory

### shipping
