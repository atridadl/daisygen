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
	binary, err := getTailwindBinary()
	if err != nil {
		return err
	}
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

func getTailwindBinary() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("failed to get executable path: %w", err)
	}

	execDir := filepath.Dir(execPath)
	twDir := filepath.Join(execDir, "tw")

	osName := runtime.GOOS // Changed from 'os' to 'osName'
	arch := runtime.GOARCH

	var binaryName string
	switch osName { // Use 'osName' here
	case "darwin":
		binaryName = fmt.Sprintf("macos-%s", arch)
	case "linux":
		binaryName = fmt.Sprintf("linux-%s", arch)
	case "windows":
		binaryName = fmt.Sprintf("windows-%s.exe", arch)
	default:
		return "", fmt.Errorf("unsupported OS: %s", osName)
	}

	binaryPath := filepath.Join(twDir, binaryName)
	if _, err := os.Stat(binaryPath); os.IsNotExist(err) {
		return "", fmt.Errorf("Tailwind binary not found: %s", binaryPath)
	}

	return binaryPath, nil
}
