window.define("pages/page/search/search",function(t,e,a,r,o,n,s,i,d,c,u,l,h,g,f,p){"use strict";function v(t,e){var a={};for(var r in t)e.indexOf(r)>=0||Object.prototype.hasOwnProperty.call(t,r)&&(a[r]=t[r]);return a}var w=function(){function t(t,e){var a=[],r=!0,o=!1,n=void 0;try{for(var s,i=t[Symbol.iterator]();!(r=(s=i.next()).done)&&(a.push(s.value),!e||a.length!==e);r=!0);}catch(t){o=!0,n=t}finally{try{!r&&i.return&&i.return()}finally{if(o)throw n}}return a}return function(e,a){if(Array.isArray(e))return e;if(Symbol.iterator in Object(e))return t(e,a);throw new TypeError("Invalid attempt to destructure non-iterable instance")}}(),y=t("../../../utils/es6-promise.min.js"),b=function(t){return t&&t.__esModule?t:{default:t}}(y),m=t("../../../network/api");Page({data:{tabs:[{title:"全部",current:!0},{title:"店铺"}],keyword:"",currentTab:0,products:{},stores:{}},onLoad:function(){},onTabSwitch:function(t){var e=t.currentTarget.dataset.index,a=this.data.tabs;console.log(e,a),a.forEach(function(t,a){t.current=e==a}),this.setData({tabs:a,currentTab:e})},onSearch:function(t){var e=this;console.log(t,"eeee"),o.showLoading({mask:!0,title:"请稍后"}),this.setData({keyword:t.value});var a={keyword:t.value,sort:"sales:descend"};b.default.all([(0,m.queryGoodsProperties)(a),(0,m.queryIndexGoods)(a),(0,m.queryIndexStores)(a)]).then(function(t){o.hideLoading();var a=w(t,3),r=a[0],n=a[1],s=a[2];console.log(t),e.setData({serverTime:n.serverTime,filters:r.data,products:n.data,stores:s.data})}).catch(function(t){o.hideLoading(),console.log(t)})},handleQueryProducts:function(t){var e=this;o.showLoading({mask:!0,title:"请稍后"});var a=t.sort,r=t.query,n=this.data.keyword,s=Object.assign({keyword:n,sort:a},r),i=this.data.products;(0,m.queryIndexGoods)(s).then(function(t){o.hideLoading();var a=t.data,r=v(t,["data"]);console.log(r,"+++++++++++++++,--------------"),a.pageNo>1?e.setData({products:Object.assign({},a,{data:i.data.concat(a.data)})}):e.setData({products:a,filters:a.data.length>0?e.data.filters:{}})}).catch(function(t){o.hideLoading(),console.log(t)})},handleQueryStores:function(t){var e=this;o.showLoading({mask:!0,title:"请稍后"});var a=t.sort,r=this.data.keyword,n={keyword:r,sort:a},s=this.data.stores;(0,m.queryIndexStores)(n).then(function(t){o.hideLoading();var a=t.data;a.pageNo>1?e.setData({stores:Object.assign({},a,{data:s.data.concat(a)})}):e.setData({stores:a})}).catch(function(t){o.hideLoading(),console.log(t)})},makeproject:function(t){var e=t.currentTarget.dataset.project;o.redirectTo({url:"/pages/services/service_detail/index?shopId="+e.storeId+"&serviceId="+e.id})}})}),window.__swanRoute="pages/page/search/search",window.usingComponents=["dist/minui/wxc-search/dist/index","dist/minui/wxc-tab/dist/index","dist/minui/wxc-tab/dist/panel","pages/page/search/product-tab/index","pages/page/search/store-tab/index"],require("pages/page/search/search");