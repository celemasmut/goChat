package main

import "fmt"

// Interface que define el comportamiento de un producto
type IProduct interface {
	setStock(stock int)
	getStock()int
	setName(name string)
	getName()string
}

// Implementacion de la interfaz IProduct para el producto de tipo "Computadora"
type Computer struct {
	name string
	stock int
}

// Implementando de forma implicita la interfaz IProduct
func (c *Computer)setStock(stock int)  {
	c.stock=stock
}

func (c *Computer)setName(name string)  {
	c.name=name
}
func (c *Computer)getName()string  {
	return c.name
}

func (c *Computer)getStock()int{
	return c.stock
}

// Creando clase base de computadora, por composicion sobre herencia
type Laptop struct { //composicion sobre herencia en go . se le asigna propiedades de computadora
	Computer
}

func newLaptop() IProduct{ //utilizando factory .
	return &Laptop{
		Computer: Computer{
			name: "LAptop Computer",
			stock: 25,
		},
	}
}

type Desktop struct {
	Computer
}

func NewDesktop() IProduct  {
	return &Desktop{
		Computer{
			name: "Desktop Computer",
			stock: 35,
		},
	}
}

// Creando fabrica de productos: Factory pattern
func GetComputerFactory(computerType string)(IProduct,error){
	if computerType == "laptop"{
		return newLaptop(),nil
	}
	if computerType == "desktop"{
		return NewDesktop(),nil
	}
	return nil,fmt.Errorf("invalid computer type")
}

// Trying polymorphism
func printNAmeAndStock(p IProduct)  {
	fmt.Printf("Product name: %s, with stock %d\n", p.getName(), p.getStock())
}

func main()  {
	laptop, _ := GetComputerFactory("laptop")
	desktop, _:= GetComputerFactory("desktop")
	printNAmeAndStock(laptop)
	printNAmeAndStock(desktop)
}