package repository

import (
	"context"
	"time"

	"github.com/rseigha/goecomapi/internal/database"
	"github.com/rseigha/goecomapi/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/zap"
)

type productRepo struct {
	coll   *mongo.Collection
	logger *zap.Logger
}

func NewProductRepository(db *database.MongoDB, logger *zap.Logger) ProductRepository {
	c := db.Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mod := mongo.IndexModel{
		Keys:    bson.D{{Key: "sku", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := c.Indexes().CreateOne(ctx, mod)
	if err != nil {
		logger.Warn("could not create sku index", zap.Error(err))
	}
	return &productRepo{coll: c, logger: logger}
}

func (r *productRepo) Create(ctx context.Context, p *domain.Product) error {
	now := time.Now().UTC()
	p.CreatedAt = now
	p.UpdatedAt = now
	res, err := r.coll.InsertOne(ctx, p)
	if err != nil {
		return err
	}
	oid := res.InsertedID.(bson.ObjectID)
	p.ID = oid.Hex()
	return nil
}

func (r *productRepo) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var p domain.Product
	if err := r.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&p); err != nil {
		return nil, err
	}
	p.ID = oid.Hex()
	return &p, nil
}

func (r *productRepo) Update(ctx context.Context, p *domain.Product) error {
	oid, err := bson.ObjectIDFromHex(p.ID)
	if err != nil {
		return err
	}
	p.UpdatedAt = time.Now().UTC()
	_, err = r.coll.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": p})
	return err
}

func (r *productRepo) Delete(ctx context.Context, id string) error {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.coll.DeleteOne(ctx, bson.M{"_id": oid})
	return err
}

func (r *productRepo) List(ctx context.Context, limit, page int) ([]*domain.Product, int64, error) {
	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}
	skip := int64((page - 1) * limit)
	findOptions := options.Find().SetLimit(int64(limit)).SetSkip(skip).SetSort(bson.D{{Key: "created_at", Value: -1}})
	cur, err := r.coll.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)
	var out []*domain.Product
	for cur.Next(ctx) {
		var p domain.Product
		if err := cur.Decode(&p); err != nil {
			return nil, 0, err
		}
		out = append(out, &p)
	}
	total, err := r.coll.CountDocuments(ctx, bson.M{})
	if err != nil {
		return out, 0, err
	}
	return out, total, nil
}
