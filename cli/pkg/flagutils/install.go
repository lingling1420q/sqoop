package flagutils

import (
	"github.com/solo-io/sqoop/cli/pkg/options"
	"github.com/solo-io/sqoop/pkg/defaults"
	"github.com/spf13/pflag"
)

func AddInstallFlags(set *pflag.FlagSet, install *options.Install) {
	set.BoolVarP(&install.DryRun, "dry-run", "d", false, "Dump the raw installation yaml instead of applying it to kubernetes")
	set.StringVarP(&install.HelmChartOverride, "file", "f", "", "Install sqoop from this helm archive file rather than from a release")
	set.StringVarP(&install.Namespace, "namespace", "n", defaults.GlooSystem, "which namespace to install sqoop into")
}

func AddUninstallFlags(set *pflag.FlagSet, uninstall *options.Uninstall) {
	set.StringVarP(&uninstall.Namespace, "namespace", "n", defaults.GlooSystem, "which namespace to uninstall sqoop from")
}
