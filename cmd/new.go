package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

func init() {
	rootCmd.AddCommand(new)
	new.PersistentFlags().StringVarP(&outputDir, "output", "o", "./", "Directory for certificate to be saved to.")
}

// Template for openssl.conf:
var opensslConfTemplate = `[req]
distinguished_name = req_distinguished_name
x509_extensions = v3_ca
prompt = no
[req_distinguished_name]
CN = {{index .Domains 0}}
[v3_ca]
keyUsage = digitalSignature, keyEncipherment
extendedKeyUsage = serverAuth
subjectAltName = @alternate_names
[alternate_names]
{{range $index, $host := .Domains -}}
DNS.{{$index}} = {{$host}}
{{end -}}
`

// OpenSSLConfig defines the alt names and main common name defined
// on the certificate.
type OpenSSLConfig struct {
	Domains []string
}

var outputDir string

var new = &cobra.Command{
	Use:   "new [certificate-name] [domains...]",
	Short: "Generate a new certificate key and crt",
	Long:  "Generate a new certificate for multiple domains: new mycert mydomain.com myotherdomain.com",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		certName := args[0]
		config := OpenSSLConfig{args[1:]}

		fmt.Printf(chalk.Yellow.Color("Generating certificate called: %s with %s domains, will be saved to: %s\n"), certName, args[1:], outputDir)
		GenerateCert(outputDir, certName, config)

	},
}

// GenerateCert will generate a certificate.
func GenerateCert(saveDir, certName string, config OpenSSLConfig) {
	var openSSLConf, err = template.New("openssl.conf").Parse(opensslConfTemplate)
	if err != nil {
		panic(err)
	}
	var tpl bytes.Buffer
	openSSLConf.Execute(&tpl, config)

	mkdirErr := os.MkdirAll(saveDir, 0700)
	if mkdirErr != nil {
		fmt.Println("Could not create directory!" + saveDir)
		panic(mkdirErr)
	}
	certPath := strings.TrimRight(saveDir, "/") + "/" + certName

	writeErr := ioutil.WriteFile("/tmp/openssl.conf", tpl.Bytes(), 0644)
	if writeErr != nil {
		panic(writeErr)
	}

	command := exec.Command("openssl", "req", "-x509", "-newkey", "rsa:4096", "-days", "365", "-sha256", "-nodes",
		"-keyout", certPath+".key",
		"-out", certPath+".crt",
		"-config", "/tmp/openssl.conf")
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	genErr := command.Run()
	if genErr != nil {

		panic(genErr)
	}
	os.Remove("/tmp/openssl.conf")

	fmt.Println(chalk.Yellow.Color("Adding the certificate into Keychain so it is trusted by your machine! [requires sudo]"))

	// Now we need to trust the certificate....
	trustCert := exec.Command("sudo", "security", "add-trusted-cert", "-d", "-r", "trustRoot", "-k", "/Library/Keychains/System.keychain", certPath+".crt")
	trustCert.Stdout = os.Stdout
	trustCert.Stderr = os.Stderr
	trustErr := trustCert.Run()
	if trustErr != nil {
		panic(trustErr)
	}
	fmt.Println(chalk.Green.Color("All done! 🎉"))
}
