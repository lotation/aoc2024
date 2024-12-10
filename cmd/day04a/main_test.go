package main

import (
	"testing"
)

var table = []string{
	"ABCD12",
	"EFGH34",
	"ILMN56",
}

func TestCheckRight(t *testing.T) {
	word := "CD1"
	i := 0
	j := 2
	if !checkRight(word, table, i, j) {
		t.Errorf("Expected true, got false")
	}
}

func TestCheckLeft(t *testing.T) {
	word := "GFE"
	i := 1
	j := 2
	if !checkLeft(word, table, i, j) {
		t.Errorf("Expected true, got false")
	}
}

func TestCheckDown(t *testing.T) {
	word := "246"
	i := 0
	j := 5
	if !checkDown(word, table, i, j) {
		t.Errorf("Expected true, got false")
	}
}

func TestCheckUp(t *testing.T) {
	word := "NHD"
	i := 2
	j := 3
	if !checkUp(word, table, i, j) {
		t.Errorf("Expected true, got false")
	}
}

func TestCheckRightDown(t *testing.T) {
	word := "CH5"
	i := 0
	j := 2
	if !checkRightDown(word, table, i, j) {
		t.Errorf("Expected true, got false")
	}
}

// TODO
func TestCheckRightUp(t *testing.T) {

}

func TestCheckLeftDown(t *testing.T) {

}

func TestCheckLeftUp(t *testing.T) {
	word := "MFA"
	i := 2
	j := 2
	if !checkLeftUp(word, table, i, j) {
		t.Errorf("Expected true, got false")
	}
}
