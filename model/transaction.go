package model

import "time"

type Transaction struct {
	BaseModel
	TransactionDate time.Time
	VehicleID       string
	Vehicle         Vehicle `gorm:"foreignKey:VehicleID"`
	CustomerID      string
	Customer        Customer `gorm:"foreignKey:CustomerID"`
	EmployeeID      string
	Employee        Employee `gorm:"foreignKey:EmployeeID"`
	Type            string   `gorm:"check:type IN ('online', 'offline')"`
	Qty             int
	PaymentAmount   int64
}

func (t *Transaction) IsValidType() bool {
	return t.Type == "online" || t.Type == "offline"
}

func (Transaction) TableName() string {
	return "trx_transaction"
}
