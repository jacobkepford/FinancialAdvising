package main

type User struct {
	userName       string
	investmentList []Investment
}

func (u *User) CreateInvestment(investment Investment) {
	if !investment.isEmployerMatch && investment.employerMatchPercent > 0 {
		return
	}

	u.investmentList = append(u.investmentList, investment)
}

type Investment struct {
	currentAmount        int
	isEmployerMatch      bool
	employerMatchPercent float32
	investmentPercent    float32
}
