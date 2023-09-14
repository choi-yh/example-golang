package main

import (
	"fmt"

	gun2 "github.com/choi-yh/example-golang/TIL/design_pattern/creational_pattern/factory_method/gun"
)

func getGun(name string) (gun2.IGun, error) {
	switch name {
	case "AK47":
		return gun2.NewAK47(), nil
	case "musket":
		return gun2.NewMusket(), nil
	}

	return nil, fmt.Errorf("invalid gun name")
}

func printGun(g gun2.IGun) {
	fmt.Printf("gun name = %v \n", g.GetName())
	fmt.Printf("current bullet = %v \n", g.CountBullet())

	for i := 0; i <= 15; i++ {
		g.Shoot()
	}

	g.Reload()
	fmt.Printf("current bullet = %v \n", g.CountBullet())

	fmt.Println()
}
