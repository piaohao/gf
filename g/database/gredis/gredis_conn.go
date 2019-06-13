// Copyright 2019 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gredis

import (
	"errors"
	"github.com/gogf/gf/g/container/gvar"
)

// DoVar returns value from Do as gvar.Var.
func (c *Conn) DoVar(command string, args ...interface{}) (*gvar.Var, error) {
	v, err := c.Do(command, args...)
	return gvar.New(v, true), err
}

// ReceiveVar receives a single reply as gvar.Var from the Redis server.
func (c *Conn) ReceiveVar() (*gvar.Var, error) {
	v, err := c.Receive()
	return gvar.New(v, true), err
}

func (c *Conn) Ping(message ...string) (bool, error) {
	if len(message) > 1 {
		return false, errors.New("parameter must less than 2")
	}
	v, err := c.DoVar(PING, message)
	if err != nil {
		return false, err
	}
	if len(message) == 0 {
		return v.String() == "PONG", err
	}
	return v.Strings()[0] == message[0], err
}
func (c *Conn) Set(key, value string) {
	panic("implement me")
}

func (c *Conn) SetWithParams(key, value string, params SetParams) {
	panic("implement me")
}

func (c *Conn) Get(key string) {
	panic("implement me")
}

func (c *Conn) Exists(keys ...string) {
	panic("implement me")
}

func (c *Conn) Del(keys ...string) {
	panic("implement me")
}

func (c *Conn) Unlink(keys ...string) {
	panic("implement me")
}

func (c *Conn) Type(key string) {
	panic("implement me")
}

func (c *Conn) Keys(pattern string) {
	panic("implement me")
}

func (c *Conn) Rename(oldkey, newkey string) {
	panic("implement me")
}

func (c *Conn) Renamenx(oldkey, newkey string) {
	panic("implement me")
}

func (c *Conn) Expire(key string, seconds int) {
	panic("implement me")
}

func (c *Conn) ExpireAt(key string, unixtime int64) {
	panic("implement me")
}

func (c *Conn) Ttl(key string) {
	panic("implement me")
}

func (c *Conn) Pttl(key string) {
	panic("implement me")
}

func (c *Conn) Touch(keys ...string) {
	panic("implement me")
}

func (c *Conn) Setrange(key string, offset int64, value string) {
	panic("implement me")
}

func (c *Conn) Getrange(key string, startOffset, endOffset int64) {
	panic("implement me")
}

func (c *Conn) GetSet(key string, value string) {
	panic("implement me")
}

func (c *Conn) Mget(keys ...string) {
	panic("implement me")
}

func (c *Conn) Setnx(key string, value string) {
	panic("implement me")
}

func (c *Conn) Setex(key string, seconds int, value string) {
	panic("implement me")
}

func (c *Conn) Mset(keysvalues ...string) {
	panic("implement me")
}

func (c *Conn) Msetnx(keysvalues ...string) {
	panic("implement me")
}

func (c *Conn) DecrBy(key string, decrement int64) {
	panic("implement me")
}

func (c *Conn) Decr(key string) {
	panic("implement me")
}

func (c *Conn) IncrBy(key string, increment int64) {
	panic("implement me")
}

func (c *Conn) IncrByFloat(key string, increment float64) {
	panic("implement me")
}

func (c *Conn) Incr(key string) {
	panic("implement me")
}

func (c *Conn) Append(key string, value string) {
	panic("implement me")
}

func (c *Conn) Substr(key string, start, end int64) {
	panic("implement me")
}

func (c *Conn) Hset(key string, field string, value string) {
	panic("implement me")
}

func (c *Conn) Hget(key string, field string) {
	panic("implement me")
}

func (c *Conn) HsetBatch(key string, hash map[string]string) {
	panic("implement me")
}

func (c *Conn) Hsetnx(key string, field string, value string) {
	panic("implement me")
}

func (c *Conn) Hmset(key string, hash map[string]string) {
	panic("implement me")
}

func (c *Conn) Hmget(key string, fields ...string) {
	panic("implement me")
}

func (c *Conn) HincrBy(key string, field string, value int64) {
	panic("implement me")
}

