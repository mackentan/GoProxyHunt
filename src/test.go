package main

import (
	"proxyhunt"
	"fmt"
)

func main() {
	ips := proxyhunt.GetCnProxy()

	fmt.Println(len(ips))

	for _,ip := range ips {
		fmt.Printf("%s\n",ip.Addr)
	}
}

