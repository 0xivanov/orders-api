package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/0xivanov/orders-api/application"
)

func main() {
	app := application.New()
	ctxt, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	err := app.Start(ctxt)

	if err != nil {
		fmt.Println("cannot start app: ", err)
	}
	fmt.Println("end of main")
}
