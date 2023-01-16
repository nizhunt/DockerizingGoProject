package database

import (
	"gorm.io/gorm"
)

type Fact struct {
	gorm.Model
	Question string `json:"question" gorm:"text;not null;default:null"`
	Answer string `json:"answer" gorm:"text;not null;default:null"`
}

func GetAllFacts() ([]Fact, error) {
	var facts []Fact
	tx:= DB.Find(&facts)
	if tx.Error != nil{
		return []Fact{}, tx.Error
	}
	return facts,nil
}

func GetFact(id uint64) (Fact, error) {
	var fact Fact
	tx := DB.Where("id=?",id).First(&fact)
	if tx.Error != nil {
		return Fact{}, tx.Error
	}
	return fact, nil
}

func CreateFact(fact Fact) (error) {
	tx:= DB.Create(&fact)
	return tx.Error
}

func DeleteFact(id uint64) (error) {
	tx:= DB.Unscoped().Delete(&Fact{}, id)
	return tx.Error
}
