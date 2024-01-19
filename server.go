package main

import (
	"fmt"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/khangle880/share_room/graph"
	"github.com/khangle880/share_room/graph/dataloader"
	"github.com/khangle880/share_room/utils"

	customMiddlerware "github.com/khangle880/share_room/middleware"
	"github.com/khangle880/share_room/postgres"
	"github.com/khangle880/share_room/postgres/query"
)

const defaultPort = "8080"

// Defining the Graphql handler
func graphqlHandler(db *pg.DB) gin.HandlerFunc {
	usersRepo := query.UsersRepo{DB: db}
	c := graph.Config{Resolvers: &graph.Resolver{
		UsersRepo:      usersRepo,
		CategoriesRepo: query.CategoriesRepo{DB: db},
		BudgetsRepo:    query.BudgetRepo{DB: db},
	}}
	c.Directives.Auth = graph.Auth
	h := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
func main() {
	option, err := pg.ParseURL(
		"postgres://postgres:admin@localhost:5432/share_room?sslmode=disable",
		// os.Getenv("POSTGRESQL_URL"),
	)
	if err != nil {
		fmt.Printf("run PostgreSQL failed %v", err)
		os.Exit(1)
	}
	db := postgres.New(option)
	defer db.Close()
	db.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	utils.Get()

	// Setting up gin
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(customMiddlerware.AuthMiddleware(query.UsersRepo{DB: db}))
	r.POST("/query", dataloader.DataLoaderMiddleware(db), graphqlHandler(db))
	r.GET("/", playgroundHandler())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(r.Run())
}
