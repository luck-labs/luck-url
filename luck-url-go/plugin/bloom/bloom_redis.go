package bloom

import (
	"context"
	"github.com/demdxx/gocast"
	"github.com/luck-labs/luck-url/plugin/redis"
)

const (
	MSetBitScriptTemplate = `
local sum=0
for i, v in ipairs(ARGV) do
	sum = sum + redis.call('setbit', KEYS[1], v, 1)
end
return sum
`
	MGetBitScriptTemplate = `
for i, v in ipairs(ARGV) do
	if redis.call('getbit', KEYS[1], v) == 0 then
		return 0
	end
end
return 1
`
)

func MGetBit(ctx context.Context, key string, offsets []uint32) (uint, error) {
	args := make([]interface{}, 0)
	for _, offset := range offsets {
		args = append(args, offset)
	}
	result, err := redis.RedisClient.Eval(ctx, MGetBitScriptTemplate, []string{key}, args).Result()
	if err != nil {
		return 0, err
	}
	return gocast.ToUint(result), nil
}

func MSetBit(ctx context.Context, key string, offsets []uint32) (uint, error) {
	args := make([]interface{}, 0)
	for _, offset := range offsets {
		args = append(args, offset)
	}
	result, err := redis.RedisClient.Eval(ctx, MSetBitScriptTemplate, []string{key}, args).Result()
	if err != nil {
		return 0, err
	}
	return gocast.ToUint(result), nil
}
