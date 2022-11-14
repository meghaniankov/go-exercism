package cars


// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	totalCars := float64(productionRate)
	return totalCars * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	const CostPerTenCars = 95000
	const CostPerSingleCar = 10000

	var groupsOfTen = carsCount / 10
	var remainingCars = carsCount - (groupsOfTen * 10)

	return uint(groupsOfTen * CostPerTenCars) + uint(remainingCars * CostPerSingleCar)
}
