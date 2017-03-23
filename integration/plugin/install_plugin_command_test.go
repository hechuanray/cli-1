package plugin

import (
	"code.cloudfoundry.org/cli/integration/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = FDescribe("install-plugin", func() {
	BeforeEach(func() {
		helpers.RunIfExperimental("app command refactor is still experimental")
	})

	Describe("help", func() {
		Context("when --help flag is set", func() {
			It("Displays command usage to output", func() {
				session := helpers.CF("install-plugin", "--help")
				Eventually(session).Should(Say("NAME:"))
				Eventually(session).Should(Say("install-plugin - Install CLI plugin"))

				Eventually(session).Should(Say("USAGE:"))
				Eventually(session).Should(Say("install-plugin \\(LOCAL-PATH/TO/PLUGIN \\| URL \\| -r REPO_NAME PLUGIN_NAME\\) \\[-f\\]"))

				Eventually(session).Should(Say("EXAMPLES:"))
				Eventually(session).Should(Say("cf install-plugin ~/Downloads/plugin-foobar"))
				Eventually(session).Should(Say("cf install-plugin https://example.com/plugin-foobar_linux_amd64"))
				Eventually(session).Should(Say("cf install-plugin -r My-Repo plugin-echo"))

				Eventually(session).Should(Say("OPTIONS:"))
				Eventually(session).Should(Say("-f      Force install of plugin without confirmation"))
				Eventually(session).Should(Say("-r      Name of a registered repository where the specified plugin is located"))
				Eventually(session).Should(Say("SEE ALSO:"))
				Eventually(session).Should(Say("add-plugin-repo, list-plugin-repos, plugins"))
				Eventually(session).Should(Exit(0))
			})
		})
	})

	Context("when installing a plugin from the file system", func() {
		Context("when installing a new plugin", func() {
			Context("when the -f flag is provided", func() {
			})

			Context("when the -f flag is not provided", func() {
			})
		})

		Context("when the plugin already exists", func() {
		})

		Context("when the filepath doesn't exist", func() {
		})
	})

	Context("when installing a plugin from a URL", func() {
		Context("when installing a new plugin", func() {
			Context("when the -f flag is provided", func() {
			})

			Context("when the -f flag is not provided", func() {
			})
		})

		Context("when the plugin already exists", func() {
		})

		Context("when the URL 404s", func() {
		})

		Context("when the URL isn't a plugin", func() {
		})

		Context("when the URL isn't a binary/zipfile", func() {
		})

		Context("ssl-validation failures", func() {
		})
	})

	Context("when installing a plugin from a repo", func() {
		Context("when installing a new plugin", func() {
			Context("when the -f flag is provided", func() {
			})

			Context("when the -f flag is not provided", func() {
			})
		})

		Context("when the plugin already exists", func() {
		})

		Context("when the repo 404s", func() {
		})

		Context("when the plugin download 404s", func() {
		})

		Context("when the URL isn't a binary/zipfile", func() {
		})

		Context("ssl-validation failures", func() {
		})
	})

	Context("when installing the plugin fails", func() {
	})
})
