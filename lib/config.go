package lib

import "os"

var (
	outdir = os.Getenv("HOME") + "/.license/"
	outfile = outdir + "license.json"
	httpsource = "http://osl.kevin.codes/"
	version string
	buildDate string
	commit string
)

func GetVersion() string {
	return version
}

func GetCommit() string {
	return commit
}

func GetBuildDate() string {
	return buildDate
}

func GetOutputFilePath() string {
	return outfile
}

func GetOutputPath() string {
	return outdir
}

func GetBaseURL() string {
	return httpsource
}

func GetSingleLicenseURL() string {
	return httpsource + "licenses/"
}

func GetUpdateURL() string {
	return httpsource + "licenses/"
}

func GetUserAgentHeader() string {
	return "license v" + version
}