package yelp

import "github.com/guregu/null"

//
type Float struct {
	null.Float
}

func NewFloat(f float64, valid bool) Float { return Float{null.NewFloat(f, valid)} }
func FloatFrom(f float64) Float            { return Float{null.FloatFrom(f)} }
func FloatFromPtr(f *float64) Float        { return Float{null.FloatFromPtr(f)} }

//
type Int struct {
	null.Int
}

func NewInt(f int64, valid bool) Int { return Int{null.NewInt(f, valid)} }
func IntFrom(f int64) Int            { return Int{null.IntFrom(f)} }
func IntFromPtr(f *int64) Int        { return Int{null.IntFromPtr(f)} }

//
type Bool struct {
	null.Bool
}

func NewBool(f bool, valid bool) Bool { return Bool{null.NewBool(f, valid)} }
func BoolFrom(f bool) Bool            { return Bool{null.BoolFrom(f)} }
func BoolFromPtr(f *bool) Bool        { return Bool{null.BoolFromPtr(f)} }
