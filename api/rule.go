package handler

import (
	"net/http"
)

const configContent = "[custom]\n" +
	";不要随意改变关键字，否则会导致出错\n" +
	";acl4SSR规则-在线版\n\n" +
	";去广告：支持\n" +
	";自动测速：不支持\n" +
	";微软分流：不支持\n" +
	";苹果分流：不支持\n" +
	";增强中国IP段：不支持\n" +
	";增强国外GFW：不支持\n\n" +
	"ruleset=🎯 全球直连,https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/LocalAreaNetwork.list\n" +
	"ruleset=🎯 全球直连,https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/UnBan.list\n" +
	"ruleset=🛑 全球拦截,https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/BanAD.list\n" +
	"ruleset=🛑 全球拦截,https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/BanProgramAD.list\n" +
	"ruleset=🎯 全球直连,https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/GoogleCN.list\n" +
	"ruleset=🎯 全球直连,https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/SteamCN.list\n" +
	"ruleset=🚀 节点选择,https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Telegram.list\n" +
	"ruleset=🚀 节点选择,https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/ProxyMedia.list\n" +
	"ruleset=🚀 节点选择,https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/ProxyLite.list\n" +
	"ruleset=🎯 全球直连,https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/ChinaDomain.list\n" +
	"ruleset=🎯 全球直连,https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/ChinaCompanyIp.list\n" +
	";ruleset=🎯 全球直连,[]GEOIP,LAN\n" +
	"ruleset=🎯 全球直连,[]GEOIP,CN\n" +
	"ruleset=🐟 漏网之鱼,[]FINAL\n\n" +
	"ruleset=🎯 全球直连,DOMAIN-SUFFIX,ipinfo.io\n" +
	"ruleset=🎯 全球直连,DOMAIN-SUFFIX,qihoo.net\n" +
	"\n\n" +
	"custom_proxy_group=🚀 节点选择\\`select\\`[]DIRECT`.*\n" +
	"custom_proxy_group=🎯 全球直连\\`select\\`[]DIRECT\\`[]🚀 节点选择\n" +
	"custom_proxy_group=🛑 全球拦截\\`select\\`[]REJECT\\`[]DIRECT\n" +
	"custom_proxy_group=🐟 漏网之鱼\\`select\\`[]🚀 节点选择\\`[]🎯 全球直连\\`.*\n\n" +
	"enable_rule_generator=true\n" +
	"overwrite_original_rules=true\n"

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, configContent)
}
