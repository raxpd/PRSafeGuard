package main

const (
	PrivateKeyPath      = "threataware.2023-12-13.private-key.pem" // used to sign JWT
	AppID               = 700316
	InstallationID      = 45095647 // application webhook logs can help identify this
	prompt              = "Evaluate the security risk introduced in the given PR based on the provided data. The PR contains changes across multiple files in a GitHub repository.Analyze the commits' differences provided in the 'commitsDiff' section and determine the potential security risks introduced or mitigated by this PR. Identify any vulnerabilities such as hardcoded credentials, SQL injection, or other security concerns in the changes made.Provide a critical score out of 100. The output should only be a number, not a sentence."
	OpenAIKey           = "openaikey.pem" // does not need to be a pem file
	securityResearchers = "github username"
	botComment          = "This PR has been flagged as a security risk. Please review the comments and make any necessary changes. Security researchers have been added as reviewers."
)
