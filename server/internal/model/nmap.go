package model

type NmapScanParameters struct {
	TargetSpecification []string                 `json:"targetSpecification"`
	HostDiscovery       HostDiscoveryOptions     `json:"hostDiscovery,omitempty"`
	PortSpecification   PortSpecificationOptions `json:"portSpecification,omitempty"`
	MiscOptions         MiscOptions              `json:"miscOptions,omitempty"`
}

type HostDiscoveryOptions struct {
	ListScan          bool `json:"listScan,omitempty"`
	PingScan          bool `json:"pingScan,omitempty"`
	SkipHostDiscovery bool `json:"skipHostDiscovery,omitempty"`
}

type PortSpecificationOptions struct {
	Ports          string `json:"ports,omitempty"`
	ExcludePorts   string `json:"excludePorts,omitempty"`
	FastMode       bool   `json:"fastMode,omitempty"`
	SequentialScan bool   `json:"sequentialScan,omitempty"`
}

type MiscOptions struct {
	EnableIPv6Scan         bool `json:"enableIPv6Scan,omitempty"`
	EnableOSDetection      bool `json:"enableOSDetection,omitempty"`
	EnableVersionDetection bool `json:"enableVersionDetection,omitempty"`
	EnableScriptScanning   bool `json:"enableScriptScanning,omitempty"`
	EnableTraceroute       bool `json:"enableTraceroute,omitempty"`
	PrintVersionNumber     bool `json:"printVersionNumber,omitempty"`
}

func (p *NmapScanParameters) ToCmd() []string {
	var cmdArray []string

	cmdArray = append(cmdArray, "-oX", "-")

	if p.HostDiscovery.ListScan {
		cmdArray = append(cmdArray, "-sL")
	}
	if p.HostDiscovery.PingScan {
		cmdArray = append(cmdArray, "-sn")
	}
	if p.HostDiscovery.SkipHostDiscovery {
		cmdArray = append(cmdArray, "-Pn")
	}

	if p.PortSpecification.Ports != "" {
		cmdArray = append(cmdArray, "-p", p.PortSpecification.Ports)
	}
	if p.PortSpecification.ExcludePorts != "" {
		cmdArray = append(cmdArray, "--exclude-ports", p.PortSpecification.ExcludePorts)
	}
	if p.PortSpecification.FastMode {
		cmdArray = append(cmdArray, "-F")
	}
	if p.PortSpecification.SequentialScan {
		cmdArray = append(cmdArray, "-r")
	}

	if p.MiscOptions.EnableIPv6Scan {
		cmdArray = append(cmdArray, "-6")
	}
	if p.MiscOptions.EnableOSDetection {
		cmdArray = append(cmdArray, "-A")
	}
	if p.MiscOptions.PrintVersionNumber {
		cmdArray = append(cmdArray, "-V")
	}

	cmdArray = append(cmdArray, p.TargetSpecification...)

	return cmdArray
}
