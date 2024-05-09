function Invoke-TelemetryCollection {
    param (
        [CmdletBinding(HelpUri = 'https://pwsh.dev.tatux.co.uk/tatux.telemetry/docs/Invoke-TelemetryCollection.html')]

        [string]$ModuleName = 'UnknownModule',

        [string]$ModuleVersion = 'UnknownModuleVersion',

        [string]$ModulePath = 'UnknownModulePath',

        [string]$CommandName = 'UnknownCommand',

        [Parameter(Mandatory = $true)]
        [ValidateSet('Start', 'In-Progress', 'End', 'Module-Load')]
        [string]$Stage,

        [bool]$Failed = $false,

        [string]$Exception,

        [switch]$Minimal,

        [switch]$ClearTimer,

        [Parameter(Mandatory = $true)]
        [string]$URI
    )

    $CurrentTime = (Get-Date).ToString("yyyy-MM-ddTHH:mm:sszzz")

    $WebRequestArgs = @{
        Uri             = $URI
        Method          = 'Put'
        ContentType     = 'application/json'
        UseBasicParsing = $true
    }

    switch ($Stage) {
        'Module-Load' {
            if ((Get-Variable -Name 'GlobalExecutionDuration' -Scope script -ErrorAction SilentlyContinue) -and (-Not $ClearTimer)) {
                $script:GlobalExecutionDuration = $GlobalExecutionDuration
            }
            else {
                $script:GlobalExecutionDuration = Get-Date
            }
        }
        'Start' {
            if ((Get-Variable -Name 'GlobalExecutionDuration' -Scope script -ErrorAction SilentlyContinue) -and (-Not $ClearTimer)) {
                $script:GlobalExecutionDuration = $GlobalExecutionDuration
            }
            else {
                $script:GlobalExecutionDuration = Get-Date
            }
        }
        'In-Progress' {
            
        }
        'End' {
            # Generate hardware specific but none identifying telemetry data for the output
            $Hardware = Get-WmiObject -Class Win32_ComputerSystem
            $bootPartition = Get-WmiObject -Class Win32_DiskPartition | Where-Object -Property bootpartition -eq True
            $bootDriveSerial = $(Get-WmiObject -Class Win32_DiskDrive | Where-Object -Property index -eq $bootPartition.diskIndex)
            if ([string]::IsNullOrEmpty($bootDriveSerial.SerialNumber) -and ($bootDriveSerial.Model -like '*Virtual*')) {
                $bootDriveSerial = "VirtualDrive-$($bootDriveSerial.size)"
            }
            else {
                $bootDriveSerial = $bootDriveSerial.SerialNumber.Trim()
            }

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
                OSVersion      = $OS.Version
                OSBuildNumber  = $OS.BuildNumber
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
                ModuleName    = if ([string]::IsNullOrEmpty($ModuleName)) { 'UnknownModule' } else { $ModuleName }
                ModuleVersion = if ([string]::IsNullOrEmpty($ModuleVersion)) { 'UnknownModuleVersion' } else { $ModuleVersion }
                ModulePath    = if ([string]::IsNullOrEmpty($ModulePath)) { 'UnknownModulePath' } else { $ModulePath }
                CommandName   = if ([string]::IsNullOrEmpty($CommandName)) { 'UnknownCommand' } else { $CommandName }
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
            $AllData += @{Exception = $Exception.ToString() }
            if ($Minimal) {
                $AllData | ForEach-Object {
                    if ($_.Name -notin @('ID', 'CommandName', 'ModuleName', 'ModuleVersion', 'LocalDateTime', 'ExecutionDuration', 'Stage', 'Failed')) {
                        $_.Value = 'Minimal'
                    }
                    $body = $AllData | ConvertTo-Json
                    Invoke-WebRequest @WebRequestArgs -Body $body > $null
                }
            }
            else {
                $body = $AllData | ConvertTo-Json
                Invoke-WebRequest @WebRequestArgs -Body $body > $null
            }        
        }
    }
}