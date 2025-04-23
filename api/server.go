package main

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"slices"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vektah/gqlparser/v2/ast"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"api/graph"
)

func startServer(port string) error {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOriginFunc: func(origin string) (bool, error) {
			u, err := url.Parse(origin)
			if err != nil {
				return false, err
			}
			host := u.Hostname()
			return host == "localhost" || host == "127.0.0.1", nil
		}, AllowHeaders: []string{"*"},
		AllowMethods:     []string{"GET", "HEAD", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))

	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. ")
	}

	c, err := mongo.Connect(options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := c.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	gql(e, c)

	return e.Start(port)
}

func gql(e *echo.Echo, c *mongo.Client) {
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Client: c}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,

		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				origin := r.Header.Get("Origin")
				log.Println(origin)
				if origin == "" || origin == r.Header.Get("Host") || origin == "http://localhost:1345" || origin == "http://localhost:5173" {
					return true
				}

				return slices.Contains([]string{":3000", "https://ui.mysite.com"}, origin)
			},
		},
	})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	playgroundHandler := playground.Handler("GraphQL", "/query")

	e.GET("/query", func(c echo.Context) error {
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	})
	e.POST("/query", func(c echo.Context) error {
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

}
