package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Genre struct {
	GenreID   int    `bson:"genre_id" json:"genre_id" validate:"required"`
	GenreName string `bson:"genre_name" json:"genre_name" validate:"required,min=2,max=500"`
}

type Ranking struct {
	RankingValue int    `bson:"ranking_value" json:"ranking_value" validate:"required"`
	RankingName  string `bson:"ranking_name" json:"ranking_name" validate:"required"`
}

type Movie struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`                     //omitempty - to create a unique ID in the object ID in the collection
	ImdbID      string        `bson:"imdb_id" json:"imdb_id" validate:"required"`             //add validation that this field is required
	Title       string        `bson:"title" json:"title" validate:"required,min=2,max=500"`   //validation: title needs to be from 2 - 500 characters
	PosterPath  string        `bson:"poster_path" json:"poster_path" validate:"required,url"` //it needs to be a url
	YouTubeID   string        `bson:"youtube_id" json:"youtube_id" validate:"required"`
	Genre       []Genre       `bson:"genre" json:"genre" validate:"required,dive"` //dive means that the genre struct will also be validated
	AdminReview string        `bson:"admin_review" json:"admin_review"`
	Ranking     Ranking       `bson:"ranking" json:"ranking" validate:"required"`
}
