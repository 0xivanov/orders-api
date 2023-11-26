package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/0xivanov/orders-api/model"
	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	Client *redis.Client
}

func getOrderId(id uint64) string {
	return fmt.Sprintf("order:%v", id)
}

func (redis *RedisRepo) Insert(ctx context.Context, order *model.Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("failed to parse to json: %v", err)
	}
	result := redis.Client.SetNX(ctx, getOrderId(order.OrderId), string(data), 0)
	if result.Err() != nil {
		return fmt.Errorf("failed to insert order: %v", err)
	}
	return nil
}

func (r *RedisRepo) List(ctx context.Context, offset uint64, limit int) ([]model.Order, error) {
	// Assuming you have a pattern for your order keys, adjust it accordingly
	orderPattern := "order:*"

	// Use the KEYS command to get all order keys that match the pattern
	orderKeys, err := r.Client.Keys(ctx, orderPattern).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get order keys: %v", err)
	}

	// Use the SORT command to get a range of items by their score (order key).
	// The LIMIT and OFFSET options allow for pagination.
	orderKeys, err = r.Client.Sort(ctx, orderKeys[0], &redis.Sort{
		Offset: int64(offset),
		Count:  int64(limit),
		Order:  "ASC",
	}).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to sort orders: %v", err)
	}

	var orders []model.Order

	// Retrieve each order by key
	for _, key := range orderKeys {
		orderJSON, err := r.Client.Get(ctx, key).Result()
		if err != nil {
			log.Printf("Error retrieving order from Redis: %v", err)
			continue
		}

		// Convert JSON to order struct
		var order model.Order
		err = json.Unmarshal([]byte(orderJSON), &order)
		if err != nil {
			log.Printf("Error unmarshalling order JSON: %v", err)
			continue
		}

		orders = append(orders, order)
	}

	return orders, nil
}

var ErrNotExist = errors.New("order does not exist")

func (r *RedisRepo) FindById(ctx context.Context, id uint64) (model.Order, error) {
	key := getOrderId(id)

	value, err := r.Client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return model.Order{}, ErrNotExist
	} else if err != nil {
		return model.Order{}, fmt.Errorf("get order: %w", err)
	}

	var order model.Order
	err = json.Unmarshal([]byte(value), &order)
	if err != nil {
		return model.Order{}, fmt.Errorf("failed to decode order json: %w", err)
	}

	return order, nil
}

func (r *RedisRepo) Update(ctx context.Context, order model.Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("failed to encode order: %w", err)
	}

	key := getOrderId(order.OrderId)

	err = r.Client.SetXX(ctx, key, string(data), 0).Err()
	if errors.Is(err, redis.Nil) {
		return ErrNotExist
	} else if err != nil {
		return fmt.Errorf("set order: %w", err)
	}

	return nil
}

func (r *RedisRepo) DeleteOrder(ctx context.Context, id uint64) error {
	orderKey := getOrderId(id)

	// Delete order from Redis
	err := r.Client.Del(ctx, orderKey).Err()
	if err != nil {
		return fmt.Errorf("failed to delete order: %v", err)

	}
	return nil
}
