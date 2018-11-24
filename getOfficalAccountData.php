<?php
    header("Content-Type:text/html;charset=UTF-8");

    global $appkey = 'wpa8689835980';
    function http_post_data($url, $data_string) {
        $ch = curl_init();
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
    /**
     * [getAppMsgStat description]
     * @param  [type] &$old_data 获取点赞数 并且将点赞相应的数据插入到原来的数组中
     * @return [type]            [description]
     */
    function getAppMsgStat(&$old_data){
        if(!empty($old_data['content_url'])){
            $url  = "http://182.61.31.212:8888/geText";
            $arr = array('appkey' => $appkey,'url' => $old_data['content_url']);
            $data = json_encode($arr); 
            list($return_code, $return_content) = http_post_data($url, $data);
            if(!empty($return_content['appmsgstat'])){
                $return_content = json_decode($return_content,true);
                $app_msg_stat = $return_content['appmsgstat'];
                $old_data['appmsgstat'] = $app_msg_stat;
            }else{
                // 打失败日志
                // 从新执行一次
                Log::info($old_data['content_url'].'获取点赞数失败');
            }
        }
    }

    function getAppMsgExtInfo(&$app_msg_ext_info){
        // 判定是否有附加文章
        if($app_msg_ext_info['is_multi'] == 1){
            foreach ($$app_msg_ext_info['multi_app_msg_item_list'] as $key => $value) {
                getAppMsgStat($value);
            }
        }
        if(!empty($app_msg_ext_info['content_url'])){
            getAppMsgStat($app_msg_ext_info);
        }
    }

    function getCommMsgInfo($comm_msg_info){
        if(!empty($comm_msg_info)){
            return $comm_msg_info;
        }
        return false;
    }

    function insertPublicOffilcalContent($data){
        // 组装数据 和 将数据 json_encode
        $insert_data = array(
            'biz' => $data['biz'],
            'app_msg_ext_info' => json_encode($data['app_msg_ext_info']),
            'comm_msg_info' => json_encode($data['comm_msg_info']),
            'datatime' => $data['datatime'],
        );
        
        // 插入数据到数据库中
        $res = DB::table('publicOfficialContent')->insert($insert_data);
        if($res == 1){
            // Log::info($data['biz']."插入一条数据");
        }
    }

    function getOfficalAccountData($biz,$first_time = false){
        $url  = "http://182.61.31.212:8888/getArticleList";
        $arr = array('appkey' => $appkey,'biz' => $biz);
        $data = json_encode($arr); 
        list($return_code, $return_content) = http_post_data($url, $data);

        $res = array();
        $tmp = array();

        if($first_time){
            if($return_code == 200){
                $return_content = json_decode($return_content,true);
                // 对list里面的数据进行遍历处理
                foreach ($return_content['list'] as $key => $value) {
                    foreach ($value as $k => $v) {
                        if($k == 'app_msg_ext_info'){
                            $tmp['app_msg_ext_info'] = $v;
                        }else if($k == 'comm_msg_info'){
                            $tmp['comm_msg_info'] = getCommMsgInfo($v);
                            $tmp['datatime'] = $v['datatime'];
                        }
                        $tmp['biz'] = $biz;
                    }
                    // 插入数据
                    insertPublicOffilcalContent($tmp);

                }

            }
        }
    }


/*
$url  = "http://182.61.31.212:8888/geText";
$arr = array('appkey'=>'wpa8689835980','url'=>'http://mp.weixin.qq.com/s?__biz=MzI0MDI3NDM3NA==&mid=2247484957&idx=3&sn=8c77eb4e8558b8ea99c66f680c929375#rd');
$data = json_encode($arr); 

list($return_code, $return_content) = http_post_data($url, $data);
//print_r($return_code);
print_r($return_content);
  */
// $url  = "http://182.61.31.212:8888/getArticleList";
// $arr = array('appkey'=>'wpa8689835980','biz' => "MzI0MDI3NDM3NA==");
// $data = json_encode($arr); 

// list($return_code, $return_content) = http_post_data($url, $data);
// //print_r($return_code);
// print_r($return_content);

 






