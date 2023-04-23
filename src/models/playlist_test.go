package models

import (
	"log"
	"testing"
)

func TestPlaylistInsert(t *testing.T) {
	pl := &Playlist{
		Name:         "test",
		CreateUserId: 1,
	}
	if err := pl.Insert(); err != nil {
		log.Fatal(err.Error())
	} else {
		log.Println("insert success")
	}
}

func TestPlaylistGet(t *testing.T) {
	pl := &Playlist{
		Id: 1,
	}
	if err := pl.Get(); err != nil {
		log.Fatal(err.Error())
	} else {
		log.Printf("%v", pl)
		log.Println("get success")
	}
}

func TestPlaylistUpdate(t *testing.T) {
	pl := &Playlist{
		Id: 1,
	}
	if err := pl.Get(); err != nil {
		log.Fatal(err.Error())
	} else {
		pl.Name = "test2"
		if err := pl.Update(); err != nil {
			log.Fatal(err.Error())
		} else {
			log.Println("update success")
		}
	}
}

func TestPlaylistExits(t *testing.T) {
	pl := &Playlist{
		Id: 1,
	}
	if pl.Exists() {
		log.Println("exists")
	} else {
		log.Fatal("not exists")
	}
}

func TestPlaylistGetPlaylists(t *testing.T) {
	pl, err := GetPlaylists(1)
	if err != nil {
		log.Fatal(err.Error())
	} else {
		log.Printf("%v", pl)
		log.Println("get success")
	}
}

func TestPlaylistDelete(t *testing.T) {
	pl := &Playlist{
		Id: 1,
	}
	if ok := pl.Exists(); ok != true {
		log.Fatal("this playlist not exists")
	} else {
		if err := pl.Delete(); err != nil {
			log.Fatal(err.Error())
		} else {
			log.Println("delete success")
		}
	}

}
