package main

import "testing"

func TestCreateInvestment(t *testing.T) {
	t.Run("Validate proper ID was created for investment", func(t *testing.T) {
		user := User{userName: "Jacob Kepford"}
		investment := Investment{currentAmount: 100000, isEmployerMatch: true, employerMatchPercent: 7.5, investmentPercent: 11}

		user.CreateInvestment(investment)

		got := user.investmentList[0].id
		want := 1

		if got != want {
			t.Errorf("Expected %d got %d", want, got)
		}
	})

	t.Run("Validate employer match set when isEmployerMatch is true", func(t *testing.T) {
		user := User{userName: "Jacob Kepford"}
		investment := Investment{currentAmount: 100000, isEmployerMatch: true, employerMatchPercent: 7.5, investmentPercent: 11}

		user.CreateInvestment(investment)

		got := user.investmentList[0].employerMatchPercent
		want := investment.employerMatchPercent

		if got != want {
			t.Errorf("Expected %f got %.2f", want, got)
		}
	})

	t.Run("Test creating investment with no employer match but employer match percent", func(t *testing.T) {
		user := User{userName: "Jacob Kepford"}
		investment := Investment{currentAmount: 100000, isEmployerMatch: false, employerMatchPercent: 7.5, investmentPercent: 11}

		user.CreateInvestment(investment)

		got := user.investmentList[0].employerMatchPercent

		if got != 0 {
			t.Errorf("Expected 0 got %.2f", got)
		}
	})

}
