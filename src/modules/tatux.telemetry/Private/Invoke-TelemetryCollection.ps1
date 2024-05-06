function Invoke-TelemetryCollection {
    param (
        [CmdletBinding(HelpUri = 'https://pwsh.dev.tatux.co.uk/tatux.telemetry/docs/Invoke-TelemetryCollection.html')]

        [Parameter(Mandatory = $false)]
        [array]$ExecutionContextInput,

        [Parameter(Mandatory = $true)]
        [ValidateSet('Start', 'In-Progrss', 'End')]
        [string]$Stage,

        [bool]$Failed = $false,

        [switch]$Minimal,

        [switch]$ClearTimer,

        [string]$URI = 'http://localhost:9000/api/telemetry'
    )
    if ((Get-Variable -Name 'GlobalExecutionDuration' -Scope script -ErrorAction SilentlyContinue) -and (-Not $ClearTimer)) {
        $script:GlobalExecutionDuration = $GlobalExecutionDuration
    }
    else {
        $script:GlobalExecutionDuration = Get-Date
    }
    $CurrentTime = (Get-Date).ToString("yyyy-MM-ddTHH:mm:sszzz")

    # Generate hardware specific but none identifying telemetry data for the output
    $Hardware = Get-WmiObject -Class Win32_ComputerSystem
    $bootPartition = Get-WmiObject -Class Win32_DiskPartition | Where-Object -Property bootpartition -eq True
    $bootDriveSerial = $(Get-WmiObject -Class Win32_DiskDrive | Where-Object -Property index -eq $bootPartition.diskIndex | Select-Object -ExpandProperty SerialNumber).Trim()
    $HardwareData = @{
        Manufacturer              = $Hardware.Manufacturer
        Model                     = $Hardware.Model
        TotalPhysicalMemory       = $Hardware.TotalPhysicalMemory
        NumberOfProcessors        = $Hardware.NumberOfProcessors
        NumberOfLogicalProcessors = $Hardware.NumberOfLogicalProcessors
        PartOfDomain              = $Hardware.PartOfDomain
        HardwareSerialNumber      = $((Get-WmiObject -Class Win32_BIOS).SerialNumber)
        BootDriveSerial           = $bootDriveSerial
    }

    # Generate OS specific but none identifying telemetry data for the output
    $OS = Get-WmiObject -Class Win32_OperatingSystem

    $OSData = @{
        OSType         = $OS.Caption
        OSArchitecture = $OS.OSArchitecture
        Version        = $OS.Version
        BuildNumber    = $OS.BuildNumber
        SerialNumber   = $OS.SerialNumber
    }

    # Generate PowerShell specific but none identifying telemetry data for the output

    $PSData = @{
        PowerShellVersion = $PSVersionTable.PSVersion.ToString()
        HostVersion       = $Host.Version.ToString()
        HostName          = $Host.Name.ToString()
        HostUI            = $Host.UI.ToString()
        HostCulture       = $Host.CurrentCulture.ToString()
        HostUICulture     = $Host.CurrentUICulture.ToString()
    }

    # Generate module specific but none identifying telemetry data for the output

    $ModuleData = @{
        ModuleName    = $ExecutionContextInput.ModuleName
        ModuleVersion = $ExecutionContextInput.ModuleVersion
        ModulePath    = $ExecutionContextInput.ModulePath
        CommandName   = $ExecutionContextInput.CommandName
    }
    # Create a new hashtable
    $AllData = @{}

    # Add each hashtable to the new hashtable
    $AllData += $HardwareData
    $AllData += $OSData
    $AllData += $PSData
    $AllData += $ModuleData
    $AllData += @{ID = $AllData.BootDriveSerial + "_" + $AllData.SerialNumber } 
    $AllData += @{LocalDateTime = $CurrentTime }
    $AllData += @{ExecutionDuration = [Int64]$($(New-TimeSpan -Start $script:GlobalExecutionDuration -End $(Get-Date)).TotalMilliseconds * 1e6) }
    $AllData += @{Stage = $Stage }
    $AllData += @{Failed = $Failed }
    # Generate the telemetry data
    Write-Output $AllData
    if ($Minimal) {
        $AllData | ForEach-Object {
            if ($_.Name -notin @('ID', 'CommandName', 'ModuleName', 'ModuleVersion', 'LocalDateTime', 'ExecutionDuration', 'Stage', 'Failed')) {
                $_.Value = 'Minimal'
            }
            $body = $AllData | ConvertTo-Json
            Invoke-WebRequest -Uri $URI -Method Put -Body $body -ContentType 'application/json'
        }
    }
    else {
        $body = $AllData | ConvertTo-Json
        Invoke-WebRequest -Uri $URI -Method Put -Body $body -ContentType 'application/json'
    }

}