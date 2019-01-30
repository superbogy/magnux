package models

type Photo struct {
    ID uint `gorm:"primary_key"`
    Name string `gorm:"size:100""`
    UserId int
    Size int
    Attr string
    Path string `gorm:"size:150"`
    TiTle string
    Nickname string

}

func (photo *Photo) Get(filter Filter) *Photo  {
    DB.Where(filter).Find(&photo)

    return photo
}



func (photo *Photo) update() error  {
    err := DB.Save(&photo).Error

    return err
}



func (photo *Photo) insert() error  {
    err := DB.Create(&photo).Error

    return err
}


func (photo *Photo) delete() error  {
    err := DB.Delete(&photo, "id=?", photo.ID).Error

    return err
}


func GetPhoto (page int, size int) []*Photo {
    var photos []*Photo
    offset := (page - 1) * size
    DB.Debug().Offset(offset).Limit(size).Find(&photos)

    return photos
}
