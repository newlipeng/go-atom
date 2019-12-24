package dcache

import (
    "fmt"
    "log"
    "github.com/bradfitz/gomemcache/memcache"
    "github.com/pquerna/ffjson/ffjson"
    "go-atom/dutil"

    "strings"
    "time"
)

type memcached struct {
    client *memcache.Client
    threshold int
}

func (this *memcached) Call(ret map[string]interface{}, f func(ret map[string]interface{}, cacheRet map[string][]byte, emptyKey []string, params ...interface{}), prefix string, key []string, params ...interface{}) {

    cacheRet, emptyKey := this.GetMulti(key, prefix)
    f(ret, cacheRet, emptyKey, params...)
    var keyArr []string
    valArr := make(map[string][]byte)
    for k, v := range ret {
        _, ok := cacheRet[k]
        if !ok {
            keyArr = append(keyArr, k)
            tmpVal, ok := ffjson.Marshal(v)
            if nil == ok {
                valArr[k] = tmpVal
            } else {
                log.Println("error")
            }
        }
    }
    fmt.Println("begin", keyArr, valArr, "end")
    this.SyncSetMulti(keyArr, valArr, 60, prefix)
}

func (this *memcached) GetProxyItem(ret map[string]interface{}, f func(ret map[string]interface{}, cacheRet map[string][]byte, emptyKey []string, params ...interface{}), prefix string, key []string, params ...interface{}) {


}

func (this *memcached) GetProxyMulti(ret map[string]interface{}, f func(ret map[string]interface{}, cacheRet map[string][]byte, emptyKey []string, params ...interface{}), prefix string, key []string, params ...interface{}) {

    cacheRet, emptyKey := this.GetMulti(key, prefix)
    f(ret, cacheRet, emptyKey, params...)
    var keyArr []string
    valArr := make(map[string][]byte)
    for _, v := range emptyKey {
        data, ok := ret[v]
        if ok {

        }
    }
    for k, v := range ret {
        _, ok := cacheRet[k]
        if !ok {
            keyArr = append(keyArr, k)
            tmpVal, ok := ffjson.Marshal(v)
            if nil == ok {
                valArr[k] = tmpVal
            } else {
                log.Println("error")
            }
        }
    }
    fmt.Println("begin", keyArr, valArr, "end")
    this.SyncSetMulti(keyArr, valArr, 60, prefix)
}

func MemcachedInit(memcachedServer string, memcachedThreshold int, maxIdleConns int, timeout int) *memcached {
    cache := &memcached{}
    serverList := strings.Split(memcachedServer, ",")
    if len(serverList) < 0 {
        log.Fatal("memcached conf is incorrect, please check it")
    }
    cache.client = memcache.New(serverList...)
    cache.client.MaxIdleConns = maxIdleConns
    cache.client.Timeout = time.Duration(timeout) * time.Millisecond
    cache.threshold = memcachedThreshold
    return cache
}

func (this *memcached) SetServer(memcachedServer string, maxIdleConns int, timeout int) {
    serverList := strings.Split(memcachedServer, ",")
    if len(serverList) > 0 {
        log.Fatal("memcached conf is incorrect, please check it")
    }
    this.client = memcache.New(serverList...)
    this.client.MaxIdleConns = maxIdleConns
    this.client.Timeout = time.Duration(timeout) * time.Second
}

func (this *memcached) SetCompressThreshold(val int) {
    this.threshold = val
}

func (this *memcached) SetItem(key string, val []byte, expire int32, prefix string) {
    var flags uint32 = 1
    if len(val) > this.threshold {
        flags = 2
        val = dutil.SnappyEncode(val)
    }
    item := &memcache.Item{Key: prefix + key, Flags: flags, Value: val, Expiration: expire}
    err := this.client.Set(item)
    if (nil != err) {
        log.Println("memcached set key %s error", prefix + key)
    }
}

func (this *memcached) AsyncSetItem(key string, val []byte, expire int32, prefix string) {
    go this.SetItem(key, val, expire, prefix);
}

func (this *memcached) SetMulti(key []string, val map[string][]byte, expire int32, prefix string) {
    for _, item := range key {
        this.SetItem(item, val[item], expire, prefix)
    }
}

func (this *memcached) AsyncSetMulti(key []string, val map[string][]byte, expire int32, prefix string) {
    go this.SetMulti(key, val, expire, prefix)
}

func (this *memcached) GetItem(key string, prefix string) (ret string) {
    item, err := this.client.Get(prefix + key)
    if (nil != err) {
    } else {
        switch item.Flags {
        case 2:
            item.Value = dutil.SnappyDecode(item.Value)
            ret = string(item.Value[:])
        case 1:
            fallthrough
        default:
            ret = string(item.Value[:])
        }
    }
    return
}

func (this *memcached) GetMulti(key []string, prefix string) (ret map[string][]byte, emptyKey []string) {
    ret = make(map[string][]byte)
    keyArr := []string{}
    tmpMap := make(map[string]string)
    for _, v := range key {
        keyArr = append(keyArr, prefix + v)
        tmpMap[prefix+v] = v
    }
    itemMap, err := this.client.GetMulti(keyArr)
    if nil != err {
        emptyKey = key
    } else {
        for _, v:= range itemMap {
            if v.Flags == 2 {
                v.Value = dutil.SnappyDecode(v.Value)
            }
            ret[tmpMap[v.Key]] = v.Value
            delete(tmpMap, v.Key)
        }
        for _, v:= range tmpMap {
            emptyKey = append(emptyKey, v)
        }
    }
    return
}

