package main

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/disk"
	"github.com/urfave/cli"
)

type Volume struct {
	Name       string
	Total      uint64
	Used       uint64
	Available  uint64
	UsePercent float64
	Mount      string
}

func ActionVolumes(c *cli.Context) error {
	stats, err := disk.Partitions(true)
	if err != nil {
		return err
	}

	var vols []*Volume

	for _, stat := range stats {
		usage, err := disk.Usage(stat.Mountpoint)
		if err != nil {
			continue
		}

		vol := &Volume{
			Name:       stat.Device,
			Total:      usage.Total,
			Used:       usage.Used,
			Available:  usage.Free,
			UsePercent: usage.UsedPercent,
		}

		vols = append(vols, vol)
	}

	volsByteArr, err := json.MarshalIndent(vols, "", "\t")
	if err != nil {
		return err
	}

	fmt.Println(string(volsByteArr))
	return nil
}
