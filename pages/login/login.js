window.define("pages/login/login",function(o,t,e,n,a,i,s,l,u,d,r,c,h,f,g,p){"use strict";function v(o){return o&&o.__esModule?o:{default:o}}var m=o("../../utils/util"),w=v(m),y=o("../../network/api"),T=v(y),D=o("../../utils/const.js"),C=v(D),S=o("../../utils/notification.js"),b=v(S),R=o("../../dist/index"),N={data:{loading:!1,verify:!1,count:60,countDown:!1,steps:[{current:!0,done:!0,text:"步骤一",desc:"10.01"},{done:!1,current:!1,text:"步骤二",desc:"10.02"}]}};w.default.registComp(N,R.Toast,R.TopTips,{onLoad:function(o){var t=o.type||"",e=o.from||"",n=o.channelCode||"",a=o.zhaomuName||"";this.setData({type:t,from:e,zhaomuName:a,channelCode:n})},getCode:function(){var o=this;if(this.data.verify){var t=this,e={mobile:this.data.phone};this.data.loading||this.data.countDown||(a.showLoading({title:"请稍候"}),this.setData({loading:!0}),T.default.checkBindMobileIsExist(e).then(function(n){a.hideLoading(),a.stopPullDownRefresh(),o.setData({loading:!1}),n.data.exist?o.showZanTopTips("该手机号已经被绑定过"):(o.setData({countDown:!0}),T.default.getVerifyCode(e).then(function(e){a.stopPullDownRefresh(),console.log("bingooooooooooooo",e),e.data.result?t.setCount():o.showZanTopTips("获取验证码失败")}).catch(function(o){return console.log(o)}))}).catch(function(t){console.log(t),a.hideLoading(),o.setData({loading:!1})}))}},verifyPhone:function(o){this.setData({verify:11==o.detail.value.length&&/^1[0-9]{10}$/.test(o.detail.value),phone:o.detail.value})},formSubmit:function(o){var t=this;console.log(o.detail.value);var e=o.detail.value,n=void 0;return e.phone&&0!=e.phone.trim().length?/^1[0-9]{10}$/.test(e.phone)?e.verifyCode&&0!=e.verifyCode.trim().length?void T.default.bindMobile({verifyCode:e.verifyCode,mobile:e.phone}).then(function(o){if(a.stopPullDownRefresh(),console.log("bingooooooooooooo",o),o.data.result){a.showToast({title:"手机号绑定成功",duration:1e3});var n=a.getStorageSync("userinfo");console.log(n),n.mobile=e.phone,a.setStorageSync("userinfo",n),i().globalData.userinfo=n,console.log("this.data.type="+t.data.type),"zhaomu"===t.data.type?t.applyRecruitment():a.navigateBack({delta:1})}else t.showZanTopTips(o.data.msg?o.data.msg:"验证码不正确")}).catch(function(o){return console.log(o)}):(n="验证码不能为空",void this.showZanTopTips(n)):(n="请输入正确手机号",void this.showZanTopTips(n)):(n="手机号不能为空",void this.showZanTopTips(n))},applyRecruitment:function(){var o=this;T.default.applyRecruitment({from:this.data.from,channelCode:this.data.channelCode}).then(function(){var t=a.getStorageSync("userinfo");t.isDist=!0,a.setStorageSync("userinfo",t),i().globalData.userinfo=t,console.log("----userInfo--54545-------",t),o.data.zhaomuName||a.showToast({title:"恭喜您，已成功申请为"+o.data.zhaomuName+"！",duration:1e3}),b.default.post(C.default.action.BIND_PHONE_SUCCESS,{}),setTimeout(function(){a.navigateBack({delta:1})},2e3)}).catch(function(o){console.log("apply err=>",o),a.hideLoading(),a.showToast({title:T.default.errors(o.error)||"网络连接错误，请稍后重试",duration:1e3}),setTimeout(function(){a.navigateBack({delta:1})},2e3)})},setCount:function(){var o=this.data,t=o.count,e=o.phone,n=this;if(0==t&&11==e.length&&/^1[0-9]{10}$/.test(e))return void this.setData({verify:!0,phone:e,count:60,countDown:!1});var a=t-1;this.setData({count:a}),setTimeout(function(){n.setCount()},1e3)}}),Page(N)}),window.__swanRoute="pages/login/login",window.usingComponents=[],require("pages/login/login");