package composition

type AnimalType string

const (
	MAMMAL  AnimalType = "mammal"
	REPTILE AnimalType = "reptile"
	BIRD    AnimalType = "bird"
	FISH    AnimalType = "fish"
)

// Animal base struct
type Animal struct {
	Name string
}

func (a *Animal) IsWarmBlood() bool {
	return true
}

// Dog struct embedding Animal
type Dog struct {
	Animal
	Breed string
}

// Cat struct embedding Animal
type Cat struct {
	Animal
	IsDomestic bool
}

// Bird struct embedding Animal
type Bird struct {
	Animal
	CanFly bool
}

// Override IsWarmBlood with additional functionality
func (b *Bird) IsWarmBlood() bool {
	// Call base method first
	baseResult := b.Animal.IsWarmBlood()

	return baseResult
}

// Fish struct embedding Animal
type Fish struct {
	Animal
	WaterType string
}

// Completely override IsWarmBlood
func (f *Fish) IsWarmBlood() bool {
	// Fish are always cold-blooded regardless of type
	return false
}
