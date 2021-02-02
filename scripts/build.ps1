$SCRIPT_LOCATION = Get-Location
$PROJECT_ROOT = Split-Path $SCRIPT_LOCATION -Parent
$BIN_FOLDER = $PROJECT_ROOT+"\"+"bin"
$EXEC_NAME = "demo.exe"

Write-Host -ForegroundColor Blue "--Script location is:" $SCRIPT_LOCATION
Write-Host -ForegroundColor Blue "--Root (executable) location is:" $PROJECT_ROOT
Write-Host -ForegroundColor Blue "--Bin path is:" $BIN_FOLDER


#In powershell functions goes first before usage
function ClearAllFiles([string]$directory) {
    Write-Host("-Cleaning bin folder") -ForegroundColor Cyan
    Get-ChildItem $directory -Include *.* -Recurse | foreach($_){
        Write-Host -ForegroundColor Cyan "--Cleaning:" + $_.fullname
        remove-item $_.fullname -Force -Recurse
    }
}

function CleanDirectory([string]$dir){
    Write-Host "--Cleaning $dir" -ForegroundColor Cyan
    Remove-Item -Path $dir\* -Force -Recurse
}

function Retry-Command{
    param (
        [Parameter(Mandatory=$true)][string]$command,
        [Parameter(Mandatory=$true)][hashtable]$args,
        [Parameter(Mandatory=$false)][int]$retries = 10,
        [Parameter(Mandatory=$false)][int]$secondsDelay = 5
    )

    $retrycount = 0
    $completed = $false

    while (-not $completed) {
        try {
            & $command @args
            Write-Verbose ("Command [{0}] succeeded after $retrycount retries" -f $command)
            $completed = $true
        } catch {
            if ($retrycount -ge $retries) {
                Write-Verbose ("Command [{0}] failed the maximum number of {1} times." -f $command, $retrycount)
                throw
            } else {
                $retrycount++
                Write-Verbose ("Command [{0}] failed. Retrying in {1} seconds. RetryCount $retrycount / $retries" -f $command, $secondsDelay)
                Start-Sleep $secondsDelay
            }
        }
    }
}

function EnsureDir([string]$dir, [switch]$RemoveContents = $false){
    if (-not (test-path $dir))
    {
        Write-Host -ForegroundColor Cyan "--Creating bin folder"
        mkdir $dir | out-null
    }

    if ($RemoveContents)
    {
        Retry-Command -command 'CleanDirectory' -Args @{dir="$dir"} -Verbose
    }
}

function BuildGoExec([string]$execDir){
    Write-Host -ForegroundColor Cyan "-Building the exec package from"$execDir "to"$BIN_FOLDER
    go build -o $BIN_FOLDER\$EXEC_NAME $execDir
}

Write-Host -ForegroundColor Red "-Starting windows build"
EnsureDir -dir $BIN_FOLDER -RemoveContents $true
BuildGoExec -execDir $PROJECT_ROOT
Write-Host -ForegroundColor Red "-Finished windows build"

#GO must be installed and the workspace created properly
#The executable file will be read automatically. A GO project cannot have multiple 'main' packages at the same level

