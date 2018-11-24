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

$opts = array( 
	'http'=>array( 
		'method'=>"GET", 
		'header'=>"Accept-language: en\r\n" . 
		"Cookie: foo=bar\r\n" 
	) 
); 
$context = stream_context_create($opts); 

p($context);