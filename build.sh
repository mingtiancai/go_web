go build -o ./target/server main.go
# go build -o ./target/chitchat/chitchat ./ChitChat/*.go
go build -o ./target/c3/handle ./c3/handle/*.go
go build -o ./target/c3/handlefunc ./c3/handlefunc/*.go
go build -o ./target/c3/chain_handlefunc ./c3/chain_handlefunc/*.go
go build -o ./target/c3/chain_handler ./c3/chain_handler/*.go
go build -o ./target/c3/httprouter ./c3/httprouter/*.go
go build -o ./target/c3/https ./c3/https/*.go