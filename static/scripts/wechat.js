const
    appId = 'xxx'
apiHost = 'http://xxx.com/api/'
host = window.location.href
tokenKey = 'token'

function getQueryString(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i")
    var r = window.location.search.substr(1).match(reg)
    if (r != null) return unescape(r[2]); return null
}

function axiosGet({ url, params, headers, success, err }) {
    headers = headers || {}
    if (headers && !headers['token']) {
        var token = localStorage.getItem(tokenKey)
        if (!!token) {
            headers['token'] = token
        }
    }
    axios.get(url, {
        params: params,
        headers: headers
    }).then(function (response) {
        res = response.data
        if (res && res.code == 1) {
            if (typeof success == 'function') {
                success(response.data)
            }
            return true
        } else {
            if (typeof err == 'function') {
                err(response.data)
            }
            return false
        }
    }).catch(function (error) {
        if (typeof err == 'function') {
            err(error)
        }
        return false
    });
}

function auth(appId, froceRefresh) {
    var token = localStorage.getItem(tokenKey)
    if (froceRefresh){
        token = null
    }
    if (!token) {
        var code = getQueryString('code')
        if (!code) {
            var uri = window.location
            var encodeUri = encodeURIComponent(uri)
            var redirect = 'https://open.weixin.qq.com/connect/oauth2/authorize?appid=' + appId + '&redirect_uri=' + encodeUri + '&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect'
            window.location = redirect
        }
        else {
            axiosGet({
                url: apiHost + 'wxsession',
                params: { 'code': code },
                success: function (res) {
                    localStorage.setItem(tokenKey, res.data.token)
                },
                err: function (res) {
                    console.error(res)
                }
            })
        }
    }
    return true
}

function getDirectory() {
    axiosGet({
        url: apiHost + 'v1/directory',
        success: function (res) {
            console.log(res.data)
        },
        err: function (res) {
            console.error(res)
        }
    })
}

function init() {
    if (!auth(appId)) return
    getDirectory()
}

init()
