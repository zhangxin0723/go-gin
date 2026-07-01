package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	Id              int              `gorm:"primaryKey"`
	Name            string           `gorm:"size:16;not null;unique"`
	Age             int              `gorm:"default:18"`
	UserDetailModel *UserDetailModel `gorm:"foreignKey:UserId"`
	CreatedAt       time.Time
}

type UserDetailModel struct {
	Id        int        `gorm:"primaryKey"`
	UserId    int        `gorm:"not null;unique"` // 一对一情况下，需要加上唯一约束
	Email     string     `gorm:"size:32"`
	UserModel *UserModel `gorm:"foreignKey:UserId"`
}

func (u *UserModel) BeforeCreate(tx *gorm.DB) error {
	fmt.Println("创建的钩子函数")
	return nil
}

func (u *UserModel) BeforeUpdate(tx *gorm.DB) error {
	fmt.Println("更新的钩子函数")
	return nil
}

func (u *UserModel) BeforeDelete(tx *gorm.DB) error {
	fmt.Println("删除的钩子函数")
	return nil
}

func (u *UserModel) AfterFind(tx *gorm.DB) error {
	fmt.Println("查询后的钩子函数")
	return nil
}
