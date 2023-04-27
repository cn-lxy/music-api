package models

import (
	"context"
	"fmt"
	"log"

	"github.com/cn-lxy/music-api/tools/config"
	"github.com/cn-lxy/music-api/tools/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PlaylistSongIds struct {
	PlaylistId uint64 `json:"id" bson:"id"`
	SongIds    bson.A `json:"songs" bson:"songs"`
}

// CreatePlaylistSongIds create playlist songs ids array in mongodb
func (p *PlaylistSongIds) CreatePlaylistSongIds() error {
	collection := db.Client.Database(config.Cfg.Mongo.Db).Collection(config.Cfg.Mongo.Collection)
	doc := bson.D{
		{
			Key:   "id",
			Value: p.PlaylistId,
		},
		{
			Key:   "songs",
			Value: bson.A{},
		},
	}
	res, err := collection.InsertOne(context.Background(), doc)
	if err != nil {
		return err
	}
	log.Println(res.InsertedID)
	return nil
}

// DeletePlaylistSongIdx delete document of id is p.PlaylistId in mongodb
func (p *PlaylistSongIds) DeletePlaylistSongIds() error {
	collection := db.Client.Database(config.Cfg.Mongo.Db).Collection(config.Cfg.Mongo.Collection)
	// filter is query condition: id = p.PlaylistId, the type is bson.M
	filter := bson.M{"id": p.PlaylistId}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (p *PlaylistSongIds) AddSong(id uint64) error {
	collection := db.Client.Database(config.Cfg.Mongo.Db).Collection(config.Cfg.Mongo.Collection)

	findPS := PlaylistSongIds{}
	// step one: check playlist exists in documents
	err := collection.FindOne(
		context.Background(),
		bson.D{
			{
				Key:   "id",
				Value: p.PlaylistId,
			},
		},
	).Decode(&findPS)
	log.Printf("find playlist song ids: %+v\n", findPS)
	if err != nil {
		return fmt.Errorf("playlist is not existed")
	}

	// step two: check input id whether exist in playlist's songs
	if len(findPS.SongIds) > 0 {
		for _, v := range findPS.SongIds {
			if uint64(v.(int64)) == id {
				return fmt.Errorf("song id %d already exists in playlist %d", id, p.PlaylistId)
			}
		}
	}

	filter := bson.D{{
		Key:   "id",
		Value: p.PlaylistId,
	}}
	update := bson.D{{
		Key:   "$push",
		Value: bson.D{{Key: "songs", Value: id}},
	}}
	updateRes, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	log.Println(updateRes.ModifiedCount)
	return nil
}

func (p *PlaylistSongIds) DelSong(id uint64) error {
	collection := db.Client.Database(config.Cfg.Mongo.Db).Collection(config.Cfg.Mongo.Collection)

	filter := bson.M{"id": p.PlaylistId}
	update := bson.M{"$pull": bson.M{"songs": id}}
	options := options.Update().SetUpsert(false)
	res, err := collection.UpdateMany(context.Background(), filter, update, options)
	if err != nil {
		return err
	}
	log.Printf("matched count: %d, modified count: %d", res.MatchedCount, res.ModifiedCount)
	return nil
}

// GetAllsong get all songs in playlist from mongodb
func (p *PlaylistSongIds) GetAllSong() error {
	collection := db.Client.Database(config.Cfg.Mongo.Db).Collection(config.Cfg.Mongo.Collection)
	// filter is a bson.M type
	filter := bson.M{"id": p.PlaylistId}
	err := collection.FindOne(context.Background(), filter).Decode(p)
	if err != nil {
		return err
	}
	return nil
}
