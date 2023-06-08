package zooKeeper

import (
	"../animals"
	"../cage"
	"fmt"
)

type ZooKeeper struct {
	Name   string
	Animal *animals.Animal
}

func (k *ZooKeeper) CatchRabbit(a *animals.Rabbit, c *cage.RabbitCage) {
	if c.Animal != nil {
		fmt.Println("The cage is not empty!")
		return
	}
	c.PutRabbit(a)
}

func (k *ZooKeeper) CatchLizard(a *animals.Lizard, c *cage.LizardCage) {
	if c.Animal != nil {
		fmt.Println("The cage is not empty!")
		return
	}
	c.PutLizard(a)
}

func (k *ZooKeeper) CatchChicken(a *animals.Chicken, c *cage.ChickenCage) {
	if c.Animal != nil {
		fmt.Println("The cage is not empty!")
		return
	}
	c.PutChicken(a)
}

func NewZooKeeper(name string) ZooKeeper {
	return ZooKeeper{
		Name:   name,
		Animal: nil,
	}
}
