package model

import "github.com/jinzhu/gorm"

type SkeletonRequest struct {
	gorm.Model
	Param string
}
