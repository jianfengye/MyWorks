<?php
include_once("Snoopy.class.php");

$snoopy = new Snoopy;

global $domainLinks;

$domainLinks = array();

function getDomainLinks($url, $domain) {
	global $domainLinks;

	$snoopy = new Snoopy;
	$snoopy->agent = "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.17 (KHTML, like Gecko) Chrome/24.0.1312.57 Safari/537.17";

	$snoopy->rawheaders['Accept'] = 'text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8';
	$snoopy->rawheaders['Accept-Charset'] = 'GBK,utf-8;q=0.7,*;q=0.3';
	$snoopy->rawheaders['Connection'] = 'keep-alive';
	$snoopy->rawheaders['Accept-Language'] = 'zh-CN,zh;q=0.8';
	$snoopy->rawheaders['Cache-Control'] = 'max-age=0';

	$links = array();
	if ($snoopy->fetchlinks($url)) {
		foreach ($snoopy->results as $link) {
			if (stripos($link, $domain) === false) {
				continue;
			}
			if (in_array($link, $domainLinks)) {
				continue;
			}
			$domainLinks[] = $link;
			echo $link . "\r\n";
			getDomainLinks($link, $domain);
		}
	}
}

getDomainLinks('http://www.fjjxxy.cn/jxxy/index.aspx', 'www.fjjxxy.cn');