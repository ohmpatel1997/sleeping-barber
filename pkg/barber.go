package pkg

import (
	"fmt"
	"time"
)

type barber struct {
	ID   int
	Shop *Shop
	Stop chan bool
}

func newBarber(id int, shop *Shop) *barber {

	return &barber{
		ID:   id,
		Stop: make(chan bool),
		Shop: shop,
	}
}

func (b *barber) Start() {
	for {
		select {
		case cus := <-b.Shop.waitingLounge: // get the customer from waiting list
			fmt.Printf("\nbarber %d found the customer. doing the awesome haircut of %s...!\n", b.ID, cus.Name)
			time.Sleep(time.Second)
			// haircut is done, ask client to leave
			cus.HairCutDone()

		case <-b.Stop: // if closing the shop
			return
		}
	}
}

func (b *barber) Close() {
	fmt.Printf("stopping barber %d.......", b.ID)
	select {
	case b.Stop <- true:
		fmt.Printf("\nbarber %d stopped.......\n", b.ID)
	}
}
