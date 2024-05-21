package domain

import (
	"context"
	"time"
	log "telemetry/pkg/logging"
)

// TelemetryService provides an interface for telemetry services.

type PowerShellData struct {
	ID                        string        
	CommandName               string        
	LocalDateTime             time.Time     
	ExecutionDuration         time.Duration 
	ExecutionID               string        
	Failed                    bool          
	Exception                 string        
	Manufacturer              string        
	Model                     string        
	TotalPhysicalMemory       int           
	NumberOfProcessors        int           
	NumberOfLogicalProcessors int           
	PartOfDomain              bool          
	HardwareSerialNumber      string        
	BootDriveSerial           string        
	OSType                    string        
	OSArchitecture            string        
	OSVersion                 string        
	OSBuildNumber             string        
	PowerShellVersion         string        
	HostVersion               string        
	HostName                  string        
	HostUI                    string        
	HostCulture               string        
	HostUICulture             string        
	ModuleName                string        
	ModuleVersion             string        
	ModulePath                string        
}

type PipelineData struct {
	ID                        string
	PipelineName			  string
	LocalDateTime             time.Time
	ExecutionDuration         time.Duration
	ExecutionID               string
	Failed                    bool
	Exception                 string
	RunnerOS				  string
	RunnerArchitecture		  string
	SourceControlProvider	  string
}

type ServiceHealthData struct {
	ID                        string
	ServiceName				  string
	ServiceEndpoint			  string
	Uptime					  time.Duration
	CheckInterval			  time.Duration
	LastCheck				  time.Time
	NextCheck				  time.Time
	Failed					  bool
	Exception				  string
	FailureCount			  int
	LocalDateTime             time.Time
}