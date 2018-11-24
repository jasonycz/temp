<?php

function p($str){
	echo "<pre>";
	print_r($str);
	echo "</pre>";
}
function ps($str){
	p($str);
	die();
}

function grab_image($url, $base_dir='',$section_dir='', $filename='', &$i = 0){
        if(empty($url)){
          return false;
        }

        // 为空就当前目录
        if(empty($base_dir)){
            $base_dir = './pic/';
        }
        $dir = $base_dir.$section_dir;
        
        // 判断是否是文件夹 不是则创建
        if(!is_dir($dir)){
            echo "目录".$dir."不存在，将创建目录".$dir."<br>";
            $mode = 0777; //创建目录的模式
            $res = mkdir($dir,$mode,true);
            if($res){
                echo "目录".$dir."创建成功<br>";
            }else{
                echo "目录".$dir."创建失败,请注意<br>";
                return;
            }
        }

        $dir = realpath($dir);
        // 目录+文件
        $filename = $dir . (empty($filename) ? '/'.time().'.jpeg' : '/'.$filename);

        // 开始捕捉 
        ob_start(); 
        readfile($url); 
        $img = ob_get_contents(); 
        ob_end_clean(); 
        $size = strlen($img); 
        $fp2 = fopen($filename , "a"); 
        fwrite($fp2, $img); 
        fclose($fp2); 

        echo ++$i.'|';
        return $filename; 
    } 
function  download_pic_file($res){
        if(empty($res)){
                echo "下载到本地时,数据为空";
                return;
            }
        $i = 0;
        foreach ($res as  $value) {
            $url = $value['pic_url'];
            $base_dir = './pic/';
            grab_image($url);
        }
        echo "<br>文件下载成功,一共".$i."张图片!<br>";

        return $i;
    }
function unescape($str) {
 $str = rawurldecode($str);
 preg_match_all("/(?:%u.{4})|.{4};|&#\d+;|.+/U",$str,$r);
 $ar = $r[0];
 #print_r($ar);
 foreach($ar as $k=>$v) {
 if(substr($v,0,2) == "%u")
 $ar[$k] = iconv("UCS-2","GB2312",pack("H4",substr($v,-4)));
 elseif(substr($v,0,3) == "")
 $ar[$k] = iconv("UCS-2","GB2312",pack("H4",substr($v,3,-1)));
 elseif(substr($v,0,2) == "&#") {
 echo substr($v,2,-1)."";
 $ar[$k] = iconv("UCS-2","GB2312",pack("n",substr($v,2,-1)));
 }
 }
 return join("",$ar);
}
/**
 * 读取文件 并处理内容
 * @param  $filename 文件路径
 * @param  $time 日期
 * @param  $table 表名
 * @param  $type kpi区分
 */
function readfromfile($filename) {
    $goods_list = array();
    if(!file_exists($filename)){
        echo "文件：$time----读取失败：文件不存在！----\r\n";
        return;
    }
    $list_fir = scandir($filename);

    foreach ($list_fir as $value) {
        if($value == '.' || $value == '..'){
            continue;
        }
        $file = fopen($filename.'/'.$value,'r'); 
        while ($data = fgetcsv($file)) {     //每次读取CSV里面的一行内容
            //此为一个数组，要获得每一个数据，访问数组下标即可
            $goods_list[] = $data;
        }
        //判断非空文件进行
        $arr = array();
        if(!empty($goods_list)){
            foreach ($goods_list as $v) {
                $arr[] = explode("\x01",$v[0]);
            }
            
        }else{
            echo "文件：$time----文件内容为空----\r\n";
        }
        fclose($file);
        return $arr;
    }
}
$a = array(
    array(
        'a' => 1,
    )
);
// $url ='http://mp.weixin.qq.com/mp/getmasssendmsg?__biz=MTA3NDM1MzUwMQ%3D%3D&mid=2651934761&idx=1&sn=322d1778ed95104cf81a275c5087594b&chksm=73d3521b44a4db0d295fc87a9eb60a2c3b4562b162692f2936e363a6736c115e3980b681646e&scene=0&key=9fb316639d9e8fea69107a4dbe8a2a959848c2421ba3186956a8eb1b279302811a99d6f172deb6dadce8e870809c37aca36190b83db71451f318558afd59593028b98e2aea6979a54f885c8d505faf91&ascene=7&uin=MjA4ODI2ODU4NA%3D%3D&devicetype=android-23&version=2605033d&nettype=WIFI&abtest_cookie=AQABAAgAAQAghh4AAAA%3D&pass_ticket=T7ZHh2sOzb8YwxMX1YZW%2BHuSJ7DJera1qX7kKHhKd9QMe9n%2Fx49U4D4EQZ9X%2FRbL&wx_header=1';
// $url = 'http://mp.weixin.qq.com/mp/getmasssendmsg?__biz=MTA3NDM1MzUwMQ%3D%3D&mid=2651934761&idx=1';
// echo file_get_contents($url);

echo round(1121.23432,1);

// http://mp.weixin.qq.com/profile?src=3&timestamp=1484663529&ver=1&
// signature=CVIt88FW2zhB7F-vLKiKD7jZihSkAlXdz8k5xcSwSvhj8oL
// *aSTvqkYN5R0W8FNg9*
// Z4h7OIcoh1W1*mvUJhaA==

// MjM5NTIyMzkwMg==

// $str ='neihanapp';
// echo urlencode($str);
// echo json_encode($str);
// echo urldecode('Z4h7OIcoh1W1*mvUJhaA==');
// echo utf8_decode('Z4h7OIcoh1W1*mvUJhaA==');
// echo base64_decode('Z4h7OIcoh1W1*mvUJhaA==');
// echo base64_encode($str);

// http://mp.weixin.qq.com/mp/getmasssendmsg?__biz=MTA3NDM1MzUwMQ%3D%3D&mid=2651934761&idx=1&sn=322d1778ed95104cf81a275c5087594b&chksm=73d3521b44a4db0d295fc87a9eb60a2c3b4562b162692f2936e363a6736c115e3980b681646e&scene=0&key=9fb316639d9e8fea69107a4dbe8a2a959848c2421ba3186956a8eb1b279302811a99d6f172deb6dadce8e870809c37aca36190b83db71451f318558afd59593028b98e2aea6979a54f885c8d505faf91&ascene=7&uin=MjA4ODI2ODU4NA%3D%3D&devicetype=android-23&version=2605033d&nettype=WIFI&abtest_cookie=AQABAAgAAQAghh4AAAA%3D&pass_ticket=T7ZHh2sOzb8YwxMX1YZW%2BHuSJ7DJera1qX7kKHhKd9QMe9n%2Fx49U4D4EQZ9X%2FRbL&wx_header=1



