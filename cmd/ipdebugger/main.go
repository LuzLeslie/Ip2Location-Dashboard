package main

import (
	"fmt"
	simpleDB "github.com/LuzLeslie/Ip2Location-Dashboard/infrastructure/datastore/simpleDB"
	graphqlApp "github.com/LuzLeslie/Ip2Location-Dashboard/infrastructure/delivery/graphql"
	analyzeRepository "github.com/LuzLeslie/Ip2Location-Dashboard/package/analyze/repository"
	analyzeUseCase "github.com/LuzLeslie/Ip2Location-Dashboard/package/analyze/usecase"
	locationRepository "github.com/LuzLeslie/Ip2Location-Dashboard/package/location/repository"
	locationUseCase "github.com/LuzLeslie/Ip2Location-Dashboard/package/location/usecase"
	"github.com/graphql-go/graphql"
	"github.com/spf13/cobra"
	"net/http"
	"os"

	proxyRepository "github.com/LuzLeslie/Ip2Location-Dashboard/package/proxy/repository"
	proxyUseCase "github.com/LuzLeslie/Ip2Location-Dashboard/package/proxy/usecase"

	"github.com/graphql-go/handler"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})
	var ip2location, ip2proxy string
	rootCmd.Flags().StringVarP(&ip2location, "ip2location", "l", "", "Ip2 Location database file .bin")
	rootCmd.MarkFlagRequired("ip2location")
	rootCmd.Flags().StringVarP(&ip2proxy, "ip2proxy", "p", "", "Ip2 Proxy database file .bin")
}

var ip2location, ip2proxy string

var rootCmd = &cobra.Command{
	Use:   "ipdebugger",
	Short: "analyze and visualize your connections",
	Long:  `analyze and visualize your connections`,

	Run: func(cmd *cobra.Command, args []string) {

		ip2location, _ = cmd.Flags().GetString("ip2location")
		ip2proxy, _ = cmd.Flags().GetString("ip2proxy")

		createApp(ip2location,ip2proxy)

		// h, err := hanlderGraphql.NewHandler()
		// if err != nil {
		// 	log.Error(err)
		// }

		// h := handler.New(&handler.Config{
		// 	Schema: &graphqlServer.Schema,
		// })

		// http.Handle("/graphql", corsMiddleware(h))
		// http.HandleFunc("/subscriptions", subscriptions.Handler)
		// fmt.Println("graphql api server is started at: http://localhost:4000/graphql")
		// fmt.Println("subscriptions api server is started at: http://localhost:4000/subscriptions")
		// http.ListenAndServe(":4000", nil)
	},
}

func createApp(pathLocationBin, pathProxyBin  string) {
	log.Debug("App starting...")

	dbConn := simpleDB.New()
	analyzeRepo := analyzeRepository.New(dbConn)

	locationRepo := locationRepository.New(pathLocationBin)
	locationUcase := locationUseCase.New(locationRepo)

	proxyRepo := proxyRepository.New(pathProxyBin)
	proxyUcase := proxyUseCase.New(proxyRepo)

	analyzeUCase := analyzeUseCase.New(analyzeRepo, locationUcase, proxyUcase)

	app := graphqlApp.New(analyzeUCase)

	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:        app.Query(),
		Mutation:     app.Mutation(),
		Subscription: app.Subscription(),
	})

	app.RootSchema = &graphqlSchema

	if err != nil {
		log.Fatal(err)
	}

	graphQLHandler := handler.New(&handler.Config{
		Schema:   &graphqlSchema,
		GraphiQL: false,
		Pretty:   false,
	})

	// app.CreateRootSchema()
	// log.Debug("Pointer schema echo : ", app.RootSchema)
	// log.Warn("Pointer mariana main:", app.Mariana)
	// log.Warn("Pointer mariana main:", &app.Mariana)

	// http.HandleFunc("/graphql", app.Handler)
	http.Handle("/graphql", corsMiddleware(graphQLHandler))
	fmt.Println("graphql api server is started at: http://localhost:4000/graphql")

	http.HandleFunc("/subscriptions", graphqlApp.SubscriptionHandler)
	fmt.Println("subscriptions api server is started at: http://localhost:4000/subscriptions")

	http.ListenAndServe(":4000", nil)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST,OPTIONS")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
