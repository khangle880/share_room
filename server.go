package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/khangle880/share_room/graph"
	"github.com/khangle880/share_room/graph/dataloader"
	"github.com/khangle880/share_room/utils"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"

	customMiddlerware "github.com/khangle880/share_room/middleware"
)

const defaultPort = "8080"

type ApiConfig struct {
	*database.Queries
}

// Defining the Graphql handler
func graphqlHandler(db *database.Queries) gin.HandlerFunc {
	utils.GetLog()
	c := graph.Config{Resolvers: &graph.Resolver{
		DBQueries: db,
	}}
	c.Directives.Auth = graph.Auth
	c.Directives.HasRole = graph.HasRole
	h := handler.NewDefaultServer(graph.NewExecutableSchema(c))
	// h.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	// 	if !userForContext(ctx).IsAdmin {
	// 		graphql.GetOperationContext(ctx).DisableIntrospection = true
	// 	}

	// 	return next(ctx)
	// })
	// ?Error Hooks
	// h.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
	// 	err := graphql.DefaultErrorPresenter(ctx, e)
	// 	return err
	// })
	// h.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
	// 	return gqlerror.Errorf("Internal server error!")
	// })

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
	dbURL := os.Getenv("DB_URL")
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		utils.GetLog().Info().Msgf("Can't connect to database:", err)
	}
	if err != nil {
		fmt.Printf("run PostgreSQL failed %v", err)
		os.Exit(1)
	}

	db := sqldblogger.OpenDriver(dbURL, conn.Driver(), zerologadapter.New(*utils.GetLog()))
	db.Ping()
	dbQueries := database.New(conn)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Setting up gin
	r := gin.New()
	r.Use(customMiddlerware.DefaultStructuredLogger())
	r.Use(gin.Recovery())
	r.Use(customMiddlerware.GinContextToContextMiddleware())
	r.Use(customMiddlerware.AuthMiddleware(dbQueries))
	r.POST("/query", dataloader.DataLoaderMiddleware(dbQueries), graphqlHandler(dbQueries))
	r.GET("/", playgroundHandler())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(r.Run())
}
