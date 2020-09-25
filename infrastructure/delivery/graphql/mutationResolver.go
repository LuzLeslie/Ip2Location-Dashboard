package graphql

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"
)

func (s Schema) AnalyzeIp(params graphql.ResolveParams) (interface{}, error) {
	log.Debug("Resolver mutation-AnalyzeIp")
	ip := params.Args["ip"].(string)
	log.Debug("Analyzing: ", ip)

	result, err := s.AnalyzeUseCase.GetInfoIp(ip)
	if err != nil {
		return nil, err
	}

	s.IpPublisher()
	return result, nil
}

func (s Schema) IpPublisher() {

	graphqlSchema, _ := graphql.NewSchema(graphql.SchemaConfig{ // fix this reference Schema
		Query:        s.Query(),
		Mutation:     s.Mutation(),
		Subscription: s.Subscription(),
	})

	Subscribers.Range(func(key, value interface{}) bool {
		subscriber, ok := value.(*Subscriber)
		if !ok {
			return true
		}
		payload := graphql.Do(graphql.Params{
			Schema:        graphqlSchema, // fix this reference Schema
			RequestString: subscriber.RequestString,
		})
		message, err := json.Marshal(map[string]interface{}{
			"type":    "data",
			"id":      subscriber.OperationID,
			"payload": payload,
		})
		if err != nil {
			log.Printf("failed to marshal message: %v", err)
			return true
		}
		if err := subscriber.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
			if err == websocket.ErrCloseSent {
				Subscribers.Delete(key)
				return true
			}
			log.Printf("failed to write to ws connection: %v", err)
			return true
		}
		return true
	})
}
