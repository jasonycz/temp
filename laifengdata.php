<?php

$list = array();
$articles = array(
	array(
	    "articleId" => "1",
        "title" => "【锐评】篮球改革，看点不仅仅是姚明",
        "imgUrl" => "http://mmbiz.qpic.cn/mmbiz_png/3qGBpicMFJ41D573LHYmcFEHFcC73XUJcxuSBouycXFkRV8iaKEbuxP0w9LUceqkKnKKTwekqaAasKMbwxebu5pw/0?wx_fmt=png",
        "author" => "薛原",
        "readNum" => "250",
        "voteNum" => "100",
        "date" => "1484132343"
	),
	array(
	    "articleId" => "2",
        "title" => "人民时评】农村土地整合是个潜力股",
        "imgUrl" => "http://mmbiz.qpic.cn/mmbiz_png/3qGBpicMFJ41D573LHYmcFEHFcC73XUJc5IsvuGic7kzrBdf815NKfzXn6MAZbOVCiaM7xSZ0eDGiaYel0iavqxlowg/0?wx_fmt=png",
        "author" => "伍振军",
        "readNum" => "1000",
        "voteNum" => "250",
        "date" => "1484139601"
	)
);
$list['list']['articles'] = $articles;
// var_dump($list);
// var_dump(json_encode($articles));
$str = json_encode($list);
echo $str;
// return $str;







