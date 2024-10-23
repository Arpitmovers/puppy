package puppy

import (
	"github.com/Arpitmovers/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! woof!"
}

func BigBark() string {
	return dog.WhenGrowsUp(Barks())
}
