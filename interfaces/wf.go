package main 

type WellsFargo struct {
	balance int
}

func NewWellsFargo() *WellsFargo {
	return &WellsFargo{
		balance: 0,
	}
}

func(w *WellsFargo) GetBalance() int {
	panic("Implement me")
	return w.balance 
}
func(w *WellsFargo) Deposit(amount int) {
	panic("Implement me")
}
func(w *WellsFargo) WithDraw(amount int) error {
	panic("Implement me")
}