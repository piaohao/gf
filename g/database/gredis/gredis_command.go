package gredis

const (
	XX = "xx"
	NX = "nx"
	PX = "px"
	EX = "ex"

	MATCH = "match"
	COUNT = "count"
)

// all redis commands
const (
	PING                 = "PING"
	SET                  = "SET"
	GET                  = "GET"
	QUIT                 = "QUIT"
	EXISTS               = "EXISTS"
	DEL                  = "DEL"
	UNLINK               = "UNLINK"
	TYPE                 = "TYPE"
	FLUSHDB              = "FLUSHDB"
	KEYS                 = "KEYS"
	RANDOMKEY            = "RANDOMKEY"
	RENAME               = "RENAME"
	RENAMENX             = "RENAMENX"
	RENAMEX              = "RENAMEX"
	DBSIZE               = "DBSIZE"
	EXPIRE               = "EXPIRE"
	EXPIREAT             = "EXPIREAT"
	TTL                  = "TTL"
	SELECT               = "SELECT"
	MOVE                 = "MOVE"
	FLUSHALL             = "FLUSHALL"
	GETSET               = "GETSET"
	MGET                 = "MGET"
	SETNX                = "SETNX"
	SETEX                = "SETEX"
	MSET                 = "MSET"
	MSETNX               = "MSETNX"
	DECRBY               = "DECRBY"
	DECR                 = "DECR"
	INCRBY               = "INCRBY"
	INCR                 = "INCR"
	APPEND               = "APPEND"
	SUBSTR               = "SUBSTR"
	HSET                 = "HSET"
	HGET                 = "HGET"
	HSETNX               = "HSETNX"
	HMSET                = "HMSET"
	HMGET                = "HMGET"
	HINCRBY              = "HINCRBY"
	HEXISTS              = "HEXISTS"
	HDEL                 = "HDEL"
	HLEN                 = "HLEN"
	HKEYS                = "HKEYS"
	HVALS                = "HVALS"
	HGETALL              = "HGETALL"
	RPUSH                = "RPUSH"
	LPUSH                = "LPUSH"
	LLEN                 = "LLEN"
	LRANGE               = "LRANGE"
	LTRIM                = "LTRIM"
	LINDEX               = "LINDEX"
	LSET                 = "LSET"
	LREM                 = "LREM"
	LPOP                 = "LPOP"
	RPOP                 = "RPOP"
	RPOPLPUSH            = "RPOPLPUSH"
	SADD                 = "SADD"
	SMEMBERS             = "SMEMBERS"
	SREM                 = "SREM"
	SPOP                 = "SPOP"
	SMOVE                = "SMOVE"
	SCARD                = "SCARD"
	SISMEMBER            = "SISMEMBER"
	SINTER               = "SINTER"
	SINTERSTORE          = "SINTERSTORE"
	SUNION               = "SUNION"
	SUNIONSTORE          = "SUNIONSTORE"
	SDIFF                = "SDIFF"
	SDIFFSTORE           = "SDIFFSTORE"
	SRANDMEMBER          = "SRANDMEMBER"
	ZADD                 = "ZADD"
	ZRANGE               = "ZRANGE"
	ZREM                 = "ZREM"
	ZINCRBY              = "ZINCRBY"
	ZRANK                = "ZRANK"
	ZREVRANK             = "ZREVRANK"
	ZREVRANGE            = "ZREVRANGE"
	ZCARD                = "ZCARD"
	ZSCORE               = "ZSCORE"
	MULTI                = "MULTI"
	DISCARD              = "DISCARD"
	EXEC                 = "EXEC"
	WATCH                = "WATCH"
	UNWATCH              = "UNWATCH"
	SORT                 = "SORT"
	BLPOP                = "BLPOP"
	BRPOP                = "BRPOP"
	AUTH                 = "AUTH"
	SUBSCRIBE            = "SUBSCRIBE"
	PUBLISH              = "PUBLISH"
	UNSUBSCRIBE          = "UNSUBSCRIBE"
	PSUBSCRIBE           = "PSUBSCRIBE"
	PUNSUBSCRIBE         = "PUNSUBSCRIBE"
	PUBSUB               = "PUBSUB"
	ZCOUNT               = "ZCOUNT"
	ZRANGEBYSCORE        = "ZRANGEBYSCORE"
	ZREVRANGEBYSCORE     = "ZREVRANGEBYSCORE"
	ZREMRANGEBYRANK      = "ZREMRANGEBYRANK"
	ZREMRANGEBYSCORE     = "ZREMRANGEBYSCORE"
	ZUNIONSTORE          = "ZUNIONSTORE"
	ZINTERSTORE          = "ZINTERSTORE"
	ZLEXCOUNT            = "ZLEXCOUNT"
	ZRANGEBYLEX          = "ZRANGEBYLEX"
	ZREVRANGEBYLEX       = "ZREVRANGEBYLEX"
	ZREMRANGEBYLEX       = "ZREMRANGEBYLEX"
	SAVE                 = "SAVE"
	BGSAVE               = "BGSAVE"
	BGREWRITEAOF         = "BGREWRITEAOF"
	LASTSAVE             = "LASTSAVE"
	SHUTDOWN             = "SHUTDOWN"
	INFO                 = "INFO"
	MONITOR              = "MONITOR"
	SLAVEOF              = "SLAVEOF"
	CONFIG               = "CONFIG"
	STRLEN               = "STRLEN"
	SYNC                 = "SYNC"
	LPUSHX               = "LPUSHX"
	PERSIST              = "PERSIST"
	RPUSHX               = "RPUSHX"
	ECHO                 = "ECHO"
	LINSERT              = "LINSERT"
	DEBUG                = "DEBUG"
	BRPOPLPUSH           = "BRPOPLPUSH"
	SETBIT               = "SETBIT"
	GETBIT               = "GETBIT"
	BITPOS               = "BITPOS"
	SETRANGE             = "SETRANGE"
	GETRANGE             = "GETRANGE"
	EVAL                 = "EVAL"
	EVALSHA              = "EVALSHA"
	SCRIPT               = "SCRIPT"
	SLOWLOG              = "SLOWLOG"
	OBJECT               = "OBJECT"
	BITCOUNT             = "BITCOUNT"
	BITOP                = "BITOP"
	SENTINEL             = "SENTINEL"
	DUMP                 = "DUMP"
	RESTORE              = "RESTORE"
	PEXPIRE              = "PEXPIRE"
	PEXPIREAT            = "PEXPIREAT"
	PTTL                 = "PTTL"
	INCRBYFLOAT          = "INCRBYFLOAT"
	PSETEX               = "PSETEX"
	CLIENT               = "CLIENT"
	TIME                 = "TIME"
	MIGRATE              = "MIGRATE"
	HINCRBYFLOAT         = "HINCRBYFLOAT"
	SCAN                 = "SCAN"
	HSCAN                = "HSCAN"
	SSCAN                = "SSCAN"
	ZSCAN                = "ZSCAN"
	WAIT                 = "WAIT"
	CLUSTER              = "CLUSTER"
	ASKING               = "ASKING"
	PFADD                = "PFADD"
	PFCOUNT              = "PFCOUNT"
	PFMERGE              = "PFMERGE"
	READONLY             = "READONLY"
	GEOADD               = "GEOADD"
	GEODIST              = "GEODIST"
	GEOHASH              = "GEOHASH"
	GEOPOS               = "GEOPOS"
	GEORADIUS            = "GEORADIUS"
	GEORADIUS_RO         = "GEORADIUS_RO"
	GEORADIUSBYMEMBER    = "GEORADIUSBYMEMBER"
	GEORADIUSBYMEMBER_RO = "GEORADIUSBYMEMBER_RO"
	MODULE               = "MODULE"
	BITFIELD             = "BITFIELD"
	HSTRLEN              = "HSTRLEN"
	TOUCH                = "TOUCH"
	SWAPDB               = "SWAPDB"
	MEMORY               = "MEMORY"
	XADD                 = "XADD"
	XLEN                 = "XLEN"
	XDEL                 = "XDEL"
	XTRIM                = "XTRIM"
	XRANGE               = "XRANGE"
	XREVRANGE            = "XREVRANGE"
	XREAD                = "XREAD"
	XACK                 = "XACK"
	XGROUP               = "XGROUP"
	XREADGROUP           = "XREADGROUP"
	XPENDING             = "XPENDING"
	XCLAIM               = "XCLAIM"
)

