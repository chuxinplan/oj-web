package models

import (
	. "github.com/open-fightcoder/oj-web/common/store"
)

type SysConfig struct {
	Id       int64  `form:"id" json:"id"`
	SysKey   string `form:"sysKey" json:"sysKey"`     //键
	SysValue string `form:"sysValue" json:"sysValue"` //值
}

func (this SysConfig) Create(sysConfig *SysConfig) (int64, error) {
	_, err := OrmWeb.Insert(sysConfig)
	if err != nil {
		return 0, err
	}
	return sysConfig.Id, nil
}

func (this SysConfig) Remove(id int64) error {
	sysConfig := new(SysConfig)
	_, err := OrmWeb.Id(id).Delete(sysConfig)
	return err
}

func (this SysConfig) Update(sysConfig *SysConfig) error {
	_, err := OrmWeb.AllCols().ID(sysConfig.Id).Update(sysConfig)
	return err
}

func (this SysConfig) GetById(id int64) (*SysConfig, error) {
	sysConfig := new(SysConfig)
	has, err := OrmWeb.Id(id).Get(sysConfig)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return sysConfig, nil
}
