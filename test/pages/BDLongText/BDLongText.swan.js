function _classCallCheck(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}function _possibleConstructorReturn(e,t){if(!e)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!t||"object"!=typeof t&&"function"!=typeof t?e:t}function _inherits(e,t){if("function"!=typeof t&&null!==t)throw new TypeError("Super expression must either be null or a function, not "+typeof t);e.prototype=Object.create(t&&t.prototype,{constructor:{value:e,enumerable:!1,writable:!0,configurable:!0}}),t&&(Object.setPrototypeOf?Object.setPrototypeOf(e,t):e.__proto__=t)}!function(global,san){function processTemplateModule(filterTemplateArrs,filterModule){eval(filterModule);var modules={},templateFiltersObj={};return filterTemplateArrs&&filterTemplateArrs.forEach(function(element){var filterName=element.filterName,func=element.func,module=element.module;modules[module]=eval(module),templateFiltersObj[filterName]=function(){var e;return(e=modules[module])[func].apply(e,arguments)}}),templateFiltersObj}global.errorMsg=[];var templateComponents=Object.assign({},{}),param={},filterArr=JSON.parse("[]");try{filterArr&&filterArr.forEach(function(item){param[item.module]=eval(item.module)});var pageContent='<to-home s-if="showToHome"></to-home><view class="error-page-box" s-if=" !startLoading && error "><network-toast s-if=" !networkAvailable " on-bindloadData="eventHappen(\'loadData\', $event, \'loadData\', \'\', \'bind\')"></network-toast><network-toast s-else="" text="{{ errorMsg }}" on-bindloadData="eventHappen(\'loadData\', $event, \'loadData\', \'\', \'bind\')"></network-toast></view><view class="bd-long-text" s-elif=" !startLoading "><view class="bd-long-text-title"><text >{{title}}\n\t </text></view><view class="bd-long-text-userinfo"><image class="bd-long-text-header" src="{{header}}"></image><text class="bd-long-text-name">{{userName}}\n\t\t</text><text class="bd-long-text-time">{{month}}月{{day}}日\n\t\t</text></view><view class="bd-long-rich-text"><rich-text nodes="{{BDContent}}" on-bindtap="eventHappen(\'tap\', $event, \'tap\', \'\', \'bind\')"></rich-text></view></view><view class="ad"><ad appid="bb1adc43" apid="6029463" type="banner"></ad></view>',renderPage=function renderPage(filters,modules){var componentFactory=global.componentFactory,componentFragments=componentFactory.getAllComponents(),user_self_compo_45a2=function(global,componentFragments){var templateComponents={},usingComponentMap=JSON.parse("{}"),filterCustomArr=JSON.parse("[]"),modules={},filtersObj={};filterCustomArr&&filterCustomArr.forEach(function(element){modules[element.module]=eval(element.module);var func=element.func,module=element.module;filtersObj[element.filterName]=function(){var e;return(e=modules[module])[func].apply(e,arguments)}});var components=Object.assign({},componentFragments,templateComponents);return global.componentFactory.componentDefine("common/NetworkErrorToast/NetworkErrorToast.swan",Object.assign({},componentFactory.getProtos("super-custom-component"),{template:'<swan-network-toast><view class="network-error-box"><image class="error-img" src="/static/resources/images/empty_failed.png"></image><text class="error-toast">{{text || "网络出错啦，请点击按钮重新加载"}}</text><button class="error-btn" s-if=" !text " on-bindtap="eventHappen(\'tap\', $event, \'retryRequest\', \'\', \'bind\')">重新加载</button></view></swan-network-toast>',componentPath:"common/NetworkErrorToast/NetworkErrorToast.swan",componentName:"network-toast",customComponentCss:"swan-network-toast .network-toast__network-error-box {\n  display: flex;\n  flex-direction: column;\n  align-items: center;\n  justify-content: center;\n}\n\nswan-network-toast .network-toast__error-img {\n  width: 26.666666666666668vw;\n  height: 26.666666666666668vw;\n}\n\nswan-network-toast .network-toast__error-toast {\n  margin-top: 3.2vw;\n  font-family: PingFangSC-Regular;\n  font-size: 3.7333333333333334vw;\n  color: #999999;\n  text-align: center;\n}\n\nswan-network-toast .network-toast__error-btn {\n  border: 0.26666666666666666vw solid #D2D2D2;\n  border-radius: 13.333333333333334vw;\n  font-family: PingFangSC-Medium;\n  font-size: 4vw;\n  color: #636363;\n  text-align: center;\n  margin-top: 6.4vw;\n  width: 49.06666666666667vw;\n  height: 9.866666666666667vw;\n}\n\nswan-network-toast swan-button::after{\n  border: none;\n}"}),{classProperties:{components:components,filters:Object.assign({},filtersObj)}}),Promise.resolve().then(function(){for(var e in usingComponentMap)components[e]=customAbsolutePathMap[usingComponentMap[e]]}),global.componentFactory.getComponents("common/NetworkErrorToast/NetworkErrorToast.swan")}(global,componentFragments),user_self_compo_45a1=function(global,componentFragments){var templateComponents={},usingComponentMap=JSON.parse("{}"),filterCustomArr=JSON.parse("[]"),modules={},filtersObj={};filterCustomArr&&filterCustomArr.forEach(function(element){modules[element.module]=eval(element.module);var func=element.func,module=element.module;filtersObj[element.filterName]=function(){var e;return(e=modules[module])[func].apply(e,arguments)}});var components=Object.assign({},componentFragments,templateComponents);return global.componentFactory.componentDefine("common/ToHomeButton/ToHomeButton.swan",Object.assign({},componentFactory.getProtos("super-custom-component"),{template:"<swan-to-home><cover-image class=\"float-home-low\" src=\"/static/resources/images/feed_button_gohome.png\" on-bindtap=\"eventHappen('tap', $event, 'toHome', '', 'catch')\"></cover-image></swan-to-home>",componentPath:"common/ToHomeButton/ToHomeButton.swan",componentName:"to-home",customComponentCss:"/* common/ToHomeButton/ToHomeButton.css */\nswan-to-home .to-home__float-home,swan-to-home  .to-home__float-home-low {\n  width: 50px;\n  height: 50px;\n}\nswan-to-home .to-home__float-home {\n  position: fixed;\n  bottom: 45.86666666666667vw;\n  right: 5.333333333333333vw;\n  z-index: 1000000;\n}\nswan-to-home .to-home__float-home-low {\n  position: fixed;\n  bottom: 31.733333333333334vw;\n  right: 3.466666666666667vw;\n  z-index: 1000000;\n}"}),{classProperties:{components:components,filters:Object.assign({},filtersObj)}}),Promise.resolve().then(function(){for(var e in usingComponentMap)components[e]=customAbsolutePathMap[usingComponentMap[e]]}),global.componentFactory.getComponents("common/ToHomeButton/ToHomeButton.swan")}(global,componentFragments),customAbsolutePathMap={"common/ToHomeButton/ToHomeButton":user_self_compo_45a1,"common/NetworkErrorToast/NetworkErrorToast":user_self_compo_45a2},pageUsingComponentMap=JSON.parse('{"network-toast":"common/NetworkErrorToast/NetworkErrorToast","to-home":"common/ToHomeButton/ToHomeButton"}'),customComponents=Object.keys(pageUsingComponentMap).reduce(function(e,t){return e[t]=customAbsolutePathMap[pageUsingComponentMap[t]],e},{});global.pageRender(pageContent,templateComponents,customComponents,filters,modules)},oldPatch=function(e){Object.assign(e.components,{});var t=function(t){function o(t){_classCallCheck(this,o);var n=_possibleConstructorReturn(this,(o.__proto__||Object.getPrototypeOf(o)).call(this,t));return n.components=e.components,n}return _inherits(o,t),o}(e);t.template='<swan-wrapper tabindex="-1">'+pageContent+"</swan-wrapper>";var o=new t;o.slaveLoaded(),o.communicator.onMessage("initData",function(e){o.setInitData(e),o.attach(document.body)},{listenPreviousEvent:!0})};global.pageRender?renderPage(filterArr,param):oldPatch(window.PageComponent)}catch(e){throw global.errorMsg.execError=e,e}}(window,window.san);