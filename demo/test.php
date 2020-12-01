<?php
//$body_digest = 'SHA-256=' . base64_encode(hash('sha256', $body, true));

$body = '{"lng":"116.23128","lat":"40.22077","ip":"36.112.75.34","token":"ad95e6983a197eb448de26d1f5db60b6","bu":{"union_id":"1","media_id":"1","posid":"test4maidian","pack_name":"ad.union.package","website":"www.baidu.com","media_type":"1","key_word":"玻尿酸"},"device":{"device_id":"123123test123123","os":"iOS","osv":"10.3.1","connection_type":1,"mac":"00-16-EA-AE-3C-40","imei":"123123test123123","idfa":"123123test123123","oaid":"123123test123123","androidid":"123123test123123"},"content":{"page":1,"limit":5},"api_version":"1.0","request_id":""}';
$h = hash('sha256', $body, 'raw');
$h = base64_encode($h);
$h = base64_encode(hash('sha256', $body, true));
var_dump($h);

//2eU1DdfR3ZqX+A7TxwbnqD1RlYI2QPZMwB7JPO9KNhs=

//$str_to_sign = "host: open.sy.soyoung.com\ndate: Tue, 03 Nov 2020 07:47:30 GMT\nPOST /union/union/getList HTTP/1.1\ndigest: SHA-256=i3QdI5LFcjX5h9LwJ6/SsP/AVyTbutwOBKvrj5WbTk0=";

$str_to_sign = "host: open.sy.soyoung.com
date: Tue, 03 Nov 2020 08:00:43 GMT
POST /union/union/getList HTTP/1.1
digest: SHA-256=s8ywYJz3H69SIhwkmViNzRdQ8+v2QKmH7amOWkOPp3E=";
$str_to_header = "host date request-line digest";
$str_to_sign = trim($str_to_sign);
$str_to_header = trim($str_to_header);

$ak = "zYhNVeuGLbesGRFM";
$sk = "l4XItYeVzvvYpYdZAdDR9ZlRYttwtuK7";
$signature = hash_hmac('sha256', $str_to_sign, $sk, true);
$signature = base64_encode($signature);
$Authorization = sprintf('hmac username="%s", algorithm="hmac-sha256", headers="%s", signature="%s"', $ak, $str_to_header, $signature);
print_r(array(
    "ak" => $ak,
    "sk" => $sk,
    "str_to_sign" => $str_to_sign,
    "str_to_header" => $str_to_header,
    "signature" => $signature,
    "Authorization" => $Authorization,
));

//ppUIH5aQMtv24NZoLLJjGMd2SiHNkvVk4XIGLDWby2M=