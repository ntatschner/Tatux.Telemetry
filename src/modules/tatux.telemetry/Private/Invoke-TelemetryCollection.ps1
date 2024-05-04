function Invoke-TelemetryCollection {
    param (
        [CmdletBinding(OutputType = 'json', HelpUri = 'https://pwsh.dev.tatux.co.uk/tatux.telemetry/docs/Invoke-TelemetryCollection.html')]
        [Parameter(Mandatory = $true)]
        [array]$ExecutionContext,

        [Parameter(Mandatory = $true)]
        [switch]$Minimal
    )

    # Generate hardware specific but none identifying telemetry data for the output
    $Hardware = Get-WmiObject -Class Win32_ComputerSystem
    $HardwareData = @{
        Manufacturer = $Hardware.Manufacturer
        Model = $Hardware.Model
        TotalPhysicalMemory = $Hardware.TotalPhysicalMemory
        NumberOfProcessors = $Hardware.NumberOfProcessors
        NumberOfLogicalProcessors = $Hardware.NumberOfLogicalProcessors
    }

    # Generate OS specific but none identifying telemetry data for the output
    $OS = Get-WmiObject -Class Win32_OperatingSystem

    $OSData = @{
        OSArchitecture = $OS.OSArchitecture
        Version = $OS.Version
        BuildNumber = $OS.BuildNumber
        ServicePackMajorVersion = $OS.ServicePackMajorVersion
        ServicePackMinorVersion = $OS.ServicePackMinorVersion
    }

    # Generate PowerShell specific but none identifying telemetry data for the output

    $PSData = @{
        PowerShellVersion = $PSVersionTable.PSVersion
        HostVersion = $Host.Version
        HostName = $Host.Name
        HostUI = $Host.UI
        HostCulture = $Host.CurrentCulture
        HostUICulture = $Host.CurrentUICulture
    }

    # Generate module specific but none identifying telemetry data for the output

    $ModuleData = @{
        ModuleName = $ExecutionContext.ModuleName
        ModuleVersion = $ExecutionContext.ModuleVersion
        ModulePath = $ExecutionContext.ModulePath
    }
# Create a new hashtable
$AllData = @{}

# Add each hashtable to the new hashtable
$AllData += $HardwareData
$AllData += $OSData
$AllData += $PSData
$AllData += $ModuleData

    # Generate the telemetry data

    
    
}