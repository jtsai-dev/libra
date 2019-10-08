/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-12 15:01:48
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-16 18:08:13
 */
package models

import (
	"time"
)

const (
	DirNode = iota + 1
	OptNode
)

type Node struct {
	Id          int64
	WxAccountId int64
	Name        string
	Parent      int64
	Weight      float64
	NodeType    int
	Remark      string
	Status      int
	Created     time.Time
	Updated     time.Time
	Deleted     time.Time
}

type NodeIn struct {
	NodeType int `binding:"required"` // 1: dir; 2: opt;
	Parent   int64
	Name     string  `binding:"required"`
	Weight   float64 `binding:"min=0"`
}

type NodePutIn struct {
	Id     int64 `binding:"required"`
	Parent int64
	Name   string  `binding:"required"`
	Weight float64 `binding:"min=0"`
}

type NodeOut struct {
	Id       int64   `json:"id", binding:"required"`
	Name     string  `json:"name", binding:"required"`
	Parent   int64   `json:"parent", binding:"required"`
	Weight   float64 `json:"weight", binding:"min=0"`
	NodeType int     `json:"nodeType"` // 1: dir; 2: opt;
}