// redis key words
const (
	AGGREGATE   = "AGGREGATE"
	ALPHA       = "ALPHA"
	ASC         = "ASC"
	BY          = "BY"
	DESC        = "DESC"
	LIMIT       = "LIMIT"
	MESSAGE     = "MESSAGE"
	NO          = "NO"
	NOSORT      = "NOSORT"
	PMESSAGE    = "PMESSAGE"
	OK          = "OK"
	ONE         = "ONE"
	QUEUED      = "QUEUED"
	STORE       = "STORE"
	WEIGHTS     = "WEIGHTS"
	WITHSCORES  = "WITHSCORES"
	RESETSTAT   = "RESETSTAT"
	REWRITE     = "REWRITE"
	RESET       = "RESET"
	FLUSH       = "FLUSH"
	LOAD        = "LOAD"
	KILL        = "KILL"
	LEN         = "LEN"
	REFCOUNT    = "REFCOUNT"
	ENCODING    = "ENCODING"
	IDLETIME    = "IDLETIME"
	GETNAME     = "GETNAME"
	SETNAME     = "SETNAME"
	LIST        = "LIST"
	PONG        = "PONG"
	UNLOAD      = "UNLOAD"
	REPLACE     = "REPLACE"
	PAUSE       = "PAUSE"
	DOCTOR      = "DOCTOR"
	BLOCK       = "BLOCK"
	NOACK       = "NOACK"
	STREAMS     = "STREAMS"
	KEY         = "KEY"
	CREATE      = "CREATE"
	MKSTREAM    = "MKSTREAM"
	SETID       = "SETID"
	DESTROY     = "DESTROY"
	DELCONSUMER = "DELCONSUMER"
	MAXLEN      = "MAXLEN"
	GROUP       = "GROUP"
	IDLE        = "IDLE"
	RETRYCOUNT  = "RETRYCOUNT"
	FORCE       = "FORCE"
)

