package redis

import (
	"blog-for-go/util"
	"github.com/go-redis/redis/v8"
	"log"
)

var rdb *redis.ClusterClient
func GetRedisClient(mode string) *redis.ClusterClient {
	if rdb != nil {
		log.Println("cached rdb")
		return rdb
	}

	addr1 := [4] string { "172.17.0.2:6379", "172.17.0.2:6380", "172.21.42.9:10007", "172.21.42.9:10006" }
	addr2 := [4] string { "172.17.0.3:6379", "172.17.0.3:6380", "172.21.42.9:10005", "172.21.42.9:10004" }
	addr3 := [4] string { "172.17.0.4:6379", "172.17.0.4:6380", "172.21.42.9:10009", "172.21.42.9:10008" }

	var addr11 [2]string
	var addr22 [2]string
	var addr33 [2]string
	if mode == "local" {
		addr11 = [2] string { addr1[0], addr1[1] }
		addr22 = [2] string { addr2[0], addr2[1] }
		addr33 = [2] string { addr3[0], addr3[1] }
	} else {
		addr11 = [2] string { addr1[2], addr1[3] }
		addr22 = [2] string { addr2[2], addr2[3] }
		addr33 = [2] string { addr3[2], addr3[3] }
	}

	clusterSlots := func ()([]redis.ClusterSlot, error) {
		slots :=[]redis.ClusterSlot{
			{
				Start:0,
				End: 5460,
				Nodes: []redis.ClusterNode{
					{
						Addr: addr11[0],
					},
					{
						Addr: addr11[1],
					},
				},
			},
			{
				Start: 5461,
				End: 10922,
				Nodes: []redis.ClusterNode{
					{
						Addr: addr22[0],
					},
					{
						Addr: addr22[1],
					},
				},
			},
			{
				Start: 10923,
				End: 16383,
				Nodes: []redis.ClusterNode{
					{
						Addr: addr33[0],
					},
					{
						Addr: addr33[1],
					},
				},
			},
		}

		log.Printf("slots:%s\n", util.PrettyJSON(slots))

		return slots,nil
	}

	rdb =redis.NewClusterClient(&redis.ClusterOptions{
		ClusterSlots: clusterSlots,
		RouteRandomly: true,
	})

	log.Println("new created rdb")

	return rdb
}
