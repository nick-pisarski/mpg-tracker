package fill_up_repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"mpg-tracker/api/database"
	. "mpg-tracker/api/models"
	"mpg-tracker/api/models/errors"
)

type FillUpRepository struct {
	Database *gorm.DB
}

func New(connectionString string) *FillUpRepository {
	db := database.New(connectionString)
	return &FillUpRepository{Database: db}
}

func Make(connectionString string) FillUpRepository {
	db := database.New(connectionString)
	return FillUpRepository{Database: db}
}

func (r *FillUpRepository) Create(entity FillUpEntity) (FillUpEntity, error) {
	log.Printf("called FillUpRepository.Create().\n")

	item := &entity
	if err := r.Database.Create(item).Error; err != nil {
		return *item, err
	}
	return *item, nil
}

func (r *FillUpRepository) Get() ([]FillUpEntity, error) {
	var results []FillUpEntity
	if err := r.Database.Find(&results).Error; err != nil {
		return results, err
	}
	return results, nil
}

func (r *FillUpRepository) GetById(id uint) (FillUpEntity, error) {
	var result FillUpEntity
	if err := r.Database.First(&result, id).Error; err != nil {
		return result, err
	}
	return result, nil
}

func (r *FillUpRepository) Put(id uint, entity FillUpEntity) (FillUpEntity, error) {
	result := &FillUpEntity{Model: gorm.Model{ID: id}}
	if err := r.Database.Model(&result).Clauses(clause.Returning{}).Where("deleted_at = ?", nil).Updates(entity).Error; err != nil {
		return *result, err
	}
	return *result, nil
}

func (r *FillUpRepository) Delete(id uint) (FillUpEntity, error) {
	var result FillUpEntity
	//if err := r.Database.Clauses(clause.Returning{}).Where("deleted_at = ?", nil).Delete(&result, id).Error; err != nil {
	//	return result, err
	//}
	return result, &errors.NotImplementedError{}
}
