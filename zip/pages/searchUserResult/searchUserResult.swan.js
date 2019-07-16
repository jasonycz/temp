function _classCallCheck(e,n){if(!(e instanceof n))throw new TypeError("Cannot call a class as a function")}function _possibleConstructorReturn(e,n){if(!e)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!n||"object"!=typeof n&&"function"!=typeof n?e:n}function _inherits(e,n){if("function"!=typeof n&&null!==n)throw new TypeError("Super expression must either be null or a function, not "+typeof n);e.prototype=Object.create(n&&n.prototype,{constructor:{value:e,enumerable:!1,writable:!0,configurable:!0}}),n&&(Object.setPrototypeOf?Object.setPrototypeOf(e,n):e.__proto__=n)}!function(global,san){function processTemplateModule(filterTemplateArrs,filterModule){eval(filterModule);var modules={},templateFiltersObj={};return filterTemplateArrs&&filterTemplateArrs.forEach(function(element){var filterName=element.filterName,func=element.func,module=element.module;modules[module]=eval(module),templateFiltersObj[filterName]=function(){var e;return(e=modules[module])[func].apply(e,arguments)}}),templateFiltersObj}global.errorMsg=[];var templateComponents=Object.assign({},{}),param={},filterArr=JSON.parse("[]");try{filterArr&&filterArr.forEach(function(item){param[item.module]=eval(item.module)});var pageContent='<view s-for="item, index in userList" s-key="{{index}}"><user-list user="{{item}}" show-follow-button="{{showFollowButton}}" show-is-follow="{{showIsFollow}}" showloading="{{showloading}}" page="user" index="{{index}}"></user-list></view><view class="loading" s-if="loadMoreButton && !hasMoreUser" on-bindtap="eventHappen(\'tap\', $event, \'loadMore\', \'\', \'bind\')">更多...</view><view class="loading" s-if="hasMoreUser">加载中...</view>',renderPage=function renderPage(filters,modules){var componentFactory=global.componentFactory,componentFragments=componentFactory.getAllComponents(),user_self_compo_45a3=function(global,componentFragments){var templateComponents={},usingComponentMap=JSON.parse("{}"),filterCustomArr=JSON.parse("[]"),modules={},filtersObj={};filterCustomArr&&filterCustomArr.forEach(function(element){modules[element.module]=eval(element.module);var func=element.func,module=element.module;filtersObj[element.filterName]=function(){var e;return(e=modules[module])[func].apply(e,arguments)}});var components=Object.assign({},componentFragments,templateComponents);return global.componentFactory.componentDefine("common/UserList/UserList.swan",Object.assign({},componentFactory.getProtos("super-custom-component"),{template:'<swan-user-list><view class="user {{index==0? \'first-user\' :\'no-first\'}}"><image class="user-image" src="{{user.userSrc}}" data-name="{{user.userName}}" data-uid="{{user.uid}}" on-bindtap="eventHappen(\'tap\', $event, \'goToProfile\', \'\', \'bind\')"></image><view class="content-box"><view class="user-detail" data-name="{{user.userName}}" data-uid="{{user.uid}}" on-bindtap="eventHappen(\'tap\', $event, \'goToProfile\', \'\', \'bind\')"><view class="user-name-long">{{user.userName}}</view><view class="information-long">{{user.verified ? \'微博认证：\'+user.userInformation : user.userInformation}}</view></view></view></view></swan-user-list>',componentPath:"common/UserList/UserList.swan",componentName:"user-list",customComponentCss:"swan-user-list .user-list__user{\n  display: flex;\n}\nswan-user-list .user-list__first-user{\n  padding-left: 3.2vw;  \n}\nswan-user-list .user-list__no-first{\n  padding: 4.266666666666667vw 0 0 3.2vw;\n}\nswan-user-list .user-list__user-image{\n  width: 12.8vw;\n  height: 12.8vw;\n  border-radius: 6.4vw;\n  border: 0.13333333333333333vw solid rgba(230,230,230,0.50);\n}\nswan-user-list .user-list__content-box{\n  display: flex;\n  padding-right:3.2vw;\n  margin-left: 3.2vw;\n  border-bottom: 0.13333333333333333vw solid #E6E6E6;\n  width:77.6vw;\n  position: relative;\n}\nswan-user-list .user-list__user-detail{\n  display: flex;\n  flex-direction: column;\n  background: #FFFFFF;\n  padding-bottom: 4vw;\n}\nswan-user-list .user-list__user-name{\n  width: 55.2vw;\n  text-overflow: ellipsis;\n  white-space: nowrap;\n  overflow: hidden;\n  font-family: PingFangSC-Medium;\n  font-size: 4.266666666666667vw;\n  color: #333333;\n  margin-bottom: 1.2vw;\n}\nswan-user-list .user-list__information{\n  width: 55.2vw;\n  text-overflow: ellipsis;\n  white-space: nowrap;\n  overflow: hidden;\n  font-family: PingFangSC-Regular;\n  font-size: 3.7333333333333334vw;\n  color: #999999;\n  height: 5.333333333333333vw;\n}\nswan-user-list .user-list__user-name-long{\n  width: 77.6vw;\n  text-overflow: ellipsis;\n  white-space: nowrap;\n  overflow: hidden;\n  font-family: PingFangSC-Medium;\n  font-size: 4.266666666666667vw;\n  color: #333333;\n  margin-bottom: 1.2vw;\n}\nswan-user-list .user-list__information-long{\n  width: 77.6vw;\n  text-overflow: ellipsis;\n  white-space: nowrap;\n  overflow: hidden;\n  font-family: PingFangSC-Regular;\n  font-size: 3.7333333333333334vw;\n  color: #999999;\n  height: 5.333333333333333vw;\n}\nswan-user-list .user-list__follow-box{\n  position: absolute;\n  right: 3.2vw;\n  top: 2.6666666666666665vw;\n}"}),{classProperties:{components:components,filters:Object.assign({},filtersObj)}}),Promise.resolve().then(function(){for(var e in usingComponentMap)components[e]=customAbsolutePathMap[usingComponentMap[e]]}),global.componentFactory.getComponents("common/UserList/UserList.swan")}(global,componentFragments),customAbsolutePathMap={"common/UserList/UserList":user_self_compo_45a3},pageUsingComponentMap=JSON.parse('{"user-list":"common/UserList/UserList"}'),customComponents=Object.keys(pageUsingComponentMap).reduce(function(e,n){return e[n]=customAbsolutePathMap[pageUsingComponentMap[n]],e},{});global.pageRender(pageContent,templateComponents,customComponents,filters,modules)},oldPatch=function(e){Object.assign(e.components,{});var n=function(n){function t(n){_classCallCheck(this,t);var o=_possibleConstructorReturn(this,(t.__proto__||Object.getPrototypeOf(t)).call(this,n));return o.components=e.components,o}return _inherits(t,n),t}(e);n.template='<swan-wrapper tabindex="-1">'+pageContent+"</swan-wrapper>";var t=new n;t.slaveLoaded(),t.communicator.onMessage("initData",function(e){t.setInitData(e),t.attach(document.body)},{listenPreviousEvent:!0})};global.pageRender?renderPage(filterArr,param):oldPatch(window.PageComponent)}catch(e){throw global.errorMsg.execError=e,e}}(window,window.san);