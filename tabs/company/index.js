window.define("tabs/company/index",function(e,o,n,t,a,i,c,d,s,r,l,p,u,m,f,y){"use strict";Page({data:{selectArea:{},tab:"tab",dataType:1},onLoad:function(){this.selectComponent("#company").handleLoad()},onShow:function(){var e=i().getFlash("address");if(e){console.log(e);var o=e.province,n=e.city,t=e.district;this.setData({selectArea:{provinceCode:o.code,cityCode:n.code||0,districtCode:t.code||0,city:""+(o.fullName||"")+(n.fullName||"")+(t.fullName||""),code:((o.code||0)+","+(n.code||0)+","+(t.code||0)).split(",").filter(function(e){return 0!=e}).join(",")}})}}})}),window.__swanRoute="tabs/company/index",window.usingComponents=["component/decoration_company/index"],require("tabs/company/index");