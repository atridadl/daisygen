package generator

import (
	"embed"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

//go:embed tw
var embeddedTW embed.FS

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
	defer os.RemoveAll(filepath.Dir(binary))

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
	osName := runtime.GOOS
	arch := runtime.GOARCH

	var binaryName string
	switch osName {
	case "darwin":
		binaryName = fmt.Sprintf("macos-%s", arch)
	case "linux":
		binaryName = fmt.Sprintf("linux-%s", arch)
	case "windows":
		binaryName = fmt.Sprintf("windows-%s.exe", arch)
	default:
		return "", fmt.Errorf("unsupported OS: %s", osName)
	}

	// Create a temporary directory to extract the binary
	tempDir, err := os.MkdirTemp("", "daisygen-tw")
	if err != nil {
		return "", fmt.Errorf("failed to create temp directory: %w", err)
	}

	binaryPath := filepath.Join(tempDir, binaryName)

	// Read the embedded file
	embeddedFile, err := embeddedTW.Open(fmt.Sprintf("tw/%s", binaryName))
	if err != nil {
		return "", fmt.Errorf("failed to open embedded file: %w", err)
	}
	defer embeddedFile.Close()

	// Create the binary file
	outputFile, err := os.OpenFile(binaryPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	// Copy the content
	_, err = io.Copy(outputFile, embeddedFile)
	if err != nil {
		return "", fmt.Errorf("failed to write binary: %w", err)
	}

	return binaryPath, nil
}
