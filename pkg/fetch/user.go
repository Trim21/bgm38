package fetch

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"

	"bgm38/pkg/db"
	"bgm38/pkg/model"
	"bgm38/pkg/utils"
)

const _cachePrefix = "bgm38:user:info:"

func requestUser(userID string) (model.User, error) {
	u := model.User{}
	res, err := client.R().Get(utils.StrConcat("https://mirror.api.bgm.rin.cat/user/", userID))
	if err != nil {
		return u, err
	}
	err = json.Unmarshal(res.Body(), &u)
	return u, err
}

func User(userID string) (model.User, error) {
	u, err := tryGetCache(userID)
	if err == redis.Nil {
		u, err = requestUser(userID)
		if err != nil {
			setCache(userID, u)
		}
	} else if err != nil {
		return u, nil
	}
	return u, err
}

func Users(userIDs ...string) (map[string]model.User, error) {
	var result = make(map[string]model.User)
	for _, userID := range userIDs {
		u, err := User(userID)
		if err != nil {
			return nil, err
		}
		result[userID] = u
	}
	return result, nil
}

func tryGetCache(userID string) (model.User, error) {
	var u = model.User{}
	var data, err = db.Redis.Get(_cachePrefix + userID).Bytes()
	if err != nil {
		return u, err
	}
	err = json.Unmarshal(data, &u)
	return u, err
}

func setCache(userID string, u model.User) {
	var data, err = json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("set cache")
	db.Redis.Set(_cachePrefix+userID, data, time.Hour*24*30)
}
