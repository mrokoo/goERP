package task

import (
	"testing"

	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/record"
	"github.com/stretchr/testify/assert"
)

func TestTask_UpdateTaskItems(t *testing.T) {
	assert := assert.New(t)
	a := NewTask("W001", IN_PURCHASE, "PO001", []TaskItem{
		NewTaskItem("P002", 300),
	})
	r := record.NewRecord("W001", "U001", []record.RecordItem{
		record.NewRecordItem("P002", 100),
	})
	a.AddRecord(r)
	assert.Equal(100, a.Items[0].Quantity)

}
