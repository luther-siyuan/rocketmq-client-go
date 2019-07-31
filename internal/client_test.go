/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package internal

import (
	"context"
	"github.com/apache/rocketmq-client-go/internal/remote"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const (
	testTopic  = "TestTopic"
	testGroup = "TestGroup"
)

var defaultOptions = ClientOptions{
	GroupName:       testGroup,
	NameServerAddrs: []string{"127.0.0.1:9876"},
	ClientIP:        "127.0.0.1",
	InstanceName:    "siYuan",
}

func TestPullMessage(t *testing.T) {
	Convey("Given a starting client", t, func() {
		client := GetOrNewRocketMQClient(defaultOptions)
		req := &PullMessageRequest{
			ConsumerGroup:  "please_rename_unique_group_name_4",
			Topic:          testTopic,
			QueueId:        0,
			QueueOffset:    0,
			MaxMsgNums:     32,
			SysFlag:        0x1 << 2,
			SubExpression:  "*",
			ExpressionType: "TAG",
		}
		_, err := client.PullMessage(context.Background(), "127.0.0.1:10911", req)
		So(err, ShouldBeNil)
	})
}

func TestGetOrNewRocketMQClient(t *testing.T) {
	Convey("Given a starting client", t, func() {
		client := &rmqClient{
			option:       defaultOptions,
			remoteClient: remote.NewRemotingClient(),
		}
		rmqClient := GetOrNewRocketMQClient(defaultOptions)

		Convey("client from clientMap should not be nil", func() {
			expectedClient, ok := clientMap.Load(client.ClientID())
			So(ok, ShouldBeTrue)
			So(expectedClient, ShouldNotBeNil)
		})

		Convey("processor ReqNotifyConsumerIdsChanged have been registered", func() {
			So(rmqClient.remoteClient, ShouldNotEqual, client)
			t.Logf("rmqClient: %+v", rmqClient)
		})
	})
}
