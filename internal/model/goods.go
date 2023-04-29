package model

type Unit struct {
	ID   string `gorm:"primaryKey;<-:create;size:191"`
	Name string `gorm:"not null"`
	Note string
}

type Category struct {
	ID   string `gorm:"primaryKey;<-:create;size:191"`
	Name string `gorm:"not null"`
	Note string
}

type Product struct {
	ID           string         `gorm:"primaryKey"`
	Name         string         `gorm:"not null"`
	CategoryID   *string        `gorm:"default:null"`
	Category     Category       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UnitID       *string        `gorm:"default:null"`
	Unit         Unit           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OpeningStock []OpeningStock `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	State        string         `gorm:"default:active"`
	Note         string
	Img          string
	Intro        string
	Purchase     float64
	Retail       float64
	Grade1       float64
	Grade2       float64
	Grade3       float64
}

type OpeningStock struct {
	ProductID string `gorm:"size:191;primaryKey"`
	WarehouseID string    `gorm:"not null; size:191;primaryKey"`
	Warehouse   Warehouse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Amount      int
}
