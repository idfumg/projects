package main

import "fmt"

func main() {
	ns := GetNutritionalScore(NutritionalData{
		Energy:              EnergyFromKcal(500),
		Sugars:              SugarGram(20),
		SaturatedFattyAcids: SaturatedFattyAcids(5),
		Sodium:              SodiumMilligram(500),
		Fruits:              FruitsPercent(90),
		Fibre:               FibreGram(5),
		Protein:             ProteinGram(2),
	}, Food)

	fmt.Printf("Nutritional Score: %d\n", ns.Value)
	fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())
}
