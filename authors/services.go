package authors

import (
	"ibsen-back-go/common"
	"ibsen-back-go/util/log"
)

func getAuthorByID(id string) {
	cypher := "MATCH (a:author) WHERE a.id=$id RETURN a"
	params := map[string]interface{}{"id": id}
	res := common.Run(cypher, params)
	log.Logger("res %v", res)
}
