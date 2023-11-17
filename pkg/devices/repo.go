package devices

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	UpsertDevice(*Device) error
	GetDevicesByOwner(string) ([]Device, error)
	GetDeviceById(uint32) (*Device, error)
}

type repoImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repoImpl{
		db: db,
	}
}

func (r *repoImpl) UpsertDevice(device *Device) error {
	return r.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(device).Error
}

func (r *repoImpl) GetDevicesByOwner(owner string) ([]Device, error) {
	devices := make([]Device, 0)
	err := r.db.Model(&Device{}).Where("owner = ?", owner).Scan(&devices).Error
	return devices, err
}

func (r *repoImpl) GetDeviceById(id uint32) (*Device, error) {
	device := &Device{}
	err := r.db.Model(&Device{}).Where("id = ?", id).First(device).Error
	return device, err
}
