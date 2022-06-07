package fill_up_repository

import (
	"fmt"
	"mpg-tracker/api/models"
	"reflect"
	"testing"
	"time"
)

var sut FillUpRepository

func createFillUpEntityInDb(price float32, gallons float32, miles float32, dateFilled time.Time) *models.FillUpEntity {
	item := &models.FillUpEntity{
		Price:      price,
		Gallons:    gallons,
		Miles:      miles,
		DateFilled: dateFilled,
	}
	if err := sut.Database.Create(item).Error; err != nil {
		panic("failed to create test data")
	}
	return item
}
func assertDeepEquals(t *testing.T, expected any, result any) {
	if reflect.DeepEqual(expected, result) {
		t.Errorf("expected: %v, received: %v", expected, result)
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Initializing Subject Under Test: FillUpRepository")
	sut = Make("mpgtracker_test.db")

	m.Run()

	fmt.Println("Tearing down: deleting contents of db")
	sut.Database.Unscoped().Where("ID > 0").Delete(&[]models.FillUpEntity{})
}

func TestFillUpRepository_Create(t *testing.T) {
	t.Run("it should creates a FillUp", func(t *testing.T) {
		fillup := models.FillUpEntity{
			Price:      999.00,
			Gallons:    99.00,
			Miles:      999.99,
			DateFilled: time.Time{}.Local(),
		}
		result, err := sut.Create(fillup)
		if err != nil {
			t.Error(err)
		}
		if result.ID == 0 {
			t.Error("expected result.ID to NOT be 0")
		}
	})
}

func TestFillUpRepository_Get(t *testing.T) {
	fillUpEntity := createFillUpEntityInDb(100.00, 12.22, 333.33, time.Time{}.Local())

	t.Run("it should return a list of FillUpEntities", func(t *testing.T) {
		expected := []models.FillUpEntity{*fillUpEntity}

		result, err := sut.Get()
		if err != nil {
			t.Error(err)
		}

		assertDeepEquals(t, expected, result)
	})
}

func TestFillUpRepository_GetById(t *testing.T) {
	fillUpEntity := createFillUpEntityInDb(100.00, 12.22, 333.33, time.Time{}.Local())

	t.Run("it should return a FillUpEntity", func(t *testing.T) {
		result, err := sut.GetById(fillUpEntity.ID)
		if err != nil {
			t.Error(err)
		}

		assertDeepEquals(t, fillUpEntity, result)
	})
}
