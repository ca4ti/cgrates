/*
Real-time Online/Offline Charging System (OCS) for Telecom & ISP environments
Copyright (C) ITsysCOM GmbH

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>
*/
package config

import (
	"reflect"
	"testing"

	"github.com/cgrates/cgrates/utils"
)

func TestHTTPCfgloadFromJsonCfg(t *testing.T) {
	cfgJSONStr := &HTTPJsonCfg{
		Json_rpc_url:              utils.StringPointer("/jsonrpc"),
		Ws_url:                    utils.StringPointer("/ws"),
		Dispatchers_registrar_url: utils.StringPointer("/randomUrl"),
		Freeswitch_cdrs_url:       utils.StringPointer("/freeswitch_json"),
		Http_Cdrs:                 utils.StringPointer("/cdr_http"),
		Use_basic_auth:            utils.BoolPointer(false),
		Auth_users:                utils.MapStringStringPointer(map[string]string{}),
	}
	expected := &HTTPCfg{
		HTTPJsonRPCURL:          "/jsonrpc",
		HTTPWSURL:               "/ws",
		DispatchersRegistrarURL: "/randomUrl",
		HTTPFreeswitchCDRsURL:   "/freeswitch_json",
		HTTPCDRsURL:             "/cdr_http",
		HTTPUseBasicAuth:        false,
		HTTPAuthUsers:           map[string]string{},
	}
	if cfgJsn, err := NewDefaultCGRConfig(); err != nil {
		t.Error(err)
	} else if err = cfgJsn.httpCfg.loadFromJsonCfg(cfgJSONStr); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(expected, cfgJsn.httpCfg) {
		t.Errorf("Expected %+v \n, received %+v", utils.ToJSON(expected), utils.ToJSON(cfgJsn.httpCfg))
	}
}

func TestHTTPCfgAsMapInterface(t *testing.T) {
	cfgJSONStr := `{
	"http": {},
}`
	eMap := map[string]interface{}{
		utils.HTTPJsonRPCURLCfg:          "/jsonrpc",
		utils.DispatchersRegistrarURLCfg: "/dispatchers_registrar",
		utils.HTTPWSURLCfg:               "/ws",
		utils.HTTPFreeswitchCDRsURLCfg:   "/freeswitch_json",
		utils.HTTPCDRsURLCfg:             "/cdr_http",
		utils.HTTPUseBasicAuthCfg:        false,
		utils.HTTPAuthUsersCfg:           map[string]interface{}{},
	}
	if cgrCfg, err := NewCGRConfigFromJsonStringWithDefaults(cfgJSONStr); err != nil {
		t.Error(err)
	} else if rcv := cgrCfg.httpCfg.AsMapInterface(); !reflect.DeepEqual(rcv, eMap) {
		t.Errorf("Expected %+v, received %+v", eMap, rcv)
	}
}

func TestHTTPCfgAsMapInterface1(t *testing.T) {
	cfgJSONStr := `{
	"http": {
       "json_rpc_url": "/rpc",					
	   "ws_url": "",	
	   "use_basic_auth": true,					
	   "auth_users": {"user1": "authenticated", "user2": "authenticated"},
     },
}`
	eMap := map[string]interface{}{
		utils.HTTPJsonRPCURLCfg:          "/rpc",
		utils.DispatchersRegistrarURLCfg: "/dispatchers_registrar",
		utils.HTTPWSURLCfg:               "",
		utils.HTTPFreeswitchCDRsURLCfg:   "/freeswitch_json",
		utils.HTTPCDRsURLCfg:             "/cdr_http",
		utils.HTTPUseBasicAuthCfg:        true,
		utils.HTTPAuthUsersCfg: map[string]interface{}{
			"user1": "authenticated",
			"user2": "authenticated",
		},
	}
	if cgrCfg, err := NewCGRConfigFromJsonStringWithDefaults(cfgJSONStr); err != nil {
		t.Error(err)
	} else if rcv := cgrCfg.httpCfg.AsMapInterface(); !reflect.DeepEqual(rcv, eMap) {
		t.Errorf("Expected %+v, received %+v", eMap, rcv)
	}
}
