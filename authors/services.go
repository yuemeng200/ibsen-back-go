package authors

import (
	"awesomeProject/common"
	"common"
)

func getAuthorByID(id) {
	cypher := "MATCH (a:author) WHERE a.id=$id RETURN a"
	common.Run
}