// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gredis provides convenient client for redis server.
//
// Redis Client.
//
// Redis Commands Official: https://redis.io/commands
//
// Redis Chinese Documentation: http://redisdoc.com/
package gredis

import (
	"github.com/gogf/gf/g/container/garray"
	"github.com/gogf/gf/g/util/gconv"
)

const (
	setCmd          = "SET"
	delCmd          = "DEL"
	getCmd          = "GET"
	keysCmd         = "KEYS"
	pingCmd         = "PING"
	echoCmd         = "ECHO"
	infoCmd         = "INFO"
	hSetCmd         = "HSET"
	hGetCmd         = "HGET"
	hDelCmd         = "HDEL"
	hKeysCmd        = "HKEYS"
	scanCmd         = "SCAN"
	hScanCmd        = "HSCAN"
	appendCmd       = "APPEND"
	getRangeCmd     = "GETRANGE"
	setRangeCmd     = "SETRANGE"
	expireCmd       = "EXPIRE"
	flushDbCmd      = "FLUSHDB"
	flushAllCmd     = "FLUSHALL"
	existsCmd       = "EXISTS"
	hExistsCmd      = "HEXISTS"
	hGetAllCmd      = "HGETALL"
	incrByCmd       = "INCRBY"
	incrCmd         = "INCR"
	incrByFloatCmd  = "INCRBYFLOAT"
	hIncrCmd        = "HINCR"
	hIncrByCmd      = "HINCRBY"
	hIncrByFloatCmd = "HINCRBYFLOAT"
)

// Ping 使用客户端向 Redis 服务器发送一个 PING ，如果服务器运作正常的话，会返回一个 PONG 。
// 通常用于测试与服务器的连接是否仍然生效，或者用于测量延迟值。
func (r *Redis) Ping() (bool, error) {
	v, err := r.DoVar(pingCmd)
	if err != nil {
		return false, err
	}
	return v.String() == "PONG", err
}

// FlushDb 清空当前数据库中的所有 key。
func (r *Redis) FlushDb() error {
	_, err := r.DoVar(flushDbCmd)
	if err != nil {
		return err
	}
	return nil
}

// FlushAll 清空整个 Redis 服务器的数据(删除所有数据库的所有 key )。
func (r *Redis) FlushAll() error {
	_, err := r.DoVar(flushAllCmd)
	if err != nil {
		return err
	}
	return nil
}

