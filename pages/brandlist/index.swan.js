function _classCallCheck(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}function _possibleConstructorReturn(e,t){if(!e)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!t||"object"!=typeof t&&"function"!=typeof t?e:t}function _inherits(e,t){if("function"!=typeof t&&null!==t)throw new TypeError("Super expression must either be null or a function, not "+typeof t);e.prototype=Object.create(t&&t.prototype,{constructor:{value:e,enumerable:!1,writable:!0,configurable:!0}}),t&&(Object.setPrototypeOf?Object.setPrototypeOf(e,t):e.__proto__=t)}!function(global,san){function processTemplateModule(filterTemplateArrs,filterModule){eval(filterModule);var modules={},templateFiltersObj={};return filterTemplateArrs&&filterTemplateArrs.forEach(function(element){var filterName=element.filterName,func=element.func,module=element.module;modules[module]=eval(module),templateFiltersObj[filterName]=function(){var e;return(e=modules[module])[func].apply(e,arguments)}}),templateFiltersObj}global.errorMsg=[];var templateComponents=Object.assign({},{}),param={},filterArr=JSON.parse("[]");try{filterArr&&filterArr.forEach(function(item){param[item.module]=eval(item.module)});var pageContent='<view class="container"><scroll-view on-bindscrolltolower="eventHappen(\'scrolltolower\', $event, \'lower\', \'\', \'bind\')" enable-back-to-top="" scroll-top="{{scrollTop}}" scroll-y="" style="width:100%;height: 100vh;"><template s-for="item, index in list"><view class="zan-panel  zan-panel__no-margin-top" on-bindtap="eventHappen(\'tap\', $event, \'makeproduct\', \'\', \'bind\')" data-item="{{item}}"><view class="zan-card" style="padding-top: 2.6666666666666665vw;display: flex;align-items: center;"><view class="" style="width:15.466666666666667vw;"><image class="" src="{{item.logo || \'http://imgcdnali.wfenxiao.com.cn/xcx/c/icons/brand.png\' }}" style="width:15.466666666666667vw;height: 14.4vw;" mode="{{scaleToFill}}"></image></view><view class="" style="flex:1;margin-left: 2.1333333333333333vw;"><view class=""><view class="zan-card__left-col" style="font-size:3.466666666666667vw;margin-bottom: 1.3333333333333333vw;color: #333;"><text >{{item.brandName}}</text></view></view><view class=" " style="margin-bottom:0;height:8vw;"><view class="" style="font-size: 3.3333333333333335vw;color: #999999;"><text class="zan-ellipsis--l2" style="line-height: 135%;">{{item.description}}</text></view></view></view></view></view></template><view id="loadmore"><template-7e2382 data="{{{loading: loadmore}}}" data-origin-name="zan-loadmore"></template-7e2382></view></scroll-view></view>',renderPage=function(e,t){var n=global.componentFactory,o=(n.getAllComponents(),{}),l=JSON.parse("{}"),a=Object.keys(l).reduce(function(e,t){return e[t]=o[l[t]],e},{});global.pageRender(pageContent,templateComponents,a,e,t)},oldPatch=function(e){Object.assign(e.components,{});var t=function(t){function n(t){_classCallCheck(this,n);var o=_possibleConstructorReturn(this,(n.__proto__||Object.getPrototypeOf(n)).call(this,t));return o.components=e.components,o}return _inherits(n,t),n}(e);t.template='<swan-wrapper tabindex="-1">'+pageContent+"</swan-wrapper>";var n=new t;n.slaveLoaded(),n.communicator.onMessage("initData",function(e){n.setInitData(e),n.attach(document.body)},{listenPreviousEvent:!0})};global.pageRender?renderPage(filterArr,param):oldPatch(window.PageComponent)}catch(e){throw global.errorMsg.execError=e,e}}(window,window.san);