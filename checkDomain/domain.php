<?php

$pattern = $argv[1];

// 获取*可能的字符
$list = array();
for ($i = 97; $i <= 122; $i++) {
	$list[] = chr($i);
}
for($i = 0; $i <= 9; $i++) {
	$list[] = "$i";
}

checkPatten($pattern);

function checkPatten($pattern) {
	global $list;
	$pos = strpos($pattern, '*');
	if ($pos === false) {
		$isExist = isDomainExist($pattern);
		echo "{$pattern} :";
		echo isDomainExist($pattern)? "domain is empty" : "domain had exist";
		echo PHP_EOL;
	} else {
		foreach ($list as $item) {
			$pattern[$pos] = $item;
			checkPatten($pattern);
		}
	}
}

function isDomainExist($domain) {
	$api = 'http://pandavip.www.net.cn/check/check_ac1.cgi?callback=jQuery&domain=';
	$api .= urlencode($domain);

	$content = file_get_contents($api);

	$items = explode('|', $content);
	$code = $items[2];

	return $code == 210;
}