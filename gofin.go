package gofin

import (
	"math"
)

func FutureValue(presentValue, interestRate float64, periods int) float64 {
	futureValue := presentValue * math.Pow(1 + interestRate, float64(periods))
	return futureValue
}

func NetPresentValue(interestRate float64, periods int, cashFlows []float64) float64 {
	npv := 0.0
	for i := 0; i < len(cashFlows); i++ {
		npv += cashFlows[i] / math.Pow(1 + interestRate, float64(i))
	}
	return npv
}

func PresentValue(futureValue, interestRate float64, periods int) float64 {
	presentValue := futureValue / math.Pow(1 + interestRate, float64(periods))
	return presentValue
}

func PresentValueAnnuity(interestRate float64, periods int, cashFlows []float64) float64 {
	presentValueAnnuity := 0.0
	for i := 0; i < len(cashFlows); i++ {
		presentValueAnnuity += cashFlows[i] / math.Pow(1 + interestRate, float64(i))
	}
	return presentValueAnnuity
}

func PresentValueAnnuityDue(interestRate float64, periods int, cashFlows []float64) float64 {
	presentValueAnnuityDue := 0.0
	for i := 0; i < len(cashFlows); i++ {
		presentValueAnnuityDue += cashFlows[i] / math.Pow(1 + interestRate, float64(i))
	}
	return presentValueAnnuityDue
}

func PresentValuePerpetuity(interestRate, cashFlow float64) float64 {
	presentValuePerpetuity := cashFlow / interestRate
	return presentValuePerpetuity
}

func PresentValuePerpetuityDue(interestRate, cashFlow float64) float64 {
	presentValuePerpetuityDue := cashFlow / interestRate
	return presentValuePerpetuityDue
}

func InterestRateGrowingPerpetuity(presentValue, cashFlow, growthRate float64) float64 {
	interestRateGrowingPerpetuity := cashFlow / presentValue + growthRate
	return interestRateGrowingPerpetuity
}

func InterestRateGrowingAnnuity(presentValue, cashFlows, growthRate float64) float64 {
	interestRateGrowingAnnuity := cashFlows / presentValue + growthRate
	return interestRateGrowingAnnuity
}

func InterestRateGrowingAnnuityDue(presentValue, cashFlows, growthRate float64) float64 {
	interestRateGrowingAnnuityDue := cashFlows / presentValue + growthRate
	return interestRateGrowingAnnuityDue
}

func InterestRateAnnuity(presentValue, cashFlows float64) float64 {
	interestRateAnnuity := cashFlows / presentValue
	return interestRateAnnuity
}

func InterestRateAnnuityDue(presentValue, cashFlows float64) float64 {
	interestRateAnnuityDue := cashFlows / presentValue
	return interestRateAnnuityDue
}

func InterestRatePerpetuityDue(presentValue, cashFlow float64) float64 {
	interestRatePerpetuityDue := cashFlow / presentValue
	return interestRatePerpetuityDue
}

func InterestRateGrowingPerpetuityDue(presentValue, cashFlow, growthRate float64) float64 {
	interestRateGrowingPerpetuityDue := cashFlow / presentValue + growthRate
	return interestRateGrowingPerpetuityDue
}

func InterestRate(presentValue, futureValue float64, periods int) float64 {
	return math.Pow(futureValue/presentValue, 1/float64(periods)) - 1
}

func InterestRateContinuousCompounding(presentValue, futureValue float64, periods int) float64 {
	return math.Log(futureValue/presentValue) / float64(periods)
}

func InterestRatePerpetuity(presentValue, cashFlow float64) float64 {
	return cashFlow / presentValue
}

func PresentValueGrowingAnnuity(interestRate, growthRate float64, periods int, cashFlows []float64) float64 {
	pv := 0.0
	for i := 0; i < len(cashFlows); i++ {
		pv += cashFlows[i] / math.Pow(1+interestRate, float64(i+1))
	}
	return pv
}

func PresentValueGrowingAnnuityDue(interestRate, growthRate float64, periods int, cashFlows []float64) float64 {
	pv := 0.0
	for i := 0; i < len(cashFlows); i++ {
		pv += cashFlows[i] / math.Pow(1+interestRate, float64(i))
	}
	return pv
}

func PresentValueGrowingPerpetuity(interestRate, growthRate, cashFlow float64) float64 {
	return cashFlow / (interestRate - growthRate)
}

func PresentValueGrowingPerpetuityDue(interestRate, growthRate, cashFlow float64) float64 {
	return cashFlow / (interestRate - growthRate)
}