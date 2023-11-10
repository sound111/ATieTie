package snowflakes

import (
	"Web/settings"
	"time"

	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake
var id uint16

func getMachineId() (uint16, error) {
	return id, nil
}

func Init() (err error) {
	//time.Parse 必须用2006,01,02,15,04,05这个时间
	t, err := time.Parse("2006-01-02", settings.Conf.SnowConfig.StartTime)
	if err != nil {
		return
	}
	id = settings.Conf.SnowConfig.MachineID

	sonySettings := sonyflake.Settings{
		StartTime: t,
		MachineID: getMachineId,
	}

	sf = sonyflake.NewSonyflake(sonySettings)

	return
}

func GetID() (id uint64, err error) {
	id, err = sf.NextID()
	return
}
