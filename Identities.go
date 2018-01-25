package exmodels

import (
    "time"
    "golang.org/x/crypto/bcrypt"
)

type Identities struct {
    Id                uint   `gorm:"primary_key"`
    Email             string `gorm:"size:255"`
    PasswordDigest    string `gorm:"size:255"`
    IsActive          bool   `sql:"default: null"`
    RetryCount        int    `sql:"default: null"`
    IsLocked          bool   `sql:"default: null"`
    IsTotpEnabled     bool   `sql:"default: false"`
    IsPhone2FAEnabled bool   `sql:"default: false, name:is_phone_2fa_enabled"`
    TotpSecret        string `gorm:"size:255"`

    LockedAt     *time.Time `sql:"default: null"`
    LastVerifyAt *time.Time `sql:"default: null"`
    CreatedAt    *time.Time `sql:"default: null"`
    UpdatedAt    *time.Time `sql:"default: null"`
}

func (this *Identities) SetHashedPassword(password string) error {
    hashByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    this.PasswordDigest = string(hashByte)
    return err
}

func (this *Identities) CheckPassword(password string) error {
    return bcrypt.CompareHashAndPassword([]byte(this.PasswordDigest), []byte(password))
}