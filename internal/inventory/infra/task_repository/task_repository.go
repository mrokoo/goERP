package task_repository

import (
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/task"
	"github.com/mrokoo/goERP/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (t *TaskRepository) GetAll() ([]*task.Task, error) {
	var list []*model.Task
	if err := t.db.Preload(clause.Associations).Preload("Recrods.Items").Find(&list).Error; err != nil {
		return nil, err
	}
	var tasks []*task.Task
	for _, task := range list {
		tasks = append(tasks, toDomain(task))
	}
	return tasks, nil
}

func (t *TaskRepository) GetByID(ID string) (*task.Task, error) {
	var task model.Task
	task.ID = ID
	if err := t.db.Preload(clause.Associations).Preload("Recrods.Items").First(&task).Error; err != nil {
		return nil, err
	}
	return toDomain(&task), nil
}

// Save 函数将任务保存到数据库中。
// 如果保存失败，则返回一个错误。
func (t *TaskRepository) Save(task *task.Task) error {
	// 将任务转换为 MySQL 格式。
	task_ := toModel(task)

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

func (t *TaskRepository) GetTaskByPurchaseID(ID string, kind task.Kind) (*task.Task, error) {
	var task_ model.Task
	switch kind {
	case task.IN_PURCHASE:
		if err := t.db.Preload(clause.Associations).Preload("Recrods.Items").First(&task_, "purchase_order_id = ?", ID).Error; err != nil {
			return nil, err
		}
	case task.OUT_PURCHASE:
		if err := t.db.Preload(clause.Associations).Preload("Recrods.Items").First(&task_, "purchase_return_order_id = ?", ID).Error; err != nil {
			return nil, err
		}
	}
	return toDomain(&task_), nil
}
