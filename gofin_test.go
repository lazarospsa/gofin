package gofin

import (
	"testing"
	"math"
)

func TestNetPresentValue(t *testing.T) {
	var interestRate float64 = 0.1
	var periods int = 2
	var cashFlows []float64 = []float64{100, 100}
	var expected float64 = 121
	actual := netPresentValue(interestRate, periods, cashFlows)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestPresentValue(t *testing.T) {
	var futureValue float64 = 100
	var interestRate float64 = 0.1
	var periods int = 2
	var expected float64 = 121
	actual := presentValue(futureValue, interestRate, periods)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestPresentValueAnnuity(t *testing.T) {
	var interestRate float64 = 0.1
	var periods int = 2
	cashFlows := []float64{100, 100}
	var expected float64 = 121
	actual := presentValueAnnuity(interestRate, periods, cashFlows)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestPresentValueAnnuityDue(t *testing.T) {
	var interestRate float64 = 0.1
	var periods int = 2
	cashFlows := []float64{100, 100}
	var expected float64 = 121
	actual := presentValueAnnuityDue(interestRate, periods, cashFlows)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestPresentValuePerpetuity(t *testing.T) {
	var interestRate float64 = 0.1
	var cashFlow float64 = 100
	var expected float64 = 1000
	actual := presentValuePerpetuity(interestRate, cashFlow)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestPresentValuePerpetuityDue(t *testing.T) {
	var interestRate float64 = 0.1
	var cashFlow float64 = 100
	var expected float64 = 1000
	actual := presentValuePerpetuityDue(interestRate, cashFlow)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestPresentValueGrowingPerpetuity(t *testing.T) {
	var interestRate float64 = 0.1
	var growthRate float64 = 0.1
	var cashFlow float64 = 100
	var expected float64 = 1000
	actual := presentValueGrowingPerpetuity(interestRate, growthRate, cashFlow)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestPresentValueGrowingPerpetuityDue(t *testing.T) {
	var interestRate float64 = 0.1
	var growthRate float64 = 0.1
	var cashFlow float64 = 100
	var expected float64 = 1000
	actual := presentValueGrowingPerpetuityDue(interestRate, growthRate, cashFlow)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestFutureValue(t *testing.T) {
	var presentValue float64 = 100
	var interestRate float64 = 0.1
	var periods int = 2
	var expected float64 = 121
	actual := futureValue(presentValue, interestRate, periods)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestPresentValueGrowingAnnuity(t *testing.T) {
	var interestRate float64 = 0.1
	var growthRate float64 = 0.1
	var periods int = 2
	cashFlows := []float64{100, 100}
	var expected float64 = 121
	actual := presentValueGrowingAnnuity(interestRate, growthRate, periods, cashFlows)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestPresentValueGrowingAnnuityDue(t *testing.T) {
	var interestRate float64 = 0.1
	var growthRate float64 = 0.1
	var periods int = 2
	cashFlows := []float64{100, 100}
	var expected float64 = 121
	actual := presentValueGrowingAnnuityDue(interestRate, growthRate, periods, cashFlows)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}


func TestInterestRate(t *testing.T) {
	var presentValue float64 = 100
	var futureValue float64 = 121
	var periods int = 2
	var expected float64 = 0.1
	actual := interestRate(presentValue, futureValue, periods)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestInterestRateContinuousCompounding(t *testing.T) {
	var presentValue float64 = 100
	var futureValue float64 = 121
	var periods int = 2
	var expected float64 = 0.1
	actual := interestRateContinuousCompounding(presentValue, futureValue, periods)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestInterestRatePerpetuity(t *testing.T) {
	var presentValue float64 = 100
	var cashFlow float64 = 100
	var expected float64 = 1
	actual := interestRatePerpetuity(presentValue, cashFlow)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestInterestRateGrowingPerpetuity(t *testing.T) {
	var presentValue float64 = 100
	var cashFlow float64 = 100
	var growthRate float64 = 0.1
	var expected float64 = 1.1
	actual := interestRateGrowingPerpetuity(presentValue, cashFlow, growthRate)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestInterestRateGrowingAnnuity(t *testing.T) {
	var presentValue float64 = 100
	var cashFlows float64 = 100
	var growthRate float64 = 0.1
	var expected float64 = 1.1
	actual := interestRateGrowingAnnuity(presentValue, cashFlows, growthRate)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestInterestRateGrowingAnnuityDue(t *testing.T) {
	var presentValue float64 = 100
	var cashFlows float64
	var growthRate float64 = 0.1
	var expected float64 = 1.1
	actual := interestRateGrowingAnnuityDue(presentValue, cashFlows, growthRate)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestInterestRateAnnuity(t *testing.T) {
	var presentValue float64 = 100
	var cashFlows float64
	var expected float64 = 1
	actual := interestRateAnnuity(presentValue, cashFlows)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestInterestRateAnnuityDue(t *testing.T) {
	var presentValue float64 = 100
	var cashFlows float64
	var expected float64 = 1
	actual := interestRateAnnuityDue(presentValue, cashFlows)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestInterestRatePerpetuityDue(t *testing.T) {
	var presentValue float64 = 100
	var cashFlow float64 = 100
	var expected float64 = 1
	actual := interestRatePerpetuityDue(presentValue, cashFlow)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestInterestRateGrowingPerpetuityDue(t *testing.T) {
	var presentValue float64 = 100
	var cashFlow float64 = 100
	var growthRate float64 = 0.1
	var expected float64 = 1.1
	actual := interestRateGrowingPerpetuityDue(presentValue, cashFlow, growthRate)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func compareFloat64(a, b float64) bool {
	return math.Abs(a - b) < 0.000
}
