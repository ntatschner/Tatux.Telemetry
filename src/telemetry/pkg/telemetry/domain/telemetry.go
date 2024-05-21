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