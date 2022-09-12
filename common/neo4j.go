package common

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"ibsen-back-go/util/log"
)

var driver neo4j.Driver

func InitDB(uri, username, password string) {
	var err error
	driver, err = neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	fmt.Println(err)
}

func Run(cypher string, params map[string]interface{}) interface{} {
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	greeting, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"CREATE (a:Greeting) SET a.message = $message RETURN a.message + ', from node ' + id(a)",
			map[string]interface{}{"message": "hello, world"})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		log.Logger("%s", err)
	}

	return greeting
}
