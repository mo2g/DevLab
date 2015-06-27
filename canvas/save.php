<?php
preg_match('#^data:image/(\w+);base64,(.+)#', $_POST['data'], $result);

$type = $result[1];//图片类型
$data = base64_decode($result[2]);//解码图片数据
$time = time();
$file = "{$time}.{$type}";//图片存储名称
file_put_contents($file, $data);//保存图片

//显示图片
header('Content-Type:image/png');
echo $data;