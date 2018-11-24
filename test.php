<?php

function p($str)
{
    echo '<pre>';
    print_r($str);
    echo '</pre>';
}
function ps($str)
{
    p($str);
    die();
}

function _calculateScore($strVersionCode)
{
    if (empty($strVersionCode)) {
        return 0;
    }
    $arrCode = explode('.', $strVersionCode);
    $intScore = 0;
    foreach ($arrCode as $intCode) {
        $intScore = $intCode | $intScore << 10;
    }
    return $intScore;
}
$version = "12.0.0.0";
$score = _calculateScore($version);

$sql = "insert into host_version(version_code,score)values('" . $version . "'," . $score . ");";

echo $sql;
