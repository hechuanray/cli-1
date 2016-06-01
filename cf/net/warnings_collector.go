package net

import (
	"os"
	"strings"

	"github.com/cloudfoundry/cli/cf/terminal"
)

const DeprecatedEndpointWarning = "Endpoint deprecated"

type WarningsCollector struct {
	ui               terminal.UI
	warningProducers []WarningProducer
}

//go:generate counterfeiter . WarningProducer

type WarningProducer interface {
	Warnings() []string
}

func NewWarningsCollector(ui terminal.UI, warningsProducers ...WarningProducer) WarningsCollector {
	return WarningsCollector{
		ui:               ui,
		warningProducers: warningsProducers,
	}
}

func (warningsCollector WarningsCollector) PrintWarnings() {
	warnings := []string{}
	for _, warningProducer := range warningsCollector.warningProducers {
		for _, warning := range warningProducer.Warnings() {
			if warning == DeprecatedEndpointWarning {
				continue
			}
			warnings = append(warnings, warning)
		}
	}

	if os.Getenv("CF_RAISE_ERROR_ON_WARNINGS") != "" {
		if len(warnings) > 0 {
			panic(strings.Join(warnings, "\n"))
		}
	}

	warnings = warningsCollector.removeDuplicates(warnings)

	for _, warning := range warnings {
		warningsCollector.ui.Warn(warning)
	}
}

func (warningsCollector WarningsCollector) removeDuplicates(stringArray []string) []string {
	length := len(stringArray) - 1
	for i := 0; i < length; i++ {
		for j := i + 1; j <= length; j++ {
			if stringArray[i] == stringArray[j] {
				stringArray[j] = stringArray[length]
				stringArray = stringArray[0:length]
				length--
				j--
			}
		}
	}
	return stringArray
}
