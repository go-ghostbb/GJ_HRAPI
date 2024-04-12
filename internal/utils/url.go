package utils

import (
	"context"
	"fmt"
	"ghostbb.io/gb/frame/g"
)

func GetURL(route string) (string, error) {
	if string(route[0]) == "/" {
		route = route[1:]
	}

	// 取得domain
	domain, err := g.Cfg().Get(context.TODO(), "domain")
	if err != nil {
		return "", err
	}

	if domain.IsEmpty() {
		// 如果資料為空，給一個預設網域
		domain.Set("127.0.0.1:5180")
	}

	return fmt.Sprintf("http://%s/%s", domain.String(), route), nil
}
