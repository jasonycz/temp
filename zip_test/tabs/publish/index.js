window.define("tabs/publish/index",function(n,e,o,i,t,a,s,u,l,r,d,g,b,h,p,w){"use strict";Page({data:{},bindPublish:function(n){var e=n.currentTarget.dataset.url;console.log(n);var o=t.getStorageSync("userinfo");this.setData({userinfo:o}),o.mobile?t.navigateTo({url:e}):this.setData({showModal:!0,url:"/pages/login/login"})}})}),window.__swanRoute="tabs/publish/index",window.usingComponents=["component/phonemodal/phonemodal"],require("tabs/publish/index");