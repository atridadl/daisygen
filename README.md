# DaisyUI + TailwindCSS Generator built with Go

## Usage

### As a CLI
1. Run ```go installgithub.com/atridadl/daisygen/generator```
2. Ensure your GO bin directory is in your PATH
3. Run ```daisygen -extensions html -directory ./pages/templates -output-dir ./public/css``` (this is an example... use the values that make sense for your project)
4. ???
5. Profit
   
### As a Go package
1. Import ```github.com/atridadl/daisygen/generator``` using ```go get github.com/atridadl/daisygen/generator```
2. Add ```daisygen "github.com/atridadl/daisygen/generator"``` to your imports
3. Use ```daisygen.Generate("extention","directory","outputdirectory"``` directly in your go code (ideally your mail.go file)
4. ???
5. Profit
