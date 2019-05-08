package ws

import (
	"log"
	"strconv"
	"time"

	"github.com/onnidev/api/infra"
	"github.com/spf13/viper"
	redis "gopkg.in/redis.v5"
)

// GetRedisAddr dskjfnsdf
func GetRedisAddr() (string, string) {
	addr := viper.GetString("redishost") + ":" + strconv.Itoa(viper.GetInt("redisport"))
	password := "onnirules"
	if viper.GetString("env") == "homolog" || viper.GetString("env") == "production" {
		addr = "redis-master.default.svc.cluster.local:6379"
		password = "p84axRMOVk"
	}
	return addr, password

}

// Redis is a function that subscribe to a Redis pub/sub channel
func Redis(channelname string, done chan bool) {
	addr, password := GetRedisAddr()
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
	pubsub, err := client.Subscribe(channelname)
	for err != nil {
		time.Sleep(10 * time.Second)
		log.Println(err.Error())
		pubsub, err = client.Subscribe(channelname)
	}
	infra.NewLogger("REDIS").Printf("Subscribed to Redis channel named %s\n", channelname)
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage()
		if err != nil {
			panic(err)
		}

		log.Println("[REDIS] Channel " + channelname + " received " + msg.Payload)
		switch channelname {
		case "onni/dashboard":
			DashboardHub.Broadcast <- []byte(msg.Payload)
		case "onni/staff":
			StaffHub.Broadcast <- []byte(msg.Payload)
		case "onni/app":
			AppHub.Broadcast <- []byte(msg.Payload)
		}

	}
}

// Publish TODO: NEEDS COMMENT INFO
func Publish(channelname, msg string) {
	log.Printf("publishing "+msg+" to REDIS at channel %s", channelname)
	addr, password := GetRedisAddr()
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
	err := client.Publish(channelname, msg).Err()
	if err != nil {
		log.Println(err.Error())
	}
}
