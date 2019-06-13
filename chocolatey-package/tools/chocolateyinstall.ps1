
$ErrorActionPreference = 'Stop'

$toolsDir = "$(Split-Path -parent $MyInvocation.MyCommand.Definition)"

$packageArgs = @{
    PackageName    = $env:ChocolateyPackageName
    unzipLocation  = $toolsDir
    url        = 'https://github.com/giansalex/pluralsight-transcripter/releases/download/v1.0.0/transcripter-win32.zip' 
    url64      = 'https://github.com/giansalex/pluralsight-transcripter/releases/download/v1.0.0/transcripter-win64.zip' 

    checksum       = '33608C525B044267607774A6F2E86F1921E0F6134A8D4132BBE983905C56B388'
    checksumType   = 'sha256'
    checksum64     = '4D5E8A585D6D0E1A9852DF4576E90F77FAEDB1DFB7CC2E6486C5AB1BD937E1FD'
    checksumType64 = 'sha256'
}
 
Install-ChocolateyZipPackage @packageArgs
