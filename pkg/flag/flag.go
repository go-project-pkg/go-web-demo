package flag

import (
	goFlag "flag"
	"fmt"
	"strings"

	"github.com/spf13/pflag"

	"go-web-demo/pkg/log"
)

// WordSepNormalizeFunc changes all flags that contain "_" separators.
func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.ReplaceAll(name, "_", "-"))
	}

	return pflag.NormalizedName(name)
}

// WarnWordSepNormalizeFunc changes and warns for flags that contain "_" separators.
func WarnWordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		normalizedName := strings.ReplaceAll(name, "_", "-")
		log.Warn(
			fmt.Sprintf(
				"%s is DEPRECATED and will be removed in a future version. Use %s instead.",
				name,
				normalizedName,
			),
		)

		return pflag.NormalizedName(normalizedName)
	}

	return pflag.NormalizedName(name)
}

// InitFlags normalizes, parses, then logs the command line flags.
func InitFlags(fs *pflag.FlagSet) {
	fs.SetNormalizeFunc(WordSepNormalizeFunc)
	fs.AddGoFlagSet(goFlag.CommandLine)
}

// PrintFlags logs the flags in the flagset.
func PrintFlags(flags *pflag.FlagSet) {
	flags.VisitAll(func(flag *pflag.Flag) {
		log.Debug("Flag value has been parsed")
	})
}