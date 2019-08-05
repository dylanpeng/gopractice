package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

func GetPlaceLocationCacheKey(placeName, placeDesc string) string {
	data := []byte(placeName + placeDesc)
	serialKey := fmt.Sprintf("%x", sha1.Sum(data))
	return serialKey
}

type NearbyResultLocation struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func main() {
	cl := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})

	defer cl.Close()

	s := GetPlaceLocationCacheKey("新街口商场", "xinjiekoushangchang")
	fmt.Println(s)

	startLngLat := &NearbyResultLocation{
		Lat: 116.372350,
		Lng: 39.940646,
	}
	serialStartLngLat, _ := json.Marshal(startLngLat)
	fmt.Println(serialStartLngLat)
	_ = cl.Set(s, string(serialStartLngLat), 0).Err()
	//fmt.Println(err)
	r,_:=cl.Get(s).Result()
	//fmt.Println(err)
	fmt.Println(r)
	st := &NearbyResultLocation{}
	json.Unmarshal([]byte(r),st)
	fmt.Println(st)
}