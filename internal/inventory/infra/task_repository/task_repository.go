package task_repository

import (
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/task"
	"gorm.io/gorm"
)

type InTaskRepository struct {
	db *gorm.DB
}

/*
NewInTaskRepository initializes and returns an InTaskRepository struct
that uses the given *gorm.DB object to interact with the database. This
function also automatically migrates the InTask, InRecord, and InItem
tables to ensure they match their corresponding Go struct definitions.

Parameters:
- db (*gorm.DB): the GORM database object to use for interactions with
  the database.

Returns:
- (*InTaskRepository): a pointer to a new InTaskRepository struct that
  uses the given *gorm.DB object.
*/
func NewInTaskRepository(db *gorm.DB) *InTaskRepository {
	db.AutoMigrate(&InTask{})
	db.AutoMigrate(&InRecord{})
	db.AutoMigrate(&InItem{})
	return &InTaskRepository{
		db: db,
	}
}

/*
GetAll retrieves all InTask objects from the repository.

Returns:
- []*task.InTask: A slice of InTask objects, converted to task.InTask objects.
- error: An error if the operation failed, nil otherwise.
*/
func (r InTaskRepository) GetAll() ([]*task.InTask, error) {
	var inTasks []InTask
	result := r.db.Preload("Records").Find(&inTasks)
	if err := result.Error; err != nil {
		return nil, err
	}
	var inTaskList []*task.InTask
	for _, inTask := range inTasks {
		inTaskList = append(inTaskList, inTask.toTask())
	}
	return inTaskList, nil
}

func (r InTaskRepository) GetByID(ID string) (*task.InTask, error) {
	inTask := InTask{
		ID: ID,
	}
	result := r.db.First(&inTask)
	if err := result.Error; err != nil {
		return nil, err
	}
	return inTask.toTask(), nil
}

func (r InTaskRepository) Save(inTask *task.InTask) error {
	inTask_ := toMySQLInTask(inTask)
	result := r.db.Create(inTask_)
	return result.Error
}

type OutTaskRepository struct {
	db *gorm.DB
}

func NewOutTaskRepository(db *gorm.DB) *OutTaskRepository {
	db.AutoMigrate(&OutTask{})
	db.AutoMigrate(&OutRecord{})
	db.AutoMigrate(&OutItem{})
	return &OutTaskRepository{
		db: db,
	}
}

func (r OutTaskRepository) GetAll() ([]*task.OutTask, error) {
	var outTasks []OutTask
	result := r.db.Preload("Records").Find(&outTasks)
	if err := result.Error; err != nil {
		return nil, err
	}
	var outTaskList []*task.OutTask
	for _, outTask := range outTasks {
		outTaskList = append(outTaskList, outTask.toTask())
	}
	return outTaskList, nil
}

func (r OutTaskRepository) GetByID(ID string) (*task.OutTask, error) {
	outTask := OutTask{
		ID: ID,
	}
	result := r.db.First(&outTask)
	if err := result.Error; err != nil {
		return nil, err
	}
	return outTask.toTask(), nil
}

func (r OutTaskRepository) Save(outTask *task.OutTask) error {
	outTask_ := toMySQLOutTask(outTask)
	result := r.db.Create(outTask_)
	return result.Error
}
