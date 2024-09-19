package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/MaranhaoTRT22/13-GraphQL/graph"
	"github.com/MaranhaoTRT22/13-GraphQL/internal/database"
    _ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8099"

func main() {
    // Inicializando conexão com banco de dados
    db, err:= sql.Open("sqlite3", "./cmd/database/data.db")
    if err != nil {
        log.Fatalf("Falha ao tentar abrir o banco de dados: %v", err)
    }
    // Finaliza conexão com banco de dados
    defer db.Close()

    categoryDb:= database.NewCategory(db)
    courseDb:= database.NewCourse(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
        CategoryDB: categoryDb,
        CourseDB: courseDb,
    }}))

	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("conectado ao http://localhost:%s/ para o GraphQL Playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
