// Copyright 2019 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gredis_test

import (
	"github.com/gogf/gf/g/database/gredis"
	"github.com/gogf/gf/g/test/gtest"
	redis2 "github.com/gogf/gf/third/github.com/gomodule/redigo/redis"
	"strings"
	"testing"
	"time"
)

var (
	config = gredis.Config{
		Host: "127.0.0.1",
		Port: 6379,
		Db:   1,
	}
)

func Test_NewClose(t *testing.T) {
	gtest.Case(t, func() {
		redis := gredis.New(config)
		gtest.AssertNE(redis, nil)
		err := redis.Close()
		gtest.Assert(err, nil)
	})
}

func Test_Do(t *testing.T) {
	gtest.Case(t, func() {
		redis := gredis.New(config)
		defer redis.Close()
		_, err := redis.Do("SET", "k", "v")
		gtest.Assert(err, nil)

		r, err := redis.Do("GET", "k")
		gtest.Assert(err, nil)
		gtest.Assert(r, []byte("v"))

		_, err = redis.Do("DEL", "k")
		gtest.Assert(err, nil)
		r, err = redis.Do("GET", "k")
		gtest.Assert(err, nil)
		gtest.Assert(r, nil)
	})
}

func Test_Send(t *testing.T) {
	gtest.Case(t, func() {
		redis := gredis.New(config)
		defer redis.Close()
		err := redis.Send("SET", "k", "v")
		gtest.Assert(err, nil)

		r, err := redis.Do("GET", "k")
		gtest.Assert(err, nil)
		gtest.Assert(r, []byte("v"))
	})
}

func Test_Stats(t *testing.T) {
	gtest.Case(t, func() {
		redis := gredis.New(config)
		defer redis.Close()
		redis.SetMaxIdle(2)
		redis.SetMaxActive(100)
		redis.SetIdleTimeout(500 * time.Millisecond)
		redis.SetMaxConnLifetime(500 * time.Millisecond)

		array := make([]*gredis.Conn, 0)
		for i := 0; i < 10; i++ {
			array = append(array, redis.Conn())
		}
		stats := redis.Stats()
		gtest.Assert(stats.ActiveCount, 10)
		gtest.Assert(stats.IdleCount, 0)
		for i := 0; i < 10; i++ {
			array[i].Close()
		}
		stats = redis.Stats()
		gtest.Assert(stats.ActiveCount, 2)
		gtest.Assert(stats.IdleCount, 2)
		//time.Sleep(3000*time.Millisecond)
		//stats  = redis.Stats()
		//fmt.Println(stats)
		//gtest.Assert(stats.ActiveCount,  0)
		//gtest.Assert(stats.IdleCount,    0)
	})
}

func Test_Conn(t *testing.T) {
	gtest.Case(t, func() {
		redis := gredis.New(config)
		defer redis.Close()
		conn := redis.Conn()
		defer conn.Close()

		r, err := conn.Do("GET", "k")
		gtest.Assert(err, nil)
		gtest.Assert(r, []byte("v"))

		_, err = conn.Do("DEL", "k")
		gtest.Assert(err, nil)
		r, err = conn.Do("GET", "k")
		gtest.Assert(err, nil)
		gtest.Assert(r, nil)
	})
}

func Test_Instance(t *testing.T) {
	gtest.Case(t, func() {
		group := "my-test"
		gredis.SetConfig(config, group)
		defer gredis.RemoveConfig(group)
		redis := gredis.Instance(group)
		defer redis.Close()

		conn := redis.Conn()
		defer conn.Close()

		_, err := conn.Do("SET", "k", "v")
		gtest.Assert(err, nil)

		r, err := conn.Do("GET", "k")
		gtest.Assert(err, nil)
		gtest.Assert(r, []byte("v"))

		_, err = conn.Do("DEL", "k")
		gtest.Assert(err, nil)
		r, err = conn.Do("GET", "k")
		gtest.Assert(err, nil)
		gtest.Assert(r, nil)
	})
}

func Test_Basic(t *testing.T) {
	gtest.Case(t, func() {
		config1 := gredis.Config{
			Host: "127.0.0.2",
			Port: 6379,
			Db:   1,
		}
		redis := gredis.New(config1)
		_, err := redis.Do("info")
		gtest.Assert(err == nil, false)

		config1 = gredis.Config{
			Host: "127.0.0.1",
			Port: 6379,
			Db:   1,
			Pass: "666666",
		}
		redis = gredis.New(config1)
		_, err = redis.Do("info")
		gtest.Assert(err == nil, false)

		config1 = gredis.Config{
			Host: "127.0.0.1",
			Port: 6379,
			Db:   100,
		}
		redis = gredis.New(config1)
		_, err = redis.Do("info")
		gtest.Assert(err == nil, false)

		redis = gredis.Instance("gf")
		gtest.Assert(redis == nil, true)
		gredis.ClearConfig()

		redis = gredis.New(config)
		defer redis.Close()
		_, err = redis.DoVar("SET", "k", "v")
		gtest.Assert(err, nil)

		v, err := redis.DoVar("GET", "k")
		gtest.Assert(err, nil)
		gtest.Assert(v.String(), "v")

		conn := redis.GetConn()
		_, err = conn.DoVar("SET", "k", "v")
		gtest.Assert(err, nil)

		//v, err = conn.ReceiveVar()
		//gtest.Assert(err, nil)
		//gtest.Assert(v.String(), "v")

		psc := redis2.PubSubConn{Conn: conn}
		psc.Subscribe("gf")
		redis.DoVar("PUBLISH", "gf", "gf test")
		go func() {
			for {
				v, _ := conn.ReceiveVar()
				switch obj := v.Val().(type) {
				case redis2.Message:
					gtest.Assert(string(obj.Data), "gf test")
				case redis2.Subscription:

				}
			}
		}()

		time.Sleep(time.Second)
	})
}

