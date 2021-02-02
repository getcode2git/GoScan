// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// UserLog is the golang structure for table user_log.
type UserLog struct {
    Id        int         `orm:"id"         json:"id"`         //   
    Username  string      `orm:"username"   json:"username"`   //   
    Ip        string      `orm:"ip"         json:"ip"`         //   
    UserAgent string      `orm:"user_agent" json:"user_agent"` //   
    CreateAt  *gtime.Time `orm:"create_at"  json:"create_at"`  //   
    UpdateAt  *gtime.Time `orm:"update_at"  json:"update_at"`  //   
}