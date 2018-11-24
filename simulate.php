<?php
class requester {
 /**
  *
  * @var String 请求的完整URL
  */
 public $url;
 /**
  *
  * @var String 请求方式 GET,POST
  */
 public $method = 'GET';
 /**
  *
  * @var String Content-Type eg. 'application/json'
  *      ,'application/xml','application/x-www-form-urlencoded'
  */
 public $content_type = 'application/x-www-form-urlencoded';
 /**
  *
  * @var String 请求过程使用的字符集编码
  */
 public $charset = 'UTF-8';
 /**
  *
  * @var String 请求数据
  */
 public $data;
 /**
  *
  * @var boolean 是否启用cookie
  */
 public $enableCookie;
 /**
  *
  * @var resource 启用cookie发送请求时需要的cookie文件
  */
 public $cookieFile;
 /**
  *
  * @var boolean 启用时会将响应头信息作为数据流输出
  */
 public $enableHeaderOutput;
 
 /**
  *
  * @param String $charset
  *         请求URL.
  */
  function __construct($url) {
    $this->url = $url;
 }
 
 /**
  * 模拟浏览器发送请求
  *
  * @return array
  *         包含封装的http状态码,响应内容和cookie的数组array(retCode,retContent,retCookieFile).
  */
 public final function request() {
  ini_set ( 'max_execution_time', '0' );
  
  $ch = curl_init ();
  
  $header = array (
    'Content-Type: ' . $this->content_type . '; charset=' . strtoupper ( $this->charset ) . '',
    'Content-Length: ' . strlen ( $this->c ) 
  );
  $phone = "Mozilla/5.0 (Linux; Android 4.4.2; sdk Build/KK) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36 MicroMessenger/6.0.0.61_r920612.501 NetType/epc.tmobile.com";
    // curl_setopt($ch, CURLOPT_USERAGENT, "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0)");
  curl_setopt($ch, CURLOPT_USERAGENT, $phone);
  curl_setopt ( $ch, CURLOPT_URL, $this->url );
  curl_setopt ( $ch, CURLOPT_HTTPHEADER, $header );
  curl_setopt ( $ch, CURLOPT_HEADER, $this->enableHeaderOutput );
  curl_setopt ( $ch, CURLOPT_RETURNTRANSFER, 1 );
  
  $cookieJar = null;


  if ($this->enableCookie) {
   // 带cookie请求服务器
   curl_setopt ( $ch, CURLOPT_COOKIEFILE, $this->cookieFile );
   // 保存服务器发送的cookie
   $cookieJar = tempnam ( 'tmp', 'cookie' );
   curl_setopt ( $ch, CURLOPT_COOKIEJAR, $cookieJar );
  }
  
  if (strtoupper ( $this->method ) == 'POST') {
   curl_setopt ( $ch, CURLOPT_POST, 1 );
   curl_setopt ( $ch, CURLOPT_POSTFIELDS, $this->data );
  }
  
  $return_content = curl_exec ( $ch );
  
  $return_code = curl_getinfo ( $ch, CURLINFO_HTTP_CODE );
  
  curl_close ( $ch );
  
  return array (
    $return_code,
    $return_content,
    $cookieJar 
  );
 }
}

function http_post_data($url, $data_string) {

    $ch = curl_init();
    $phone = "Mozilla/5.0 (Linux; Android 4.4.2; sdk Build/KK) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36 MicroMessenger/6.0.0.61_r920612.501 NetType/epc.tmobile.com";
    // curl_setopt($ch, CURLOPT_USERAGENT, "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0)");
    curl_setopt($ch, CURLOPT_USERAGENT, $phone);
    curl_setopt($ch, CURLOPT_POST, 1);
    curl_setopt($ch, CURLOPT_URL, $url);
    curl_setopt($ch, CURLOPT_POSTFIELDS, $data_string);
    curl_setopt($ch, CURLOPT_HTTPHEADER, array(
    'Content-Type: application/json; charset=utf-8',
    'Content-Length: ' . strlen($data_string))
    );
    ob_start();
    curl_exec($ch);
    $return_content = ob_get_contents();
    ob_end_clean();

    $return_code = curl_getinfo($ch, CURLINFO_HTTP_CODE);
    return array($return_code, $return_content);
}

$url = "http://www.baidu.com";
$res = http_post_data($url,"hello wold");
print_r($res);

// $requester = new requester($url);

// $requester->url  ='https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxgeticon?seq=1861471339&username=@57c78b57adb3ebd20e03e52b36aa06ce79d58d092705692c04f4f78ac6ea78bd&skey=@crypt_2ade0da7_ba40b759fc9acb70a22a3463759da3c7';
// // $requester->url = 'www.laifengzixun.com';

// // $requester->url = 'http://chat.gettyimages.cn/chat/get_searchurl1.php?ask=可以问问你把&ip=106.120.217.164&start_time=&userid=244';
// $requester->charset= "utf-8";
// $requester->content_type='application/x-www-form-urlencoded';
// // $requester->data="username=xzg&passwd=xzg";
// $requester->enableCookie= true;
// $requester->enableHeaderOutput=false;
// $requester->method="get";
// // $requester->method="post";
// $requester->cookieFile=$retCookieFile;
// $arr = $requester->request();
// list ( $retCode, $retContent, $retCookieFile ) = $arr;
// echo $retCode . '<br>' . $retContent . '<br>' . $retCookieFile . '<br>';


