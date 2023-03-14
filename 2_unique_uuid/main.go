package main

import (
	"encoding/json"
	"github.com/google/uuid"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"log"
)

func makeUUIDHandler(n *maelstrom.Node) func(maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		var body map[string]interface{}
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		id := uuid.New()
		body["id"] = id.String()
		body["type"] = "generate_ok"
		return echo(n, msg, body)
	}
}

func echo(n *maelstrom.Node, msg maelstrom.Message, body map[string]interface{}) error {
	return n.Reply(msg, body)
}

func main() {
	n := maelstrom.NewNode()
	n.Handle("generate", makeUUIDHandler(n))

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
