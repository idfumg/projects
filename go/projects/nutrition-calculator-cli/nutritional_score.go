package main

import "fmt"

type ScoreType int

const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType ScoreType
}

var scoreToLetter = []string{"A", "B", "C", "D", "E"}

type EnergyKJ float64
type SugarGram float64
type SaturatedFattyAcids float64
type SodiumMilligram float64
type FruitsPercent float64
type FibreGram float64
type ProteinGram float64

type NutritionalData struct {
	Energy              EnergyKJ
	Sugars              SugarGram
	SaturatedFattyAcids SaturatedFattyAcids
	Sodium              SodiumMilligram
	Fruits              FruitsPercent
	Fibre               FibreGram
	Protein             ProteinGram
	IsWater             bool
}

var energyLevels = []float64{3350, 3015, 2680, 2345, 2010, 1675, 1340, 1005, 670, 335}
var sugarsLevels = []float64{45, 60, 36, 31, 27, 22.5, 18, 13.5, 9, 4.5}
var saturatedFattyAcids = []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
var sodiumLevels = []float64{900, 810, 720, 630, 540, 450, 360, 270, 180, 90}
var fibreLevels = []float64{4.7, 3.7, 2.8, 1.9, 0.9}
var proteinLevels = []float64{8, 6.4, 4.8, 3.2, 1.6}
var beverageEnergyLevels = []float64{270, 240, 210, 180, 150, 120, 90, 60, 30, 0}
var sugarsBeverageLevels = []float64{13.5, 12, 10.5, 9, 7.5, 6, 4.5, 3, 1.5, 0}

func (e EnergyKJ) GetPoints(st ScoreType) int {
	if st == Beverage {
		return GetPointsFromRange(float64(st), beverageEnergyLevels)
	}
	return GetPointsFromRange(float64(st), energyLevels)
}

func (e SugarGram) GetPoints(st ScoreType) int {
	if st == Beverage {
		return GetPointsFromRange(float64(st), sugarsBeverageLevels)
	}
	return GetPointsFromRange(float64(st), sugarsLevels)
}

func (e SaturatedFattyAcids) GetPoints(st ScoreType) int {
	return GetPointsFromRange(float64(st), saturatedFattyAcids)
}

func (e SodiumMilligram) GetPoints(st ScoreType) int {
	return GetPointsFromRange(float64(st), sodiumLevels)
}

func (f FruitsPercent) GetPoints(st ScoreType) int {
	if st == Beverage {
		if f > 80 {
			return 10
		} else if f > 60 {
			return 4
		} else if f > 40 {
			return 2
		}
	}
	if f > 80 {
		return 5
	} else if f > 60 {
		return 2
	} else if f > 40 {
		return 1
	}
	return 0
}

func (e FibreGram) GetPoints(st ScoreType) int {
	return GetPointsFromRange(float64(st), fibreLevels)
}

func (e ProteinGram) GetPoints(st ScoreType) int {
	return GetPointsFromRange(float64(st), proteinLevels)
}

func EnergyFromKcal(kcal float64) EnergyKJ {
	return EnergyKJ(kcal * 4.184)
}

func SodiumFromSalt(saltMg float64) SodiumMilligram {
	return SodiumMilligram(saltMg / 2.5)
}

func GetNutritionalScore(nd NutritionalData, scoreType ScoreType) NutritionalScore {
	value := 0
	positive := 0
	negative := 0

	if scoreType != Water {
		fruitPoints := nd.Fruits.GetPoints(scoreType)
		fibrePoints := nd.Fibre.GetPoints(scoreType)
		negative = nd.Energy.GetPoints(scoreType) + nd.Sugars.GetPoints(scoreType) + nd.SaturatedFattyAcids.GetPoints(scoreType) + nd.Sodium.GetPoints(scoreType)
		positive = fruitPoints + fibrePoints + nd.Protein.GetPoints(scoreType)

		fmt.Println("fruitPoints:", fruitPoints)
		fmt.Println("fibrePoints:", fibrePoints)
		fmt.Println("negative:", negative)
		fmt.Println("positive:", positive)

		if scoreType == Cheese {
			value = positive - negative
			fmt.Printf("1. value = %d\n", value)
		} else {
			if negative >= 11 && fruitPoints < 5 {
				fmt.Printf("2. value = %d\n", value)
				value = negative - positive - fruitPoints
			} else {
				fmt.Printf("3. value = %d\n", value)
				value = negative - positive
				fmt.Printf("3. value = %d\n", value)
			}
		}
	}

	return NutritionalScore{
		Value:     value,
		Positive:  positive,
		Negative:  negative,
		ScoreType: scoreType,
	}
}

func (ns NutritionalScore) GetNutriScore() string {
	if ns.ScoreType == Food {
		fmt.Printf("ns.Value = %d\n", ns.Value)
		return scoreToLetter[GetPointsFromRange(float64(ns.Value), []float64{18, 10, 2, -1})]
	} else if ns.ScoreType == Water {
		return scoreToLetter[0]
	}
	return scoreToLetter[GetPointsFromRange(float64(ns.Value), []float64{9, 5, 1, -2})]
}

func GetPointsFromRange(v float64, steps []float64) int {
	n := len(steps)
	for i, value := range steps {
		if v > value {
			return n - i
		}
	}
	return 0
}
