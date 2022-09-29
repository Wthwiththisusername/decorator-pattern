package decorator

import "fmt"

func main() {
	var beverage iBeverage = newEspresso()
	fmt.Println("Your order is", beverage.getDescription())
	fmt.Println("Total price: %d/n", beverage.cost())
	var beverage2 iBeverage = newDarkRoast()
	beverage2 = newMocha(beverage2)
	beverage2 = newMocha(beverage2)
	beverage2 = newWhip(beverage2)
	fmt.Println("Your order is ", beverage2.getDescription())
	fmt.Println("Total price: ", beverage2.cost())
}
