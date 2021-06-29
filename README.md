# Installation
Generate the Prisma Client using :
```shell
go run github.com/prisma/prisma-client-go generate
```

Sync the SQLite database with your schema :
```shell
go run github.com/prisma/prisma-client-go db push
```
# Usage
Start the server :
```shell
go run .
```
Browse to http://localhost:8080 to interact with the GraphQL playground.
# Resources
* Optimizing N+1 database queries using Dataloaders : https://gqlgen.com/reference/dataloaders/
* go generate based DataLoader : https://github.com/vektah/dataloaden
* Part 3 - Adding Dataloader with Dataloaden : https://www.youtube.com/watch?v=D3SrjuTaUQU
* Prisma Client Go : https://github.com/prisma/prisma-client-go