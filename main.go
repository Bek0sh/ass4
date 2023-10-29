package main

func main() {

	shirtItem := newChat("Hello there")

	observerFirst := &User{Name: "Bekarys"}
	observerSecond := &User{Name: "Saiat"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()
}
