package types

const (
	// BridgeNetworkDriver defines the bridge driver
	BridgeNetworkDriver = "bridge"
	// DefaultNetworkDriver is the default network type used
	DefaultNetworkDriver = BridgeNetworkDriver
	// MacVLANNetworkDriver defines the macvlan driver
	MacVLANNetworkDriver = "macvlan"
	// MacVLANNetworkDriver defines the macvlan driver
	IPVLANNetworkDriver = "ipvlan"

	// IPAM drivers
	// HostLocalIPAMDriver store the ip
	HostLocalIPAMDriver = "host-local"
	// DHCPIPAMDriver get subnet and ip from dhcp server
	DHCPIPAMDriver = "dhcp"

	// DefaultSubnet is the name that will be used for the default CNI network.
	DefaultNetworkName = "podman"
	// DefaultSubnet is the subnet that will be used for the default CNI network.
	DefaultSubnet = "10.88.0.0/16"

	// valid macvlan driver mode values
	MacVLANModeBridge   = "bridge"
	MacVLANModePrivate  = "private"
	MacVLANModeVepa     = "vepa"
	MacVLANModePassthru = "passthru"

	// valid ipvlan driver modes
	IPVLANModeL2  = "l2"
	IPVLANModeL3  = "l3"
	IPVLANModeL3s = "l3s"
)

// ValidMacVLANModes is the list of valid mode options for the macvlan driver
var ValidMacVLANModes = []string{MacVLANModeBridge, MacVLANModePrivate, MacVLANModeVepa, MacVLANModePassthru}

// ValidIPVLANModes is the list of valid mode options for the ipvlan driver
var ValidIPVLANModes = []string{IPVLANModeL2, IPVLANModeL3, IPVLANModeL3s}
