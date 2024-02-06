package main

import (
	"context"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/khangle880/share_room/dataloaders"
	"github.com/khangle880/share_room/graph"
	pg "github.com/khangle880/share_room/pg/sqlc"
	"github.com/khangle880/share_room/utils"
	"github.com/rs/zerolog"

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

type MyQueryTracer struct {
	log zerolog.Logger
}

func (t *MyQueryTracer) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	t.log.Trace().Msgf("Executing command: %v,\nargs: %v", data.SQL, data.Args)
	return ctx
}

func (t *MyQueryTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	if data.Err != nil {
		t.log.Error().Msgf("Query failed: %v", data.Err)
	} else {
		t.log.Debug().Msgf("Query succeeded: %v", data.CommandTag)
	}
}

func main() {
	// dsn := os.Getenv("POSTGRESQL_URL")
	utils.GetLog()
	config, err := pgxpool.ParseConfig("postgres://postgres:admin@localhost:5432/share_room")
	if err != nil {
		utils.GetLog().Err(err).Msg("Unable to parse connString")
		os.Exit(1)
	}
	// config.ConnConfig.Tracer = &MyQueryTracer{log: *utils.GetLog()}
	conn, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		utils.GetLog().Err(err).Msg("Can't connect to database:")
		os.Exit(1)
	}
	defer conn.Close()

	repo := pg.NewRepository(conn)
	dl := dataloaders.NewRetriever()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	gin.Logger()

	// Setting up gin
	r := gin.New()
	r.Use(customMiddlerware.DefaultStructuredLogger())
	gin.ForceConsoleColor()
	r.Use(gin.Recovery())
	r.Use(customMiddlerware.GinContextToContextMiddleware())
	r.Use(customMiddlerware.AuthMiddleware(repo))
	r.POST("/query", dataloaders.Middleware(repo), graphqlHandler(repo, dl))
	r.GET("/", playgroundHandler())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(r.Run())
}

// srv := graphqlHandler(repo, dl)
// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
// http.Handle("/query", srv)
// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
// log.Fatal(http.ListenAndServe(":"+port, nil))