https://wx.qq.com//cgi-bin/mmwebwx-bin/webwxcheckurl?requrl=https%3A%2F%2Fmp.weixin.qq.com%2Fmp%2Fgetmasssendmsg%3F__biz%3DMjM5ODIyMTE0MA%3D%3D%26from%3D1%26uin%3DMjA4ODI2ODU4NA%253D%253D%26key%3D417d265a523a5e43a86884b3718fb06f608031f9b950afcc4194e438d17f291efd29a505f94385d31f5efdd54a8541717cd82705f3725e1d5393fcbc4cea34b51ec73e570f382f54cabe95aac2afda20%26devicetype%3DiOS10.2.1%26version%3D16050520%26lang%3Dzh_CN%26nettype%3DWIFI%26ascene%3D7%26fontScale%3D94%26pass_ticket%3DCm0Fp1aTTFSq62%252BJodErD9844NBVm2nuhmK%252BueYG7ErXwXl%252BUGSxWGYY%252BpKaoYdN%26wx_header%3D1%23wechat_webview_type%3D1&skey=%40crypt_2ade0da7_0e5bd1c64ef39447ad1aa852443186da&deviceid=e721046265143044&pass_ticket=50%252BqYdOR%252Bj7liK2MUe2ESapXaANV7qtzSAExNyS9vp6zQMGBPWhUUx4MZRmTg7Y3&opcode=2&scene=1&username=@4bb8d75f901dd10b5b867d15b155c6417c729de7110f52f7b794ea8620649d37


https://wx.qq.com//cgi-bin/mmwebwx-bin/webwxcheckurl?requrl=https%3A%2F%2Fmp.weixin.qq.com%2Fmp%2Fgetmasssendmsg%3F__biz%3DMjM5ODIyMTE0MA%3D%3D%26from%3D1%26uin%3DMjA4ODI2ODU4NA%253D%253D%26key%3D417d265a523a5e43a86884b3718fb06f608031f9b950afcc4194e438d17f291efd29a505f94385d31f5efdd54a8541717cd82705f3725e1d5393fcbc4cea34b51ec73e570f382f54cabe95aac2afda20%26devicetype%3DiOS10.2.1%26version%3D16050520%26lang%3Dzh_CN%26nettype%3DWIFI%26ascene%3D7%26fontScale%3D94%26pass_ticket%3DCm0Fp1aTTFSq62%252BJodErD9844NBVm2nuhmK%252BueYG7ErXwXl%252BUGSxWGYY%252BpKaoYdN%26wx_header%3D1%23wechat_webview_type%3D1&skey=%40crypt_2ade0da7_0e5bd1c64ef39447ad1aa852443186da&deviceid=e721046265143044&pass_ticket=50%252BqYdOR%252Bj7liK2MUe2ESapXaANV7qtzSAExNyS9vp6zQMGBPWhUUx4MZRmTg7Y3&opcode=2&scene=1&username=@4bb8d75f901dd10b5b867d15b155c6417c729de7110f52f7b794ea8620649d37


https://wx.qq.com//cgi-bin/mmwebwx-bin/webwxcheckurl?skey=@crypt_2ade0da7_0e5bd1c64ef39447ad1aa852443186da&deviceid=e721046265143044&pass_ticket=50%2BqYdOR%2Bj7liK2MUe2ESapXaANV7qtzSAExNyS9vp6zQMGBPWhUUx4MZRmTg7Y3&opcode=2&scene=1&username=@4bb8d75f901dd10b5b867d15b155c6417c729de7110f52f7b794ea8620649d37



https://wx.qq.com//cgi-bin/mmwebwx-bin/webwxcheckurl?requrl=https://mp.weixin.qq.com/mp/getmasssendmsg?__biz=MjM5ODIyMTE0MA==&from=1&uin=MjA4ODI2ODU4NA%3D%3D&key=417d265a523a5e43a86884b3718fb06f608031f9b950afcc4194e438d17f291efd29a505f94385d31f5efdd54a8541717cd82705f3725e1d5393fcbc4cea34b51ec73e570f382f54cabe95aac2afda20&devicetype=iOS10.2.1&version=16050520&lang=zh_CN&nettype=WIFI&ascene=7&fontScale=94&pass_ticket=Cm0Fp1aTTFSq62%2BJodErD9844NBVm2nuhmK%2BueYG7ErXwXl%2BUGSxWGYY%2BpKaoYdN&wx_header=1#wechat_webview_type=1&skey=@crypt_2ade0da7_0e5bd1c64ef39447ad1aa852443186da&deviceid=e721046265143044&pass_ticket=50%2BqYdOR%2Bj7liK2MUe2ESapXaANV7qtzSAExNyS9vp6zQMGBPWhUUx4MZRmTg7Y3&opcode=2&scene=1&username=@4bb8d75f901dd10b5b867d15b155c6417c729de7110f52f7b794ea8620649d37



?>