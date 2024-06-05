package model

import (
	"go-web-new/global"
	"go-web-new/model/schemas"
	"go-web-new/utils/errmsg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Avatar   string `gorm:"type:varchar(40);not null" json:"avatar" label:"头像"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required" label:"角色"`
}

type Users []Users

func CheckUser(name string) (code int) {
	var users User
	//查询用户是否存在
	if users.ID > 0 {
		// 如果存在则引出错误
		return errmsg.ERROR_USERNAME_USED
	}
	// 否则正确
	return errmsg.SUCCSE
}

// 唯一性判断，如果是在编辑用户的时候，使用之前的逻辑，如果不修改username 只修改其他字段，则会一直卡主，因此重写解耦
func UniqUser(name string, id int) int {
	var user User
	global.Db.Select("id").Where("username=?", name).First(&user)
	if user.ID > 0 {
		if int(user.ID) == id {
			return errmsg.SUCCSE
		}
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCSE
}

// Add User
func (model *User) CreateUser() int {
	err := global.Db.Create(&model).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// get one user

func (user *User) GetUser(id int) int {
	err := global.Db.Select("username,avatar,role").Where("ID=?", id).First(&user).Error
	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCSE

}

// 用户列表
func (users *Users) GetUsers(username string, Size int, Page int) int64 {
	var total int64

	if username != "" {
		global.Db.Select("id,username,role").Limit(Size).Offset((Page - 1) * Size).Find(&users)
		global.Db.Model(&users).Where(
			"username LIKE ?", username+"%",
		).Count(&total)
		return total
	}

	// 分页
	err := global.Db.Select("id,username,role").Offset((Page - 1) * Size).Limit(Size).Find(&users).Error
	global.Db.Model(&users).Count(&total)
	if err != nil {
		return 0
	}
	// 返回用户的列表
	return total
}

// 登陆验证
func (user *User) CheckLogin(data schemas.Login) int {
	var PasswordErr error
	global.Db.Where("username=?", data.Username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if PasswordErr != nil {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCSE
}