type SetParams struct {
	params map[string]interface{}
}

func NewSetParams() *SetParams {
	return &SetParams{
		params: make(map[string]interface{}),
	}
}

/**
 * Set the specified expire time, in seconds.
 * @param secondsToExpire
 * @return SetParams
 */
func (p SetParams) ex(secondsToExpire int) {
	p.params[EX] = secondsToExpire
}

/**
 * Set the specified expire time, in milliseconds.
 * @param millisecondsToExpire
 * @return SetParams
 */
func (p SetParams) px(millisecondsToExpire int64) {
	p.params[PX] = millisecondsToExpire
}

/**
 * Only set the key if it does not already exist.
 * @return SetParams
 */
func (p SetParams) nx() {
	p.params[NX] = true
}

/**
 * Only set the key if it already exist.
 * @return SetParams
 */
func (p SetParams) xx() {
	p.params[XX] = true
}

type ScanParams struct {
	params map[string]interface{}
}

func (p ScanParams) match(pattern string) {
	p.params[MATCH] = pattern
}

func (p ScanParams) count(count int) {
	p.params[COUNT] = count
}

func (p ScanParams) GetParams() []interface{} {
	arr := make([]interface{}, 0)
	for k, v := range p.params {
		arr = append(arr, k)
		arr = append(arr, v)
	}
	return arr
}

