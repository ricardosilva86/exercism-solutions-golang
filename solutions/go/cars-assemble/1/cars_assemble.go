package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsPerMinute := (productionRate / 60)
	rate := float64(carsPerMinute) * successRate / 100
	return int(rate)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	left_overs := carsCount % 10
	whole := (carsCount - left_overs) / 10
	return uint((left_overs * 10000) + (whole * 95000))
}
