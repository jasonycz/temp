function _classCallCheck(e,n){if(!(e instanceof n))throw new TypeError("Cannot call a class as a function")}function _possibleConstructorReturn(e,n){if(!e)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!n||"object"!=typeof n&&"function"!=typeof n?e:n}function _inherits(e,n){if("function"!=typeof n&&null!==n)throw new TypeError("Super expression must either be null or a function, not "+typeof n);e.prototype=Object.create(n&&n.prototype,{constructor:{value:e,enumerable:!1,writable:!0,configurable:!0}}),n&&(Object.setPrototypeOf?Object.setPrototypeOf(e,n):e.__proto__=n)}!function(global,san){function processTemplateModule(filterTemplateArrs,filterModule){eval(filterModule);var modules={},templateFiltersObj={};return filterTemplateArrs&&filterTemplateArrs.forEach(function(element){var filterName=element.filterName,func=element.func,module=element.module;modules[module]=eval(module),templateFiltersObj[filterName]=function(){var e;return(e=modules[module])[func].apply(e,arguments)}}),templateFiltersObj}global.errorMsg=[];var filterTemplateArrs0=JSON.parse("[]"),templateFiltersObj0={};filterTemplateArrs0.length&&(templateFiltersObj0=processTemplateModule(filterTemplateArrs0,""));var template7e2455=san.defineComponent({components:PageComponent.components,template:'<swan-template data-origin-name="zan-toast"><text class="zan-toast" s-if=" zanToast.show " style="width:70%" on-bindtap="eventHappen(\'tap\', $event, \'clearZanToast\', \'\', \'bind\')">{{ zanToast.title }}</text></swan-template>',filters:Object.assign({},templateFiltersObj0),inited:function(){var e=this,n=function(n){for(var t in n)e.data.set(t,n[t])};n(this.data.get("data")),this.watch("data",n)},eventHappen:function(){var e;(e=this.owner).eventHappen.apply(e,arguments)}});PageComponent.components["template-7e2455"]=template7e2455;var filterTemplateArrs1=JSON.parse("[]"),templateFiltersObj1={};filterTemplateArrs1.length&&(templateFiltersObj1=processTemplateModule(filterTemplateArrs1,""));var template7e2456=san.defineComponent({components:PageComponent.components,template:'<swan-template data-origin-name="mineUser"><view on-bindtap="eventHappen(\'tap\', $event, \'navToProfile\', \'\', \'bind\')" style="display: flex;text-align: center;padding: 7.2vw 0vw;background-color:{{ color || \'#FF9300\' }}"><view style="padding: 0 4vw;"><image s-if="avatarUrl" src="{{avatarUrl}}" class="user-img"></image><view s-else="" class="user-img"><open-data type="userAvatarUrl" /></view></view><view style="font-size: 4vw;color:#fff;text-align: left;display: flex;flex:1; justify-content: center; flex-direction:column;"><view style="margin-bottom: 3.466666666666667vw;" class="zan-ellipsis"><template s-if="nickName">{{nickName}}</template><open-data s-else="" type="userNickName" /></view><view data-url="{{mobile ? \'/pages/login/verify_cell?mobile=\'+mobile+\'&fomatMobile=\'+mobile : \'/pages/login/login\'}}"><view s-if="mobile" style="font-size: 2.933333333333333vw;">{{mobile}}>></view><view s-else="" style="font-size: 2.933333333333333vw;">未绑定手机</view></view></view><view style="display:flex;width:5.333333333333333vw;align-items:center;justify-content:center;color:#fff;padding-right:3.2vw">></view></view></swan-template>',filters:Object.assign({},templateFiltersObj1),inited:function(){var e=this,n=function(n){for(var t in n)e.data.set(t,n[t])};n(this.data.get("data")),this.watch("data",n)},eventHappen:function(){var e;(e=this.owner).eventHappen.apply(e,arguments)}});PageComponent.components["template-7e2456"]=template7e2456;var filterTemplateArrs2=JSON.parse("[]"),templateFiltersObj2={};filterTemplateArrs2.length&&(templateFiltersObj2=processTemplateModule(filterTemplateArrs2,""));var template7e2457=san.defineComponent({components:PageComponent.components,template:'<swan-template data-origin-name="mineOrders"><view s-if="myOrderIsShow"><view class="mine-order-all"><view data-url="/pages/orders/list/list?type=0" on-bindtap="eventHappen(\'tap\', $event, \'handleMakeUrl\', \'\', \'bind\')" class="zan-cell zan-cell--access " style="background-color:#fff;margin-top: 2.6666666666666665vw;border-top: 0.13333333333333333vw solid #e5e5e5;border-bottom: 0.13333333333333333vw solid #e5e5e5; "><view class="zan-cell__bd" style="display:flex;align-items: center ;"><view class="menu-name">我的订单</view></view><view class="zan-cell__ft">查看全部</view></view></view><view style="display:flex;padding:2.6666666666666665vw 0; border-bottom: 0.13333333333333333vw solid #e5e5e5;background-color:#fff; "><view class="order-item" data-url="/pages/orders/list/list?type=1" on-bindtap="eventHappen(\'tap\', $event, \'handleMakeUrl\', \'\', \'bind\')"><view class="zan-badge"><view class="zan-badge__count" s-if=" order[10] ">{{ order[10] }}</view><image src="http://imgcdnali.wfenxiao.com.cn/xcx/c/order/ic_order_no_delivery.png" class="order-item-img"></image></view><view >待付款</view></view><view class="order-item" data-url="/pages/orders/list/list?type=2" on-bindtap="eventHappen(\'tap\', $event, \'handleMakeUrl\', \'\', \'bind\')"><view class="zan-badge"><view class="zan-badge__count" s-if=" order[30] ">{{ order[30] }}</view><image src="http://imgcdnali.wfenxiao.com.cn/xcx/c/order/ic_order_nopay.png" class="order-item-img"></image></view><view >待发货</view></view><view class="order-item" data-url="/pages/orders/list/list?type=3" on-bindtap="eventHappen(\'tap\', $event, \'handleMakeUrl\', \'\', \'bind\')"><view class="zan-badge"><view class="zan-badge__count" s-if=" order[61] ">{{ order[61] }}</view><image src="http://imgcdnali.wfenxiao.com.cn/xcx/c/order/ic_order_delivery.png" class="order-item-img"></image></view><view >待收货</view></view><view class="order-item" data-url="/pages/orders/list/list?type=4" on-bindtap="eventHappen(\'tap\', $event, \'handleMakeUrl\', \'\', \'bind\')"><view class="zan-badge"><view class="zan-badge__count" s-if=" order[70] ">{{ order[70] }}</view><image src="http://imgcdnali.wfenxiao.com.cn/xcx/c/order/ic_order_done.png" class="order-item-img"></image></view><view >已完成</view></view><view s-if="false" class="order-item" data-url="/pages/returngoods/list/index" on-bindtap="eventHappen(\'tap\', $event, \'handleMakeUrl\', \'\', \'bind\')"><view class="zan-badge"><image src="http://imgcdnali.wfenxiao.com.cn/xcx/c/1120/ic_order_refund.png" class="order-item-img"></image></view><view >退款</view></view></view></view></swan-template>',filters:Object.assign({},templateFiltersObj2),inited:function(){var e=this,n=function(n){for(var t in n)e.data.set(t,n[t])};n(this.data.get("data")),this.watch("data",n)},eventHappen:function(){var e;(e=this.owner).eventHappen.apply(e,arguments)}});PageComponent.components["template-7e2457"]=template7e2457;var filterTemplateArrs3=JSON.parse("[]"),templateFiltersObj3={};filterTemplateArrs3.length&&(templateFiltersObj3=processTemplateModule(filterTemplateArrs3,""));var template7e2458=san.defineComponent({components:PageComponent.components,template:'<swan-template data-origin-name="cell"><view style="position:relative;display:flex;flex-direction:column;align-items:center;justify-content:center;width:25vw;height:25vw;"><button s-if="item.type==\'contact\'" plain="" open-type="contact" style="border:0px;left:0;top:0;width:100%;height:100%;position:absolute;"></button><image style="width:6.4vw;height:6.4vw;" mode="scaleToFill" src="{{item.icon}}"></image><view style="color:#000;font-size:3.2vw;margin-top:2.1333333333333333vw;">{{item.title}}</view></view></swan-template>',filters:Object.assign({},templateFiltersObj3),inited:function(){var e=this,n=function(n){for(var t in n)e.data.set(t,n[t])};n(this.data.get("data")),this.watch("data",n)},eventHappen:function(){var e;(e=this.owner).eventHappen.apply(e,arguments)}});PageComponent.components["template-7e2458"]=template7e2458;var filterTemplateArrs4=JSON.parse("[]"),templateFiltersObj4={};filterTemplateArrs4.length&&(templateFiltersObj4=processTemplateModule(filterTemplateArrs4,""));var template7e2459=san.defineComponent({components:PageComponent.components,template:'<swan-template data-origin-name="mineAddress"><view class="zan-panel  zan-panel__with-top-border " s-if=" addressIsShow "><view data-url="/pages/address/list/list" on-bindtap="eventHappen(\'tap\', $event, \'handleMakeUrl\', \'\', \'bind\')" class="zan-cell zan-cell--access"><view class="zan-cell__bd" style="display:flex; align-items: center ;"><view class="menu-name">地址管理</view></view><view class="zan-cell__ft" style="width:2.6666666666666665vw;height:2.6666666666666665vw"></view></view></view><view style="display:flex;flex-flow: row wrap;" class="zan-panel  zan-panel__with-top-border "><template s-key="index" s-for="item, index in menus"><view s-if="item.bindtap" on-bindtap="eventHappen(\'tap\', $event, item.bindtap, \'\', \'bind\')"><template-7e2458 data="{{{item}}}" data-origin-name="cell" /></view><view s-elif="item.url" data-url="{{item.url}}" on-bindtap="eventHappen(\'tap\', $event, \'handleMakeUrl\', \'\', \'bind\')"><template-7e2458 data="{{{item}}}" data-origin-name="cell" /></view><view s-else=""><template-7e2458 data="{{{item}}}" data-origin-name="cell" /></view></template></view></swan-template>',filters:Object.assign({},templateFiltersObj4),inited:function(){var e=this,n=function(n){for(var t in n)e.data.set(t,n[t])};n(this.data.get("data")),this.watch("data",n)},eventHappen:function(){var e;(e=this.owner).eventHappen.apply(e,arguments)}});PageComponent.components["template-7e2459"]=template7e2459;var templateComponents=Object.assign({},{"template-7e2455":template7e2455,"template-7e2456":template7e2456,"template-7e2457":template7e2457,"template-7e2458":template7e2458,"template-7e2459":template7e2459}),param={},filterArr=JSON.parse("[]");try{filterArr&&filterArr.forEach(function(item){param[item.module]=eval(item.module)});var pageContent='<template s-if="userInfo"><view class="container"><template-7e2456 data="{{{...userInfo, color:color}}}" data-origin-name="mineUser" /><template-7e2457 data="{{{myOrderIsShow:mine.myOrderIsShow, order}}}" data-origin-name="mineOrders" /><template-7e2459 data="{{{addressIsShow:mine.addressIsShow,menus}}}" data-origin-name="mineAddress" /><template-7e2455 data="{{{ zanToast }}}" data-origin-name="zan-toast"></template-7e2455><phonemodal show-modal="{{showModal}}" url="{{url}}" on-bindchangeuser="eventHappen(\'changeuser\', $event, \'changeuser\', \'\', \'bind\')"></phonemodal></view><phone-action class="phoneaction" list="{{phoneList}}" /></template>',renderPage=function renderPage(filters,modules){var componentFactory=global.componentFactory,componentFragments=componentFactory.getAllComponents(),user_self_compo_934433=function(global,componentFragments){var templateComponents={},usingComponentMap=JSON.parse('{"phonebutton":"component/phonebutton/index","toast":"component/toast/index"}'),filterCustomArr=JSON.parse("[]"),modules={},filtersObj={};filterCustomArr&&filterCustomArr.forEach(function(element){modules[element.module]=eval(element.module);var func=element.func,module=element.module;filtersObj[element.filterName]=function(){var e;return(e=modules[module])[func].apply(e,arguments)}});var components=Object.assign({},componentFragments,templateComponents);return global.componentFactory.componentDefine("component/phonemodal/phonemodal.swan",Object.assign({},componentFactory.getProtos("super-custom-component"),{template:"<swan-phonemodal><view class=\"modal-wrap\" on-bindtap=\"eventHappen('tap', $event, 'onCancel', '', 'bind')\" style=\"display: {{ show ? 'block': 'none' }}\"><view class=\"modal-container\"><view class=\"modal-header\">{{title}}</view><view class=\"modal-body\">更多服务登陆后即可体验</view><view class=\"modal-footer\"><phonebutton class=\"modal-footer\" loginUrl=\"{{loginUrl}}\" on-bindchangeuser=\"eventHappen('changeuser', $event, 'changeuser', '', 'bind')\" on-bindshowtoast=\"eventHappen('showtoast', $event, 'showtoast', '', 'bind')\"></phonebutton></view></view></view><toast show=\"{{toast.show}}\" title=\"{{toast.title}}\" /></swan-phonemodal>",componentPath:"component/phonemodal/phonemodal.swan",componentName:"phonemodal",customComponentCss:"/* component/phonemodal/phonemodal.css */\n\n\n.phonemodal__modal-wrap{\n  width: 100%;\n  height: 100%;\n  background-color: rgba(0,0,0,0.3);\n  position: fixed;\n  z-index: 1000;\n  top: 0;\n  right: 0;\n  left: 0;\n  bottom: 0;\n}\n\n\n.phonemodal__modal-container{\n  position: absolute;\n  height: 46.13333333333333vw;\n  width: 73.6vw;\n  z-index: 9999;\n  border-radius: 2.1333333333333333vw;\n  top: 0;\n  right: 0;\n  left: 0;\n  bottom: 0;\n  margin: auto;\n  background-color: #fff;\n}\n\n\n.phonemodal__modal-header{\n  text-align: center;\n  font-size: 4.266666666666667vw;\n  margin: 4.266666666666667vw 0 6.4vw 0;\n}\n\n\n.phonemodal__modal-body{\n  font-size: 3.7333333333333334vw;\n  color: #000;\n  opacity: 0.43;\n  margin-bottom: 9.066666666666666vw;\n  text-align: center;\n}\n\n\n.phonemodal__modal-footer{\n  display: flex;\n  justify-content: center;\n}\n"}),{classProperties:{components:components,filters:Object.assign({},filtersObj)}}),Promise.resolve().then(function(){for(var e in usingComponentMap)components[e]=customAbsolutePathMap[usingComponentMap[e]]}),global.componentFactory.getComponents("component/phonemodal/phonemodal.swan")}(global,componentFragments),user_self_compo_934434=function(global,componentFragments){var templateComponents={},usingComponentMap=JSON.parse('{"wxc-popup":"dist/minui/wxc-popup/dist/index"}'),filterCustomArr=JSON.parse("[]"),modules={},filtersObj={};filterCustomArr&&filterCustomArr.forEach(function(element){modules[element.module]=eval(element.module);var func=element.func,module=element.module;filtersObj[element.filterName]=function(){var e;return(e=modules[module])[func].apply(e,arguments)}});var components=Object.assign({},componentFragments,templateComponents);return global.componentFactory.componentDefine("component/phone_action/index.swan",Object.assign({},componentFactory.getProtos("super-custom-component"),{template:"<swan-phone-action><view class=\"zan-dialog {{ showDialog ? 'zan-dialog--show' : '' }}\" style=\"width: 100%;\" on-bindtap=\"eventHappen('tap', $event, 'hideDialog', '', 'bind')\" s-if=\"showDialog\"><view on-bindtap=\"eventHappen('tap', $event, 'hide', '', 'bind')\" class=\"zan-dialog__mask\"></view><view class=\"zan-dialog__container\"><view class=\"phone-action__container\"><view s-for=\"item, index in list\" class=\"phone-action__item\" on-bindtap=\"eventHappen('tap', $event, 'call', '', 'bind')\" data-phone=\"{{item}}\">{{item}}</view></view></view></view></swan-phone-action>",componentPath:"component/phone_action/index.swan",componentName:"phone-action",customComponentCss:".phone-action__zan-dialog__mask {\n  position: fixed;\n  top: 0;\n  left: 0;\n  right: 0;\n  bottom: 0;\n  z-index: 10;\n  background: rgba(0, 0, 0, 0.7);\n  display: none;\n}\n.phone-action__zan-dialog__container {\n  position: fixed;\n  bottom: 0;\n  width: 100vw;\n  background: white;\n  transform: translateY(150%);\n  transition: all 0.4s ease;\n  z-index: 11;\n}\n.phone-action__zan-dialog--show .phone-action__zan-dialog__container {\n  transform: translateY(0);\n}\n.phone-action__zan-dialog--show .phone-action__zan-dialog__mask {\n  display: block;\n}\n.phone-action__container{\n    background: white;\n}\n.phone-action__item{\n    padding: 2.6666666666666665vw;\n    display: flex;\n    justify-content: center;\n}\n.phone-action__item:not(:last-child){\n    border-bottom: 0.13333333333333333vw solid #b2b2b2;\n}"}),{classProperties:{components:components,filters:Object.assign({},filtersObj)}}),Promise.resolve().then(function(){for(var e in usingComponentMap)components[e]=customAbsolutePathMap[usingComponentMap[e]]}),global.componentFactory.getComponents("component/phone_action/index.swan")}(global,componentFragments),user_self_compo_934435=function(global,componentFragments){var templateComponents={},usingComponentMap=JSON.parse('{"toast":"component/toast/index"}'),filterCustomArr=JSON.parse("[]"),modules={},filtersObj={};filterCustomArr&&filterCustomArr.forEach(function(element){modules[element.module]=eval(element.module);var func=element.func,module=element.module;filtersObj[element.filterName]=function(){var e;return(e=modules[module])[func].apply(e,arguments)}});var components=Object.assign({},componentFragments,templateComponents);return global.componentFactory.componentDefine("component/phonebutton/index.swan",Object.assign({},componentFactory.getProtos("super-custom-component"),{template:"<swan-phonebutton><view class=\"button-wrap\"><button on-bindtap=\"eventHappen('tap', $event, 'gotoLogin', '', 'catch')\" class=\"{{size == 'big'?'modal-btn-big':'modal-btn-primary'}}\">极速登录</button></view><toast show=\"{{toast.show}}\" title=\"{{toast.title}}\" /></swan-phonebutton>",componentPath:"component/phonebutton/index.swan",componentName:"phonebutton",customComponentCss:"/* component/phonebutton/index.css */\n.phonebutton__button-wrap{\n  display: flex;\n  justify-content: center;\n}\n.phonebutton__modal-btn-primary{\n  width: 29.333333333333332vw;\n  height: 8.533333333333333vw;\n  background: #f32828;\n  border-radius: 1.0666666666666667vw;\n  line-height: 8.533333333333333vw;\n  color: #fff;\n  font-size: 3.7333333333333334vw;\n}\n.phonebutton__modal-btn-big{\n  width:48vw;\n  height:11.733333333333333vw;\n  line-height:11.733333333333333vw;\n  background-color:#f22222;\n  border-color:#f22222;\n  font-size:3.7333333333333334vw;\n  color:#fff;\n\n}\n"}),{classProperties:{components:components,filters:Object.assign({},filtersObj)}}),Promise.resolve().then(function(){for(var e in usingComponentMap)components[e]=customAbsolutePathMap[usingComponentMap[e]]}),global.componentFactory.getComponents("component/phonebutton/index.swan")}(global,componentFragments),user_self_compo_934436=function(global,componentFragments){var templateComponents={},usingComponentMap=JSON.parse("{}"),filterCustomArr=JSON.parse("[]"),modules={},filtersObj={};filterCustomArr&&filterCustomArr.forEach(function(element){modules[element.module]=eval(element.module);var func=element.func,module=element.module;filtersObj[element.filterName]=function(){var e;return(e=modules[module])[func].apply(e,arguments)}});var components=Object.assign({},componentFragments,templateComponents);return global.componentFactory.componentDefine("component/toast/index.swan",Object.assign({},componentFactory.getProtos("super-custom-component"),{template:"<swan-toast><view class=\"toast\" s-if=\" show \" style=\"width:70%\" on-bindtap=\"eventHappen('tap', $event, 'clearToast', '', 'bind')\">{{ title }}</view></swan-toast>",componentPath:"component/toast/index.swan",componentName:"toast",customComponentCss:".toast__toast {\n  position: fixed;\n  top: 35%;\n  left: 15%;\n  transform: translateZ(0) translateY(-100%);\n  background: rgba(0, 0, 0, 0.7);\n  color: #fff;\n  font-size: 14px;\n  line-height: 1.5em;\n  margin: 0 auto;\n  box-sizing: border-box;\n  padding: 10px;\n  text-align: center;\n  border-radius: 4px;\n  z-index: 900;\n}\n"}),{classProperties:{components:components,filters:Object.assign({},filtersObj)}}),Promise.resolve().then(function(){for(var e in usingComponentMap)components[e]=customAbsolutePathMap[usingComponentMap[e]]}),global.componentFactory.getComponents("component/toast/index.swan")}(global,componentFragments),user_self_compo_934437=function(global,componentFragments){var templateComponents={},usingComponentMap=JSON.parse('{"wxc-mask":"dist/minui/wxc-mask/dist/index"}'),filterCustomArr=JSON.parse("[]"),modules={},filtersObj={};filterCustomArr&&filterCustomArr.forEach(function(element){modules[element.module]=eval(element.module);var func=element.func,module=element.module;filtersObj[element.filterName]=function(){var e;return(e=modules[module])[func].apply(e,arguments)}});var components=Object.assign({},componentFragments,templateComponents);return global.componentFactory.componentDefine("dist/minui/wxc-popup/dist/index.swan",Object.assign({},componentFactory.getProtos("super-custom-component"),{template:"<swan-wxc-popup><wxc-mask status=\"{{maskStatus||'hide'}}\" _pid=\"{{_pid}}\" locked=\"{{locked}}\" opacity=\"0.6\"></wxc-mask><view class=\"popup popup__status--{{status}} popup__content-align--{{align}} popup__status--{{status}}--animotion-{{animationMode}}\" on-bindtap=\"eventHappen('tap', $event, 'handleTap', '', 'bind')\" on-bindtouchmove=\"eventHappen('touchmove', $event, 'catchmove', '', 'catch')\"><slot ></slot></view></swan-wxc-popup>",componentPath:"dist/minui/wxc-popup/dist/index.swan",componentName:"wxc-popup",customComponentCss:".wxc-popup__popup {\n    position: fixed;\n    z-index: 200000;\n    width: 100%;\n    height: 100%;\n    top: 50%;\n    left: 50%;\n    transform: translate(-50%, -50%);\n    display: flex;\n    overflow: hidden\n}\n.wxc-popup__popup__status {}\n.wxc-popup__popup__status--hide {\n    display: none\n}\n.wxc-popup__popup__status--show {\n    display: flex;\n    transition: all 300ms linear\n}\n.wxc-popup__popup__status--fadeOut {\n    display: flex;\n    transition: all 300ms linear\n}\n.wxc-popup__popup__status--fadeOut--animotion-none {}\n.wxc-popup__popup__status--fadeOut--animotion-center {\n    width: 0;\n    height: 0\n}\n.wxc-popup__popup__status--fadeOut--animotion-top {\n    top: -100%\n}\n.wxc-popup__popup__status--fadeOut--animotion-bottom {\n    transform: translate(-50%, 50%)\n}\n.wxc-popup__popup__status--fadeOut--animotion-left {\n    left: -100%\n}\n.wxc-popup__popup__status--fadeOut--animotion-right {\n    transform: translate(50%, -50%)\n}\n.wxc-popup__popup__status--fadeIn {\n    display: flex\n}\n.wxc-popup__popup__status--fadeIn--animotion-none {}\n.wxc-popup__popup__status--fadeIn--animotion-center {\n    width: 0;\n    height: 0\n}\n.wxc-popup__popup__status--fadeIn--animotion-top {\n    top: -100%\n}\n.wxc-popup__popup__status--fadeIn--animotion-bottom {\n    transform: translate(-50%, 50%)\n}\n.wxc-popup__popup__status--fadeIn--animotion-left {\n    left: -100%\n}\n.wxc-popup__popup__status--fadeIn--animotion-right {\n    transform: translate(50%, -50%)\n}\n.wxc-popup__popup__content-animotion {}\n.wxc-popup__popup__content-align {}\n.wxc-popup__popup__content-align--center {\n    align-items: center;\n    justify-content: center\n}\n.wxc-popup__popup__content-align--top {\n    align-items: flex-start;\n    justify-content: center\n}\n.wxc-popup__popup__content-align--bottom {\n    align-items: flex-end;\n    justify-content: center\n}\n.wxc-popup__popup__content-align--left {\n    align-items: center;\n    justify-content: flex-start\n}\n.wxc-popup__popup__content-align--right {\n    align-items: center;\n    justify-content: flex-end\n}"}),{classProperties:{components:components,filters:Object.assign({},filtersObj)}}),Promise.resolve().then(function(){for(var e in usingComponentMap)components[e]=customAbsolutePathMap[usingComponentMap[e]]}),global.componentFactory.getComponents("dist/minui/wxc-popup/dist/index.swan")}(global,componentFragments),user_self_compo_934438=function(global,componentFragments){var templateComponents={},usingComponentMap=JSON.parse("{}"),filterCustomArr=JSON.parse("[]"),modules={},filtersObj={};filterCustomArr&&filterCustomArr.forEach(function(element){modules[element.module]=eval(element.module);var func=element.func,module=element.module;filtersObj[element.filterName]=function(){var e;return(e=modules[module])[func].apply(e,arguments)}});var components=Object.assign({},componentFragments,templateComponents);return global.componentFactory.componentDefine("dist/minui/wxc-mask/dist/index.swan",Object.assign({},componentFactory.getProtos("super-custom-component"),{template:'<swan-wxc-mask><view class="mask mask__status--{{status}}"><view class="mask__bg" style="background:rgba({{backgroundColor}},{{opacity}});"></view><view class="mask__content mask__content--{{cAlign}}" style="{{__positionStyle}}"><slot ></slot></view></view></swan-wxc-mask>',componentPath:"dist/minui/wxc-mask/dist/index.swan",componentName:"wxc-mask",customComponentCss:".wxc-mask__mask {\n    width: 100%;\n    height: 100%;\n    position: fixed;\n    top: 0;\n    left: 0;\n    z-index: 100000\n}\n.wxc-mask__mask__bg {\n    width: 100%;\n    height: 100%\n}\n.wxc-mask__mask__content {\n    display: inline-block;\n    position: absolute;\n    z-index: 1\n}\n.wxc-mask__mask__status {}\n.wxc-mask__mask__status--show {}\n.wxc-mask__mask__status--hide {\n    opacity: 0;\n    display: none;\n    z-index: -10\n}"}),{classProperties:{components:components,filters:Object.assign({},filtersObj)}}),Promise.resolve().then(function(){for(var e in usingComponentMap)components[e]=customAbsolutePathMap[usingComponentMap[e]]}),global.componentFactory.getComponents("dist/minui/wxc-mask/dist/index.swan")}(global,componentFragments),customAbsolutePathMap={"component/phonemodal/phonemodal":user_self_compo_934433,"component/phone_action/index":user_self_compo_934434,"component/phonebutton/index":user_self_compo_934435,"component/toast/index":user_self_compo_934436,"dist/minui/wxc-popup/dist/index":user_self_compo_934437,"dist/minui/wxc-mask/dist/index":user_self_compo_934438},pageUsingComponentMap=JSON.parse('{"phonemodal":"component/phonemodal/phonemodal","phone-action":"component/phone_action/index"}'),customComponents=Object.keys(pageUsingComponentMap).reduce(function(e,n){return e[n]=customAbsolutePathMap[pageUsingComponentMap[n]],e},{});global.pageRender(pageContent,templateComponents,customComponents,filters,modules)},oldPatch=function(e){Object.assign(e.components,{"template-7e2455":template7e2455,"template-7e2456":template7e2456,"template-7e2457":template7e2457,"template-7e2458":template7e2458,"template-7e2459":template7e2459});var n=function(n){function t(n){_classCallCheck(this,t);var o=_possibleConstructorReturn(this,(t.__proto__||Object.getPrototypeOf(t)).call(this,n));return o.components=e.components,o}return _inherits(t,n),t}(e);n.template='<swan-wrapper tabindex="-1">'+pageContent+"</swan-wrapper>";var t=new n;t.slaveLoaded(),t.communicator.onMessage("initData",function(e){t.setInitData(e),t.attach(document.body)},{listenPreviousEvent:!0})};global.pageRender?renderPage(filterArr,param):oldPatch(window.PageComponent)}catch(e){throw global.errorMsg.execError=e,e}}(window,window.san);