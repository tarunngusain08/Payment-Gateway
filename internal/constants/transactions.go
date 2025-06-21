package constants

type TransactionStatus string

const (
	StatusPending TransactionStatus = "PENDING"
	StatusSuccess TransactionStatus = "SUCCESS"
	StatusFailed  TransactionStatus = "FAILED"
)

type TransactionType string

const (
	TypeDeposit    TransactionType = "DEPOSIT"
	TypeWithdrawal TransactionType = "WITHDRAWAL"
)
