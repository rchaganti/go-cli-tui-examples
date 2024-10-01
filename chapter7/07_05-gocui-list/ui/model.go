package ui

import "github.com/awesome-gocui/gocui"

type item struct {
	name       string
	isSelected bool
}

var namespaces = []item{
	{name: "default", isSelected: true},
	{name: "kube-system", isSelected: false},
	{name: "kube-public", isSelected: false},
	{name: "kube-node-lease", isSelected: false},
	{name: "ingress-nginx", isSelected: false},
	{name: "cert-manager", isSelected: false},
	{name: "monitoring", isSelected: false},
	{name: "logging", isSelected: false},
	{name: "kubeapps", isSelected: false},
	{name: "kubevirt", isSelected: false},
}

func InvokeList() error {
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		return err
	}

	defer g.Close()

	g.SetManagerFunc(layout)
	setKeybindings(g)

	g.SelFrameColor = gocui.ColorGreen
	g.Highlight = true

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}

	return nil
}
