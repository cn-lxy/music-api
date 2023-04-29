package models

import (
	"fmt"
	"strconv"

	tools "github.com/cn-lxy/music-api/tools/db"
)

type Playlist struct {
	Id           uint64 `json:"id"`
	Name         string `json:"name"`
	CreateUserId uint64 `json:"createUserId"`
	CreateTime   string `json:"createTime"`
	UpdateTime   string `json:"updateTime"`
	PlayCount    uint64 `json:"playCount"`
}

// Insert playlist into the database
// only insert the name, create_user_id
func (p *Playlist) Insert() error {
	if p.CreateUserId == 0 || p.Name == "" {
		return fmt.Errorf("insufficient required parameters")
	}
	res, err := tools.Query("SELECT id FROM playlist WHERE name =? AND create_user_id =?", p.Name, p.CreateUserId)
	if err != nil {
		return err
	}
	if len(res) != 0 {
		return fmt.Errorf("this playlist is existed")
	}
	err = tools.Update("INSERT INTO playlist (name, create_user_id) VALUES (?, ?)", p.Name, p.CreateUserId)
	if err != nil {
		return err
	}
	resId, err := tools.Query("SELECT id FROM playlist WHERE name =? AND create_user_id =?", p.Name, p.CreateUserId)
	if err != nil {
		return err
	}
	p.Id, _ = strconv.ParseUint(resId[0]["id"].(string), 10, 64)
	return err
}

// Update playlist in the database
func (p *Playlist) Update() error {
	if p.Id == 0 && p.CreateUserId == 0 {
		return fmt.Errorf("insufficient required parameters")
	}
	err := tools.Update("UPDATE playlist SET name =?, play_count =? WHERE id =?", p.Name, p.PlayCount, p.Id)
	return err
}

// Delete playlist from the database
func (p *Playlist) Delete() error {
	if p.CreateUserId == 0 || p.Id == 0 {
		return fmt.Errorf("insufficient required parameters")
	}
	// make sure the playlist exists
	if !p.exists() {
		return fmt.Errorf("playlist with id %v does not exist", p.Id)
	}
	err := tools.Update("DELETE FROM playlist WHERE id =? and create_user_id =?", p.Id, p.CreateUserId)
	return err
}

// Check if a playlist exists in the database
func (p *Playlist) exists() bool {
	res, err := tools.Query("SELECT id FROM playlist WHERE id = ? AND create_user_id = ?", p.Id, p.CreateUserId)
	if err != nil {
		return false
	}
	return len(res) > 0
}

// Get a playlist from the database
func (p *Playlist) Get() error {
	if p.CreateUserId == 0 || p.Id == 0 {
		return fmt.Errorf("insufficient required parameters")
	}
	res, err := tools.Query("SELECT id, name, create_user_id, create_time, update_time, play_count FROM playlist WHERE id = ? AND create_user_id = ?", p.Id, p.CreateUserId)
	if err != nil {
		return err
	}
	if len(res) == 0 {
		return fmt.Errorf("playlist with id %v does not exist", p.Id)
	}
	p.Id, _ = strconv.ParseUint(res[0]["id"].(string), 10, 64)
	p.Name = res[0]["name"].(string)
	p.CreateUserId, _ = strconv.ParseUint(res[0]["create_user_id"].(string), 10, 64)
	p.CreateTime = res[0]["create_time"].(string)
	p.UpdateTime = res[0]["update_time"].(string)
	p.PlayCount, _ = strconv.ParseUint(res[0]["play_count"].(string), 10, 64)
	return nil
}

// GetPlaylists get user's all playlist
func GetPlaylists(id uint64) ([]Playlist, error) {
	res, err := tools.Query("SELECT id, name, create_user_id, create_time, update_time, play_count FROM playlist WHERE create_user_id = ?", id)
	if err != nil {
		return nil, err
	}
	var playlists []Playlist
	for _, v := range res {
		var playlist Playlist
		playlist.Id, _ = strconv.ParseUint(v["id"].(string), 10, 64)
		playlist.Name = v["name"].(string)
		playlist.CreateUserId, _ = strconv.ParseUint(v["create_user_id"].(string), 10, 64)
		playlist.CreateTime = v["create_time"].(string)
		playlist.UpdateTime = v["update_time"].(string)
		playlist.PlayCount, _ = strconv.ParseUint(v["play_count"].(string), 10, 64)
		playlists = append(playlists, playlist)
	}
	return playlists, nil
}
