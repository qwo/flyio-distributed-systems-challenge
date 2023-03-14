alias maelstrom=~/Code/maelstrom/maelstrom


```
1. maelstrom test -w echo --bin ~/go/bin/maelstrom-echo --node-count 1 --time-limit 10

2. maelstrom test -w unique-ids --bin ~/go/bin/maelstrom-unique-ids --time-limit 30 --rate 1000 --node-count 3 --availability total --nemesis partition

```