package cage

import "go_course/hw2/animals"

type cage struct {
	Name string
}

type RabbitCage struct {
	cage
	Animal *animals.Rabbit
}

type LizardCage struct {
	cage
	Animal *animals.Lizard
}

type ChickenCage struct {
	cage
	Animal *animals.Chicken
}

func (rc *RabbitCage) PutRabbit(a *animals.Rabbit) {
	rc.Animal = a
}

func (lc LizardCage) PutLizard(a *animals.Lizard) {
	lc.Animal = a
}

func (cc ChickenCage) PutChicken(a *animals.Chicken) {
	cc.Animal = a
}

func NewRabbitCage(name string) RabbitCage {
	return RabbitCage{cage{Name: name}, nil}
}

func NewLizardCage(name string) LizardCage {
	return LizardCage{cage{Name: name}, nil}
}

func NewChickenCage(name string) ChickenCage {
	return ChickenCage{cage{Name: name}, nil}
}
