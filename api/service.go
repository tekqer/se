package api

import (
	"context"

	"github.com/tekqer/se/proto"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	db *mongo.Database
}

func (s *Service) CreateApp(ctx context.Context, app *proto.App) error {
	// TODO: we should first get user id from ctx.
	return nil
}