func (c *Conn) HincrByFloat(key string, field string, value float64) {
	panic("implement me")
}

func (c *Conn) Hexists(key string, field string) {
	panic("implement me")
}

func (c *Conn) Hdel(key string, fields ...string) {
	panic("implement me")
}

func (c *Conn) Hlen(key string) {
	panic("implement me")
}

func (c *Conn) Hkeys(key string) {
	panic("implement me")
}

func (c *Conn) Hvals(key string) {
	panic("implement me")
}

func (c *Conn) HgetAll(key string) {
	panic("implement me")
}

func (c *Conn) Rpush(key string, strings ...string) {
	panic("implement me")
}

func (c *Conn) Lpush(key string, strings ...string) {
	panic("implement me")
}

func (c *Conn) Llen(key string) {
	panic("implement me")
}

func (c *Conn) Lrange(key string, start int64, stop int64) {
	panic("implement me")
}

func (c *Conn) Ltrim(key string, start int64, stop int64) {
	panic("implement me")
}

func (c *Conn) Lindex(key string, index int64) {
	panic("implement me")
}

func (c *Conn) Lset(key string, index int64, value string) {
	panic("implement me")
}

func (c *Conn) Lrem(key string, count int64, value string) {
	panic("implement me")
}

func (c *Conn) Lpop(key string) {
	panic("implement me")
}

func (c *Conn) Rpop(key string) {
	panic("implement me")
}

func (c *Conn) Rpoplpush(srckey string, dstkey string) {
	panic("implement me")
}

func (c *Conn) Sadd(key string, members ...string) {
	panic("implement me")
}

func (c *Conn) Smembers(key string) {
	panic("implement me")
}

func (c *Conn) Srem(key string, member ...string) {
	panic("implement me")
}

func (c *Conn) Spop(key string) {
	panic("implement me")
}

func (c *Conn) SpopBatch(key string, count int64) {
	panic("implement me")
}

func (c *Conn) Smove(srckey string, dstkey string, member string) {
	panic("implement me")
}

func (c *Conn) Scard(key string) {
	panic("implement me")
}

func (c *Conn) Sismember(key string, member string) {
	panic("implement me")
}

func (c *Conn) Sinter(keys ...string) {
	panic("implement me")
}

func (c *Conn) Sinterstore(dstkey string, keys ...string) {
	panic("implement me")
}

func (c *Conn) Sunion(keys ...string) {
	panic("implement me")
}

func (c *Conn) Sunionstore(dstkey string, keys ...string) {
	panic("implement me")
}

func (c *Conn) Sdiff(keys ...string) {
	panic("implement me")
}

func (c *Conn) Sdiffstore(dstkey string, keys ...string) {
	panic("implement me")
}

func (c *Conn) Srandmember(key string) {
	panic("implement me")
}

func (c *Conn) Watch(keys ...string) {
	panic("implement me")
}

func (c *Conn) Strlen(key string) {
	panic("implement me")
}

func (c *Conn) Lpushx(key string, string ...string) {
	panic("implement me")
}

func (c *Conn) Persist(key string) {
	panic("implement me")
}

func (c *Conn) Rpushx(key string, string ...string) {
	panic("implement me")
}

func (c *Conn) Echo(String string) {
	panic("implement me")
}

func (c *Conn) ConfigSet(parameter string, value string) {
	panic("implement me")
}

func (c *Conn) ConfigGet(pattern string) {
	panic("implement me")
}

func (c *Conn) ConfigResetStat() {
	panic("implement me")
}

func (c *Conn) Multi() {
	panic("implement me")
}

func (c *Conn) Exec() {
	panic("implement me")
}

func (c *Conn) Discard() {
	panic("implement me")
}

func (c *Conn) Scan(cursor int64, params ScanParams) {
	panic("implement me")
}

func (c *Conn) Hscan(key string, cursor int64, params ScanParams) {
	panic("implement me")
}

func (c *Conn) Sscan(key string, cursor int64, params ScanParams) {
	panic("implement me")
}

func (c *Conn) Zscan(key string, cursor int64, params ScanParams) {
	panic("implement me")
}
