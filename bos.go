package main

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/bos"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	AK, SK := "3a9861ff96dd472f875791b14cb0b6a1", "b3bc7ead16ef4780b928d013ca9ce2b7"

	// 用户指定的Endpoint
	ENDPOINT := "bj.bcebos.com"
	BUCKETNAME := "searchbox"
	// CDN := "b.bdstatic.com"

	// 初始化一个BosClient
	bosClient, err := bos.NewClient(AK, SK, ENDPOINT)

	if err != nil {
		fmt.Printf("%+v", err)
		return
	}

	// fmt.Printf("%+v", bosClient)
	// result, err := bosClient.GetBucketAcl(BUCKETNAME)

	etag, err := bosClient.PutObjectFromFile(BUCKETNAME, "1", fileName, nil)
	fmt.Printf("%+v", result)
	return

}
