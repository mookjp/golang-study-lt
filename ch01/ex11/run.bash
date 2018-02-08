#!/bin/bash
SCRIPT_DIR=$(cd $(dirname $0);pwd)

# Directory to save output
OUTPUT_DIR="${SCRIPT_DIR}/tmp"
mkdir -p ${OUTPUT_DIR}

# Url
# Got by google chrome developer console on https://www.alexa.com/topsites
# var out = ""; $(".site-listing .DescriptionCell p a").each(function (i, a) { out += a.getAttribute("href").replace("/siteinfo/", "") + " " }); console.log(out)
URLS="google.com youtube.com facebook.com baidu.com wikipedia.org reddit.com yahoo.com google.co.in qq.com taobao.com amazon.com twitter.com tmall.com google.co.jp live.com instagram.com vk.com sohu.com sina.com.cn jd.com weibo.com 360.cn google.de google.co.uk google.com.br google.fr google.ru yandex.ru netflix.com linkedin.com twitch.tv google.it t.co google.com.hk google.es pornhub.com alipay.com xvideos.com yahoo.co.jp google.ca google.com.mx ebay.com microsoft.com bing.com ok.ru imgur.com office.com bongacams.com hao123.com aliexpress.com"

go run ${SCRIPT_DIR}/fetchall.go ${OUTPUT_DIR} ${URLS}
