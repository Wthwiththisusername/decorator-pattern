package decorator

type iBeverage interface {
	cost() float32
	getDescription() string
}
