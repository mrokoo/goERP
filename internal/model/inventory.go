package model

import "time"

type InventoryFlow struct {
	ID          string    `gorm:"primaryKey;size:191;"`
	TaskID      *string   `gorm:"size:191;"`
	TakeID      *string   `gorm:"size:191;"`
	ProductID   string    `gorm:"size:191;"`
	Product     Product   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	WarehouseID string    `gorm:"size:191;"`
	Warehouse   Warehouse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Flow        string
	Previous    int // Previous Quantity
	Change      int // Change Quantity
	Present     int // Present Quantity
	CreatedAt   time.Time
}

type Task struct {
	ID              string    `gorm:"primaryKey;size:191;"`
	WarehouseID     string    `gorm:"size:191; not null"`
	Warehouse       Warehouse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"` // 外键约束
	Kind            string
	State           string
	Items           []TaskItem   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Recrods         []TaskRecord `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;default:null;"`
	IO              bool
	PurchaseOrderID *string `gorm:"size:191;default:null;"`
	// PurchaseOrder         PurchaseOrder       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"` // 外键约束
	PurchaseReturnOrderID *string `gorm:"size:191;default:null;"`
	// PurchaseReturnOrder   PurchaseReturnOrder `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"` // 外键约束
	SaleOrderID       *string `gorm:"size:191;"`
	SaleReturnOrderID *string `gorm:"size:191;"`
	AllotID           *string `gorm:"size:191;"`
	Allot             Allot   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	CreatedAt         time.Time
}

type TaskItem struct {
	ID        string  `gorm:"size:191;primaryKey"`
	TaskID    string  `gorm:"size:191;primaryKey"` // 外键约束
	ProductID string  `gorm:"size:191;primaryKey"`
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Total     int
	Quantity  int
}

type TaskRecord struct {
	ID          string    `gorm:"size:191;primaryKey"`
	TaskID      string    `gorm:"size:191;primaryKey"` // 外键约束
	WarehouseID string    `gorm:"size:191;"`
	Warehouse   Warehouse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"` //外键约束
	UserID      string    `gorm:"size:191;"`
	User        User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"` // 外键约束
	State       string
	Items       []TaskRecordItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time
}

type TaskRecordItem struct {
	ID           string  `gorm:"size:191;primaryKey"`
	TaskRecordID string  `gorm:"size:191;primaryKey"` // 外键约束
	ProductID    string  `gorm:"size:191;primaryKey"`
	Product      Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"` // 外键约束
	Quantity     int
}

type Allot struct {
	ID             string `gorm:"primaryKey;"`
	InWarehouseID  string
	OutWarehouseID string
	UserID         string `gorm:"size:191;"`
	User           User   `gorm:"contraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	CreatedAt      time.Time
	Items          []AllotItem
}

type AllotItem struct {
	ID        string  `gorm:"primaryKey;"`
	AllotID   string  `gorm:"size:191;"`
	ProductID string  `gorm:"size:191;"`
	Product   Product `gorm:"contraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Quantity  int
}

type Take struct {
	ID          string    `gorm:"primaryKey"`
	WarehouseID string    `gorm:"size:191;"`
	Warehouse   Warehouse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	UserID      string    `gorm:"size:191;"`
	User        User      `gorm:"contraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	CreateAt    time.Time
	Items       []TakeItem
}

type TakeItem struct {
	ID        string  `gorm:"primaryKey"`
	TakeID    string  `gorm:"size:191;"`
	ProductID string  `gorm:"size:191;"`
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Quantity  int
}
