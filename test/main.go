package main

import(
	"fmt"
	// "time"
	// "github.com/kamidzuua/decktracker-linux/server"
	"github.com/kamidzuua/decktracker-linux/client"
)

func main(){
	fmt.Println("hui")

	// go server.Server_run()
	// time.Sleep(20 * time.Second)
	client.Client_run()

}
