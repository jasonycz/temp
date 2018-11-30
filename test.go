package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

const MAX_MERGE_DEPTH = 5

func main() {

	extJson := `
	{"extEnable":true,"extAppid":"14851328","directCommit":false,"ext":{"appName":"\u6570\u5b57\u95e8\u5e97\u6388\u6743\u6d4b\u8bd5","appId":8006},"extPages":{"pages\/home\/home":{"navigationBarTitleText":"\u6570\u5b57\u95e8\u5e97\u6388\u6743\u6d4b\u8bd5"},"pages\/map\/map":{"navigationBarTitleText":"\u6570\u5b57\u95e8\u5e97\u6388\u6743\u6d4b\u8bd5"}},"window":{"navigationBarBackgroundColor":"#ffffff","navigationBarTitleText":"\u6570\u5b57\u95e8\u5e97\u6388\u6743\u6d4b\u8bd5","navigationBarTextStyle":"black","backgroundColor":"#f4f4f4","backgroundTextStyle":"light","enablePullDownRefresh":false},"tabBar":{"color":"#666666","selectedColor":"#E93E4E","borderStyle":"white","backgroundColor":"#ffffff","list":[{"pagePath":"pages\/launch\/launch","iconPath":"assets\/images\/tabBar\/index.png","selectedIconPath":"assets\/images\/tabBar\/index_selected.png","text":"\u9996\u9875"},{"pagePath":"pages\/classify\/classify","iconPath":"assets\/images\/tabBar\/category.png","selectedIconPath":"assets\/images\/tabBar\/category_selected.png","text":"\u5546\u54c1\u5206\u7c7b"},{"pagePath":"pages\/store\/store-list\/store-list","iconPath":"assets\/images\/tabBar\/store.png","selectedIconPath":"assets\/images\/tabBar\/store_selected.png","text":"\u9644\u8fd1\u95e8\u5e97"},{"pagePath":"pages\/member\/mine\/mine","iconPath":"assets\/images\/tabBar\/profile.png","selectedIconPath":"assets\/images\/tabBar\/profile_selected.png","text":"\u6211"}]}}
	`
	appJson := `
		{"_sub_swan":{"pages/car":"4ec84c7e0bf91f439b830b93eeec7e44","pages/decoration":"01377692670aded09db0cd67f9b6467a","pages/edu":"131cc74384468c534cfeaabce54ea5c2","pages/fenxiao":"78099b483a3f7ac4f4673abc6e2317c3","pages/food":"e931dfec1380c3fc630f9a0d4a60f94c","pages/giftcard":"adf14a661e63ff4584f2d331aaffe1a9","pages/gongqiu":"175bebd5f45542b3a8181819c7040f84","pages/hotel":"67d6f26a9fea33a7c0de648e3b9f81ec","pages/live":"44acee9a56e1cb0ae32e7a2fe170da60","pages/mine":"20094f26704194ce4cae956730ef888d","pages/orders":"6877153eed97d38a336171c4bc0577e7","pages/returngoods":"ab4f546c3dbfa3cbefe37ba6593d3d8f","pages/services":"414e5f2201edae1e7c263d3f6014ee1c","pages/shop":"7e7ad462120ea7c662b7e8c6e7410a2c","pages/sources":"69c68a268eeb44a29240d02626ed9630","pages/storevalue":"f6d1ebca7807a6e54e6fc71fe0625b4f","subs/pingce":"64370a1c2ca6f90587b2c151adad6bcd"},"pages":["tabs/home/index","tabs/mine/index","tabs/order/index","tabs/about/index","tabs/map/index","tabs/brandlist/index","tabs/custome/custome","tabs/carts/index","tabs/publish/index","tabs/supplyBuy/index","tabs/info/index","tabs/fenxiao/index","tabs/giftCard/index","tabs/goodsCategory/index","tabs/company/index","tabs/picture/index","tabs/designer/index","tabs/live/index","modules/map/index","modules/custome/custome","pages/webview/index","pages/addressdetail/addressdetail","pages/addresssearch/addresssearch","pages/citys/citys","pages/login/login","pages/brandlist/index","pages/aboutmine/aboutmine","pages/page/search/search","pages/storelist/storelist","pages/address/list/list","pages/address/search/search","pages/address/new/new","pages/brandmap/brandmap","pages/products/show/show","pages/login/verify_cell","pages/cards/receive_card/index","pages/cards/result/index","pages/coupons/receive_coupon/index","pages/page/products/index","pages/tuan/mytuan/index","pages/tuan/detail/index","component/address_picker/index","pages/supplys/supplys","pages/manjian/index","pages/zhaomu/register","pages/pingjia/myevaluate/index","pages/pingjia/productevaluate/index","pages/pingjia/storeevaluate/index","pages/pingjia/storeevaluatelist/index","pages/pingjia/evaluateform/index","pages/pingjia/evaluatesuccess/index","pages/agreement/index","pages/cutprice/detail/index"],"subPackages":[{"pages":["products/products","desc/desc","search/search","classify/index","album/album","carts/index"],"root":"pages/shop"},{"pages":["new/index","choosetype/index","mygongqiu/index","supplyDetail/index"],"root":"pages/gongqiu"},{"pages":["mybalance/index","mystore/index","result/index"],"root":"pages/storevalue"},{"pages":["user/profile/index","cards/card_detail","coupons/coupons","collection/index","cards/cards","scores/scores","scores/score_mall","scores/score_new_order","scores/score_orders","scores/score_order_show","scores/score_product","cutprice/index"],"root":"pages/mine"},{"pages":["brands/index","services/index","products/index","stores/index","miaosha/index","tuan/index","picture/index","newsinfo/newsinfo","newsinfo/newsinfodetail","newsinfo/newssearch","supplyinfo/supplyinfo","coupon/index","designer/designer","company/company","cutprice/index"],"root":"pages/sources"},{"pages":["order_create/index","order_detail/index","order_result/index","select_service/index","select_worker/index","service_detail/index","orders/index","worker_detail/index","services/index","workers/index","service_textarea/index"],"root":"pages/services"},{"pages":["consult/index","apply/index","applydetail/index","logistics/index","applyform/index","list/index"],"root":"pages/returngoods"},{"pages":["sources/giftcardlist/index","detail/index","receive/index","mygiftcard/index","record/index","buy/index","choose/index","result/index","give/index"],"root":"pages/giftcard"},{"pages":["stores/index","account/index","bank_account/index","center/index","customers/index","orders/index","ordersdetail/index","products/index","products/store","withdraw/index","withdraw_histories/index","withdraw_detail/index","search/index","lowerlevel/index","lowerlevel_detail/index"],"root":"pages/fenxiao"},{"pages":["cards/cards","coupons/coupons","list/list","show/show","o2oshow/show","new/new","showstates/showstates","showlogistics/showlogistics","result/index","new/newmsg"],"root":"pages/orders"},{"pages":["myfood/index","fooddetail/index","createorder/index","foodtextarea/index","result/index"],"root":"pages/food"},{"pages":["list/index","orders/index","order_detail/index","room_detail/index","order_create/index","checkdate/index","phonedetail/index","order_result/index"],"root":"pages/hotel"},{"pages":["detail/index","push/index","list/index","favorites/index","available/index"],"root":"pages/live"},{"pages":["detail/index","my/index","sources/list/index","recorddetail/index","managerdetail/index","timescardstorelist/index","storelist/index","projectlist/index","proordercreate/index","projectorders/index","projectorderdetail/index"],"root":"pages/car"},{"pages":["designer/detail/index","company/detail/index","quotation/index","viewpager/index","search/index"],"root":"pages/decoration"},{"pages":["coursedetail/index","list/index","favorites/index","columnlist/index"],"root":"pages/edu"},{"pages":["bodytest/index","testresult/index","resultdetail/index","testcourse/index"],"root":"subs/pingce"}],"tabBar":{"backgroundColor":"#FFFFFF","borderStyle":"black","color":"#5F646E","list":[{"iconPath":"images/tab/mic_1001.png","pagePath":"tabs/home/index","selectedIconPath":"images/tab/mic_1001_focus.png","text":"首页"},{"iconPath":"images/tab/mic_2002.png","pagePath":"tabs/map/index","selectedIconPath":"images/tab/mic_2002_focus.png","text":"地图"},{"iconPath":"images/tab/mic_4001.png","pagePath":"tabs/mine/index","selectedIconPath":"images/tab/mic_4001_focus.png","text":"我的"},{"iconPath":"images/tab/mic_9003.png","pagePath":"tabs/about/index","selectedIconPath":"images/tab/mic_9003_focus.png","text":"关于"},{"iconPath":"images/tab/mic_9002.png","pagePath":"tabs/map/index","selectedIconPath":"images/tab/mic_9002_focus.png","text":"地图"}],"selectedColor":"#ca2626"},"window":{"backgroundColor":"#f1f1f1","backgroundTextStyle":"red","navigationBarBackgroundColor":"#ffffff","navigationBarTextStyle":"light","navigationBarTitleText":"WeChat"}}
	`

	var extJsonStruct map[string]interface{}
	err := json.Unmarshal([]byte(extJson), &extJsonStruct)
	if err != nil {
		panic(err)
	}

	var appJsonStruct map[string]interface{}
	err = json.Unmarshal([]byte(appJson), &appJsonStruct)
	if err != nil {
		panic(err)
	}

	if extEnable, ok := extJsonStruct["extEnable"]; ok {
		if extEnable == true {
			// 删除 extEnable
			delete(extJsonStruct, "extEnable")
		}

		for key, value := range extJsonStruct {
			// 非extPages 节点下数据处理
			if key != "extPages" {
				// 如果有就 merge 没有就添加到 appJsonStruct[key]
				aJKeyMap, aJKeyMapOk := getMap(appJsonStruct[key])
				eJKeyMap, eJKeyMapOk := getMap(extJsonStruct[key])

				if aJKeyMapOk && eJKeyMapOk {
					appJsonStruct[key] = MapMerge(aJKeyMap, eJKeyMap)
				}
				// extPage 节点下的数据处理
			} else {
				vMap, vMapOk := getMap(value)
				if vMapOk {
					for k, v := range vMap {
						// 获取文件加下 k.json
						// 读取文件
						kJson, err := os.Open(k + ".json")
						if err != nil {
							panic(err)
							// return errors.New("open " + k + ".json fail " + err.Error())
						}
						// 解析 app.json
						var kJsonStruct map[string]interface{}
						jsonParser := json.NewDecoder(kJson)
						if err = jsonParser.Decode(&kJsonStruct); err != nil {
							panic(err)
							// return errors.New(k + ".json content is not valid json " + err.Error())
						}

						vMap, vMapOk := getMap(v)
						kJsonStructMap, kJsonStructMapOk := getMap(kJsonStruct)

						if vMapOk && kJsonStructMapOk {
							kJsonStruct = MapMerge(kJsonStructMap, vMap)
							kJson, err := json.Marshal(kJsonStruct)
							if err != nil {
								panic(err)
							}
							// 写入到文件中
							err = ioutil.WriteFile(k+".json", kJson, 0644)
							if err != nil {
								panic(err)
							}
						}
					}
				}

			}
		}
		// 写入到文件app.json 中
		appJsonByte, err := json.Marshal(appJsonStruct)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile("app.json", appJsonByte, 0644)
		if err != nil {
			panic(err)
		}

	}

	fmt.Printf("%+v", extJsonStruct)
}

func MapMerge(dst, src map[string]interface{}) map[string]interface{} {
	return merge(dst, src, 0)
}

func merge(dst, src map[string]interface{}, depth int) map[string]interface{} {
	if depth > MAX_MERGE_DEPTH {
		return dst
	}

	for key, srcVal := range src {
		if dstVal, ok := dst[key]; ok {
			srcMap, srcMapOk := getMap(srcVal)
			dstMap, dstMapOk := getMap(dstVal)

			if srcMapOk && dstMapOk {
				srcVal = merge(dstMap, srcMap, depth+1)
			}
		}
		dst[key] = srcVal
	}

	return dst
}

func getMap(i interface{}) (map[string]interface{}, bool) {
	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Map {
		m := map[string]interface{}{}
		for _, k := range value.MapKeys() {
			m[k.String()] = value.MapIndex(k).Interface()
		}
		return m, true
	}
	return map[string]interface{}{}, false
}
