# emon-api-go
Go module to read [emon-http-arduino](https://github.com/exzz/emon-http-arduino)

## Example

```go
package main

import (
	"fmt"
	"os"

	emon "github.com/exzz/emon-api-go"
)

func main() {

	// set up
	e, err := emon.NewClient(emon.Config{
		URL: "http://192.168.168.16:5555",
	})
	if err != nil {
		fmt.Printf("Unable to connect to Energymonitor: %s", err)
		os.Exit(1)
	}

	// read data
	err = e.Read()
	if err != nil {
		fmt.Printf("Cannot fetch EnergyMonitor data: %s", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", e.Sensor)
}
```

Output should look like this :
```
{SiteName:Home ModuleName:mainPower Irms:1.15 Vrms:237.57 RealPower:197.28 ApparentPower:273.54 PowerFactor:0.72 ExecTime:1669656 Time:1479115468}
```
