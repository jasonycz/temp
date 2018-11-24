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
$source_file = './api.json';
$destination_file = './resouces.js';

// 读取 json 文件
$json = file_get_contents($source_file);

// 转换 json 为数组
$json2array = json_decode($json,true);

// 组装数据
$json2array['scheme'] = 'https';
unset($json2array['schemes']);
unset($json2array['swagger']);
unset($json2array['info']);
unset($json2array['produces']);
unset($json2array['definitions']);

$json2array['resources'] = [];

foreach ($json2array['paths'] as $path => $vvvv) {
    foreach ($vvvv as $method => $vvv) {
        if(isset($vvv['summary'])){
          unset($vvvv[$method]['summary']);
        }
        if(isset($vvv['description'])){
          unset($vvvv[$method]['description']);
        }
        if(isset($vvv['responses'])){
          unset($vvvv[$method]['responses']);
        }
        foreach ($vvvv[$method]['parameters'] as $k => $vv) {
            if(isset($vv['description'])){
              unset($vvvv[$method]['parameters'][$k]['description']);
            }
            if(isset($vv['schema'])){
              $vvvv[$method]['parameters'][$k]['type'] = 'object';
              unset($vvvv[$method]['parameters'][$k]['schema']);
            }
        }
        $tags = $vvv['tags'][0];
        unset($vvvv[$method]['tags']);
        $json2array['resources'][$tags][$path] = $vvvv;

    }
}
unset($json2array['paths']);


$responses_json = 'module.exports = ';
$responses_json .= str_replace("\\/", "/",json_encode($json2array));


file_put_contents($destination_file, $responses_json);

echo $responses_json;





// p($json2array);

