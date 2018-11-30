<?php
function p($str)
{
    echo "<pre>";
    print_r($str);
    echo "</pre>";
}
function ps($str)
{
    p($str);
    die();
}
$signaturec = "";
$md5 = "";

$outputFilePath = "./swan-core-extension.zip";
// 计算 md5
$md5 = md5_file($outputFilePath);
$pk = openssl_pkey_get_private(file_get_contents("./rsa_private_key_1024.pem"));

// p($md5);

// 计算签名
openssl_private_encrypt($md5, $signature, $pk, OPENSSL_PKCS1_PADDING);

$b64sign = strtr(base64_encode($signature), '+/', '-_');

// 计算包大小
$size = filesize($outputFilePath);
p($b64sign);
// p($size);
