package id_gen

import (
	"fmt"

	"github.com/sony/sonyflake"
)

// id生成服务

var (
	sonyFlake *sonyflake.Sonyflake
	machineID uint16 // 真正的分布式环境下必须从zookeeper或者etcd中获取Consul
)

// 获取机器ID的回调函数
func getMachineID() (uint16, error) {
	return machineID, nil
}

// Init 初始化sonyFlake
func Init(mID uint16) (err error) {
	machineID = mID
	// 配置项
	st := sonyflake.Settings{}
	st.MachineID = getMachineID
	sonyFlake = sonyflake.NewSonyflake(st)
	return
}

// GetID 获取全局唯一ID
func GetID() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("must Call Init before GetID, err:%v\n", err)
		return
	}
	return sonyFlake.NextID()
}
