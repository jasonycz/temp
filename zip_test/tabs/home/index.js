window.define("tabs/home/index",function(t,i,a,e,n,o,s,c,d,r,u,h,l,f,p,g){"use strict";function v(t){return t&&t.__esModule?t:{default:t}}var m=t("../../network/api.js"),y=v(m),C=t("../../utils/const.js"),L=v(C),w=t("../../utils/notification"),b=v(w),x=t("../../utils/location.js"),D=v(x);Page({data:{locationSuccess:!0,activityCards:[],storeType:""},onLoad:function(t){this.registeNotification();var i=o().globalData.editionType;this.init(),this.setData({code:t.scene||"",storeType:i})},onShow:function(){},onHide:function(){var t=this.data,i=t.activityCards,a=t.tabName,e=t.activityIndex;e<i.length-1?n.setStorageSync("activity",{activityCards:i,tabName:a,activityIndex:e}):n.setStorageSync("activity",{})},onShareAppMessage:function(t){var i=o().globalData.userinfo;t.from,t.target;return{imageUrl:this.data.shareImg,title:this.data.desc||this.data.title,path:"/tabs/home/index?shareCode="+i.shareCode}},onUnload:function(){this.removeNotification()},init:function(){var t=this,i=o().globalConfig;i&&(this.setData(Object.assign({},this.data,{globalConfig:i})),i.isEnableLocation?(n.showLoading({title:"定位中..."}),D.default.ensureLocation(i).then(function(){}).catch(function(i){console.log(i),t.setData({locationSuccess:!1}),n.stopPullDownRefresh(),n.hideLoading()})):this.queryData())},queryData:function(){var t=this;n.showLoading({title:"请稍候"}),y.default.init().then(function(i){t.setData(Object.assign({},i.data,{serverTime:i.serverTime})),n.stopPullDownRefresh(),n.hideLoading()}).catch(function(t){console.log(t),n.stopPullDownRefresh(),n.hideLoading()})},onPullDownRefresh:function(){var t=this.data,i=t.globalConfig,a=t.location;1==i.locationType?this.init():a&&a.location?this.queryData():this.init()},registeNotification:function(){var t=this;b.default.register(L.default.action.QUERY_CONFIG_SUCCESS,this,function(){t.init()}),b.default.register(L.default.action.LOCATION_CHANGED,this,function(i){console.log("location_changed",i.data),t.setData({location:i.data,locationSuccess:!0}),t.queryData()}),b.default.register(L.default.action.GOTO_SHOP,this,function(t){console.log("get notif从订单页面跳转到店铺首页",t),n.navigateTo({url:"/pages/shop/products/products?storeid="+t.data.storeId})})},removeNotification:function(){b.default.remove(L.default.action.QUERY_CONFIG_SUCCESS,this),b.default.remove(L.default.action.LOCATION_CHANGED,this),b.default.remove(L.default.action.GOTO_SHOP,this)},handlePhoneBtnTap:function(t){var i=t.currentTarget.dataset.phone,a=i.telephone,e=i.cellphone,n=i.extNum,o=[];a&&(n?o.push(a+"转"+n):o.push(a)),e&&o.push(e),this.setData({phoneList:o});var s=this.selectComponent(".phoneaction");s&&s.show()},handleChooseAddress:function(){var t=this.data,i=t.globalConfig,a=t.location,e=t.storeType,o=void 0;o=0==i.locationType||"mall"==e?"/pages/citys/citys":"/pages/addressdetail/addressdetail?addressid="+(a?a.addressId:""),n.navigateTo({url:o})},queryActivity:function(t,i){var a=this;y.default.queryActivity({scene:t,code:i}).then(function(t){if(t&&t.data&&t.data.list){var i=[];if(t.data.list.forEach(function(t){5===t.status&&(t.VIPCardList&&t.VIPCardList.length&&i.push({name:"card",data:t.VIPCardList}),t.couponList&&t.couponList.length&&i.push({name:"coupon",data:t.couponList.map(function(t){return t.useTime=t.useTime.substring(0,11),t.useEndTime=t.useEndTime.substring(0,11),t})}),t.giftCardList&&t.giftCardList.length&&i.push({name:"gift",data:t.giftCardList}),t.point&&i.push({name:"score",data:{score:t.point}}),"smyl"===t.type&&(t.VIPCardList&&t.VIPCardList.length||t.couponList&&t.couponList.length||t.giftCardList&&t.giftCardList.length||t.point)&&i.push({name:"share",data:t.share}))}),i.length)a.setData({activityCards:i,tabName:t.data.myName,activityIndex:0}),a.showGiftPopup();else{var e=a.data,o=e.activityCards;if(e.activityIndex<o.length-1)return;var s=n.getStorageSync("activity");s.activityCards&&s.activityCards.length&&(a.setData({activityCards:s.activityCards,tabName:s.tabName,activityIndex:s.activityIndex}),a.showGiftPopup())}}}).catch(function(t){console.log("---err---",t)})},showGiftPopup:function(){var t=this.selectComponent(".home_gift_popup");t&&t.show()},hideGiftPopup:function(){var t=this.selectComponent(".home_gift_popup");t&&t.hide()},updateindex:function(t){this.setData({activityIndex:t.detail.index})}})}),window.__swanRoute="tabs/home/index",window.usingComponents=["component/page/index","dist/minui/wxc-abnor/dist/index","component/home_gift_popup/index","dist/minui/wxc-icon/dist/index","component/phone_action/index"],require("tabs/home/index");