package etcd

import (
	"strings"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/micro/go-micro/v2/config/encoder"
)

func makeEvMap(e encoder.Encoder, data map[string]interface{}, kv []*clientv3.Event, stripPrefix string) map[string]interface{} {
	if data == nil {
		data = make(map[string]interface{})
	}

	for _, v := range kv {
		switch mvccpb.Event_EventType(v.Type) {
		case mvccpb.DELETE:
			data = update(e, data, (*mvccpb.KeyValue)(v.Kv), "delete", stripPrefix)
		default:
			data = update(e, data, (*mvccpb.KeyValue)(v.Kv), "insert", stripPrefix)
		}
	}

	return data
}

func makeMap(e encoder.Encoder, kv []*mvccpb.KeyValue, stripPrefix string) map[string]interface{} {
	data := make(map[string]interface{})

	for _, v := range kv {
		data = update(e, data, v, "put", stripPrefix)
	}

	return data
}

func update(e encoder.Encoder, data map[string]interface{}, v *mvccpb.KeyValue, action, stripPrefix string) map[string]interface{} {
	// remove prefix if non empty, and ensure leading / is removed as well
	vkey := strings.TrimPrefix(strings.TrimPrefix(string(v.Key), stripPrefix), "/")
	// split on prefix
	haveSplit := strings.Contains(vkey, "/")
	keys := strings.Split(vkey, "/")

	var vals interface{}
	e.Decode(v.Value, &vals)

	// 如果不是json
	if vals == nil {
		vals = string(v.Value)
	}

	if !haveSplit && len(keys) == 1 {

		switch action {
		case "delete":
			data = make(map[string]interface{})
		default:
			// 这部分对key进行合并原先只有最后一个key有效
			if v, ok := vals.(map[string]interface{}); ok {
				data[keys[0]] = v
			}

			if v, ok := vals.(string); ok {
				data[keys[0]] = v
			}

			if v, ok := vals.([]interface{}); ok {
				data[keys[0]] = v
			}

			if v, ok := vals.(bool); ok {
				data[keys[0]] = v
			}

			if v, ok := vals.(float64); ok {
				data[keys[0]] = v
			}
		}
		return data
	}

	// set data for first iteration
	kvals := data
	// iterate the keys and make maps
	for i, k := range keys {
		kval, ok := kvals[k].(map[string]interface{})
		if !ok {
			// create next map
			kval = make(map[string]interface{})
			// set it
			kvals[k] = kval
		}

		// last key: write vals
		if l := len(keys) - 1; i == l {
			switch action {
			case "delete":
				delete(kvals, k)
			default:
				kvals[k] = vals
			}
			break
		}

		// set kvals for next iterator
		kvals = kval
	}

	return data
}
