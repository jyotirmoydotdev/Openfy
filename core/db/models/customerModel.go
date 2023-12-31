package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID                     uint              `gorm:"column:id;primaryKey"`
	Password               string            `gorm:"column:password;omitempty"`
	DisplayName            string            `gorm:"column:display_name"`
	FirstName              string            `gorm:"column:firstName"`
	LastName               string            `gorm:"column:lastName"`
	Email                  string            `gorm:"column:email;index"`
	Locale                 string            `gorm:"column:locale"`
	TaxExempt              bool              `gorm:"column:taxExempt"`
	Phone                  int               `gorm:"column:phone"`
	State                  string            `gorm:"column:state"`
	Age                    int               `gorm:"column:age"`
	DeliveryAddresses      []DeliveryAddress `gorm:"foreignKey:UserID"`
	UserCreatTime          string            `gorm:"column:userCreateTime"`
	LifetimeDuration       string            `gorm:"column:lifetimeDuration"`
	TotalSpentAmount       float64           `gorm:"column:totalSpentAmount"`
	TotalSpentCurrencyCode string            `gorm:"column:totalSpentCurrencyCode"`
	NumberOfOrders         int               `gorm:"column:numberOfOrders"`
	LastOrderId            uint              `gorm:"column:lastOrderId"`
	LastOrderCreatedAt     string            `gorm:"column:lastOrderCreatedAt"`
}

type DeliveryAddress struct {
	ID            uint   `gorm:"column:id;primaryKey"`
	UserID        uint   `gorm:"column:user_id;index"`
	FormattedArea string `gorm:"column:formattedArea"`
	FirstName     string `gorm:"column:firstName"`
	LastName      string `gorm:"column:lastName"`
	Company       string `gorm:"column:company"`
	Address1      string `gorm:"column:address1"`
	Address2      string `gorm:"column:address2"`
	Apartment     string `gorm:"column:apartment"`
	City          string `gorm:"column:city"`
	Province      string `gorm:"column:province"`
	Country       string `gorm:"column:country"`
	Phone         int    `gorm:"column:phone"`
	Zip           int    `gorm:"column:zip"`
}

type UserSecrets struct {
	ID     uint   `gorm:"primaryKey" json:"id" `
	UserID uint   `gorm:"column:user_id"`
	Email  string `gorm:"column:email"`
	Secret string `gorm:"column:secret"`
}

type UserToken struct {
	ID                uint      `gorm:"primaryKey" json:"id" `
	Email             string    `gorm:"column:email;primaryKey"`
	Token             string    `gorm:"column:token"`
	LastUsed          time.Time `gorm:"column:last_used"`
	TokenExpiry       time.Time `gorm:"column:token_expiry"`
	IsActive          bool      `gorm:"column:is_active"`
	IPAddresses       string    `gorm:"column:ip_addresses"`
	UserAgent         string    `gorm:"column:user_agent"`
	DeviceInformation string    `gorm:"column:device_information"`
	RevocationReason  string    `gorm:"column:revocation_reason"`
}

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{db: db}
}

func (ur *UserModel) Save(user *Customer) error {
	return ur.db.Create(user).Error
}

func (ur *UserModel) SaveUserSecret(UserSecret *UserSecrets) error {
	return ur.db.Create(UserSecret).Error
}

func GetUserSecretKeyByEmail(db *gorm.DB, email string) (string, error) {
	var userSecrets UserSecrets
	err := db.Model(&UserSecrets{}).Select("secret").Where("email = ?", email).First(&userSecrets).Error
	if err != nil {
		return "", err
	}
	return userSecrets.Secret, nil
}

func GetUserHashedPasswordByEmail(db *gorm.DB, email string) (string, error) {
	var UserHashedPasswor Customer
	err := db.Model(&Customer{}).Select("Password").Where("email = ?", email).First(&UserHashedPasswor).Error
	if err != nil {
		return "", err
	}
	return UserHashedPasswor.Password, nil
}

func UserExistByEmail(db *gorm.DB, email string) (bool, error) {
	var count int64
	if err := db.Model(&Customer{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func SaveToken(db *gorm.DB, userToken *UserToken) error {
	return db.Create(userToken).Error
}

func CheckEmailExist(db *gorm.DB, email string) (bool, error) {
	var count int64
	if err := db.Model(&UserToken{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
func GetTokenByEmail(db *gorm.DB, email string) (string, error) {
	var userToken UserToken
	err := db.Model(&UserToken{}).Select("token").Where("email = ?", email).First(&userToken).Error
	if err != nil {
		return "", err
	}
	return userToken.Token, nil
}
func UpdateToken(db *gorm.DB, userToken *UserToken) error {
	updatedValues := map[string]interface{}{
		"email":              userToken.Email,
		"token":              userToken.Token,
		"last_used":          userToken.LastUsed,
		"token_expiry":       userToken.TokenExpiry,
		"is_active":          userToken.IsActive,
		"ip_addresses":       userToken.IPAddresses,
		"user_agent":         userToken.UserAgent,
		"device_information": userToken.DeviceInformation,
		"revocation_reason":  userToken.RevocationReason,
	}
	return db.Model(&UserToken{}).Where("email = ?", userToken.Email).Updates(updatedValues).Error
}

func (ad *UserModel) GetUserID(email string) (uint, error) {
	var user Customer
	if err := ad.db.Model(&Customer{}).Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, fmt.Errorf("user not found")
		}
		return 0, fmt.Errorf("error fetching user: %v", err)
	}
	return user.ID, nil
}
