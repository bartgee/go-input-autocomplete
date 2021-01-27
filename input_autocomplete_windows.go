package input_autocomplete

import (
	"fmt"
	"log"

	"golang.org/x/sys/windows/registry"
)

func getWindowsVersion() []string {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	currentVersion, _, err := key.GetStringValue("CurrentVersion")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("CurrentVersion: %s\n", currentVersion)

	productName, _, err := key.GetStringValue("ProductName")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ProductName: %s\n", productName)
	out := []string{currentVersion, productName}
	return out
}