function _classCallCheck(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}function _possibleConstructorReturn(e,t){if(!e)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!t||"object"!=typeof t&&"function"!=typeof t?e:t}function _inherits(e,t){if("function"!=typeof t&&null!==t)throw new TypeError("Super expression must either be null or a function, not "+typeof t);e.prototype=Object.create(t&&t.prototype,{constructor:{value:e,enumerable:!1,writable:!0,configurable:!0}}),t&&(Object.setPrototypeOf?Object.setPrototypeOf(e,t):e.__proto__=t)}!function(global,san){function processTemplateModule(filterTemplateArrs,filterModule){eval(filterModule);var modules={},templateFiltersObj={};return filterTemplateArrs&&filterTemplateArrs.forEach(function(element){var filterName=element.filterName,func=element.func,module=element.module;modules[module]=eval(module),templateFiltersObj[filterName]=function(){var e;return(e=modules[module])[func].apply(e,arguments)}}),templateFiltersObj}global.errorMsg=[];var templateComponents=Object.assign({},{}),param={},filterArr=JSON.parse("[]");try{filterArr&&filterArr.forEach(function(item){param[item.module]=eval(item.module)});var pageContent='<view class="container" s-if="isShow"><map id="map" longitude="{{ firstMaker.longitude }}" latitude="{{ firstMaker.latitude }}" markers="{{markers}}" scale="10" on-bindmarkertap="eventHappen(\'markertap\', $event, \'markertap\', \'\', \'bind\')" style="width: 100%; height: {{height}}vh;" /><view class="brandmap--store__detail" style="height: {{100-height}}vh;  align-items: center;" s-if="height-100"><view style="flex:1;"><view style="margin-bottom:0.4vw; font-size:3.7333333333333334vw;color: #111;">({{store.title}})</view><view style="display: flex; font-size:3.2vw"><view >营业时间: </view><view style="flex: 1;"><template s-for="time, index in store.openTimeJson"><view style="padding-right: 1.3333333333333333vw;">{{time.time}}</view></template></view></view><view style="display: flex; align-items: center;font-size:3.2vw">店铺电话:{{ store.areaNum ? store.areaNum+\'-\': \'\' }} {{store.tel}} <image on-bindtap="eventHappen(\'tap\', $event, \'makephone\', \'\', \'bind\')" style="width:4.266666666666667vw;height:4.266666666666667vw;margin-left: 4.266666666666667vw;" src="http://imgcdnali.wfenxiao.com.cn/xcx/c/icons/blue-phone.png"></image></view><view style="font-size:3.2vw">店铺地址:{{store.detailAddress}}</view></view><view on-bindtap="eventHappen(\'tap\', $event, \'createmap\', \'\', \'bind\')" style="padding-left:5.333333333333333vw;font-size:3.2vw"><image src="http://imgcdnali.wfenxiao.com.cn/xcx/c/icons/ic_nav_home_footprint.png" style="width:10.133333333333333vw;height:10.133333333333333vw;"></image></view></view></view>',renderPage=function(e,t){var n=global.componentFactory,r=(n.getAllComponents(),{}),i=JSON.parse("{}"),o=Object.keys(i).reduce(function(e,t){return e[t]=r[i[t]],e},{});global.pageRender(pageContent,templateComponents,o,e,t)},oldPatch=function(e){Object.assign(e.components,{});var t=function(t){function n(t){_classCallCheck(this,n);var r=_possibleConstructorReturn(this,(n.__proto__||Object.getPrototypeOf(n)).call(this,t));return r.components=e.components,r}return _inherits(n,t),n}(e);t.template='<swan-wrapper tabindex="-1">'+pageContent+"</swan-wrapper>";var n=new t;n.slaveLoaded(),n.communicator.onMessage("initData",function(e){n.setInitData(e),n.attach(document.body)},{listenPreviousEvent:!0})};global.pageRender?renderPage(filterArr,param):oldPatch(window.PageComponent)}catch(e){throw global.errorMsg.execError=e,e}}(window,window.san);