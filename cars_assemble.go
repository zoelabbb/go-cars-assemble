package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    workingCarsPerHour := float64(productionRate) * successRate / 100.0
    return workingCarsPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    workingCarsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
    workingCarsPerMinute := int(workingCarsPerHour / 60.0)
    return workingCarsPerMinute
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    groupsOfTen := carsCount / 10
    individualCars := carsCount % 10
    cost := uint(groupsOfTen*95000 + individualCars*10000)
    return cost
}
