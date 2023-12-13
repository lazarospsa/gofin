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
	actual := NetPresentValue(interestRate, periods, cashFlows)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestPresentValue(t *testing.T) {
	var futureValue float64 = 100
	var interestRate float64 = 0.1
	var periods int = 2
	var expected float64 = 121
	actual := PresentValue(futureValue, interestRate, periods)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestPresentValueAnnuity(t *testing.T) {
	var interestRate float64 = 0.1
	var periods int = 2
	cashFlows := []float64{100, 100}
	var expected float64 = 121
	actual := PresentValueAnnuity(interestRate, periods, cashFlows)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestPresentValueAnnuityDue(t *testing.T) {
	var interestRate float64 = 0.1
	var periods int = 2
	cashFlows := []float64{100, 100}
	var expected float64 = 121
	actual := PresentValueAnnuityDue(interestRate, periods, cashFlows)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestPresentValuePerpetuity(t *testing.T) {
	var interestRate float64 = 0.1
	var cashFlow float64 = 100
	var expected float64 = 1000
	actual := PresentValuePerpetuity(interestRate, cashFlow)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestFutureValueAnnuity(t *testing.T) {
	var payment float64 = 100
	var interestRate float64 = 0.1
	var periods int = 2
	var expected float64 = 242
	actual := FutureValueAnnuity(payment, interestRate, periods)

	if compareFloat64(actual, expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestPresentValuePerpetuityDue(t *testing.T) {
	var interestRate float64 = 0.1
	var cashFlow float64 = 100
	var expected float64 = 1000
	actual := PresentValuePerpetuityDue(interestRate, cashFlow)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestPresentValueGrowingPerpetuity(t *testing.T) {
	var interestRate float64 = 0.1
	var growthRate float64 = 0.1
	var cashFlow float64 = 100
	var expected float64 = 1000
	actual := PresentValueGrowingPerpetuity(interestRate, growthRate, cashFlow)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestPresentValueGrowingPerpetuityDue(t *testing.T) {
	var interestRate float64 = 0.1
	var growthRate float64 = 0.1
	var cashFlow float64 = 100
	var expected float64 = 1000
	actual := PresentValueGrowingPerpetuityDue(interestRate, growthRate, cashFlow)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestFutureValue(t *testing.T) {
	var presentValue float64 = 100
	var interestRate float64 = 0.1
	var periods int = 2
	var expected float64 = 121
	actual := FutureValue(presentValue, interestRate, periods)

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
	actual := PresentValueGrowingAnnuity(interestRate, growthRate, periods, cashFlows)

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
	actual := PresentValueGrowingAnnuityDue(interestRate, growthRate, periods, cashFlows)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}


func TestInterestRate(t *testing.T) {
	var presentValue float64 = 100
	var futureValue float64 = 121
	var periods int = 2
	var expected float64 = 0.1
	actual := InterestRate(presentValue, futureValue, periods)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestInterestRateContinuousCompounding(t *testing.T) {
	var presentValue float64 = 100
	var futureValue float64 = 121
	var periods int = 2
	var expected float64 = 0.1
	actual := InterestRateContinuousCompounding(presentValue, futureValue, periods)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestInterestRatePerpetuity(t *testing.T) {
	var presentValue float64 = 100
	var cashFlow float64 = 100
	var expected float64 = 1
	actual := InterestRatePerpetuity(presentValue, cashFlow)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestInterestRateGrowingPerpetuity(t *testing.T) {
	var presentValue float64 = 100
	var cashFlow float64 = 100
	var growthRate float64 = 0.1
	var expected float64 = 1.1
	actual := InterestRateGrowingPerpetuity(presentValue, cashFlow, growthRate)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestInterestRateGrowingAnnuity(t *testing.T) {
	var presentValue float64 = 100
	var cashFlows float64 = 100
	var growthRate float64 = 0.1
	var expected float64 = 1.1
	actual := InterestRateGrowingAnnuity(presentValue, cashFlows, growthRate)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestInterestRateGrowingAnnuityDue(t *testing.T) {
	var presentValue float64 = 100
	var cashFlows float64
	var growthRate float64 = 0.1
	var expected float64 = 1.1
	actual := InterestRateGrowingAnnuityDue(presentValue, cashFlows, growthRate)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestInterestRateAnnuity(t *testing.T) {
	var presentValue float64 = 100
	var cashFlows float64
	var expected float64 = 1
	actual := InterestRateAnnuity(presentValue, cashFlows)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestInterestRateAnnuityDue(t *testing.T) {
	var presentValue float64 = 100
	var cashFlows float64
	var expected float64 = 1
	actual := InterestRateAnnuityDue(presentValue, cashFlows)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestInterestRatePerpetuityDue(t *testing.T) {
	var presentValue float64 = 100
	var cashFlow float64 = 100
	var expected float64 = 1
	actual := InterestRatePerpetuityDue(presentValue, cashFlow)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func TestInterestRateGrowingPerpetuityDue(t *testing.T) {
	var presentValue float64 = 100
	var cashFlow float64 = 100
	var growthRate float64 = 0.1
	var expected float64 = 1.1
	actual := InterestRateGrowingPerpetuityDue(presentValue, cashFlow, growthRate)

	if compareFloat64(actual,expected) {
		t.Errorf("Test failed, expected: '%f', got: '%f'", expected, actual)
	}
}

func compareFloat64(a, b float64) bool {
	return math.Abs(a - b) < 0.000
}
