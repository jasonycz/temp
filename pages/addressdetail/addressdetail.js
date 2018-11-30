window.define("pages/addressdetail/addressdetail",function(t,e,a,o,s,n,i,d,r,c,u,l,g,f,h,p){"use strict";function w(t){return t&&t.__esModule?t:{default:t}}var L=t("../../network/api"),y=w(L),S=t("../../utils/const.js"),v=w(S),m=t("../../utils/notification"),T=w(m),A=t("../../utils/location.js"),C=w(A),D=t("../../utils/util.js"),k=w(D),_=t("../../dist/index"),b=k.default.registComp({},_.Toast,{onLoad:function(t){var e=t.addressid;this.setData({id:e}),s.showLoading({title:"加载中"})},onShow:function(){this.queryAddressList({pageSize:1e3,pageNo:1,gpsType:"gps"})},editaddress:function(t){console.log(t.currentTarget.dataset,"+++++++++");var e=t.currentTarget.dataset.address;s.navigateTo({url:"/pages/address/new/new?type=edit&isAddress=true&item="+k.default.stringify(e)})},newaddress:function(){s.getStorageSync("userinfo").mobile?s.navigateTo({url:"/pages/address/new/new?type=new&isAddress=true"}):this.setData({showModal:!0,url:"/pages/login/login"})},queryAddressList:function(t){var e=this;y.default.queryAddressList(t).then(function(t){s.stopPullDownRefresh(),s.hideLoading(),console.log("bingooooooooooooo",t),e.setData({list:t.data.data,isShow:!0})}).catch(function(t){return console.log(t)})},getLocation:function(){var t=this;s.showLoading({title:"定位中..."}),C.default.getLocation().then(function(){s.hideLoading(),s.navigateBack()}).catch(function(e){s.hideLoading(),console.log(e),t.showZanToast("定位失败，请重试定位或选择其他收货地址")})},makesetting:function(){var t=this;s.getSetting({success:function(e){e.authSetting["scope.userLocation"]?t.getLocation():s.openSetting({success:function(e){e.authSetting["scope.userLocation"]&&t.getLocation()}})}})},addressLocation:function(t){var e=t.currentTarget.dataset.address;y.default.getUserLocation({latitude:e.lat,longitude:e.lng,type:0}).then(function(t){var a=s.getStorageSync("location"),o=Object.assign({},a,{latitude:e.lat,longitude:e.lng,location:e.title,cityName:e.city,addressId:e.id,cityCode:t.data.cityCode});n().globalData.location=o,s.setStorageSync("location",o),T.default.post(v.default.action.LOCATION_CHANGED,o)}),s.navigateBack()},makesearch:function(){s.redirectTo({url:"/pages/addresssearch/addresssearch"})}});Page(b)}),window.__swanRoute="pages/addressdetail/addressdetail",window.usingComponents=["component/phonemodal/phonemodal"],require("pages/addressdetail/addressdetail");