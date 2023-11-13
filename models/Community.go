package models

import "time"

type Community struct {
	CommunityId   uint64    `json:"community_id" db:"community_id"`
	CommunityName string    `json:"community_name" db:"community_name"`
	Introduction  string    `json:"introduction" db:"introduction"`
	CreateTime    time.Time `json:"create_time" db:"create_time"`
}
