package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	// graphqlServer "github.com/LuzLeslie/ip-debugger/infrastructure/delivery/graphql/mutation"
	// subscriptions "github.com/LuzLeslie/ip-debugger/infrastructure/delivery/graphql/subscription"
	simpleDB "github.com/LuzLeslie/ip-debugger/infrastructure/datastore/simpleDB"
	graphqlApp "github.com/LuzLeslie/ip-debugger/infrastructure/delivery/graphql"
	analyzeRepository "github.com/LuzLeslie/ip-debugger/package/analyze/repository"
	analyzeUseCase "github.com/LuzLeslie/ip-debugger/package/analyze/usecase"
	locationRepository "github.com/LuzLeslie/ip-debugger/package/location/repository"
	locationUseCase "github.com/LuzLeslie/ip-debugger/package/location/usecase"

	proxyRepository "github.com/LuzLeslie/ip-debugger/package/proxy/repository"
	proxyUseCase "github.com/LuzLeslie/ip-debugger/package/proxy/usecase"

	"github.com/graphql-go/handler"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})
	var ip2location, ip2proxy string
	rootCmd.Flags().StringVarP(&ip2location, "ip2location", "l", "", "Ip2 Location database file .bin")
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

		createApp()

		// if ip2proxy == "" && ip2location == "" {
		// 	log.Error("Error: requires at least 1 database")
		// 	fmt.Println("Error: requires at least 1 database")
		// 	cmd.Help()
		// 	os.Exit(0)
		// }

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

func createApp() {
	log.Debug("App starting...")

	dbConn := simpleDB.New()
	analyzeRepo := analyzeRepository.New(dbConn)

	pathLocationBin := "./tmp/IP2LOCATION-LITE-DB11.BIN"
	locationRepo := locationRepository.New(pathLocationBin)
	locationUcase := locationUseCase.New(locationRepo)

	pathProxyBin := "./tmp/IP2PROXY-LITE-PX10.BIN"
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
