package models

import "github.com/jinzhu/gorm"

type Album struct {
    gorm.Model
    Name string `gorm:"size:100""`
    UserId int
    TiTle string
}

func (album *Album) Get(filter Filter) *Album  {
    DB.Where(filter).Find(&album)

    return album
}



func (album *Album) update() error  {
    err := DB.Save(&album).Error

    return err
}



func (album *Album) insert() error  {
    err := DB.Create(&album).Error

    return err
}


func (album *Album) delete() error  {
    err := DB.Delete(&album, "id=?", album.ID).Error

    return err
}


func AlbumPage (page int, size int) []*Album {
    var photos []*Album
    offset := (page - 1) * size
    DB.Debug().Offset(offset).Limit(size).Find(&photos)

    return photos
}

func AddAlbum(data interface{}) error {
    err := DB.Create(data).Error

    return err
}
