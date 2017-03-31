package isolated

import (
	"code.cloudfoundry.org/cli/integration/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = FDescribe("set-space-isolation-segment command", func() {
	var organizationName string
	var spaceName string

	BeforeEach(func() {
		organizationName = helpers.NewOrgName()
		spaceName = helpers.NewSpaceName()
	})

	Describe("help", func() {
		Context("when --help flag is set", func() {
			It("Displays command usage to output", func() {
				session := helpers.CF("reset-space-isolation-segment", "--help")
				Eventually(session).Should(Say("NAME:"))
				Eventually(session).Should(Say("reset-space-isolation-segment - Reset the isolation segment assignment of a space to the org's default"))
				Eventually(session).Should(Say("USAGE:"))
				Eventually(session).Should(Say("cf reset-space-isolation-segment SPACE_NAME"))
				Eventually(session).Should(Say("SEE ALSO:"))
				Eventually(session).Should(Say("org, restart, space"))
				Eventually(session).Should(Exit(0))
			})
		})
	})

	Context("when the environment is not setup correctly", func() {
		Context("when no API endpoint is set", func() {
			BeforeEach(func() {
				helpers.UnsetAPI()
			})

			It("fails with no API endpoint set message", func() {
				session := helpers.CF("reset-space-isolation-segment", spaceName)
				Eventually(session).Should(Say("FAILED"))
				Eventually(session.Err).Should(Say("No API endpoint set. Use 'cf login' or 'cf api' to target an endpoint."))
				Eventually(session).Should(Exit(1))
			})
		})

		Context("when not logged in", func() {
			BeforeEach(func() {
				helpers.LogoutCF()
			})

			It("fails with not logged in message", func() {
				session := helpers.CF("reset-space-isolation-segment", spaceName)
				Eventually(session).Should(Say("FAILED"))
				Eventually(session.Err).Should(Say("Not logged in. Use 'cf login' to log in."))
				Eventually(session).Should(Exit(1))
			})
		})

		Context("when there is no org set", func() {
			BeforeEach(func() {
				helpers.LoginCF()
			})

			It("fails with no targeted org error message", func() {
				session := helpers.CF("reset-space-isolation-segment", spaceName)
				Eventually(session).Should(Exit(1))
				Expect(session.Out).To(Say("FAILED"))
				Expect(session.Err).To(Say("No org targeted, use 'cf target -o ORG' to target an org."))
			})
		})
	})

	FContext("when the environment is set up correctly", func() {
		var userName string

		BeforeEach(func() {
			helpers.LoginCF()
			userName, _ = helpers.GetCredentials()
			helpers.CreateOrg(organizationName)
			helpers.TargetOrg(organizationName)
		})

		FContext("when the space does not exist", func() {
			It("fails with space not found message", func() {
				session := helpers.CF("reset-space-isolation-segment", spaceName)
				Eventually(session).Should(Say("Resetting isolation segment assignment of space %s in org %s as %s...", spaceName, organizationName, userName))
				Eventually(session).Should(Say("FAILED"))
				Eventually(session.Err).Should(Say("Space '%s' not found.", spaceName))
				Eventually(session).Should(Exit(1))
			})
		})

		Context("when the space exists", func() {
			BeforeEach(func() {
				helpers.CreateSpace(spaceName)
			})

			Context("when there is a default org isolation segment", func() {
				It("resets the space isolation segment to the default org isolation segment", func() {

				})
			})

			Context("when there is no default org isolation segment", func() {
				It("resets the space isolation segment to the shared isolation segment", func() {

				})
			})

		})
	})
})
