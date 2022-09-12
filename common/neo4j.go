package common

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var driver neo4j.Driver

func InitDB(uri, username, password string) {
	var err error
	driver, err = neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	fmt.Println(err)
}

func Run(cypher string, params map[string]interface{}) neo4j.Result {
	session:= driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	res, err := session.Run(cypher, params)
	if err!=nil{
		fmt.Println(err)
	}
	return res
}
