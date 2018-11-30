window.define("pages/products/show/show",function(t,a,e,i,o,s,n,r,d,h,l,u,c,g,m,p){"use strict";function f(t){return t&&t.__esModule?t:{default:t}}function y(t,a,e){return a in t?Object.defineProperty(t,a,{value:e,enumerable:!0,configurable:!0,writable:!0}):t[a]=e,t}var v=t("../../../network/api"),D=f(v),w=t("../../../component/open_tuan/index.js"),T=f(w),S=t("../../../utils/const.js"),k=f(S),x=t("../../../utils/notification"),P=f(x),B=t("../../../utils/util.js"),I=f(B),b=t("../../../utils/wxParse/wxParse.js"),C=t("../../../dist/index");Page(Object.assign({},C.Quantity,C.Tab,C.TopTips,C.Toast,T.default,{data:{loading:!0,showDialog:!1,showDialog1:!1,showShareDialog:!1,action:0,cartCount:0,formData:{num:{quantity:1,min:1,max:9999,hasSelect:!1},sku:{price:0,quantity:0},selectSku:{},selectSkuText:""},tabs:{list:[{id:"2",title:"商品详情"}],selectedId:"1",scroll:!1,height:42},tabs2:{list:[{id:"3",title:"购物须知"}],selectedId:"1",scroll:!1,height:42},goods:{id:0,name:"",price:0,props:[],content:"",images:[],coverImg:"",specs:[],skus:[]},relativeModule:{pttext:{listType:"type2",cartType:"type1",goods:[]}},toggle:!0,isFirstLoading:!0,isOnTapPic:!1,limit:!1},onLoad:function(t){var a=this;this.setData({id:t.id,shareCode:t.shareCode||"",isFromShop:t.isFromShop,from:t.from?"normal"==t.from?"":t.from:"",height:750,isService:t.isService||!1});var e=o.getStorageSync("config");console.log("+++++++++config",e),this.setData({config:e}),P.default.register(k.default.action.QUERY_CONFIG_SUCCESS,this,function(t){a.setData({config:t.data})})},onShow:function(){this.data.isOnTapPic||(this.data.isFirstLoading?(this.queryProductById(!0),this.setData({isFirstLoading:!1})):this.queryProductById(!1))},onUnload:function(){clearInterval(this.seckKilltimer),this.seckKilltimer=void 0,clearInterval(this.timerMore),this.timerMore=void 0,clearInterval(this.openTuanTimer),this.openTuanTimer=void 0,P.default.remove(k.default.action.QUERY_CONFIG_SUCCESS,this)},onShareAppMessage:function(){var t=s().globalData.userinfo,a=this,e="/pages/products/show/show?id="+this.data.id+"&shareCode="+t.shareCode;"distribution"==this.data.from&&1==this.data.goods.isDist&&(e+="&from=share"),console.log("===path=====onShareAppMessage",e,this.data.goods.storeId);var i=this.data.goods.name;return 2==this.data.goods.activityType&&(i="【"+this.data.goods.groupEvent.limitPeople+"人团】"+this.data.goods.groupPrice+"元拼 "+i+"，快来拼团吧!"),{title:i,path:e,success:function(){1!=a.data.goods.isShare?a.collection():D.default.onShareStore({storeId:a.data.goods.storeId})}}},startSecKillCount:function(){var t=this,a=this;this.seckKilltimer=setInterval(function(){if(0==a.data.goods.overTime)return clearInterval(t.timer),void t.queryProductById(!1);var e=a.data.goods.overTime-1;if(e<1)return clearInterval(t.timer),void(t.timer=void 0);a.setData({"goods.overTime":e,"goods.leftTime":I.default.renderTimer(e)})},1e3)},call:function(){var t=this,a=this.data.store,e=a.areaNum,i=a.tel,o=a.extNum,s=[],n=e?e+"-"+i:i;if(console.log("--tel--",i),i){o?s.push(n+"转"+o):s.push(n),console.log("--phone--",s),this.setData({phoneList:s});var r=this.selectComponent(".phoneaction");console.log("phoneaction",r),r&&r.show()}else this.setData({$toast:{show:!0}}),setTimeout(function(){t.setData({$toast:{show:!1}})},1500)},queryProductById:function(t){var a=this,e={};e.id=this.data.id,e.from=this.data.from,"distribution"==this.data.from&&(e.from="share"),e.shareCode=this.data.shareCode,t&&o.showLoading({title:"请稍候"}),console.log("获取商品详情params====>",e),D.default.getStoreGoodDetail(e).then(function(e){o.stopPullDownRefresh(),o.hideLoading();var i=e.data.imgUrl.map(function(t){return{imgUrl:t,link:t,title:t}}),s=[],n="";if(e.data.item&&e.data.item.length>0){console.log("res.data.item.length>0");for(var r=0;r<e.data.item.length;r++)if(e.data.item[r]&&e.data.item[r].values&&e.data.item[r].values.length>0){for(var d=[],h=0;h<e.data.item[r].values.length;h++)d[h]={id:e.data.item[r].values[h].itemValueId,name:e.data.item[r].values[h].itemValue};console.log("-specsList--------------\x3e",d),s[r]={id:e.data.item[r].itemId,name:e.data.item[r].name,specsList:d},0!=r?n+=","+e.data.item[r].name:n=e.data.item[r].name}}else n="数量";console.log("-tempSpecs--------------\x3e",s);var l=[];if(e.data.storeItem&&e.data.storeItem.length>0)for(var u=0;u<e.data.storeItem.length;u++){var c=e.data.storeItem[u].price,g=e.data.storeItem[u].groupPrice||0;4==e.data.activityType&&(c=e.data.storeItem[u].sharePrice),l[u]={id:e.data.storeItem[u].id,properties:e.data.storeItem[u].properties,quantity:e.data.storeItem[u].stock,price:c,groupPrice:g,minBuy:e.data.storeItem[u].minBuy,maxBuy:e.data.storeItem[u].maxBuy}}console.log("this.dealPrice(res.data.originalPrice)==>",a.dealPrice(e.data.originalPrice)),a.setData({goods:Object.assign({},e.data,{formatPrice:a.dealPrice(e.data.price),formatoriginalPrice:a.dealPrice(e.data.originalPrice),formatgroupPrice:a.dealPrice(e.data.groupPrice),formatsharePrice:a.dealPrice(e.data.sharePrice),gstock:e.data.stock,images:i,imgs:e.data.imgUrl,content:e.data.description,coverImg:e.data.imgUrl?e.data.imgUrl[0]:"",specs:s,skus:l,shareCode:e.data.shareCode||"",isShare:e.data.isShare,seckillStatus:e.data.seckillStatus,overTime:e.data.overTime||0,leftTime:1==e.data.activityType?I.default.renderTimer(e.data.overTime):{}}),store:e.data.store,loading:!1,skustexts:n,formData:{num:{quantity:0==e.data.stock?0:e.data.minBuy>0?e.data.minBuy:1,min:0==e.data.stock?0:1,max:e.data.stock,minBuy:0==e.data.stock?0:e.data.minBuy>0?e.data.minBuy:1,maxBuy:e.data.maxBuy>0&&e.data.maxBuy<e.data.stock?e.data.maxBuy:e.data.stock,hasSelect:!1},sku:{id:0,price:4==e.data.activityType?e.data.sharePrice:e.data.price,quantity:e.data.stock,groupPrice:e.data.groupPrice},selectSku:{},selectSkuText:""}}),t&&(b.wxParse("content","html",e.data.description?e.data.description:"暂无",a,5),b.wxParse("shopNotice","html",e.data.shopNotice?e.data.shopNotice:"暂无",a,5)),1==e.data.activityType&&e.data.overTime&&e.data.overTime>0?a.startSecKillCount():2==e.data.activityType&&e.data.groupEvent&&e.data.groupEvent.sponsoringList&&e.data.groupEvent.sponsoringList.length>0&&a.startOpenTuanCount()}).catch(function(t){console.log(t,";;;;;;;;;;;;;;;;;;;;;;;;;"),o.stopPullDownRefresh(),o.hideLoading()})},gotoshop:function(){console.log(this.data.isFromShop,"this.data.isFromShop"),0==this.data.config.bottomMenu?o.switchTab({url:"/tabs/home/index"}):this.data.isFromShop?o.navigateBack({delta:1}):o.redirectTo({url:"/pages/shop/products/products?storeid="+this.data.store.id})},onTapCarousel:function(t){o.previewImage({current:t.currentTarget.dataset.link,urls:this.data.goods.imgs})},toogleSpecDialog:function(){var t=this.data.showDialog?0:1,a=Object.assign({},this.data.formData);a.num.hasSelect=!0,this.setData({showDialog:!this.data.showDialog,action:t,formData:a,limit:!1})},hideDialog:function(){this.setData({showDialog:!1,action:0})},hideDialog1:function(){this.setData({showDialog1:!1})},handleZanQuantityChange:function(t){var a=parseInt(t.quantity),e=isNaN(a)?1:a;e==this.data.formData.num.max&&this.showZanTopTips("已达最大库存数量"),2==this.data.goods.activityType&&6==this.data.action?2==this.data.goods.groupEvent.isBuyLimited&&(e==this.data.goods.groupEvent.limitedCount?this.setData({limit:!0}):this.setData({limit:!1})):(0==this.data.goods.activityType&&3==this.data.action||2==this.data.goods.activityType&&3==this.data.action)&&this.data.formData.sku.maxBuy>0&&(e==this.data.formData.sku.maxBuy?this.setData({limit:!0}):this.setData({limit:!1})),this.setData(y({},"formData.num.quantity",e))},handleBuyNow:function(){this.normalBuyNow()},resetSku:function(){var t=this.data.goods.storeItem;console.log("storeItem====>",t);var a=[],e=this.data.goods;e=Object.assign({},e,{skus:a}),this.setData({goods:e})},normalBuyNow:function(){var t=this,a=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{},e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:3,i=s().globalData.userinfo;if(console.log(this.data.action,"KKKKKKKKKKKKKKK",i.accessToken),!i||!i.mobile)return o.navigateTo({url:"/pages/login/login"}),void this.hideDialog();var n=this.data.formData.selectSku,r=this.data.formData.num.hasSelect,d=this.data.formData.num.quantity,h=this.data.goods.specs.length==Object.getOwnPropertyNames(n).length;if(2==this.data.goods.activityType&&6==e?2==this.data.goods.groupEvent.isBuyLimited&&(d==this.data.goods.groupEvent.limitedCount?this.setData({limit:!0}):this.setData({limit:!1})):2==this.data.goods.activityType&&3==e&&this.setData({limit:!1}),(0==this.data.goods.activityType&&3==e||2==this.data.goods.activityType&&3==this.data.action)&&(this.setData({limit:!1}),this.data.formData.sku.maxBuy>0&&(d==this.data.formData.sku.maxBuy?this.setData({limit:!0}):this.setData({limit:!1}))),console.log(this.data.formData.num.quantity,"666",this.data.limit,"777",this.data.action),0==this.data.goods.gstock)this.showZanTopTips("库存为零");else if(r&&h)o.showLoading({title:"请稍候"}),this.setData({showDialog:!1,action:0}),D.default.checkoutSingleOrder(Object.assign({goodsId:this.data.goods.id,propertiesId:this.data.formData.sku.id,num:d,consigneeId:-1},a)).then(function(e){o.hideLoading(),t.setData({showDialog:!1,action:0});var i=e.data;i.shareCode=a.shareCode||"",console.log("checkoutSingleOrder====>",i),1==i.buyLimit?o.showModal({title:"提示",content:"活动期间每人限购"+t.data.goods.limitedCount+"件，\n超出"+t.data.goods.limitedCount+"件恢复原价",success:function(){o.navigateTo({url:"/pages/orders/new/new?order="+I.default.stringify(i)})}}):o.navigateTo({url:"/pages/orders/new/new?order="+I.default.stringify(i)})}).catch(function(a){console.log("err----\x3e",a);var e="";a.data&&a.data.data&&a.data.data.length>0?(a.data.data.forEach(function(t){var a=t.properties.split(";"),i="";a.forEach(function(t){t&&t.length>0&&(i+=t.split(":")[1]||"")});var o=t.maxBuy?"最多购买"+t.maxBuy+"件":""+t.maxBuy?"最少购买"+t.minBuy+"件":"";e+=t.goodsName+":"+i+o+"  "}),e+="超出购买范围"):e=a.data&&a.data.msg?a.data.msg:D.default.errors(a.error)||a.errorMessage||"订单创建失败",t.showZanTopTips(e),o.hideLoading(),t.setData({showDialog:!1,action:0})});else if(!this.data.showDialog){var l=Object.assign({},this.data.formData);l.num.hasSelect=!0,this.setData({showDialog:!0,action:e,formData:l})}},handleAddCart:function(){var t=this,a=this.data.formData.selectSku,e=this.data.formData.num.hasSelect,i=this.data.formData.num.quantity,s=this.data.goods.specs.length==Object.getOwnPropertyNames(a).length;if(console.log("----handleAddCart--------",s,e),0==this.data.goods.gstock)this.showZanTopTips("库存为零");else if(e&&s){o.showLoading({title:"请稍候"}),this.setData({showDialog:!1,action:0});var n=this;D.default.addToCart({goodsId:this.data.goods.id,propertiesId:this.data.formData.sku.id,quantity:i}).then(function(){o.hideLoading(),n.showZanTopTips(t.data.goods.name+"加入购物车成功"),t.setData({showDialog:!1,action:0})}).catch(function(a){console.log("err===============",a);var e=a.data&&a.data.msg?a.data.msg:"加入购物车失败";t.showZanTopTips(e),o.hideLoading(),t.setData({showDialog:!1,action:0})})}else if(console.log("--------"),!this.data.showDialog){var r=Object.assign({},this.data.formData);r.num.hasSelect=!0,this.setData({showDialog:!0,action:2,formData:r})}},handleFenXiaoBuyNow:function(){var t="";this.data.goods.isDist&&1==this.data.goods.isDist&&this.data.from&&"distribution"==this.data.from?t=this.data.goods.shareCode:this.data.goods.isDist&&1==this.data.goods.isDist&&this.data.from&&"share"==this.data.from&&(t=this.data.shareCode),this.normalBuyNow({shareCode:t},4)},handleOpenTuan:function(){this.normalBuyNow({groupId:this.data.goods.groupEvent&&this.data.goods.groupEvent.groupId||"0"},6)},handleMiaoShaBuyNow:function(){this.normalBuyNow({},6)},handleMakeSure:function(){console.log("+++++");var t=this.data.action,a=this.data.formData.selectSku,e=this.data.goods.specs.length==Object.getOwnPropertyNames(a).length;if([2,3,4,5,6].indexOf(t)>-1&&!e)return void this.showZanToast("请选择规格");switch(t){case 2:this.handleAddCart();break;case 3:this.handleBuyNow();break;case 4:this.handleFenXiaoBuyNow();break;case 5:this.handleBuyNow();break;case 6:this.handleOpenTuan()}},handleSpecClick:function(t){var a=this,e=this.data.formData.num.quantity,i=t.currentTarget.dataset,o=t.currentTarget.dataset,s=o.index,n=o.currentindex;if(!i.disabled){var r=Object.assign({},this.data.formData.selectSku);r[i.pid]={id:i.id,name:i.name};var d="";this.data.goods.specs.forEach(function(t){d+=r[t.id]?r[t.id].name:"",d+=r[t.id]?",":""});var h=Object.assign({},this.data.goods);h.specs=h.specs.map(function(t){return r[t.id]&&(t.specsList=t.specsList.map(function(e){return r[t.id].id==e.id?(e.selected=!0,0==s&&(console.log(e,"999999999999999999999999"),h.showCoverImg=a.data.goods.item[0].values[n].skuLogo,h.selectSkuText=a.data.goods.item[0].values[n].itemValue),console.log(h.selectSkuText,"llllllllllllllllllllllllll")):e.selected=!1,e})),t});var l=Object.getOwnPropertyNames(r),u=void 0;h.skus.forEach(function(t){t.properties.filter(function(t){return l.filter(function(a){return t[a]==r[a].id}).length>0}).length==t.properties.length&&(u=t)}),console.log(u),h.specs.length-l.length==1?h.specs.map(function(t){return r[t.id]||(t.specsList=t.specsList.map(function(a){var e=Object.assign({},r);e[t.id]={id:a.id,name:a.name};var i=Object.getOwnPropertyNames(e),o=void 0;return h.skus.forEach(function(t){var s=t.properties.filter(function(t){return i.filter(function(a){return t[a]==e[a].id}).length>0}).length==t.properties.length;console.log("result====>",s),s&&(o=t,a.disabled=o&&0==o.quantity)}),a})),t}):h.specs.length-l.length==0&&h.specs.map(function(t){return t.specsList=t.specsList.map(function(a){var e=Object.assign({},r);e[t.id]={id:a.id,name:a.name};var i=Object.getOwnPropertyNames(e),o=void 0;return h.skus.forEach(function(t){t.properties.filter(function(t){return i.filter(function(a){return t[a]==e[a].id}).length>0}).length==t.properties.length&&(o=t,a.disabled=o&&0==o.quantity)}),a}),t}),this.setData({goods:h,formData:{num:{quantity:u&&this.data.formData.num.quantity>u.quantity?u.quantity:this.data.formData.num.quantity>0?this.data.formData.num.quantity:1,min:this.data.formData.num.min,max:u?u.quantity:this.data.formData.num.max,minBuy:u&&u.minBuy>0?u.minBuy:1,maxBuy:u&&u.maxBuy>0&&u.maxBuy<u.quantity?u.maxBuy:u?u.quantity:99999,hasSelect:this.data.formData.num.hasSelect},sku:u?{price:u.price,groupPrice:u.groupPrice,quantity:u.quantity,id:u.id,maxBuy:u.maxBuy,minBuy:u.minBuy}:this.data.formData.sku,selectSku:r,selectSkuText:d}}),console.log("selectSku----------\x3e",r),this.data.formData.sku.maxBuy>0&&(e==this.data.formData.sku.maxBuy?this.setData({limit:!0}):this.setData(y({limit:!1},"formData.num.quantity",1)))}},collection:function(){var t=this,a=this.data.goods.id;D.default.shareGoodsList({goodsId:a,type:1==this.data.goods.isShare?"cancel":"share"}).then(function(a){if(a&&a.data.result){var e=t.data.goods,i=1==e.isShare?0:1;t.setData({goods:Object.assign({},e,{isShare:i})})}}).catch(function(t){return console.log(t)})},collectionGoods:function(){var t=this,a=this.data.goods,e=a.id;a.hasCollection?D.default.unfollowGoods({goodsId:e}).then(function(){t.setData(y({},"goods.hasCollection",0))}):D.default.followGoods({goodsId:e}).then(function(){t.setData(y({},"goods.hasCollection",1))})},wxParseTagATap:function(t){var a=t.currentTarget.dataset.src;o.redirectTo({url:"/pages/webview/index?url="+I.default.stringify(a)})},showSharePopup:function(){console.log("showSharePopup++++++++++++",this.data.showShareDialog),this.setData({showShareDialog:!0})},hideSharePopup:function(){console.log("showSharePopup++++++++++++",this.data.showShareDialog),this.setData({showShareDialog:!1})},saveImg:function(){var t=this,a=this.data.id,e={goodsId:a};e.from=this.data.from,"distribution"==this.data.from&&(e.from="share"),o.showLoading({title:"请稍候"}),D.default.queryPhotoShare(e).then(function(a){console.log(a,"res"),o.hideLoading(),t.hideSharePopup();var e=a.data.shareUrl;e?t.getWxsetting(e):t.showZanTopTips("生成海报图失败")}).catch(function(){o.hideLoading(),t.hideSharePopup(),t.showZanTopTips("生成海报图失败")})},getWxsetting:function(t){o.previewImage({current:t,urls:[t]})},onTapPic:function(){this.setData({showDialog1:!this.data.showDialog1})},goGoodDetail:function(t){console.log("e---》",t);var a=t.currentTarget.dataset.id;o.navigateTo({url:"/pages/products/show/show?id="+a+"&isFromShop=true"})},dealPrice:function(t){if(t<1e5)return t;if(t<1e8){var a=t/1e4+"",e=a.indexOf(".");return-1==e?a+".00万":(console.log("====",(a+"").length,e),(a+"").length>e+2?a.substring(0,e+3)+"万":a.substring(0,e+2)+"0万")}var a=t/1e8+"",e=a.indexOf(".");return-1==e?a+".00亿":(a+"").length>e+2?a.substring(0,e+3)+"亿":a.substring(0,e+2)+"0亿"}}))}),window.__swanRoute="pages/products/show/show",window.usingComponents=["component/phonemodal/phonemodal","dist/minui/wxc-popup/dist/index","dist/minui/wxc-toast/dist/index","component/phone_action/index"],require("pages/products/show/show");