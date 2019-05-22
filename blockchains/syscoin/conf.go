/*
	Copyright 2019 whiteblock Inc.
	This file is a part of the genesis.

	Genesis is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    Genesis is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package syscoin

import (
	"encoding/json"
	"github.com/whiteblock/genesis/blockchains/helpers"
	"github.com/whiteblock/genesis/util"
	"log"
)

type sysConf struct {
	Options []string `json:"options"`
	Extras  []string `json:"extras"`

	SenderOptions   []string `json:"senderOptions"`
	ReceiverOptions []string `json:"receiverOptions"`
	MNOptions       []string `json:"mnOptions"`

	SenderExtras   []string `json:"senderExtras"`
	ReceiverExtras []string `json:"receiverExtras"`
	MNExtras       []string `json:"mnExtras"`

	MasterNodeConns int64 `json:"masterNodeConns"`
	NodeConns       int64 `json:"nodeConns"`
	PercOfMNodes    int64 `json:"percentMasternodes"`
	Validators      int64 `json:"validators"`
}

func newConf(data map[string]interface{}) (*sysConf, error) {
	out := new(sysConf)
	err := json.Unmarshal([]byte(GetDefaults()), out)

	if data == nil {
		log.Println("No params given")
		return out, nil
	}

	err = util.GetJSONStringArr(data, "options", &out.Options)
	if err != nil {
		return nil, err
	}

	err = util.GetJSONStringArr(data, "extras", &out.Extras)
	if err != nil {
		return nil, err
	}

	err = util.GetJSONStringArr(data, "senderOptions", &out.SenderOptions)
	if err != nil {
		return nil, err
	}

	err = util.GetJSONStringArr(data, "senderExtras", &out.SenderExtras)
	if err != nil {
		return nil, err
	}

	err = util.GetJSONStringArr(data, "senderExtras", &out.SenderExtras)
	if err != nil {
		return nil, err
	}

	err = util.GetJSONStringArr(data, "receiverOptions", &out.ReceiverOptions)
	if err != nil {
		return nil, err
	}

	err = util.GetJSONStringArr(data, "receiverExtras", &out.ReceiverExtras)
	if err != nil {
		return nil, err
	}

	err = util.GetJSONStringArr(data, "mnOptions", &out.MNOptions)
	if err != nil {
		return nil, err
	}

	err = util.GetJSONStringArr(data, "mnExtras", &out.MNExtras)
	if err != nil {
		return nil, err
	}

	err = util.GetJSONInt64(data, "masterNodeConns", &out.MasterNodeConns)
	if err != nil {
		return nil, err
	}

	err = util.GetJSONInt64(data, "nodeConns", &out.NodeConns)
	if err != nil {
		return nil, err
	}

	err = util.GetJSONInt64(data, "percentMasternodes", &out.PercOfMNodes)
	if err != nil {
		return nil, err
	}
	log.Printf("%+v\n", *out)
	return out, nil
}

func (sconf *sysConf) Generate() string {
	out := ""
	for _, opt := range sconf.Options {
		out += opt + "=1\n"
	}
	out += "[regtest]\n"
	out += "rpcuser=user\n"
	out += "rpcpassword=password\n"
	for _, extra := range sconf.Extras {
		extra += extra + "\n"
	}

	return out
}

func (sconf *sysConf) GenerateReceiver() string {
	out := sconf.Generate()

	for _, opt := range sconf.ReceiverOptions {
		out += opt + "=1\n"
	}

	for _, extra := range sconf.ReceiverExtras {
		extra += extra + "\n"
	}
	return out
}

func (sconf *sysConf) GenerateSender() string {
	out := sconf.Generate()

	for _, opt := range sconf.SenderOptions {
		out += opt + "=1\n"
	}

	for _, extra := range sconf.SenderExtras {
		extra += extra + "\n"
	}
	return out
}

func (sconf *sysConf) GenerateMN() string {
	out := sconf.Generate()

	for _, opt := range sconf.MNOptions {
		out += opt + "=1\n"
	}

	for _, extra := range sconf.MNExtras {
		extra += extra + "\n"
	}
	return out
}

// GetDefaults fetchs syscoin related parameter defaults
func GetDefaults() string {
	dat, err := helpers.GetStaticBlockchainConfig(blockchain, "defaults.json")
	if err != nil {
		panic(err) //Missing required files is a fatal error
	}
	return string(dat)
}

// GetServices returns the services which are used by syscoin
func GetServices() []util.Service {
	return nil
	/*return []util.Service{
	    util.Service{
	        Name:"Alpine",
	        Image:"alpine:latest",
	        Env:map[string]string{
	            "HELLO":"HI",
	            "INFLUXDB_URL":conf.Influx,
	        },
	    },
	}*/
}
