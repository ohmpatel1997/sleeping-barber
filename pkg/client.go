package pkg

import (
	"fmt"
)

type Client struct {
	Name        string
	haircutDone chan bool
}

func NewClient(name string) *Client {
	return &Client{Name: name, haircutDone: make(chan bool)}
}

func (c *Client) EnterShop(shop *Shop) {
	if !shop.Open {
		fmt.Printf("you can not enter. shop is closed")
		return
	}

	fmt.Printf("\nclient %s entering the shop...\n", c.Name)

	select {
	// try to find the waiting list
	case shop.waitingLounge <- c:
		fmt.Printf("\nsuccessfully found the sit for client %s\n", c.Name)

		// since client have found the seat in lounge, he will wait until the hair cut is done
		select {
		case <-c.haircutDone:
			fmt.Printf("\nhair cut is done for client %s. client left the shop\n", c.Name)
		}
	default:
		fmt.Printf("\ncould not able to find seat for client %s, client left the shop.\n", c.Name)
	}
}

func (c *Client) HairCutDone() {
	c.haircutDone <- true
}
