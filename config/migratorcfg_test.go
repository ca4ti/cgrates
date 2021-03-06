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

func TestMigratorCgrCfgloadFromJsonCfg(t *testing.T) {
	cfgJSON := &MigratorCfgJson{
		Out_dataDB_type:     utils.StringPointer(utils.REDIS),
		Out_dataDB_host:     utils.StringPointer("127.0.0.1"),
		Out_dataDB_port:     utils.StringPointer("6379"),
		Out_dataDB_name:     utils.StringPointer("10"),
		Out_dataDB_user:     utils.StringPointer(utils.CGRATES),
		Out_dataDB_password: utils.StringPointer(utils.EmptyString),
		Out_dataDB_encoding: utils.StringPointer(utils.MSGPACK),
		Out_storDB_type:     utils.StringPointer(utils.MYSQL),
		Out_storDB_host:     utils.StringPointer("127.0.0.1"),
		Out_storDB_port:     utils.StringPointer("3306"),
		Out_storDB_name:     utils.StringPointer(utils.CGRATES),
		Out_storDB_user:     utils.StringPointer(utils.CGRATES),
		Out_storDB_password: utils.StringPointer(utils.EmptyString),
		Out_dataDB_opts: map[string]interface{}{
			utils.RedisClusterCfg: true,
			utils.ClusterSyncCfg:  "10s",
		},
		Out_storDB_opts: map[string]interface{}{},
	}
	expected := &MigratorCgrCfg{
		OutDataDBType:     utils.REDIS,
		OutDataDBHost:     "127.0.0.1",
		OutDataDBPort:     "6379",
		OutDataDBName:     "10",
		OutDataDBUser:     utils.CGRATES,
		OutDataDBPassword: utils.EmptyString,
		OutDataDBEncoding: utils.MSGPACK,
		OutStorDBType:     utils.MYSQL,
		OutStorDBHost:     "127.0.0.1",
		OutStorDBPort:     "3306",
		OutStorDBName:     utils.CGRATES,
		OutStorDBUser:     utils.CGRATES,
		OutStorDBPassword: utils.EmptyString,
		OutDataDBOpts: map[string]interface{}{
			utils.RedisSentinelNameCfg:  utils.EmptyString,
			utils.RedisClusterCfg:       true,
			utils.ClusterSyncCfg:        "10s",
			utils.ClusterOnDownDelayCfg: "0",
		},
		OutStorDBOpts: map[string]interface{}{},
	}
	if cfgJson, err := NewDefaultCGRConfig(); err != nil {
		t.Error(err)
	} else if err = cfgJson.migratorCgrCfg.loadFromJsonCfg(cfgJSON); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(expected, cfgJson.migratorCgrCfg) {
		t.Errorf("Expected %+v \n, received %+v", utils.ToJSON(expected), utils.ToJSON(cfgJson.migratorCgrCfg))
	}
}

