package model

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) error {

	// goods
	if err := db.AutoMigrate(&Category{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&Unit{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&OpeningStock{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&Product{}); err != nil {
		return err
	}

	// share
	if err := db.AutoMigrate(&Warehouse{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&Supplier{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&Customer{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&Budget{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&Account{}); err != nil {
		return err
	}

	// system
	if err := db.AutoMigrate(&User{}); err != nil {
		return err
	}

	// inventory
	if err := db.AutoMigrate(&InventoryFlow{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&Task{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&TaskItem{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&TaskRecord{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&TaskRecordItem{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&Allot{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&AllotItem{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&Take{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&TakeItem{}); err != nil {
		return err
	}

	// sale
	if err := db.AutoMigrate(&SaleOrder{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&SaleOrderItem{}); err != nil {
		return err
	}

	// purchase
	if err := db.AutoMigrate(&PurchaseOrder{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&PurchaseOrderItem{}); err != nil {
		return err
	}
	
	return nil
}
