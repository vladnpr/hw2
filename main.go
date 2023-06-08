package main

import (
	"./animals"
	"./cage"
	"./zooKeeper"
)

func main() {
	rabbit := animals.NewRabbit("Boris", "2 kg", 2)
	lizard := animals.NewLizard("Oleg", "300 g", 10)
	chicken := animals.NewChicken("Carlos", "1.5 kg", 1)
	keeper := zooKeeper.NewZooKeeper("Johny Cage")

	cageR := cage.NewRabbitCage("Rabbit Cage")
	keeper.CatchRabbit(&rabbit, &cageR)

	cageL := cage.NewLizardCage("Lizard Cage")
	keeper.CatchLizard(&lizard, &cageL)

	cageC := cage.NewChickenCage("ChickenCage")
	keeper.CatchChicken(&chicken, &cageC)
}
