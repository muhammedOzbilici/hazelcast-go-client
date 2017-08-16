package config

import (
	. "github.com/hazelcast/go-client/internal/serialization/api"
)

const (
	DEFAULT_GROUP_NAME     = "dev"
	DEFAULT_GROUP_PASSWORD = "dev-pass"
)

type ClientConfig struct {
	GroupConfig         GroupConfig
	ClientNetworkConfig ClientNetworkConfig
}

type SerializationConfig struct {
	IsBigEndian               bool
	DataSerializableFactories map[int32]IdentifiedDataSerializableFactory
	// portableFactories map[int32]
	PortableVersion int32
	//customSerializers []
	//globalSerializer
}

func NewSerializationConfig() *SerializationConfig {
	return &SerializationConfig{IsBigEndian: true, DataSerializableFactories: make(map[int32]IdentifiedDataSerializableFactory), PortableVersion: 0}
}

func (c *SerializationConfig) AddDataSerializableFactory(f IdentifiedDataSerializableFactory, factoryId int32) {
	c.DataSerializableFactories[factoryId] = f
}

func NewClientConfig() *ClientConfig {
	return &ClientConfig{GroupConfig: NewGroupConfig(),
		ClientNetworkConfig: NewClientNetworkConfig(),
	}
}
func (clientConfig *ClientConfig) IsSmartRouting() bool {
	return clientConfig.ClientNetworkConfig.SmartRouting
}

type GroupConfig struct {
	Name     string
	Password string
}

func NewGroupConfig() GroupConfig {
	return GroupConfig{Name: DEFAULT_GROUP_NAME, Password: DEFAULT_GROUP_PASSWORD}
}

type ClientNetworkConfig struct {
	Addresses *[]Address
	//The candidate address list that client will use to establish initial connection
	ConnectionAttemptLimit int32
	/*
		While client is trying to connect initially to one of the members in the addressList, all might be not
		available. Instead of giving up, throwing Error and stopping client, it will attempt to retry as much as defined
		by this parameter.
	*/
	ConnectionAttemptPeriod int32
	//Period for the next attempt to find a member to connect
	ConnectionTimeout int32
	/*
			Socket connection timeout is a float, giving in seconds, or None.
		    Setting a timeout of None disables the timeout feature and is equivalent to block the socket until it connects.
		    Setting a timeout of zero is the same as disables blocking on connect.
	*/
	RedoOperations bool
	/*
		If true, client will redo the operations that were executing on the server and client lost the connection.
		This can be because of network, or simply because the member died. However it is not clear whether the
		application is performed or not. For idempotent operations this is harmless, but for non idempotent ones
		retrying can cause to undesirable effects. Note that the redo can perform on any member.
	*/
	SmartRouting bool
	/*
		If true, client will route the key based operations to owner of the key at the best effort. Note that it uses a
		cached value of partition count and doesn't guarantee that the operation will always be executed on the owner.
		The cached table is updated every 10 seconds.
	*/
}

func NewClientNetworkConfig() ClientNetworkConfig {
	return ClientNetworkConfig{
		Addresses:new([]Address),
		ConnectionAttemptLimit:2,
		ConnectionAttemptPeriod:3,
		ConnectionTimeout:5.0,
		RedoOperations:false,
		SmartRouting:true,
	}
}
type Address struct {
	Host string
	Port int32
}
