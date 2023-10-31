package main

import (
	"assingment/graph"
	"assingment/graph/database"
	"assingment/graph/services"
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
	Conn, err := database.Open()
	if err != nil {
		log.Println("Instance of db not created", err)
		return
	}
	serviceConn, err := services.NewConn(Conn)
	if err != nil {
		log.Println("Error in creating service db connection", err)
		return
	}
	err = serviceConn.AutoMigrate()
	if err != nil {
		log.Println("Error Tables are not created", err)
		return
	}

	sr:= services.NewStore(serviceConn)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Store: sr}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
