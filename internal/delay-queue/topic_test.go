package delayqueue

import (
	"sort"
	"testing"

	conf "github.com/amazingchow/photon-dance-delay-queue/internal/config"
	"github.com/amazingchow/photon-dance-delay-queue/internal/redis"
	"github.com/stretchr/testify/assert"
)

func TestTopicCURD(t *testing.T) {
	fakeRedisCfg := &conf.Redis{
		SentinelEndpoints:   []string{"localhost:26379", "localhost:26380", "localhost:26381"},
		SentinelMasterName:  "mymaster",
		SentinelPassword:    "Pwd123!@",
		RedisMasterPassword: "sOmE_sEcUrE_pAsS",
		RedisPoolMaxIdle:    3,
		RedisPoolMaxActive:  64,
		RedisConnectTimeout: 500,
		RedisReadTimeout:    500,
		RedisWriteTimeout:   500,
	}

	dq := &DelayQueue{
		redisCli: redis.GetOrCreateInstance(fakeRedisCfg),
	}
	// 先清理环境
	_, err := dq.redisCli.ExecCommand("DEL", DefaultTopicSetName)
	assert.Empty(t, err)

	fakeTopis := []string{
		"shopping_cart_service_line",
		"order_service_line",
		"inventory_service_line",
	}
	for _, topic := range fakeTopis {
		err := dq.putTopic(DefaultTopicSetName, topic)
		assert.Empty(t, err)
		has, err := dq.hasTopic(DefaultTopicSetName, topic)
		assert.Empty(t, err)
		assert.Equal(t, true, has)
	}

	allTopics, err := dq.listTopic(DefaultTopicSetName)
	assert.Empty(t, err)
	assert.Equal(t, len(fakeTopis), len(allTopics))
	sort.Strings(fakeTopis)
	sort.Strings(allTopics)
	for idx := range fakeTopis {
		assert.Equal(t, fakeTopis[idx], allTopics[idx])
	}

	redis.ReleaseInstance()
}
