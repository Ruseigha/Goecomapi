package repository

import (
	"context"
	"time"

	"github.com/rseigha/goecomapi/internal/database"
	"github.com/rseigha/goecomapi/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type orderRepo struct {
	coll   *mongo.Collection
	logger *zap.Logger
}

func NewOrderRepository(db *database.MongoDB, logger *zap.Logger) OrderRepository {
	c := db.Collection("orders")
	return &orderRepo{coll: c, logger: logger}
}

func (r *orderRepo) Create(ctx context.Context, o *domain.Order) error {
	now := time.Now().UTC()
	o.CreatedAt = now
	o.UpdatedAt = now
	res, err := r.coll.InsertOne(ctx, o)
	if err != nil {
		return err
	}
	oid := res.InsertedID.(bson.ObjectID)
	o.ID = oid.Hex()
	return nil
}

func (r *orderRepo) GetByID(ctx context.Context, id string) (*domain.Order, error) {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var o domain.Order
	if err := r.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&o); err != nil {
		return nil, err
	}
	o.ID = oid.Hex()
	return &o, nil
}

func (r *orderRepo) GetByUserID(ctx context.Context, userID string) ([]*domain.Order, error) {
	// userID is stored as string (ObjectID hex)
	cur, err := r.coll.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var out []*domain.Order
	for cur.Next(ctx) {
		var o domain.Order
		if err := cur.Decode(&o); err != nil {
			return nil, err
		}
		out = append(out, &o)
	}
	return out, nil
}
