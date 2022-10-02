package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ohmpatel1997/sleeping-barber/pkg"
)

func main() {
	shop := pkg.NewShop(3, 1)
	shop.Start()

	for i := 0; i < 10; i++ {
		cl1 := pkg.NewClient(fmt.Sprintf("client-%d", i))
		go cl1.EnterShop(shop)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case <-quit:
		// blocking operation to shut down the shop
		shop.ShutDown()
	}

}
