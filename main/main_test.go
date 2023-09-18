package main

import "testing"

func TestCreateInvestment(t *testing.T) {
	t.Run("Test creating valid investment", func(t *testing.T) {
		user := User{userName: "Jacob Kepford"}
		investment := Investment{currentAmount: 100000, isEmployerMatch: true, employerMatchPercent: 7.5, investmentPercent: 11}

		user.CreateInvestment(investment)

		if len(user.investmentList) == 0 {
			t.Error("Investment was not found for provided user")
		}
	})

	t.Run("Test creating investment with invalid employer match", func(t *testing.T) {
		user := User{userName: "Jacob Kepford"}
		investment := Investment{currentAmount: 100000, isEmployerMatch: false, employerMatchPercent: 7.5, investmentPercent: 11}

		user.CreateInvestment(investment)

		if len(user.investmentList) > 0 {
			t.Error("Employer match is false but employer match percent is greater than 0")
		}
	})

}
