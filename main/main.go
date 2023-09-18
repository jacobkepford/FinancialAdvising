package main

type User struct {
	userName       string
	investmentList []Investment
}

func (u *User) CreateInvestment(investment Investment) {
	if !investment.isEmployerMatch && investment.employerMatchPercent > 0 {
		investment.employerMatchPercent = 0
	}

	investment.id = len(u.investmentList) + 1

	u.investmentList = append(u.investmentList, investment)
}

type Investment struct {
	id                   int
	currentAmount        int
	isEmployerMatch      bool
	employerMatchPercent float32
	investmentPercent    float32
}
