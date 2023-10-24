package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username   string
    Password   string
    Role       int
    Rental     []Rental
}

type UserWithDetail struct {
    gorm.Model
    Username   string
    Password   string
    Role       int
    UserDetail UserDetail
}

type Bicycle struct {
    gorm.Model
    Name        string
    Type        string
    Seat        string
    PricePerHour int
}

type Rental struct {
    gorm.Model
    BicycleID       uint
    UserID          uint
    RentalStartTime string
    RentalEndTime   string
    TotalPrice      float64
    User            User
    Bicycle         Bicycle
}

type UserDetail struct {
    gorm.Model
    UserID             uint
    Nama               string
    Alamat             string
    NoTelp             string
    JenisKelamin       string
    TanggalTempatLahir string
}