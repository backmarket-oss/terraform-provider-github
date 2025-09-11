package github

import (
	"testing"
)

func TestAccGithubRepositoryProject(t *testing.T) {

	// randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates a repository project", func(t *testing.T) {

		// config := fmt.Sprintf(`

		// 	resource "github_repository" "test" {
		// 	  name = "tf-acc-test-%s"
		// 	  has_projects = true
		// 	  visibility = "private"
		// 	}

		// 	resource "github_repository_project" "test" {
		// 	  name       = "test"
		// 	  repository = github_repository.test.name
		// 	  body       = "this is a test project"
		// 	}
		// `, randomID)

		// check := resource.ComposeTestCheckFunc(
		// 	resource.TestMatchResourceAttr(
		// 		"github_repository_project.test", "url",
		// 		regexp.MustCompile(randomID+"/projects/1"),
		// 	),
		// )

		// testCase := func(t *testing.T, mode string) {
		// 	resource.Test(t, resource.TestCase{
		// 		PreCheck:  func() { skipUnlessMode(t, mode) },
		// 		Providers: testAccProviders,
		// 		Steps: []resource.TestStep{
		// 			{
		// 				Config: config,
		// 				Check:  check,
		// 			},
		// 		},
		// 	})
		// }

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("Skipped because of: this test fails because of 410 Projects (classic) has been deprecated in favor of the new Projects experience")
		})

		t.Run("with an organization account", func(t *testing.T) {
			t.Skip("Skipped because of: this test fails because of 410 Projects (classic) has been deprecated in favor of the new Projects experience")
		})

	})
}
