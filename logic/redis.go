package main

import (
	"bytes"
	"encoding/json"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"im/libs/define"
	"im/libs/proto"
)

var (
	RedisCli *redis.Client
)

func InitRedis() (err error) {
	RedisCli = redis.NewClient(&redis.Options{
		Addr:     Conf.Redis.RedisAddr,
		Password: Conf.Redis.RedisPw,        // no password set
		DB:       Conf.Redis.RedisDefaultDB, // use default DB
	})
	if pong, err := RedisCli.Ping().Result(); err != nil {
		log.Infof("RedisCli Ping Result pong: %s,  err: %s", pong, err)
	}

	return
}

// 发布订阅消息
func RedisPublishCh(serverId int8, uid string, msg []byte) (err error) {
	var redisMsg = &proto.RedisMsg{
		Op:       define.OP_SINGLE_SEND,
		ServerId: serverId,
		UserId:   uid,
		Msg:      msg,
	}

	redisMsgStr, err := json.Marshal(redisMsg)

	log.Infof("RedisPublishCh redisMsg info : %s", redisMsgStr)

	err = RedisCli.Publish(define.REDIS_SUB, redisMsgStr).Err()
	return
}

// 发布到房间
func RedisPublishRoom(rid int32, msg []byte) (err error) {
	var redisMsg = &proto.RedisMsg{
		Op:     define.OP_ROOM_SEND,
		RoomId: rid,
		Msg:    msg,
	}
	redisMsgStr, err := json.Marshal(redisMsg)
	log.Infof("RedisPublishRoom redisMsg info : %s", redisMsgStr)
	err = RedisCli.Publish(define.REDIS_SUB, redisMsgStr).Err()
	return
}

func RedisPublishRoomCount(rid int32, count int) (err error) {
	var redisMsg = &proto.RedisRoomCount{
		Op:     define.OP_ROOM_COUNT_SEND,
		RoomId: rid,
		Count:  count,
	}
	redisMsgStr, err := json.Marshal(redisMsg)
	log.Infof("RedisPublishRoomCount redisMsg info : %s", redisMsgStr)
	err = RedisCli.Publish(define.REDIS_SUB, redisMsgStr).Err()
	return
}

func RedisPublishRoomInfo(rid int32, count int, RoomUserInfo map[string]string) (err error) {
	// , roomUserList []
	var redisMsg = &proto.RedisRoomInfo{
		Op:           define.OP_ROOM_INFO_SEND,
		RoomId:       rid,
		Count:        count,
		RoomUserInfo: RoomUserInfo,
	}
	redisMsgByte, err := json.Marshal(redisMsg)
	log.Infof("RedisPublishRoomInfo redisMsg info : %s", redisMsgByte)
	err = RedisCli.Publish(define.REDIS_SUB, redisMsgByte).Err()
	return
}

/**
减少指定在线用户信息(暂未用到)
*/
func RedisPublishRoomUserLess(rid int32, count int, RoomUserInfo map[string]string) (err error) {
	// , roomUserList []
	var redisMsg = &proto.RedisRoomInfo{
		Op:           define.OP_ROOM_INFO_SEND,
		RoomId:       rid,
		Count:        count,
		RoomUserInfo: RoomUserInfo,
	}
	redisMsgByte, err := json.Marshal(redisMsg)
	log.Infof("RedisPublishRoomInfo redisMsg info : %s", redisMsgByte)
	err = RedisCli.Publish(define.REDIS_SUB, redisMsgByte).Err()
	return
}

func getKey(key string) string {

	var returnKey bytes.Buffer
	returnKey.WriteString(define.REDIS_PREFIX)
	returnKey.WriteString(key)
	return returnKey.String()
}

func getRoomUserKey(key string) string {

	var returnKey bytes.Buffer
	returnKey.WriteString(define.REDIS_ROOM_USER_PREFIX)
	returnKey.WriteString(key)
	return returnKey.String()
}
