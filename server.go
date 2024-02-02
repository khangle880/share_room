package main

import (
	"context"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/khangle880/share_room/dataloaders"
	"github.com/khangle880/share_room/graph"
	"github.com/khangle880/share_room/pg"
	"github.com/khangle880/share_room/utils"
	zerologadapter "github.com/jackc/pgx-zerolog"

	customMiddlerware "github.com/khangle880/share_room/middleware"
)

const defaultPort = "8080"

// Defining the Graphql handler
func graphqlHandler(repo *pg.RepoSvc, dataloader dataloaders.Retriever) gin.HandlerFunc {
	utils.GetLog()
	c := graph.Config{Resolvers: &graph.Resolver{
		Repository:  repo,
		DataLoaders: dataloader,
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
	dsn := os.Getenv("POSTGRESQL_URL")
	// conn, err := pg.Open(dbURL)
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		utils.GetLog().Err(err).Msg("Can't connect to database:")
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	logger := zerologadapter.NewLogger(utils.Log)
	conn.Config().Tracer = &tracelog.TraceLog{Logger: logger, LogLevel: tracelog.LogLevelTrace}

	// db := sqldblogger.OpenDriver(dsn, conn.Driver(), zerologadapter.New(*utils.GetLog()))
	// db.Ping()
	repo := pg.NewRepository(conn)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	gin.Logger()

	// Setting up gin
	r := gin.New()
	r.Use(customMiddlerware.DefaultStructuredLogger())
	r.Use(gin.Recovery())
	r.Use(customMiddlerware.GinContextToContextMiddleware())
	r.Use(customMiddlerware.AuthMiddleware(repo))
	dl := dataloaders.NewRetriever()
	r.POST("/query", dataloaders.Middleware(repo), graphqlHandler(repo, dl))
	r.GET("/", playgroundHandler())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(r.Run())
}
