package gnet

import "github.com/bwmarrin/snowflake"

func GenId(nodeId int64) string{
	node, err := snowflake.NewNode(nodeId)
	if err != nil {
		return ""
	}
	return node.Generate().String()
}