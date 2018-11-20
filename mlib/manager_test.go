package mlib

import "testing"

func Test_Ops(t *testing.T)  {
	mm := NewMusicManager()

	if mm == nil {
		t.Error("NewMusicManager failed.")
	}
	if mm.Len() != 0 {
		t.Error("NewMusicManager failed, not empty.")
	}
	m0 := &MusicEntry{"1", "My Heart Will Go On", "Celion Dion",  "http://qbox.me/121212121", "MP3"}
	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("NewMusicManager.Add() failed.")
	}

	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("NewMusicManager.Find() failed.")
	}

	m, err := mm.Get(0)
	if m == nil {
		t.Error("NewMusicManager.Get() failed.", err)
	}

	m = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("NewMusicManager.Remove() failed.")
	}
}