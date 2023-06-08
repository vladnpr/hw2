package main

import (
	zoo "github.com/vladnpr/zoo/animal"
	cage "github.com/vladnpr/zoo/cage"
	keeper "github.com/vladnpr/zoo/zooKeeper"
)

func main() {
	rabbit := zoo.NewRabbit("Boris", "2 kg", 2)
	lizard := zoo.NewLizard("Oleg", "300 g", 10)
	chicken := zoo.NewChicken("Carlos", "1.5 kg", 1)
	kpr := keeper.NewZooKeeper("Johny Cage")

	cageR := cage.NewRabbitCage("Rabbit Cage")
	kpr.CatchRabbit(&rabbit, &cageR)

	cageL := cage.NewLizardCage("Lizard Cage")
	kpr.CatchLizard(&lizard, &cageL)

	cageC := cage.NewChickenCage("ChickenCage")
	kpr.CatchChicken(&chicken, &cageC)
}
