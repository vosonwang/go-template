package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	User struct {
		ID             primitive.ObjectID `json:"_id" bson:"_id"`
		Username       string             `json:"username"` // 用户名和手机号都不允许重复
		Password       string             `json:"password"`
		HashPassword   []byte             `json:"-"`
		Email          string             `json:"email"`
		Mobile         uint64             `json:"mobile,string"`
		Status         int8               `json:"status,string"`                          // -1：停用  0:正常
		OrganizationID primitive.ObjectID `json:"organization_id" bson:"organization_id"` // 组织ID
		RoleID         primitive.ObjectID `json:"role_id" bson:"role_id"`
	}

	Users []User
)
