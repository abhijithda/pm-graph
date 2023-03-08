// Copyright (c) 2023 Veritas Technologies LLC. All rights reserved. IP63-2828-7171-04-15-9

// Package pg (plugin graph) is used for generating the graph image.
package pg

import (
	"time"

	"github.com/VeritasOS/plugin-manager/pluginmanager"
)

// graph of plugin and its dependencies.
type graph struct {
	// fileNoExt is the name of the graph artifacts without extension.
	// 	Extensions could be added to generate input `.dot` file or output
	// 	`.svg` images.
	fileNoExt string
}

var g graph

// initGraphConfig: Initialize graph configurations
func initGraphConfig(imgNamePrefix string) {
	// Initialization should be done only once.
	if g.fileNoExt == "" {
		g.fileNoExt = imgNamePrefix + "." + time.Now().Format(time.RFC3339Nano)
	}
}

// GetImagePath returns the SVG image location.
func GetImagePath() string {
	return g.fileNoExt + ".svg"
}

func getDotFilePath() string {
	return g.fileNoExt + ".dot"
}

// InitGraph initliazes the graph data structure and invokes generateGraph.
func InitGraph(pluginType string, pluginsInfo map[string]*pluginmanager.PluginAttributes) error {
	return nil
}

// GenerateGraph generates an input `.dot` file based on the fileNoExt name,
// and then generates an `.svg` image output file as fileNoExt.svg.
func GenerateGraph( /*dotFile, svgFile string*/ ) error {
	return nil
}

// UpdateGraph updates the plugin node with the status and url.
func UpdateGraph(subgraphName, plugin, status, url string) error {
	return nil
}

// getStatusColor returns the color for a given result status.
func getStatusColor(status string) string {
	// Node color
	ncolor := "blue" // dStatusStart by default
	if status == pluginmanager.DStatusFail {
		ncolor = "red"
	} else if status == pluginmanager.DStatusOk {
		ncolor = "green"
	} else if status == pluginmanager.DStatusSkip {
		ncolor = "yellow"
	}
	return ncolor
}
