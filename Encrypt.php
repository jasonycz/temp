<?php
/**
 * Package Encrypt
 */
class Util_Encrypt
{
    const FILE_FLAG = 0xBD180704;
    private $_aesKey = "";
    private $_aesIv = "";
    /**
     * 构造文件第一段header内容
     * @param version_code
     * @param fileNum
     * @param indexLength
     * @param dataLength
     * @param appId
     * @return string
     */
    public function _structureHeader($version_code, $fileNum, $indexLength, $dataLength, $appId)
    {
        $header = "";
        // V 32位4字节,采用小端模式
        $header .= pack("V", self::FILE_FLAG);
        $header .= pack("V", $version_code);
        $header .= pack("V", $fileNum);
        $header .= pack("V", $indexLength);
        $header .= pack("V", $dataLength);

        // P 64位8字节,采用小端模式
        // HHVM 的 pack不支持64位的pack 所以将 appid(64位) 拆成2个32位
        $header .= $this->packInt64P($appId);
        // rsa加密aes的key和iv
        $header .= $this->rsaEncryptAes();
        return $header;
    }

    public function BinToStr($str)
    {
        $arr = explode(' ', $str);
        foreach ($arr as &$v) {
            // 他做了什么？？
            $v = pack("H" . strlen(base_convert($v, 2, 16)), base_convert($v, 2, 16));
        }

        return join('', $arr);
    }
    public function StrToBin($str)
    {
        //1.列出每个字符
        // 这边的分割正则也不理解
        // (?<!^) 后瞻消极断言
        // (?!$) 前瞻消极断言
        // 看意思好像说的是：不以^开头（但是这边 ^ 又没有被转义...），不以 $ 结尾（同上）
        // 然后得到的记过就是字符串一个个被分割成了数组（郁闷）
        // 求解释
        $arr = preg_split('/(?<!^)(?!$)/u', $str);
        //2.unpack字符
        foreach ($arr as &$v) {
            /**
             * unpack：将二进制字符串解包(英语原文：Unpack data from binary string)
             * H: 英语描述原文：Hex string, high nibble first
             * 这段代码做了什么？？
             */
            $temp = unpack('H*', $v); // 这边被解析出来的字符串为什么是 16进制的？？
            $v = base_convert($temp[1], 16, 2);
            unset($temp);
        }

        return join(' ', $arr);
    }

