package bloom

import (
	"context"
	"github.com/luck-labs/luck-url/plugin/log"
	"github.com/luck-labs/luck-url/plugin/redis"
	"github.com/sirupsen/logrus"
	"github.com/spaolacci/murmur3"
)

const (
	M = 65 * 1000 * 1000
	K = 6
)

type WarningCallback func(ctx context.Context, count int64)

// A BloomFilter is a representation of a set of _n_ items, where the main
// requirement is to make membership queries; _i.e._, whether an item is a
// member of a set.
type BloomFilter struct {
	m                uint
	k                uint
	filterKey        string
	cntKey           string
	warningThreshold int64
	warningCallback  WarningCallback
	err              error
}

func max(x, y uint) uint {
	if x > y {
		return x
	}
	return y
}

// New creates a new Bloom filter with _m_ bits and _k_ hashing functions
// We force _m_ and _k_ to be at least one to avoid panics.
func New(m uint, k uint, filterKey, cntKey string) *BloomFilter {
	return &BloomFilter{
		max(1, m),
		max(1, k),
		filterKey,
		cntKey,
		-1,
		nil,
		nil,
	}
}

func NewDefaultBloomFilter(filterKey, cntKey string) BloomFilter {
	return BloomFilter{
		m:                M,
		k:                K,
		filterKey:        filterKey,
		cntKey:           cntKey,
		warningThreshold: -1,
		warningCallback:  nil,
	}
}

func NewBloomFilterWithWarning(m uint, k uint, filterKey, cntKey string,
	warningThreshold int64, callback WarningCallback) *BloomFilter {
	return &BloomFilter{
		max(1, m),
		max(1, k),
		filterKey,
		cntKey,
		warningThreshold,
		callback,
		nil,
	}
}

// baseHashes returns the four hash values of data that are used to create k
// hashes
func baseHashes(data []byte) [4]uint64 {
	a1 := []byte{1} // to grab another bit of data
	hasher := murmur3.New128()
	hasher.Write(data) // #nosec
	v1, v2 := hasher.Sum128()
	hasher.Write(a1) // #nosec
	v3, v4 := hasher.Sum128()
	return [4]uint64{
		v1, v2, v3, v4,
	}
}

// location returns the ith hashed location using the four base hash values
func location(h [4]uint64, i uint) uint64 {
	ii := uint64(i)
	return h[ii%2] + ii*h[2+(((ii+(ii%2))%4)/2)]
}

// location returns the ith hashed location using the four base hash values
func (f *BloomFilter) location(h [4]uint64, i uint) uint {
	return uint(location(h, i) % uint64(f.m))
}

// Cap returns the capacity, _m_, of a Bloom filter
func (f *BloomFilter) Cap() uint {
	return f.m
}

// K returns the number of hash functions used in the BloomFilter
func (f *BloomFilter) K() uint {
	return f.k
}

// Add data to the Bloom Filter. Returns the filter (allows chaining)
func (f *BloomFilter) Add(ctx context.Context, data []byte) *BloomFilter {
	h := baseHashes(data)
	locations := make([]uint32, 0)
	for i := uint(0); i < f.k; i++ {
		locations = append(locations, uint32(f.location(h, i)))
	}
	_, err := MSetBit(ctx, f.filterKey, locations)
	if err != nil {
		log.Error(ctx, log.MduRedisBloom, log.IdxRedisBloomAdd, logrus.Fields{
			"tag": "MSetBit",
			"err": err,
		})
		return f
	}

	reply, err := redis.RedisClient.Incr(ctx, f.cntKey).Result()
	if err != nil {
		log.Error(ctx, log.MduRedisBloom, log.IdxRedisBloomAdd, logrus.Fields{
			"tag": "Incr",
			"err": err,
		})
		return f
	}
	if f.warningThreshold > 0 && reply > f.warningThreshold && f.warningCallback != nil {
		f.warningCallback(ctx, reply)
	}
	return f
}

// AddString to the Bloom Filter. Returns the filter (allows chaining)
func (f *BloomFilter) AddString(ctx context.Context, data string) *BloomFilter {
	return f.Add(ctx, []byte(data))
}

// Test returns true if the data is in the BloomFilter, false otherwise.
// If true, the result might be a false positive. If false, the data
// is definitely not in the set.
func (f *BloomFilter) Test(ctx context.Context, data []byte) bool {
	h := baseHashes(data)
	locations := make([]uint32, 0)

	for i := uint(0); i < f.k; i++ {
		locations = append(locations, uint32(f.location(h, i)))
	}
	result, err := MGetBit(ctx, f.filterKey, locations)
	if err != nil {
		log.Error(ctx, log.MduRedisBloom, log.IdxRedisBloomTest, logrus.Fields{
			"tag": "MGetBit",
			"err": err,
		})
		return true
	}
	// 如果结果为0，则认为是存在一个hash不在bitmap中
	return result > 0
}

func (f *BloomFilter) TestString(ctx context.Context, data string) bool {
	return f.Test(ctx, []byte(data))
}

// ClearAll clears all the data in a Bloom filter, removing all keys
func (f *BloomFilter) ClearAll(ctx context.Context) *BloomFilter {
	if err := redis.RedisClient.Del(ctx, f.filterKey).Err(); err != nil {
		log.Error(ctx, log.MduRedisBloom, log.IdxRedisBloomClear, logrus.Fields{
			"tag": "Del",
			"err": err,
		})
	}
	return f
}

// Locations returns a list of hash locations representing a data item.
func Locations(data []byte, k uint) []uint64 {
	locs := make([]uint64, k)
	// calculate locations
	h := baseHashes(data)
	for i := uint(0); i < k; i++ {
		locs[i] = location(h, i)
	}
	return locs
}

func (f *BloomFilter) SetErr(ctx context.Context, err error) {
	f.err = err
}

func (f *BloomFilter) GetErr(ctx context.Context) error {
	return f.err
}
