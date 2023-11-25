package repositories

import (
	"context"
)

func (r *Repository) CreateLog(ctx context.Context, action string, input interface{}) error {
	type Log struct {
		Action string      `json:"action"`
		Input  interface{} `json:"input"`
	}

	doc := Log{
		Action: action,
		Input:  input,
	}

	_, err := r.mongoDB.Database("gobeer").Collection("logs").InsertOne(ctx, doc)
	if err != nil {
		return err
	}

	return nil
}
