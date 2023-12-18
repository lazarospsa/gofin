package gofin

import (
	"math"
)

// FutureValueAnnuity returns the future value of an annuity.
// The future value of an annuity is the sum of the cash flows multiplied by 1 plus the interest rate to the power of the period number minus 1 divided by the interest rate.
// The interest rate must be greater than zero to avoid division by zero.
// FV= C * ((1 + r)^n - 1) / r
// FV is the future value,
// C is the cash flow at the end of each period,
// r is the interest rate,
// n is the number of periods.
func FutureValueAnnuity(payment, interestRate float64, periods int) float64 {
	if interestRate == 0 {
		// Avoid division by zero
		return 0.0
	}

	return payment * ((math.Pow(1+interestRate, float64(periods)) - 1) / interestRate)
}

// FutureValue returns the future value of an investment based on periodic, constant payments and a constant interest rate.
// FV= C * ((1 + r)^n - 1) / r
// FV is the future value,
// C is the cash flow at the end of each period,
// r is the interest rate,
// n is the number of periods.
func FutureValue(presentValue, interestRate float64, periods int) float64 {
	return presentValue * math.Pow(1+interestRate, float64(periods))
}

// NetPresentValue function calculates the net present value of a series of cash flows given an
// interest rate and number of periods.
// NPV = sum(C / (1 + r)^t)
// NPV is the net present value,
// C is the cash flow at the end of each period,
// r is the interest rate,
// t is the number of periods.
func NetPresentValue(interestRate float64, periods int, cashFlows []float64) float64 {
	npv := 0.0
	for i := 0; i < len(cashFlows); i++ {
		npv += cashFlows[i] / math.Pow(1+interestRate, float64(i))
	}
	return npv
}

// PresentValue function calculates the present value of a future amount given an interest rate and
// number of periods.
// PV = FV / (1 + r)^t
// PV is the present value,
// FV is the future value,
// r is the interest rate,
// t is the number of periods.
func PresentValue(futureValue, interestRate float64, periods int) float64 {
	return futureValue / math.Pow(1+interestRate, float64(periods))
}

// The function calculates the present value of an annuity given an interest rate, number of periods,
// and cash flows.
// PV = sum(C / (1 + r)^t)
// PV is the present value,
// C is the cash flow at the end of each period,
// r is the interest rate,
// t is the number of periods.
func PresentValueAnnuity(interestRate float64, periods int, cashFlows []float64) float64 {
	presentValueAnnuity := 0.0
	for i := 0; i < len(cashFlows); i++ {
		presentValueAnnuity += cashFlows[i] / math.Pow(1+interestRate, float64(i))
	}
	return presentValueAnnuity
}

// HoldingPeriodReturn calculates the holding period return (HPR)
func HoldingPeriodReturn(initialValue, finalValue float64) float64 {
	return (finalValue - initialValue) / initialValue
}

// HoldingPeriodReturn calculates the holding period return (HPR)
// as a percentage
func HoldingPeriodReturnPercentage(initialValue, finalValue float64) float64 {
	return ((finalValue - initialValue) / initialValue)*100
}

func PresentValueAnnuityDue(interestRate float64, periods int, cashFlows []float64) float64 {
	presentValueAnnuityDue := 0.0
	for i := 0; i < len(cashFlows); i++ {
		presentValueAnnuityDue += cashFlows[i] / math.Pow(1+interestRate, float64(i))
	}
	return presentValueAnnuityDue
}

// PresentValuePerpetuity returns the present value of a perpetuity.
// The present value of a perpetuity is the cash flow divided by the discount rate.
// The interest rate must be greater than zero to avoid division by zero.
// PV= (C / r)
// PV is the present value,
// C is the cash flow at the end of the first period,
// r is the discount rate.
func PresentValuePerpetuity(interestRate, cashFlow float64) float64 {
	if interestRate == 0 {
		// Avoid division by zero.
		return 0.0
	}

	return cashFlow / interestRate
}

// PresentValuePerpetuityDue returns the present value of a perpetuity due.
// The present value of a perpetuity due is the cash flow divided by the discount rate multiplied by 1 plus the discount rate.
// The interest rate must be greater than zero to avoid division by zero.
// PV= (C / r) * (1 / (1+r))
// PV is the present value,
// C is the cash flow at the end of the first period,
// r is the discount rate.
func PresentValuePerpetuityDue(interestRate, cashFlow float64) float64 {
	if interestRate == 0 {
		// Avoid division by zero.
		return 0.0
	}

	return cashFlow / interestRate * (1 / (1 + interestRate))
}

func InterestRateGrowingPerpetuity(presentValue, cashFlow, growthRate float64) float64 {
	return cashFlow/presentValue + growthRate
}

func InterestRateGrowingAnnuity(presentValue, cashFlows, growthRate float64) float64 {
	return cashFlows/presentValue + growthRate
}

// TODO: fix this
func InterestRateGrowingAnnuityDue(presentValue, cashFlows, growthRate float64) float64 {
	return cashFlows/presentValue + growthRate
}

func InterestRateAnnuity(presentValue, cashFlows float64) float64 {
	return cashFlows / presentValue
}

func InterestRateAnnuityDue(presentValue, cashFlows float64) float64 {
	return cashFlows / presentValue
}

// TODO: fix this
func InterestRatePerpetuityDue(presentValue, cashFlow float64) float64 {
	return cashFlow / presentValue
}

// TODO: fix this
func InterestRateGrowingPerpetuityDue(presentValue, cashFlow, growthRate float64) float64 {
	return cashFlow/presentValue + growthRate
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

// TODO: fix this
func PresentValueGrowingAnnuityDue(interestRate, growthRate float64, periods int, cashFlows []float64) float64 {
	pv := 0.0
	for i := 0; i < len(cashFlows); i++ {
		pv += cashFlows[i] / math.Pow(1+interestRate, float64(i))
	}
	return pv
}

// PresentValueGrowingPerpetuity returns the present value of a growing perpetuity.
// The present value of a growing perpetuity is the cash flow divided by the difference between the discount rate and the growth rate.
// The interest rate must be greater than the growth rate to avoid division by zero.
// PV= (C / r−g)
// PV is the present value,
// C is the cash flow at the end of the first period,
// r is the discount rate,
// g is the growth rate.
func PresentValueGrowingPerpetuity(interestRate, growthRate, cashFlow float64) float64 {
	if interestRate <= growthRate {
		// Ensure the interest rate is greater than the growth rate to avoid division by zero.
		return 0.0
	}

	return cashFlow / (interestRate - growthRate)
}

// PresentValueGrowingPerpetuityDue returns the present value of a growing perpetuity due.
// The present value of a growing perpetuity due is the cash flow divided by the difference between the discount rate and the growth rate multiplied by 1 plus the discount rate.
// The interest rate must be greater than the growth rate to avoid division by zero.
// PV= (C / r−g) * (1 / (1+r))
// PV is the present value,
// C is the cash flow at the end of the first period,
// r is the discount rate,
// g is the growth rate.
func PresentValueGrowingPerpetuityDue(interestRate, growthRate, cashFlow float64) float64 {
	if interestRate <= growthRate {
		// Ensure the interest rate is greater than the growth rate to avoid division by zero.
		return 0.0
	}

	return (cashFlow / (interestRate - growthRate)) * (1 / (1 + interestRate))
}
