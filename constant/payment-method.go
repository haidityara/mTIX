package constant

var (
	PaymentBank    = 1
	PaymentCC      = 2
	PaymentEWallet = 3
)

var PaymentMethods = map[int]string{
	PaymentBank:    "Bank",
	PaymentCC:      "Credit Card",
	PaymentEWallet: "eWallet",
}
