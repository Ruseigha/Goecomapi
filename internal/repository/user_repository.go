package repository

import (
	"context"
	"errors"
	"time"

	"github.com/rseigha/goecomapi/internal/database"
	"github.com/rseigha/goecomapi/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/zap"
)

type userRepo struct {
	coll   *mongo.Collection
	logger *zap.Logger
}

func NewUserRepository(db *database.MongoDB, logger *zap.Logger) UserRepository {
	c := db.Collection("users")

	// ensure email unique index
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mod := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := c.Indexes().CreateOne(ctx, mod)
	if err != nil {
		logger.Warn("Could not create email index", zap.Error(err))
	}

	return &userRepo{
		coll:   c,
		logger: logger,
	}
}

func (r *userRepo) Create(ctx context.Context, u *domain.User) error {
	now := time.Now().UTC()
	u.CreatedAt = now
	u.UpdatedAt = now

	res, err := r.coll.InsertOne(ctx, u)
	if err != nil {
		return err
	}

	oid, ok := res.InsertedID.(bson.ObjectID)
	if !ok {
		return errors.New("failed to parse inserted ID")
	}

	u.ID = oid.Hex()

	return nil
}

func (r *userRepo) GetByID(ctx context.Context, id string) (*domain.User, error) {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var u domain.User
	if err := r.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&u); err != nil {
		return nil, err
	}

	u.ID = oid.Hex()
	return &u, err
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var u domain.User
	if err := r.coll.FindOne(ctx, bson.M{"email": email}).Decode(&u); err != nil {
		return nil, err
	}
	// ensure ID is string hex if present
	return &u, nil
}

func (r *userRepo) Update(ctx context.Context, u *domain.User) error {
	oid, err := bson.ObjectIDFromHex(u.ID)
	if err != nil {
		return err
	}
	u.UpdatedAt = time.Now().UTC()
	_, err = r.coll.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": u})
	return err
}

func (r *userRepo) Delete(ctx context.Context, id string) error {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.coll.DeleteOne(ctx, bson.M{"_id": oid})
	return err
}
