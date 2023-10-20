package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username string `gorm:"unique;not null"`
    Password string
    Role     string
}

type Bicycle struct {
    gorm.Model
    Name        string
    Type        string
    Seat        int
    PricePerHour float64
}

type Rental struct {
    gorm.Model
    BicycleID       uint
    UserID          uint
    RentalStartTime string
    RentalEndTime   string
    TotalPrice      float64
}

type UserDetail struct {
    gorm.Model
    UserID          uint
    Nama            string
    Alamat          string
    NoTelp          string
    JenisKelamin    string
    TanggalTempatLahir string
}

