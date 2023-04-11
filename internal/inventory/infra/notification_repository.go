package repository

import "gorm.io/gorm"

var ErrNotFound = gorm.ErrRecordNotFound

type Notification = domain.Notification

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	db.AutoMigrate(&Notification{}) //自动迁移
	return &NotificationRepository{
		db: db,
	}
}

func (r *NotificationRepository) GetAll() ([]*Notification, error) {
	var notifications []Notification
	result := r.db.Find(&notifications)
	if err := result.Error; err != nil {
		return nil, err
	}
	var notificationsp []*Notification
	for i := range notifications {
		notificationsp = append(notificationsp, &notifications[i])
	}
	return notificationsp, nil
}

func (r *NotificationRepository) GetByID(notificationID string) (*Notification, error) {
	notification := Notification{
		ID: notificationID,
	}
	result := r.db.First(&notification)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &notification, nil
}

func (r *NotificationRepository) Save(notification *domain.Notification) error {
	result := r.db.Create(notification)
	return result.Error
}

func (r *NotificationRepository) Replace(notification *domain.Notification) error {
	result := r.db.Save(notification)
	return result.Error
}

func (r *NotificationRepository) Delete(notificationID string) error {
	result := r.db.Delete(&Notification{
		ID: notificationID,
	})
	return result.Error
}
