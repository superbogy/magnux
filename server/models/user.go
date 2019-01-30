package models

type User struct {
    ID uint `gorm:"primary_key"`
    Username string `gorm:"size:100""`
    Email string `gorm:"size:150"`
    Password string
    Nickname string

}

func (user *User) Get(filter Filter) *User  {
    DB.Where(filter).Find(&user)

    return user
}



func (user *User) update() error  {
    err := DB.Save(&user).Error

    return err
}



func (user *User) insert() error  {
    err := DB.Create(&user).Error

    return err
}


func (user *User) delete() error  {
    err := DB.Delete(&user, "id=?", user.ID).Error

    return err
}


func GetUser (filter interface{}) *User {
    var user User
    DB.Debug().Where(filter).First(&user)

    return &user
}