func TestMigratorCgrCfgAsMapInterface(t *testing.T) {
	cfgJSONStr := `{
	"migrator": {
		"out_datadb_host": "127.0.0.19",
		"out_datadb_port": "8865",
		"out_datadb_name": "12",
		"out_stordb_host": "127.0.0.19",
		"out_stordb_port": "1234",
        "users_filters":["users","filters","Account"],
        "out_datadb_opts":{	
		   "redis_cluster": true,					
		   "cluster_sync": "2s",					
		   "cluster_ondown_delay": "1",	
	    },
	},
}`
	eMap := map[string]interface{}{
		utils.OutDataDBTypeCfg:     "redis",
		utils.OutDataDBHostCfg:     "127.0.0.19",
		utils.OutDataDBPortCfg:     "8865",
		utils.OutDataDBNameCfg:     "12",
		utils.OutDataDBUserCfg:     "cgrates",
		utils.OutDataDBPasswordCfg: "",
		utils.OutDataDBEncodingCfg: "msgpack",
		utils.OutStorDBTypeCfg:     "mysql",
		utils.OutStorDBHostCfg:     "127.0.0.19",
		utils.OutStorDBPortCfg:     "1234",
		utils.OutStorDBNameCfg:     "cgrates",
		utils.OutStorDBUserCfg:     "cgrates",
		utils.OutStorDBPasswordCfg: "",
		utils.UsersFiltersCfg:      []string{"users", "filters", "Account"},
		utils.OutStorDBOptsCfg:     map[string]interface{}{},
		utils.OutDataDBOptsCfg: map[string]interface{}{
			utils.RedisSentinelNameCfg:  "",
			utils.RedisClusterCfg:       true,
			utils.ClusterSyncCfg:        "2s",
			utils.ClusterOnDownDelayCfg: "1",
		},
	}
	if cgrCfg, err := NewCGRConfigFromJsonStringWithDefaults(cfgJSONStr); err != nil {
		t.Error(err)
	} else if rcv := cgrCfg.migratorCgrCfg.AsMapInterface(); !reflect.DeepEqual(eMap, rcv) {
		t.Errorf("Expected %+v \n, received %+v", utils.ToJSON(eMap), utils.ToJSON(rcv))
	}
}
func TestMigratorCgrCfgAsMapInterface1(t *testing.T) {
	cfgJSONStr := `{
		"migrator": {
			"out_stordb_password": "out_stordb_password",
			"users_filters":["users","filters","Account"],
			"out_datadb_opts": {
				"redis_sentinel": "out_datadb_redis_sentinel",
			},
		},
	}`
	eMap := map[string]interface{}{
		utils.OutDataDBTypeCfg:     "redis",
		utils.OutDataDBHostCfg:     "127.0.0.1",
		utils.OutDataDBPortCfg:     "6379",
		utils.OutDataDBNameCfg:     "10",
		utils.OutDataDBUserCfg:     "cgrates",
		utils.OutDataDBPasswordCfg: "",
		utils.OutDataDBEncodingCfg: "msgpack",
		utils.OutStorDBTypeCfg:     "mysql",
		utils.OutStorDBHostCfg:     "127.0.0.1",
		utils.OutStorDBPortCfg:     "3306",
		utils.OutStorDBNameCfg:     "cgrates",
		utils.OutStorDBUserCfg:     "cgrates",
		utils.OutStorDBPasswordCfg: "out_stordb_password",
		utils.UsersFiltersCfg:      []string{"users", "filters", "Account"},
		utils.OutStorDBOptsCfg:     map[string]interface{}{},
		utils.OutDataDBOptsCfg: map[string]interface{}{
			utils.RedisSentinelNameCfg:  "out_datadb_redis_sentinel",
			utils.RedisClusterCfg:       false,
			utils.ClusterSyncCfg:        "5s",
			utils.ClusterOnDownDelayCfg: "0",
		},
	}
	if cgrCfg, err := NewCGRConfigFromJsonStringWithDefaults(cfgJSONStr); err != nil {
		t.Error(err)
	} else if rcv := cgrCfg.migratorCgrCfg.AsMapInterface(); !reflect.DeepEqual(eMap, rcv) {
		t.Errorf("Expected %+v \n, received %+v", utils.ToJSON(eMap), utils.ToJSON(rcv))
	}
}

func TestMigratorCgrCfgAsMapInterface2(t *testing.T) {
	cfgJSONStr := `{
		"migrator": {},
	}`
	eMap := map[string]interface{}{
		utils.OutDataDBTypeCfg:     "redis",
		utils.OutDataDBHostCfg:     "127.0.0.1",
		utils.OutDataDBPortCfg:     "6379",
		utils.OutDataDBNameCfg:     "10",
		utils.OutDataDBUserCfg:     "cgrates",
		utils.OutDataDBPasswordCfg: "",
		utils.OutDataDBEncodingCfg: "msgpack",
		utils.OutStorDBTypeCfg:     "mysql",
		utils.OutStorDBHostCfg:     "127.0.0.1",
		utils.OutStorDBPortCfg:     "3306",
		utils.OutStorDBNameCfg:     "cgrates",
		utils.OutStorDBUserCfg:     "cgrates",
		utils.OutStorDBPasswordCfg: "",
		utils.UsersFiltersCfg:      []string{},
		utils.OutStorDBOptsCfg:     map[string]interface{}{},
		utils.OutDataDBOptsCfg: map[string]interface{}{
			utils.RedisSentinelNameCfg:  "",
			utils.RedisClusterCfg:       false,
			utils.ClusterSyncCfg:        "5s",
			utils.ClusterOnDownDelayCfg: "0",
		},
	}
	if cgrCfg, err := NewCGRConfigFromJsonStringWithDefaults(cfgJSONStr); err != nil {
		t.Error(err)
	} else if rcv := cgrCfg.migratorCgrCfg.AsMapInterface(); !reflect.DeepEqual(eMap, rcv) {
		t.Errorf("Expected %+v \n, received %+v", utils.ToJSON(eMap), utils.ToJSON(rcv))
	}

}
