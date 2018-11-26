package library

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/bos"
)

const (
	// 用户的Access Key ID
	AK = "3a9861ff96dd472f875791b14cb0b6a1"
	// 用户的Secret Access Key
	SK = "b3bc7ead16ef4780b928d013ca9ce2b7"
	// 用户指定的Endpoint
	ENDPOINT = "bj.bcebos.com"
	BUCKETNAME = "searchbox"
	CDN = "b.bdstatic.com"
)

func NewBos() (*bos.Client , error){
	// 初始化一个BosClient
	bosClient, err := bos.NewClient(AK, SK, ENDPOINT)

	if err != nil {
		return nil, err
	}
	return bosClient, nil
}

func packageUploadRes(){

}

func UploadFile(filePath string, args...){
	return packageUploadRes()
}

	// fileName := "./group.go"
	// objectName := "2.smapp"
	etag, err := bosClient.PutObjectFromFile(BUCKETNAME, objectName, fileName, nil)
	fmt.Printf("%+v", etag)

	urlCDN := "https://" + CDN + "/" + objectName
	url := "https://" + BUCKETNAME + "." + ENDPOINT + "/" + objectName

	//  http://{$bucketName}.{$region}.bcebos.com/{$objectName}

	// Wiki: https://cloud.baidu.com/doc/BOS/GO-SDK.html#.BC.A8.BF.1E.0D.F7.81.08.DE.22.68.F1.1F.0A.F4.E2
	// url := bosClient.BasicGeneratePresignedUrl(BUCKETNAME, objectName, 0)
	fmt.Printf("%+v\n", url)
	fmt.Printf("%+v\n", urlCDN)
	return


