window.define("pages/cards/receive_card/index",function(a,t,e,i,n,o,s,d,r,c,u,h,l,g,f,p){"use strict";function w(a){return a&&a.__esModule?a:{default:a}}var y=a("../../../network/api"),m=w(y),S=a("../../../dist/index"),v=w(S),T=a("../../../utils/util"),_=w(T),D={data:{isPaying:!1}};_.default.registComp(D,v.default.Toast,{onLoad:function(a){this.setData({id:a.id})},onShow:function(){var a=this.data.id,t=this;n.showLoading({title:"加载中"}),setTimeout(function(){t.querycarddetail({id:a})},2e3);var e=n.getStorageSync("userinfo");this.setData({userinfo:e})},show:function(){var a=this.selectComponent("#popup");a&&a.show()},onTapReceiveCard:function(){var a=this,t=n.getStorageSync("userinfo");this.setData({userinfo:t}),t.mobile?(n.showLoading({title:"加载中"}),m.default.receiveCard({id:this.data.id}).then(function(t){n.hideLoading(),console.log(t),a.showZanToast("领取成功"),a.querycarddetail({id:a.data.id})}).catch(function(t){"222000226"==t.error?a.showZanToast(t.errorMessage):a.showZanToast(m.default.errors(t.error)||"网络链接错误"),a.querycarddetail({id:a.data.id})})):this.setData({showModal:!0,url:"/pages/login/login"})},querycarddetail:function(a){var t=this;n.showLoading({title:"加载中"}),m.default.queryCardDetail(a).then(function(a){console.log(a,"+_+_+_+_+_+_+"),n.hideLoading(),t.setData({carddata:a.data,isShow:!0}),2==a.data.del&&t.show()}).catch(function(a){return console.log(a)})},showToast:function(a){this.showZanToast(a)},onToBuyCard:function(){var a=this;if(!this.data.isPaying){this.setData({isPaying:!0});var t=n.getStorageSync("userinfo");this.setData({userinfo:t}),t.mobile?(n.showLoading({title:"加载中"}),m.default.queryCardToBuy({cardId:this.data.carddata.id}).then(function(t){n.hideLoading(),0==t.status&&n.requestPayment({timeStamp:t.data.timeStamp+"",nonceStr:t.data.nonceStr,package:t.data.package,signType:t.data.signType,paySign:t.data.paySign,success:function(){a.setData({isPaying:!1}),n.redirectTo({url:"/pages/cards/result/index?id="+a.data.carddata.id+"&result=1"})},fail:function(){a.setData({isPaying:!1})}})}).catch(function(t){a.setData({isPaying:!1}),console.log(t,"-----------"),n.hideLoading()})):this.setData({showModal:!0,url:"/pages/login/login"})}}}),Page(D)}),window.__swanRoute="pages/cards/receive_card/index",window.usingComponents=["component/phonemodal/phonemodal","dist/minui/wxc-popup/dist/index"],require("pages/cards/receive_card/index");