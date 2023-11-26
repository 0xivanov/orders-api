package repository

import (
	"context"
	"encoding/json"
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

func (r *RedisRepo) FindById(ctx context.Context, id uint64, offset, limit int) ([]model.Order, error) {
	orderKey := getOrderId(id)

	// Use the SORT command to get a range of items by their score (in this case, the order key).
	// The LIMIT and OFFSET options allow for pagination.
	orderKeys, err := r.Client.Sort(ctx, orderKey,
		&redis.Sort{Offset: int64(offset), Count: int64(limit), Order: "ASC"}).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get orders: %v", err)
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

// GetAllOrders retrieves all orders from Redis
func (r *RedisRepo) GetAllOrders(ctx context.Context) ([]model.Order, error) {
	// Retrieve all order keys from Redis
	keys, err := r.Client.Keys(ctx, "order:"+"*").Result()
	if err != nil {
		return nil, err
	}

	var orders []model.Order

	// Retrieve each order by key
	for _, key := range keys {
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

func (r *RedisRepo) DeleteOrder(ctx context.Context, id uint64) error {
	orderKey := getOrderId(id)

	// Delete order from Redis
	err := r.Client.Del(ctx, orderKey).Err()
	if err != nil {
		return fmt.Errorf("failed to delete order: %v", err)

	}
	return nil
}
