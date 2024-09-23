package generator

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func Generate(extensions, directory, outputDir string) error {
	// Find files
	files, err := findFiles(extensions, directory)
	if err != nil {
		return err
	}

	// Generate Tailwind config
	config := generateTailwindConfig(files)
	err = os.WriteFile("tailwind.config.js", []byte(config), 0644)
	if err != nil {
		return err
	}

	// Run Tailwind CLI
	binary := getTailwindBinary()
	cmd := exec.Command(binary, "build", "-i", "./base.css", "-c", "tailwind.config.js", "-o", filepath.Join(outputDir, "styles.css"), "--minify")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func findFiles(extensions, directory string) ([]string, error) {
	var files []string
	exts := strings.Split(extensions, ",")

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for _, ext := range exts {
			if strings.HasSuffix(info.Name(), "."+ext) {
				files = append(files, path)
				break
			}
		}
		return nil
	})

	return files, err
}

func generateTailwindConfig(files []string) string {
	fileList := strings.Join(files, "\",\"")
	return fmt.Sprintf(`
module.exports = {
  content: ["%s"],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: ["night"],
  },
  plugins: [require('daisyui'), require('@tailwindcss/typography')],
}`, fileList)
}

func getTailwindBinary() string {
	os := runtime.GOOS
	arch := runtime.GOARCH

	switch os {
	case "darwin":
		os = "macos"
	case "windows":
		return fmt.Sprintf("./tw/%s-%s.exe", os, arch)
	}

	return fmt.Sprintf("./tw/%s-%s", os, arch)
}
