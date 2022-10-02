package pkg

import (
	"fmt"
)

type Shop struct {
	waitingLounge chan *Client
	// tight coupling with the barbers, currently its size is one. but to make it scalable
	// for future, keep it array
	barbers []*barber
	Open    bool
}

func NewShop(waitingLoungeSize int, noOfBarbers int) *Shop {
	shop := &Shop{
		waitingLounge: make(chan *Client, waitingLoungeSize),
		Open:          false,
	}

	barbers := make([]*barber, 0, noOfBarbers)
	for i := 0; i < noOfBarbers; i++ {
		barbers = append(barbers, newBarber(i+1, shop))
	}

	shop.barbers = barbers
	return shop
}

func (s *Shop) Start() {
	fmt.Println("opening the shop......")
	for _, barb := range s.barbers {
		go barb.Start()
	}
	s.Open = true
	fmt.Println("shop is open now.")
}

func (s *Shop) shutDownBarbers() {
	for _, barb := range s.barbers {
		barb.Close() // non blocking operation until all barber woke up
	}
}

func (s *Shop) ShutDown() {
	fmt.Println("stopping the shop.....")
	// close the shop for new clients.
	s.Open = false

	fmt.Println("shop closed for new clients.")
	fmt.Println("waiting for the existing clients haircuts to be done...")
	for len(s.waitingLounge) > 0 {
		fmt.Printf("\nwating for remaining %d client haircut\n", len(s.waitingLounge))
		// wait for the waiting lounge to be empty
	}

	// wait for the barbers to stop
	s.shutDownBarbers()

	fmt.Println("\nshop successfully closed. please come tomorrow")
}
