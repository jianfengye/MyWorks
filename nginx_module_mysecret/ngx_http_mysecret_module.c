/*
 * Copyright (C) yejianfeng
 */

#include <ngx_config.h>
#include <ngx_core.h>
#include <ngx_http.h>

static void *
ngx_http_mysecret_create_loc_conf(ngx_conf_t *cf);
static char * 
ngx_http_mysecret(ngx_conf_t *cf, ngx_command_t *cmd, void *conf);
static ngx_int_t ngx_http_mysecret_handler(ngx_http_request_t *r);

// 定义配置结构
typedef struct {
	ngx_str_t secret;
} ngx_http_mysecret_conf_t;

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
	NULL,
	NGX_MODULE_V1_PADDING
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

	ngx_int_t rc = ngx_http_discard_request_body(r);
    if (rc != NGX_OK) {
        return rc;
    }

	ngx_http_mysecret_conf_t *mycf;

	mycf = ngx_http_get_module_loc_conf(r, ngx_http_mysecret_module);

	ngx_int_t isAuth = 0;

	// 比较secret
	ngx_str_t format = ngx_string("secret=%V");
	ngx_str_t match;
	match.len = format.len - 2 + mycf->secret.len;
	match.data = ngx_palloc(r->pool, match.len);
	ngx_snprintf(match.data, match.len, "secret=%V", &mycf->secret);

	//对args的每个字段进行比较
	ngx_uint_t i = 0;
	if (r->args.len >= match.len) {
		for (; i <= r->args.len - match.len; i++) {
			if (0 == ngx_strncasecmp(r->args.data + i, match.data, match.len)) {
				if (i < r->args.len - match.len && *(r->args.data + i + match.len) != '&') {
					continue;
				}

				if (i != 0 && *(r->args.data + i - 1) != '&') {
					continue;
				}

				isAuth = 1;
				break;
			}
		}
	}

	ngx_str_t response;
	if (isAuth == 1) {
		ngx_str_set(&response, "secret right");
	} else {
		ngx_str_set(&response, "secret wrong");
	}

	ngx_str_t type = ngx_string("text/plain");
	r->headers_out.status = NGX_HTTP_OK;
	r->headers_out.content_length_n = response.len;
	r->headers_out.content_type = type;

	rc = ngx_http_send_header(r);
	if (rc == NGX_ERROR || rc > NGX_OK || r->header_only) {
		return rc;
	}

	ngx_buf_t *b;
	b = ngx_create_temp_buf(r->pool, response.len);
	if (b == NULL) {
		return NGX_HTTP_INTERNAL_SERVER_ERROR;
	}

	ngx_memcpy(b->pos, response.data, response.len);
	b->last = b->pos + response.len;
	b->last_buf = 1;

	ngx_chain_t out;
	out.buf = b;
	out.next = NULL;

	return ngx_http_output_filter(r, &out);
}
