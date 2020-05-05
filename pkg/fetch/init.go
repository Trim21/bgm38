package fetch

import "github.com/go-resty/resty/v2"

func GetClient() *resty.Client {
	return client
}
