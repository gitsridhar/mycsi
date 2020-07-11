package main

import (
        "flag"
        "math/rand"
        "io/ioutil"
        "os"
        "fmt"
        "time"
        "github.com/golang/glog"

        driver "github.com/gitsridhar/mycsi/src/csiplugin"
)

func init() {
        flag.Set("logtostderr", "true")
}

var (
        endpoint      = flag.String("csi-address", "unix://tmp/csi.sock", "CSI endpoint")
        driverName    = flag.String("drivername", "my-csi-driver", "name of the driver")
        nodeID        = flag.String("nodeid", "", "node id")
        vendorVersion = "1.0.0"
)

func main() {
        flag.Parse()
        rand.Seed(time.Now().UnixNano())

        handle()
        os.Exit(0)
}

func handle() {
        driver := driver.GetCSIDriver()
        err := driver.InitializeDriver(*driverName, vendorVersion, *nodeID)
        if err != nil {
                glog.Fatalf("Failed to initialize CSI Driver: %v", err)
        }
        driver.Run(*endpoint)
}
