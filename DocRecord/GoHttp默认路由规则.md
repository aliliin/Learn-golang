# Go Default Router
#Learning

```
func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request){
    handler := sh.srv.Handler
    if handler == nil {
        handler = DefaultServeMux // 使用缺省的Router
    }
    if req.RequestURI == "*" && req.Method == "OPTIONS" {
        handler = globlOptionsHandler{}
    }
    handler.ServeHTTP(rw,req)
}
```

## 路由规则

* URL 分为两种，末尾是 /： 表示一个子树，后面可以跟其他的子路径；末尾不是 / ，表示一个叶子，固定的路径
    以 /结尾的 URL 可以匹配它的任何子路径，比如 /images 回匹配 /images/cute-cat.jpg
* 它采用最长匹配原则，如果多个匹配，一定采用匹配路径最长的那个进行处理
* 如果没有找到任何匹配项，会返回404