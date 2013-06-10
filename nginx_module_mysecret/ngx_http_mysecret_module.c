/*
 * Copyright (C) yejianfeng
 */

#include <ngx_config.h>
#include <ngx_core.h>
#include <ngx_http.h>

// 定义配置结构
typedef struct {
	ngx_str_t secret;
} ngx_http_mysecret_conf_t;

// 定义模块
ngx_module_t ngx_http_mysecret_module = {
	NGX_MODULE_V1,
	&ngx_http_mysecret_module_ctx,
	ngx_http_mysecret_commands,
	NGX_HTTP_MODULE,
	NULL,
	NULL,
	NULL,
	NULL,
	NULL,
	NULL,
	NGX_MODULE_V1_PADDING
};

// 定义上下文, 只会在location中出现，所以都为null
static ngx_http_module_t ngx_http_mysecret_module_ctx = {
	NULL,
	NULL,
	NULL,
	NULL,
	NULL,
	NULL,
	ngx_http_mysecret_create_loc_conf,  //这里有使用自己的配置结构
	NULL
};

static void *
ngx_http_mysecret_create_loc_conf(ngx_conf_t *cf)
{
	ngx_http_mysecret_conf_t * mycf;
	mycf = ngx_palloc(cf->pool, sizeof(ngx_http_mysecret_conf_t));
	if (mycf == NULL) {
		return NGX_CONF_ERROR;
	}

	mycf->secret.len = 0;
	mycf->secret.data = NULL;
	return mycf;
}

static ngx_command_t ngx_http_mysecret_commands[] = {
	{
		ngx_string("mysecret"),
		NGX_HTTP_LOC_CONF | NGX_CONF_TAKE1,
		ngx_http_mysecret,
		NGX_HTTP_LOC_CONF_OFFSET,
		0,
		NULL,
	},

	ngx_null_command
};

static char * 
ngx_http_mysecret(ngx_conf_t *cf, ngx_command_t *cmd, void *conf)
{
	ngx_http_core_loc_conf_t  *clcf;
	clcf = ngx_http_conf_get_module_loc_conf(cf, ngx_http_core_module);
	clcf->handler = ngx_http_mysecret_handler;

	ngx_conf_set_str_slot(cf, cmd, conf);
	return NGX_CONF_OK;
}

static ngx_int_t ngx_http_mysecret_handler(ngx_http_request_t *r)
{
	if (!(r->method & (NGX_HTTP_GET | NGX_HTTP_HEAD))) {
		return NGX_HTTP_NOT_ALLOWED;
	}

	ngx_http_mysecret_conf_t *mycf;

	mycf = ngx_http_get_module_loc_conf(r, ngx_http_mysecret_module);

	// 获取secret参数
	// 查找有args有没有secret参数
	u_char *args = r->args.data;
	size_t *len = r->args.len;

	ngx_str_t key = ngx_string("secret=") ;


	r->args
}