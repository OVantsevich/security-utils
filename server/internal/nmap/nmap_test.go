package nmap_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/OVantsevich/security-utils/server/internal/model"
	"github.com/OVantsevich/security-utils/server/internal/nmap"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
)

func TestNmap_Get(t *testing.T) {
	namp := nmap.NewNmap(time.Minute, "/home/oleg/GolandProjects/safqa/security-utils")
	_, err := namp.Get(context.Background(), model.NmapScanParameters{
		TargetSpecification: "localhost1",
		HostDiscovery:       model.HostDiscoveryOptions{},
		PortSpecification:   model.PortSpecificationOptions{},
		MiscOptions:         model.MiscOptions{},
	})
	require.NoError(t, err)
	//fmt.Print(string(data))

	files, err := filepath.Glob("/home/oleg/GolandProjects/safqa/security-utils/*xml")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		err = os.Remove(f)
		require.NoError(t, err)
	}
}