type Commands interface {
	Ping(message string) (bool, error)
	Set(key, value string)
	SetWithParams(key, value string, params SetParams)
	Get(key string)
	Exists(keys ...string)
	Del(keys ...string)
	Unlink(keys ...string)
	Type(key string)
	Keys(pattern string)
	Rename(oldkey, newkey string)
	Renamenx(oldkey, newkey string)
	Expire(key string, seconds int)
	ExpireAt(key string, unixtime int64)
	Ttl(key string)
	Pttl(key string)
	Touch(keys ...string)
	//Setbit(key string, offset int64, value bool)
	//Setbit(key string, offset int64, value string)
	//Getbit(key string, offset int64)
	Setrange(key string, offset int64, value string)
	Getrange(key string, startOffset, endOffset int64)
	//Move(key string, dbIndex int)
	GetSet(key string, value string)
	Mget(keys ...string)
	Setnx(key string, value string)
	Setex(key string, seconds int, value string)
	Mset(keysvalues ...string)
	Msetnx(keysvalues ...string)
	DecrBy(key string, decrement int64)
	Decr(key string)
	IncrBy(key string, increment int64)
	IncrByFloat(key string, increment float64)
	Incr(key string)
	Append(key string, value string)
	Substr(key string, start, end int64)
	Hset(key string, field string, value string)
	Hget(key string, field string)
	HsetBatch(key string, hash map[string]string)
	Hsetnx(key string, field string, value string)
	Hmset(key string, hash map[string]string)
	Hmget(key string, fields ...string)
	HincrBy(key string, field string, value int64)
	HincrByFloat(key string, field string, value float64)
	Hexists(key string, field string)
	Hdel(key string, fields ...string)
	Hlen(key string)
	Hkeys(key string)
	Hvals(key string)
	HgetAll(key string)
	Rpush(key string, strings ...string)
	Lpush(key string, strings ...string)
	Llen(key string)
	Lrange(key string, start int64, stop int64)
	Ltrim(key string, start int64, stop int64)
	Lindex(key string, index int64)
	Lset(key string, index int64, value string)
	Lrem(key string, count int64, value string)
	Lpop(key string)
	Rpop(key string)
	Rpoplpush(srckey string, dstkey string)
	Sadd(key string, members ...string)
	Smembers(key string)
	Srem(key string, member ...string)
	Spop(key string)
	SpopBatch(key string, count int64)
	Smove(srckey string, dstkey string, member string)
	Scard(key string)
	Sismember(key string, member string)
	Sinter(keys ...string)
	Sinterstore(dstkey string, keys ...string)
	Sunion(keys ...string)
	Sunionstore(dstkey string, keys ...string)
	Sdiff(keys ...string)
	Sdiffstore(dstkey string, keys ...string)
	Srandmember(key string)
	//Zadd(key string, score float64, member string)
	//Zadd(key string, score float64, member string, params ZAddParams)
	//Zadd(key string,  scoreMembers  g.Map)
	//Zadd(key string,   scoreMembers g.Map, params ZAddParams)
	//Zrange(key string, start int64,  stop int64)
	//Zrem(key string, members ...string)
	//Zincrby(key string,  increment float64, member string)
	//Zincrby(key string, increment float64, member string,  params ZIncrByParams)
	//Zrank(key string, member string)
	//Zrevrank(key string, member string)
	//Zrevrange(key string, start int64,  stop int64)
	//ZrangeWithScores(key string, start int64,  stop int64)
	//ZrevrangeWithScores(key string, start int64,  stop int64)
	//Zcard(key string)
	//Zscore(key string, member string)
	Watch(keys ...string)
	//sort(String key)
	//sort(String key, SortingParams sortingParameters)
	//blpop(String [] args)
	//sort(String key, SortingParams sortingParameters, String dstkey)
	//sort(String key, String dstkey)
	//brpop(String [] args)
	//brpoplpush(String source, String destination, int timeout)
	//Zcount(key string, min float64, max float64)
	//Zcount(key string, min string, max string)
	//ZrangeByScore(key string, min float64, max float64)
	//ZrangeByScore(key string, min string, max string)
	//ZrangeByScore(key string, min float64, max float64, offset int,count int)
	//ZrangeByScore(key string, min string, max string, offset int, count int)
	//ZrangeByScoreWithScores(key string, min float64, max float64)
	//ZrangeByScoreWithScores(key string, min float64, max float64,offset int, count int)
	//ZrangeByScoreWithScores(key string, min string, max string)
	//ZrangeByScoreWithScores(key string, min string, max string,offset int, count int)
	//ZrevrangeByScore(key string, max float64, min float64)
	//ZrevrangeByScore(key string, max string, min string)
	//ZrevrangeByScore(key string, max float64, min float64, offset int,count int)
	//ZrevrangeByScore(key string, max string, min string, offset int, count int)
	//ZrevrangeByScoreWithScores(key string, max float64, min float64)
	//ZrevrangeByScoreWithScores(key string, max float64, min float64,offset int, count int)
	//ZrevrangeByScoreWithScores(key string, max string, min string)
	//ZrevrangeByScoreWithScores(key string, max string, min string,offset int, count int)
	//ZremrangeByRank(key string, start int64,  stop int64)
	//ZremrangeByScore(key string, min float64, max float64)
	//ZremrangeByScore(key string, min string, max string)
	//Zunionstore(dstkey string, sets ...string)
	//Zunionstore(dstkey string,  params ZParams, sets ...string)
	//Zinterstore(dstkey string, sets ...string)
	//Zinterstore(dstkey string,  params ZParams, sets ...string)
	Strlen(key string)
	Lpushx(key string, string ...string)
	Persist(key string)
	Rpushx(key string, string ...string)
	Echo(String string)
	//Linsert(key string,  where ListPosition,  pivot string, value string)
	//bgrewriteaof()
	//bgsave()
	//lastsave()
	//save()
	ConfigSet(parameter string, value string)
	ConfigGet(pattern string)
	ConfigResetStat()
	Multi()
	Exec()
	Discard()
	//objectRefcount(String key)
	//objectIdletime(String key)
	//objectEncoding(String key)
	//bitcount(String key)
	//bitcount(String key, long start, long end)
	//bitop(BitOP op, String destKey, String ...srcKeys)
	//dump(String key)
	//restore(String key, int ttl, byte [] serializedValue)
	//restoreReplace(String key, int ttl, byte [] serializedValue)
	Scan(cursor int64, params ScanParams)
	Hscan(key string, cursor int64, params ScanParams)
	Sscan(key string, cursor int64, params ScanParams)
	Zscan(key string, cursor int64, params ScanParams)
	//waitReplicas(int replicas, long timeout)
	//bitfield(String key, String ...arguments)
	//hstrlen(String key, String field)
	//migrate(String host, int port, String key, int destinationDB, int timeout)
	//migrate(String host, int port, int destinationDB, int timeout, MigrateParams params, String ...keys)
	//clientKill(String ipPort)
	//clientKill(String ip, int port)
	//clientKill(ClientKillParams params)
	//clientGetname()
	//clientList()
	//clientSetname(String name)
	//memoryDoctor()
	//xadd(String key, StreamEntryID id, Map<String, String> hash, long maxLen, boolean approximateLength)
	//xlen(String key)
	//xrange(String key, StreamEntryID start, StreamEntryID end, long count)
	//xrevrange(String key, StreamEntryID end, StreamEntryID start, int count)
	//xread(int count, long block, Entry<String, StreamEntryID> ...streams)
	//xack(String key, String group, StreamEntryID ...ids)
	//xgroupCreate(String key, String consumer, StreamEntryID id, boolean makeStream)
	//xgroupSetID(String key, String consumer, StreamEntryID id)
	//xgroupDestroy(String key, String consumer)
	//xgroupDelConsumer(String key, String consumer, String consumerName)
	//xdel(String key, StreamEntryID ...ids)
	//xtrim(String key, long maxLen, boolean approximateLength)
	//xreadGroup(String groupname, String consumer, int count, long block, boolean noAck, Entry<String, StreamEntryID> ...streams)
	//xpending(String key, String groupname, StreamEntryID start, StreamEntryID end, int count, String consumername)
	//xclaim(String key, String group, String consumername, long minIdleTime, long newIdleTime, int retries,
	//	boolean force, StreamEntryID ...ids)
}
