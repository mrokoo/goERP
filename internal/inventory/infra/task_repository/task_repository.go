package task_repository

import (
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/task"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	db.AutoMigrate(&MySQLTask{})
	db.AutoMigrate(&MySQLTaskItem{})
	db.AutoMigrate(&MySQLRecord{})
	db.AutoMigrate(&MySQLRecordItem{})
	return &TaskRepository{db: db}
}

func (t *TaskRepository) GetAll() ([]*task.Task, error) {
	var tasks []*MySQLTask
	if err := t.db.Preload(clause.Associations).Preload("Recrods.Items").Find(&tasks).Error; err != nil {
		return nil, err
	}
	var tasks2 []*task.Task
	for _, ms := range tasks {
		task_ := ms.toTask()
		tasks2 = append(tasks2, &task_)
	}
	return tasks2, nil
}

func (t *TaskRepository) GetByID(ID string) (*task.Task, error) {
	var task MySQLTask
	task.ID = ID
	if err := t.db.Preload(clause.Associations).Preload("Recrods.Items").First(&task).Error; err != nil {
		return nil, err
	}
	task_ := task.toTask()
	return &task_, nil
}

// Save 函数将任务保存到数据库中。
// 如果保存失败，则返回一个错误。
func (t *TaskRepository) Save(task *task.Task) error {
	// 将任务转换为 MySQL 格式。
	task_ := toMySQLTask(*task)

	// 开始事务。
	tx := t.db.Begin()

	// 检查任务是否已存在。
	if _, err := t.GetByID(task_.ID); err != nil {
		if err == gorm.ErrRecordNotFound {
			// 如果任务不存在，则创建新任务。
			if err := tx.Create(&task_).Error; err != nil {
				tx.Rollback()
				return err
			}
			return tx.Commit().Error
		}
		return err
	}

	// 更新现有任务。
	if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&task_).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
