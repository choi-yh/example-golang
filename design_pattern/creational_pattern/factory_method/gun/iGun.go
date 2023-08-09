package gun

import (
	"fmt"
	"time"
)

type IGun interface {
	GetName() string
	Shoot()
	Reload()
	CountBullet() int
	getMaxBullet() int
}

type Gun struct {
	Name      string
	Bullet    int
	MaxBullet int
}

func (i *Gun) GetName() string {
	return i.Name
}

func (i *Gun) Shoot() {
	if i.CountBullet() <= 0 {
		fmt.Println("The magazine is empty. A reload is required")
		return
	}

	i.Bullet -= 1
	fmt.Printf("shoot! current bullet = %d \n", i.Bullet)
}

func (i *Gun) Reload() {
	fmt.Println("reloading...")
	time.Sleep(3 * time.Second)

	i.Bullet = i.getMaxBullet()
}

func (i *Gun) CountBullet() int {
	return i.Bullet
}

func (i *Gun) getMaxBullet() int {
	return i.MaxBullet
}
