package customized

import "time"

type HeroCustomized struct {
	TextureURL string    `bson:"textureURL" json:"textureURL"`
	PublishAt  time.Time `bson:"publishAt" json:"publishAt"`
}
