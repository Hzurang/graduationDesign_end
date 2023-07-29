package config

import (
	"ginStudy/utils"
	sf "github.com/sony/sonyflake"
	"go.uber.org/zap"
	"time"
)

var (
	sonyFlake     *sf.Sonyflake
	sonyMachineID uint16
)

/*
getMachineID
@author: LJR
@Description: 取MachineID
@return uint16
@return error
*/
func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

/*
InitSnowFlake
@author: LJR
@Description: 需传入当前的机器ID
@param startTime
@param machineId
@return err
*/
func InitSnowFlake(startTime string, machineId uint16) (err error) {
	sonyMachineID = machineId
	t, _ := time.Parse("2006-01-02", startTime)
	setting := sf.Settings{
		StartTime: t,
		MachineID: getMachineID,
	}
	sonyFlake = sf.NewSonyflake(setting)
	return
}

/*
GenID
@author: LJR
@Description: 返回生成的id
@return id
@return err
*/
func GenID() (id uint64, err error) {
	if sonyFlake == nil {
		zap.L().Error(utils.GetCodeMsg(90008))
		return
	}
	id, err = sonyFlake.NextID()
	return
}
