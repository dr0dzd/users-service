package user

import (
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository{
	return &Repository{DB: db}
}

func (r *Repository) CreateUser(u User) (User, error) {
	err := r.DB.Table("users").Create(&u).Error
	if err != nil{
		return User{}, status.Error(codes.AlreadyExists, "email should be unic")
	}
	return u, nil
}

func (r *Repository) GetUser(userID uuid.UUID) (User, error) {
	user := User{}
	err := r.DB.Table("users").
			Select("id", "email", "created_at", "deleted_at").
			Where("id = ?", userID).
			Where("deleted_at is null").
			First(&user)
	if err.RowsAffected == 0 || err.Error != nil{
		return User{}, status.Error(codes.NotFound, "no user with this id")
	}

	return user, nil
}

func (r *Repository) GetAllUsers() ([]User, error){
	result := []User{}
	err := r.DB.Table("users").
			Select("id", "email", "created_at", "updated_at").
			Find(&result).Error
	if err != nil{
		return []User{}, status.Error(codes.NotFound, "no users or invalid error")
	}
	return result, nil
}

func (r *Repository) UpdateUser(u User) (User, error) {
	err := r.DB.Table("users").
			Where("id = ?", u.ID).
			Where("deleted_at is null").
			Update("email", u.Email)
	if err.Error != nil || err.RowsAffected == 0{
		return User{}, status.Error(codes.NotFound, "no user with this")
	}
	return u, nil
}

func (r *Repository) DeleteUser(userID uuid.UUID) error {
	return r.DB.Table("users").Delete(&User{}, userID).Error
}