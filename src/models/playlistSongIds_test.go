package models

import "testing"

func TestCreatePlaylistSongIds(t *testing.T) {
	ps := &PlaylistSongIds{
		PlaylistId: 1,
	}
	if err := ps.CreatePlaylistSongIds(); err != nil {
		t.Fatal(err.Error())
	} else {
		t.Log("create success")
	}
}

func TestAddSong(t *testing.T) {
	t.Run("add song no exist", func(t *testing.T) {

		ps := &PlaylistSongIds{
			PlaylistId: 1,
		}
		if err := ps.AddSong(1); err != nil {
			t.Fatal(err.Error())
		} else {
			t.Log("add success")
		}
	})
	t.Run("add song exist", func(t *testing.T) {

		ps := &PlaylistSongIds{
			PlaylistId: 1,
		}
		if err := ps.AddSong(1); err != nil {
			t.Log(err.Error())
		} else {
			t.Error("test add song exist fail")
		}
	})
}

func TestGetAllSong(t *testing.T) {
	ps := &PlaylistSongIds{
		PlaylistId: 1,
	}
	if err := ps.GetAllSong(); err != nil {
		t.Fatal(err.Error())
	} else {
		t.Log(ps.SongIds)
	}
}

func TestDelSong(t *testing.T) {
	t.Run("del song exist", func(t *testing.T) {
		ps := &PlaylistSongIds{
			PlaylistId: 1,
		}
		if err := ps.DelSong(1); err != nil {
			t.Fatal(err.Error())
		} else {
			t.Log("del success")
		}
	})
	t.Run("del song no exist", func(t *testing.T) {
		ps := &PlaylistSongIds{
			PlaylistId: 1,
		}
		if err := ps.DelSong(1); err != nil {
			t.Fatal(err.Error())
		} else {
			t.Log("test del song no exist ok")
		}
	})
}

func TestDeletePlaylistSongIds(t *testing.T) {
	t.Run("no exist", func(t *testing.T) {
		ps := &PlaylistSongIds{
			PlaylistId: 1,
		}
		if err := ps.DeletePlaylistSongIds(); err != nil {
			t.Fatal(err.Error())
		} else {
			t.Log("delete success")
		}
	})
	t.Run("exist", func(t *testing.T) {
		ps := &PlaylistSongIds{
			PlaylistId: 1,
		}
		if err := ps.DeletePlaylistSongIds(); err != nil {
			t.Fatal(err.Error())
		} else {
			t.Log("test delete exist ok")
		}
	})
}
