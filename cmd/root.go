/*
Copyright © 2023 Oleg Gulevskyy <oleggulevskyy@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "create-zaf-app",
	Short: "Bootstrap your next ZAF Zendesk application",
	Long:  `Creates a bootstraped Zendesk application for Support locations, with a framework of your choice.`,
	Run:   createProject,
}

type Project struct {
	Name            string
	Framework       string
	ZendeskLocation string
	Tailwind        bool
	Debug           bool
	PackageManager  string

	selectedListItem  string
	selectedInputItem string
}

func (p *Project) resetSelectedListItem() {
	p.selectedListItem = ""
}

func (p *Project) resetSelectedInputItem() {
	p.selectedInputItem = ""
}

func createProject(cmd *cobra.Command, args []string) {
	proj := Project{}
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		panic(err)
	}

	framework, err := cmd.Flags().GetString("framework")
	if err != nil {
		panic(err)
	}

	pkgManager, err := cmd.Flags().GetString("pkg-manager")
	if err != nil {
		panic(err)
	}

	location, err := cmd.Flags().GetString("location")
	if err != nil {
		panic(err)
	}

	tailwind, err := cmd.Flags().GetBool("tailwind")
	if err != nil {
		panic(err)
	}

	debug, err := cmd.Flags().GetBool("debug")
	if err != nil {
		panic(err)
	}

	proj.Name = name
	proj.Framework = framework
	proj.PackageManager = pkgManager
	proj.ZendeskLocation = location
	proj.Tailwind = tailwind
	proj.Debug = debug

	// if any of the flags are set, skip the prompts, otherwise prompt the user
	err = createPrompts(&proj)
	if err != nil {
		panic(err)
	}

	fmt.Println("Project name: ", proj.Name)
	fmt.Println("Project framework: ", proj.Framework)
	fmt.Println("Project location: ", proj.ZendeskLocation)
	fmt.Println("Project tailwind: ", proj.Tailwind)
	fmt.Println("Project package manager: ", proj.PackageManager)
	fmt.Println("Project debug: ", proj.Debug)
}

func (p *Project) promptZendeskLocation() {
	p.promptList(
		"Where would you like your application to run?",
		[]list.Item{
			item("Ticket sidebar"),
			item("New ticket sidebar"),
			item("Organization sidebar"),
			item("User sidebar"),
			item("Top bar"),
			item("Nav bar"),
			item("Modal"),
			item("Ticket editor"),
			item("Background"),
		},
		"You chose: %s! No problems 💪\n",
	)
	p.ZendeskLocation = p.selectedListItem
	p.resetSelectedListItem()
}

func (p *Project) promptName() {
	p.promptInput(
		promptInputProps{
			placeholder: "my-zendesk-app",
			title:       "What would you like to name your project?",
		},
	)
	p.Name = p.selectedInputItem
	p.resetSelectedInputItem()
}

func (p *Project) promptFramework() {
	p.promptList(
		"Which framework would you like to use?",
		[]list.Item{
			item("react"),
			item("react-ts"),
			item("react-swc"),
			item("react-swc-ts"),
			item("vue"),
			item("vue-ts"),
			item("svelte"),
			item("svelte-ts"),
		},
		"Choosen one: %s! Almost over 🔥\n",
	)
	p.Framework = p.selectedListItem
	p.resetSelectedListItem()
}

func (p *Project) promptPackageManager() {
	p.promptList(
		"Which package manager would you like to use? (only npm and pnpm are supported at the moment)",
		[]list.Item{
			item("npm"),
			item("pnpm"),
		},
		"Choosen one: %s! Almost over 🔥\n",
	)
	p.PackageManager = p.selectedListItem
	p.resetSelectedListItem()
}

func (p *Project) promptTailwind() {
	p.promptList(
		"Would you like to use Tailwind CSS?",
		[]list.Item{
			item("Yes"),
			item("No"),
		},
		"Answered: %s! 🎉\n",
	)
	p.Tailwind = p.selectedListItem == "Yes"
	p.resetSelectedListItem()
}

func createPrompts(p *Project) error {
	// prompt user to define projects name
	if p.Name == "" {
		p.promptName()
	}

	if p.ZendeskLocation == "" {
		p.promptZendeskLocation()
	}

	// prompt user which framework to use
	if p.Framework == "" {
		p.promptFramework()
	}

	if p.PackageManager == "" {
		p.promptPackageManager()
	}

	if p.Tailwind == false {
		p.promptTailwind()
	}

	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("name", "n", "", "Name of the project")
	rootCmd.Flags().StringP("framework", "f", "", "Frontend framework to use")
	rootCmd.Flags().StringP("location", "l", "", "Location of Zendesk App")
	rootCmd.Flags().StringP("pkg-manager", "p", "", "Package manager (npm or pnpm only at the moment)")
	rootCmd.Flags().BoolP("tailwind", "t", false, "Use Tailwind CSS")
	rootCmd.Flags().Bool("debug", false, "Debug mode")
}
