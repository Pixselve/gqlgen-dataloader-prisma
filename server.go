package main

import (
	"gqlgen-dataloader-prisma/dataloader"
	"gqlgen-dataloader-prisma/db"
	"gqlgen-dataloader-prisma/graph"
	"gqlgen-dataloader-prisma/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	prismaClient := db.NewClient()
	if err := prismaClient.Prisma.Connect(); err != nil {
		panic(err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Prisma: prismaClient}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", dataloader.Middleware(prismaClient, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
