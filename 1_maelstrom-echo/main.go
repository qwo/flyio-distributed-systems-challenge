package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func makeEchoHandler(n *maelstrom.Node) func(maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		var body map[string]interface{}
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		body["type"] = "echo_ok"
		return echo(n, msg, body)
	}
}

func echo(n *maelstrom.Node, msg maelstrom.Message, body map[string]interface{}) error {
	return n.Reply(msg, body)
}

func main() {
	n := maelstrom.NewNode()
	n.Handle("echo", makeEchoHandler(n))

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