    /**
     * 生成aes key和iv
     * @param version_code
     * @param fileNum
     * @param appId
     * @return bool
     */
    private function createAesKeyIv($version_code, $fileNum, $appId)
    {
        // 根据已有数据hash生成aes的key和iv,hash返回16字节二进制
        $this->_aesKey = md5($version_code . $fileNum, true);
        $this->_aesIv = md5($appId, true);

        echo sprintf("AES Key or IV is invalid. Key: %s, IV: %s.", bin2hex($this->_aesKey), bin2hex($this->_aesIv));
        die();
        return true;
    }
    /**
     * Rsa加密Aes的key&iv
     * @return string
     */
    private function rsaEncryptAes()
    {
        $rsa_private_key_file = dirname(__FILE__) . "/rsa_private_key_384.pem";
        $pk = openssl_pkey_get_private(file_get_contents($rsa_private_key_file));
        if ($pk === false) {
            throw new Exception("private key is invalid");
        }
        $aesData = $this->_aesKey . $this->_aesIv;
        if (strlen($aesData) != 32) {
            throw new Exception("AES Data is invalid");
        }
        $ok = openssl_private_encrypt($aesData, $signature, $pk, OPENSSL_PKCS1_PADDING);
        if ($ok === false) {
            throw new Exception("Encrypted AES Data failed");
        }
        if (strlen($signature) != 48) {
            throw new Exception("Encrypted AES Data is invalid");
        }
        return $signature;
    }
    /**
     * 构造文件第二段文件索引
     * @param $filePath
     * @return string
     */
    public function _structureIndex($data, $version_code, $fileNum, $appId)
    {
        $cipher = "AES-128-CBC";
        $indexData = '';
        foreach ($data as $key => $val) {
            // V 32位4字节,采用小端模式
            $indexData .= pack("V", $val['fileOffset']);
            $indexData .= pack("V", $val['fileSize']);
            $indexData .= pack("V", $val['filePathSize']);
            $indexData .= pack("a*", $val['filePath']);
        }

        // 生成aes key和iv
        $this->createAesKeyIv($version_code, $fileNum, $appId);
        // 通过AES算法CBC模式加密
        if (strlen($this->_aesKey) != 16 || strlen($this->_aesIv) != 16) {
            throw new Exception(sprintf("AES Key or IV is invalid. Key: %s, IV: %s.", bin2hex($this->_aesKey), bin2hex($this->_aesIv)));
        }

        $aesRet = openssl_encrypt($indexData, $cipher, $this->_aesKey, true, $this->_aesIv);
        $ret['size'] = strlen($aesRet);
        $ret['data'] = $aesRet;
        unset($indexData, $aesRet);
        return $ret;
    }
    /**
     * 构造文件第三段压缩文件内容
     * @param $filePath
     * @return string
     */
    public function _structureDeflateData($filePath)
    {
        $startOffset = 0;
        $deflateData = '';
        foreach ($filePath as $key => $val) {
            $itemLength = strlen($val['content']);
            $indexData[] = array(
                "fileSize" => $itemLength,
                "fileOffset" => $startOffset,
                "filePath" => $val['path'],
                "filePathSize" => strlen($val['path']),
            );
            $deflateData .= $val['content'];
            $startOffset = $startOffset + $itemLength;
        }
        $ret['size'] = strlen($deflateData);
        $ret['data'] = gzencode($deflateData, 9);
        $ret['indexData'] = $indexData;
        unset($startOffset, $deflateData, $indexData);
        return $ret;
    }
    /**
     * pack int64 in little endian like pack("P", int64)
     *
     * @param int $int64
     * @return array
     */
    public function packInt64P($int64)
    {
        $appidHigh = ($int64 >> 32);
        $appidLow = $int64 & 0xFFFFFFFF;
        $ret = pack("V", $appidLow);
        $ret .= pack("V", $appidHigh);
        return $ret;
    }
    /**
     * 加密
     * @param fileDir
     * @param version_code
     * @param appId
     * @param saveFilePath
     * @return string
     */
    public function encrypt($fileDir, $version_code, $appId, $saveFilePath)
    {
        $appId = intval($appId);
        $files = array();
        // 获取包文件
        // 获取所有的包路径 & 内容
        $this->__dirTree($files, $fileDir, $fileDir);
        // $this->ps($files);
        // 包结构-第三段
        // 获取每个文件内容大小 offset 文件路径 文件路径的大小 gzip 压缩所有文件内容
        $thirdContent = $this->_structureDeflateData($files);
        // $this->ps($thirdContent);

        // 包结构-第二段
        // 生成 Key IV indexData 进行大小端处理 进行 CBC 模式AES加密
        $secondContent = $this->_structureIndex($thirdContent['indexData'], $version_code, count($files), $appId);
        // 包结构-第一段
        // 对第二段当中生成的 Key 和 IV 进行 RSA 加密 + 已有数据生成 header
        $firstContent = $this->_structureHeader($version_code, count($files), $secondContent['size'], $thirdContent['size'], $appId);
        // 拼接包内容
        // 将第一段 第二段 第三段 合并在一起
        $packageContent = $firstContent . $secondContent['data'] . $thirdContent['data'];
        unset($filePath, $fileDir, $firstContent, $secondContent, $thirdContent);
        return $this->createFile($saveFilePath, $packageContent);
    }
    /**
     * 生成加密文件
     * @param saveFilePath
     * @param content
     * @return string
     */
    private function createFile($saveFilePath, $content)
    {
        file_put_contents($saveFilePath, $content);
        return $saveFilePath;
    }
    /**
     * 递归文件夹
     * @param arr_file
     * @param directory
     * @param dir_name
     * @return array
     */
    private function __dirTree(&$arr_file, $directory, $root, $dir_name = '')
    {
        $mydir = dir($directory);
        while ($file = $mydir->read()) {
            if ($file === "." || $file === "..") {
                continue;
            }
            if (is_dir("$directory/$file")) {
                $this->__dirTree($arr_file, "$directory/$file", $root, "$dir_name/$file");
            } else {
                $arr_file[] = array(
                    "path" => "$dir_name/$file",
                    "content" => file_get_contents("$root/$dir_name/$file"),
                );
            }
        }
        $mydir->close();
    }

    public function ps($data)
    {
        var_dump($data);
        // echo "<pre>";
        // print_r($data);
        // echo "</pre>";
        die();
    }
    public function p($data)
    {
        var_dump($data);
        // echo "<pre>";
        // print_r($data);
        // echo "</pre>";
    }
}

$fileDir = "./zip";
$version_code = 1;
$appId = 123;
$saveFilePath = "./smapp/test.smapp";
$utilEncrypt = new Util_Encrypt();
$utilEncrypt->encrypt($fileDir, $version_code, $appId, $saveFilePath);
