package models

import (
	"testing"
)

func TestSysConfigCreate(t *testing.T) {
	InitAllInTest()

	sysConfig := &SysConfig{SysKey: "aaa", SysValue: "val"}
	if _, err := sysConfig.Create(sysConfig); err != nil {
		t.Error("Create() failed. Error:", err)
	}
}
func TestSysConfigRemove(t *testing.T) {
	InitAllInTest()

	var sysConfig SysConfig
	if err := sysConfig.Remove(3); err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}
func TestSysConfigUpdate(t *testing.T) {
	InitAllInTest()

	sysConfig := &SysConfig{Id: 1, SysKey: "aaa", SysValue: "aaa"}
	if err := sysConfig.Update(sysConfig); err != nil {
		t.Error("Update() failed. Error:", err)
	}
}
func TestSysConfigGetById(t *testing.T) {
	InitAllInTest()

	sysConfig := &SysConfig{SysKey: "key", SysValue: "val"}
	SysConfig{}.Create(sysConfig)

	getSysConfig, err := SysConfig{}.GetById(sysConfig.Id)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getSysConfig != *sysConfig {
		t.Error("GetById() failed:", "%v != %v", sysConfig, getSysConfig)
	}
}
