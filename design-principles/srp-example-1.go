package design_principles

import "time"

// 社交产品

// 第一阶段：用户信息类的设计

type Stage1UserInfo struct {
	userId            string
	username          string
	email             string
	telephone         string
	createTime        time.Time
	lastLoginTime     time.Time
	avatarURL         string
	provinceOfAddress string
	cityOfAddress     string
	regionOfAddress   string
	detailOfAddress   string
}

// ------------------------------
// 第二阶段：增加类电商功能，需要将地址信息独立为一个类

type Stage2UserInfo struct {
	userId        string
	username      string
	email         string
	telephone     string
	createTime    time.Time
	lastLoginTime time.Time
	avatarURL     string
}

type Stage2UserAddress struct {
	province string
	city     string
	region   string
	detail   string
}

// ------------------------------
// 第三阶段：与多个App统一账户登录，需要将账户信息独立类一个类

type Stage3UserInfo struct {
	userId        string
	username      string
	createTime    time.Time
	lastLoginTime time.Time
	avatarURL     string
}

type Stage3UserAddress struct {
	province string
	city     string
	region   string
	detail   string
}

type Stage3Account struct {
	email     string
	telephone string
	password  string
}

