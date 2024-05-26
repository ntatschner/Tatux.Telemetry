package domain

import (
	"time"
)

// TelemetryService provides an interface for telemetry services.

type ExecutionData struct {
	ID                string        `json:"id"`
	LocalDateTime     time.Time     `json:"localDateTime"`
	ExecutionDuration time.Duration `json:"executionDuration"`
	ExecutionID       string        `json:"executionId"`
	Failed            bool          `json:"failed"`
	Exception         string        `json:"exception"`
}

type PowerShellData struct {
	ExecutionData
	CommandName               string `json:"commandName"`
	Manufacturer              string `json:"manufacturer"`
	Model                     string `json:"model"`
	TotalPhysicalMemory       int    `json:"totalPhysicalMemory"`
	NumberOfProcessors        int    `json:"numberOfProcessors"`
	NumberOfLogicalProcessors int    `json:"numberOfLogicalProcessors"`
	PartOfDomain              bool   `json:"partOfDomain"`
	HardwareSerialNumber      string `json:"hardwareSerialNumber"`
	BootDriveSerial           string `json:"bootDriveSerial"`
	OSType                    string `json:"osType"`
	OSArchitecture            string `json:"osArchitecture"`
	OSVersion                 string `json:"osVersion"`
	OSBuildNumber             string `json:"osBuildNumber"`
	PowerShellVersion         string `json:"powerShellVersion"`
	HostVersion               string `json:"hostVersion"`
	HostName                  string `json:"hostName"`
	HostUI                    string `json:"hostUI"`
	HostCulture               string `json:"hostCulture"`
	HostUICulture             string `json:"hostUICulture"`
	ModuleName                string `json:"moduleName"`
	ModuleVersion             string `json:"moduleVersion"`
	ModulePath                string `json:"modulePath"`
}

type PipelineData struct {
	ExecutionData
	PipelineName          string `json:"pipelineName"`
	Exception             string `json:"exception"`
	RunnerOS              string `json:"runnerOs"`
	RunnerArchitecture    string `json:"runnerArchitecture"`
	SourceControlProvider string `json:"sourceControlProvider"`
}

type ServiceHealthData struct {
	ID              string
	ServiceName     string
	ServiceEndpoint string
	Uptime          time.Duration
	CheckInterval   time.Duration
	LastCheck       time.Time
	NextCheck       time.Time
	Failed          bool
	Exception       string
	FailureCount    int
	LocalDateTime   time.Time
}