func TestDumpCmds(t *testing.T) {
	str := `PING, SET, GET, QUIT, EXISTS, DEL, UNLINK, TYPE, FLUSHDB, KEYS, RANDOMKEY, RENAME, RENAMENX,
	RENAMEX, DBSIZE, EXPIRE, EXPIREAT, TTL, SELECT, MOVE, FLUSHALL, GETSET, MGET, SETNX, SETEX,
	MSET, MSETNX, DECRBY, DECR, INCRBY, INCR, APPEND, SUBSTR, HSET, HGET, HSETNX, HMSET, HMGET,
	HINCRBY, HEXISTS, HDEL, HLEN, HKEYS, HVALS, HGETALL, RPUSH, LPUSH, LLEN, LRANGE, LTRIM, LINDEX,
	LSET, LREM, LPOP, RPOP, RPOPLPUSH, SADD, SMEMBERS, SREM, SPOP, SMOVE, SCARD, SISMEMBER, SINTER,
	SINTERSTORE, SUNION, SUNIONSTORE, SDIFF, SDIFFSTORE, SRANDMEMBER, ZADD, ZRANGE, ZREM, ZINCRBY,
	ZRANK, ZREVRANK, ZREVRANGE, ZCARD, ZSCORE, MULTI, DISCARD, EXEC, WATCH, UNWATCH, SORT, BLPOP,
	BRPOP, AUTH, SUBSCRIBE, PUBLISH, UNSUBSCRIBE, PSUBSCRIBE, PUNSUBSCRIBE, PUBSUB, ZCOUNT,
	ZRANGEBYSCORE, ZREVRANGEBYSCORE, ZREMRANGEBYRANK, ZREMRANGEBYSCORE, ZUNIONSTORE, ZINTERSTORE,
	ZLEXCOUNT, ZRANGEBYLEX, ZREVRANGEBYLEX, ZREMRANGEBYLEX, SAVE, BGSAVE, BGREWRITEAOF, LASTSAVE,
	SHUTDOWN, INFO, MONITOR, SLAVEOF, CONFIG, STRLEN, SYNC, LPUSHX, PERSIST, RPUSHX, ECHO, LINSERT,
	DEBUG, BRPOPLPUSH, SETBIT, GETBIT, BITPOS, SETRANGE, GETRANGE, EVAL, EVALSHA, SCRIPT, SLOWLOG,
	OBJECT, BITCOUNT, BITOP, SENTINEL, DUMP, RESTORE, PEXPIRE, PEXPIREAT, PTTL, INCRBYFLOAT,
	PSETEX, CLIENT, TIME, MIGRATE, HINCRBYFLOAT, SCAN, HSCAN, SSCAN, ZSCAN, WAIT, CLUSTER, ASKING,
	PFADD, PFCOUNT, PFMERGE, READONLY, GEOADD, GEODIST, GEOHASH, GEOPOS, GEORADIUS, GEORADIUS_RO,
	GEORADIUSBYMEMBER, GEORADIUSBYMEMBER_RO, MODULE, BITFIELD, HSTRLEN, TOUCH, SWAPDB, MEMORY,
	XADD, XLEN, XDEL, XTRIM, XRANGE, XREVRANGE, XREAD, XACK, XGROUP, XREADGROUP, XPENDING, XCLAIM`
	for _, item := range strings.Split(str, ",") {
		println(strings.TrimSpace(item) + "=\"" + strings.TrimSpace(item) + "\"")
	}
	str = `AGGREGATE, ALPHA, ASC, BY, DESC, GET, LIMIT, MESSAGE, NO, NOSORT, PMESSAGE, PSUBSCRIBE,
    PUNSUBSCRIBE, OK, ONE, QUEUED, SET, STORE, SUBSCRIBE, UNSUBSCRIBE, WEIGHTS, WITHSCORES,
    RESETSTAT, REWRITE, RESET, FLUSH, EXISTS, LOAD, KILL, LEN, REFCOUNT, ENCODING, IDLETIME,
    GETNAME, SETNAME, LIST, MATCH, COUNT, PING, PONG, UNLOAD, REPLACE, KEYS, PAUSE, DOCTOR, 
    BLOCK, NOACK, STREAMS, KEY, CREATE, MKSTREAM, SETID, DESTROY, DELCONSUMER, MAXLEN, GROUP, 
    IDLE, TIME, RETRYCOUNT, FORCE`
	println()
	println()
	for _, item := range strings.Split(str, ",") {
		println(strings.TrimSpace(item) + "=\"" + strings.TrimSpace(item) + "\"")
	}
}

func Test_Commands(t *testing.T) {
	gtest.Case(t, func() {
		config1 := gredis.Config{
			Host: "10.1.1.63",
			Port: 6379,
			Db:   1,
			Pass: "123456",
		}
		redis := gredis.New(config1)
		ok, err := redis.Conn().Ping("hello")
		gtest.Assert(err, nil)
		gtest.Assert(ok, "hello")
		ok, err = redis.Conn().Ping()
		gtest.Assert(err, nil)
		gtest.Assert(ok, "PONG")

		ok, err = redis.Conn().Ping("a", "b")
		gtest.Assert(err == nil, false)
		gtest.Assert(ok, "")
	})
}
