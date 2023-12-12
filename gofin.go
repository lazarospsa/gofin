package gofin

import (
	"math"
)

func futureValue(presentValue, interestRate float64, periods int) float64 {
	futureValue := presentValue * math.Pow(1 + interestRate, float64(periods))
	return futureValue
}

func netPresentValue(interestRate float64, periods int, cashFlows []float64) float64 {
	npv := 0.0
	for i := 0; i < len(cashFlows); i++ {
		npv += cashFlows[i] / math.Pow(1 + interestRate, float64(i))
	}
	return npv
}

func presentValue(futureValue, interestRate float64, periods int) float64 {
	presentValue := futureValue / math.Pow(1 + interestRate, float64(periods))
	return presentValue
}

func presentValueAnnuity(interestRate float64, periods int, cashFlows []float64) float64 {
	presentValueAnnuity := 0.0
	for i := 0; i < len(cashFlows); i++ {
		presentValueAnnuity += cashFlows[i] / math.Pow(1 + interestRate, float64(i))
	}
	return presentValueAnnuity
}

func presentValueAnnuityDue(interestRate float64, periods int, cashFlows []float64) float64 {
	presentValueAnnuityDue := 0.0
	for i := 0; i < len(cashFlows); i++ {
		presentValueAnnuityDue += cashFlows[i] / math.Pow(1 + interestRate, float64(i))
	}
	return presentValueAnnuityDue
}

func presentValuePerpetuity(interestRate, cashFlow float64) float64 {
	presentValuePerpetuity := cashFlow / interestRate
	return presentValuePerpetuity
}

func presentValuePerpetuityDue(interestRate, cashFlow float64) float64 {
	presentValuePerpetuityDue := cashFlow / interestRate
	return presentValuePerpetuityDue
}

func interestRateGrowingPerpetuity(presentValue, cashFlow, growthRate float64) float64 {
	interestRateGrowingPerpetuity := cashFlow / presentValue + growthRate
	return interestRateGrowingPerpetuity
}

func interestRateGrowingAnnuity(presentValue, cashFlows, growthRate float64) float64 {
	interestRateGrowingAnnuity := cashFlows / presentValue + growthRate
	return interestRateGrowingAnnuity
}

func interestRateGrowingAnnuityDue(presentValue, cashFlows, growthRate float64) float64 {
	interestRateGrowingAnnuityDue := cashFlows / presentValue + growthRate
	return interestRateGrowingAnnuityDue
}

func interestRateAnnuity(presentValue, cashFlows float64) float64 {
	interestRateAnnuity := cashFlows / presentValue
	return interestRateAnnuity
}

func interestRateAnnuityDue(presentValue, cashFlows float64) float64 {
	interestRateAnnuityDue := cashFlows / presentValue
	return interestRateAnnuityDue
}

func interestRatePerpetuityDue(presentValue, cashFlow float64) float64 {
	interestRatePerpetuityDue := cashFlow / presentValue
	return interestRatePerpetuityDue
}

func interestRateGrowingPerpetuityDue(presentValue, cashFlow, growthRate float64) float64 {
	interestRateGrowingPerpetuityDue := cashFlow / presentValue + growthRate
	return interestRateGrowingPerpetuityDue
}

func interestRate(presentValue, futureValue float64, periods int) float64 {
	return math.Pow(futureValue/presentValue, 1/float64(periods)) - 1
}

func interestRateContinuousCompounding(presentValue, futureValue float64, periods int) float64 {
	return math.Log(futureValue/presentValue) / float64(periods)
}

func interestRatePerpetuity(presentValue, cashFlow float64) float64 {
	return cashFlow / presentValue
}

func presentValueGrowingAnnuity(interestRate, growthRate float64, periods int, cashFlows []float64) float64 {
	pv := 0.0
	for i := 0; i < len(cashFlows); i++ {
		pv += cashFlows[i] / math.Pow(1+interestRate, float64(i+1))
	}
	return pv
}

func presentValueGrowingAnnuityDue(interestRate, growthRate float64, periods int, cashFlows []float64) float64 {
	pv := 0.0
	for i := 0; i < len(cashFlows); i++ {
		pv += cashFlows[i] / math.Pow(1+interestRate, float64(i))
	}
	return pv
}

func presentValueGrowingPerpetuity(interestRate, growthRate, cashFlow float64) float64 {
	return cashFlow / (interestRate - growthRate)
}

func presentValueGrowingPerpetuityDue(interestRate, growthRate, cashFlow float64) float64 {
	return cashFlow / (interestRate - growthRate)
}