// Echo 打印一个特定的信息 message ，测试时使用
func (r *Redis) Echo(message string) (string, error) {
	v, err := r.DoVar(echoCmd)
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

// Info 以一种易于解释（parse）且易于阅读的格式，返回关于 Redis 服务器的各种信息和统计数值
func (r *Redis) Info() (string, error) {
	v, err := r.DoVar(infoCmd)
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

// Scan 迭代当前数据库中的数据库键
func (r *Redis) Scan(startIndex int64, pattern string) (int64, []string, error) {
	v, err := r.DoVar(scanCmd, startIndex, "MATCH", pattern)
	if err != nil {
		return 0, nil, err
	}
	results := v.Interfaces()
	if len(results) != 2 {
		return 0, []string{}, nil
	}
	return gconv.Int64(results[0]), gconv.Strings(results[1]), nil
}

// Append 如果 key 已经存在并且是一个字符串， APPEND 命令将 value 追加到 key 原来的值的末尾。
// 如果 key 不存在， APPEND 就简单地将给定 key 设为 value ，就像执行 SET key value 一样。
func (r *Redis) Append(key string, value string) (int64, error) {
	v, err := r.DoVar(appendCmd, key, value)
	if err != nil {
		return 0, err
	}
	return v.Int64(), nil
}

// GetRange 返回 key 中字符串值的子字符串，字符串的截取范围由 start 和 end 两个偏移量决定(包括 start 和 end 在内)。
// 负数偏移量表示从字符串最后开始计数， -1 表示最后一个字符， -2 表示倒数第二个，以此类推。
// GETRANGE 通过保证子字符串的值域(range)不超过实际字符串的值域来处理超出范围的值域请求。
func (r *Redis) GetRange(key string, start int, end int) (string, error) {
	v, err := r.DoVar(getRangeCmd, key, start, end)
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

// SetRange 用 value 参数覆写(overwrite)给定 key 所储存的字符串值，从偏移量 offset 开始。
// 不存在的 key 当作空白字符串处理。
// SETRANGE 命令会确保字符串足够长以便将 value 设置在指定的偏移量上，如果给定 key 原来储存的字符串长度比偏移量小(比如字符串只有 5 个字符长，但你设置的 offset 是 10 )，那么原字符和偏移量之间的空白将用零字节(zerobytes, "\x00" )来填充。
// 注意你能使用的最大偏移量是 2^29-1(536870911) ，因为 Redis 字符串的大小被限制在 512 兆(megabytes)以内。如果你需要使用比这更大的空间，你可以使用多个 key 。
func (r *Redis) SetRange(key string, start int, value string) (int64, error) {
	v, err := r.DoVar(setRangeCmd, key, start, value)
	if err != nil {
		return 0, err
	}
	return v.Int64(), nil
}

// Expire 为给定 key 设置生存时间，当 key 过期时(生存时间为 0 )，它会被自动删除。
// 在 Redis 中，带有生存时间的 key 被称为『易失的』(volatile)。
func (r *Redis) Expire(key string, timeout int64) (bool, error) {
	v, err := r.DoVar(expireCmd, key, timeout)
	if err != nil {
		return false, err
	}
	return v.Bool(), nil
}

// Set 将字符串值 value 关联到 key 。
// 如果 key 已经持有其他值， SET 就覆写旧值，无视类型。
// 对于某个原本带有生存时间（TTL）的键来说， 当 SET 命令成功在这个键上执行时， 这个键原有的 TTL 将被清除。
func (r *Redis) Set(key string, value string) (bool, error) {
	v, err := r.DoVar(setCmd, key, value)
	if err != nil {
		return false, err
	}
	return v.String() == "OK", nil
}

// SetNx 只在键不存在时，才对键进行设置操作。
func (r *Redis) SetNx(key string, value string) (bool, error) {
	v, err := r.DoVar(setCmd, key, value, "NX")
	if err != nil {
		return false, err
	}
	return v.String() == "OK", nil
}

// SetEx 设置键的过期时间为 second 秒
func (r *Redis) SetEx(key string, value string, timeout int64) (bool, error) {
	v, err := r.DoVar(setCmd, key, value, "EX", timeout)
	if err != nil {
		return false, err
	}
	return v.String() == "OK", nil
}

// SetPx 设置键的过期时间为 millisecond 毫秒
func (r *Redis) SetPx(key string, value string, timeout int64) (bool, error) {
	v, err := r.DoVar(setCmd, key, value, "PX", timeout)
	if err != nil {
		return false, err
	}
	return v.String() == "OK", nil
}

// Get 返回 key 所关联的字符串值。
// 如果 key 不存在那么返回特殊值 nil 。
// 假如 key 储存的值不是字符串类型，返回一个错误，因为 GET 只能用于处理字符串值。
func (r *Redis) Get(key string) (string, error) {
	v, err := r.DoVar(getCmd, key)
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

// Exists 检查给定 key 是否存在
func (r *Redis) Exists(key string) (bool, error) {
	v, err := r.DoVar(existsCmd, key)
	if err != nil {
		return false, err
	}
	return v.Bool(), nil
}

// ExistsBatch 批量检查key是否存在，返回存在的个数
func (r *Redis) ExistsBatch(keys ...string) (int64, error) {
	v, err := r.DoVar(existsCmd, gconv.Interfaces(keys)...)
	if err != nil {
		return 0, err
	}
	return v.Int64(), nil
}

// Del 删除给定的一个或多个 key 。
//不存在的 key 会被忽略。
func (r *Redis) Del(keys ...string) (int64, error) {
	v, err := r.DoVar(delCmd, gconv.Interfaces(keys)...)
	if err != nil {
		return 0, err
	}
	return v.Int64(), nil
}

// Keys retrieves keys that match a pattern
func (r *Redis) Keys(pattern string) ([]string, error) {
	v, err := r.DoVar(keysCmd, pattern)
	if err != nil {
		return nil, err
	}
	return v.Strings(), nil
}

// Incr increments the key's value
func (r *Redis) Incr(key string) (int64, error) {
	v, err := r.DoVar(incrCmd, key)
	if err != nil {
		return -1, err
	}
	return v.Int64(), nil
}

// IncrBy increments the key's value by the increment provided
func (r *Redis) IncrBy(key string, increment int64) (int64, error) {
	v, err := r.DoVar(incrByCmd, key, increment)
	if err != nil {
		return -1, err
	}
	return v.Int64(), nil
}

// IncrByFloat increments the key's value by the increment provided
func (r *Redis) IncrByFloat(key string, increment float64) (float64, error) {
	v, err := r.DoVar(incrByFloatCmd, key, increment)
	if err != nil {
		return -1, err
	}
	return v.Float64(), nil
}

// Decr decrements the key's value
func (r *Redis) Decr(key string) (int64, error) {
	return r.IncrBy(key, -1)
}

// DecrBy decrements the key's value by the decrement provided
func (r *Redis) DecrBy(key string, decrement int64) (int64, error) {
	return r.IncrBy(key, -decrement)
}

// DecrByFloat decrements the key's value by the decrement provided
func (r *Redis) DecrByFloat(key string, decrement float64) (float64, error) {
	return r.IncrByFloat(key, -decrement)
}

// HScan 迭代哈希键中的键值对
func (r *Redis) HScan(key string, startIndex int64, pattern string) (int64, []string, error) {
	v, err := r.DoVar(hScanCmd, key, startIndex, "MATCH", pattern)
	if err != nil {
		return 0, nil, err
	}
	results := v.Interfaces()
	if len(results) != 2 {
		return 0, []string{}, nil
	}
	return gconv.Int64(results[0]), gconv.Strings(results[1]), nil
}

// HSet sets a key's field/value pair
func (r *Redis) HSet(key string, field string, value string) (bool, error) {
	v, err := r.DoVar(hSetCmd, key, field, value)
	if err != nil {
		return false, err
	}
	return v.Bool(), nil
}

// HKeys retrieves a hash's keys
func (r *Redis) HKeys(key string) ([]string, error) {
	v, err := r.DoVar(hKeysCmd, key)
	if err != nil {
		return nil, err
	}
	return v.Strings(), nil
}

// HExists determine's a key's field's existence
func (r *Redis) HExists(key string, field string) (bool, error) {
	v, err := r.DoVar(hExistsCmd, key, field)
	if err != nil {
		return false, err
	}
	return v.Bool(), nil
}

// HGet retrieves a key's field's value
func (r *Redis) HGet(key string, field string) (string, error) {
	v, err := r.DoVar(hGetCmd, key, field)
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

// HGetAll retrieves the key
func (r *Redis) HGetAll(key string) (map[string]interface{}, error) {
	v, err := r.DoVar(hGetAllCmd, key)
	if err != nil {
		return nil, err
	}
	return gconv.Map(v.Val()), nil
}

// HDel deletes a key's fields
func (r *Redis) HDel(key string, fields ...string) (int64, error) {
	params := garray.New()
	params.Append(key)
	params.Append(gconv.Interfaces(fields)...)
	v, err := r.DoVar(hDelCmd, params)
	if err != nil {
		return 0, err
	}
	return v.Int64(), nil
}

// HIncr increments the key's field's value
func (r *Redis) HIncr(key string, field string) (int64, error) {
	v, err := r.DoVar(hIncrCmd, key, field)
	if err != nil {
		return 0, err
	}
	return v.Int64(), nil
}

// HIncrBy increments the key's field's value by the increment provided
func (r *Redis) HIncrBy(key string, field string, increment int64) (int64, error) {
	v, err := r.DoVar(hIncrByCmd, key, field, increment)
	if err != nil {
		return 0, err
	}
	return v.Int64(), nil
}

// HIncrByFloat increments the key's field's value by the increment provided
func (r *Redis) HIncrByFloat(key string, field string, increment float64) (float64, error) {
	v, err := r.DoVar(hIncrByFloatCmd, key, field, increment)
	if err != nil {
		return 0, err
	}
	return v.Float64(), nil
}

// HDecr decrements the key's field's value
func (r *Redis) HDecr(key string, field string) (int64, error) {
	return r.HIncrBy(key, field, -1)
}

// HDecrBy decrements the key's field's value by the decrement provided
func (r *Redis) HDecrBy(key string, field string, decrement int64) (int64, error) {
	return r.HIncrBy(key, field, -decrement)
}

// HDecrByFloat decrements the key's field's value by the decrement provided
func (r *Redis) HDecrByFloat(key string, field string, decrement float64) (float64, error) {
	return r.HIncrByFloat(key, field, -decrement)
}

func (r *Redis) LPush(key string, value ...string) (int64, error) {
	params := garray.New()
	params.Append(key)
	params.Append(gconv.Interfaces(value)...)
	v, err := r.DoVar("LPush", params)
	if err != nil {
		return 0, err
	}
	return v.Int64(), nil
}

func (r *Redis) LPushX(key string, value ...string) (int64, error) {
	params := garray.New()
	params.Append(key)
	params.Append(gconv.Interfaces(value)...)
	v, err := r.DoVar("LPushX", params)
	if err != nil {
		return 0, err
	}
	return v.Int64(), nil
}

func (r *Redis) LPop(key string) (string, error) {
	v, err := r.DoVar("LPop", key)
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

func (r *Redis) RPush(key string, value ...string) (int64, error) {
	params := garray.New()
	params.Append(key)
	params.Append(gconv.Interfaces(value)...)
	v, err := r.DoVar("RPush", params)
	if err != nil {
		return 0, err
	}
	return v.Int64(), nil
}

func (r *Redis) RPushX(key string, value ...string) (int64, error) {
	params := garray.New()
	params.Append(key)
	params.Append(gconv.Interfaces(value)...)
	v, err := r.DoVar("RPushX", params)
	if err != nil {
		return 0, err
	}
	return v.Int64(), nil
}

func (r *Redis) RPop(key string, value ...string) (string, error) {
	v, err := r.DoVar("RPop", key)
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

func (r *Redis) LLen(key string) (int64, error) {
	v, err := r.DoVar("LLen", key)
	if err != nil {
		return 0, err
	}
	return v.Int64(), nil
}

func (r *Redis) LRange(key string, start, end int64) ([]string, error) {
	v, err := r.DoVar("LRange", key, start, end)
	if err != nil {
		return nil, err
	}
	return v.Strings(), nil
}

func (r *Redis) LSet(key string, index int64, value string) (bool, error) {
	v, err := r.DoVar("LPushX", key, index, value)
	if err != nil {
		return false, err
	}
	return v.String() == "OK", nil
}
