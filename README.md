# PR-ThreatAware

PR-ThreatAware is a GitHub application developed from the ground up, designed to be installed across GitHub organizations. It efficiently tracks and analyzes pull requests (PRs) within the organization's repositories.

## How it Works

PR-ThreatAware utilizes webhooks to monitor and track PR events. Upon receiving a PR event, it collects essential context surrounding the PR, including details such as the PR description, file changes, commit diffs, user information, and other relevant parameters.

## Security Risk Analysis

The application employs OpenAI's GPT-3.5 Turbo models to evaluate the security risks introduced by each PR. Leveraging these models, it measures and assigns a risk score based on the analysis performed.

## Review Process

If the PR's risk level exceeds a predefined threshold, PR-ThreatAware takes action by adding reviewers from the security team to ensure comprehensive evaluation and mitigation of potential security risks.

#### PR that introduces Security Risk

<img width="1092" alt="290974888-366968c8-15ac-48f2-98e8-0c451c77354b" src="https://github.com/raxpd/threataware/assets/42084500/78d25187-83c7-4b83-ad37-7c9fed9d567b">
<img width="982" alt="290974953-4073d839-d9e6-4185-89d8-44f22d0e4841" src="https://github.com/raxpd/threataware/assets/42084500/a5261d44-f384-4acf-8280-a4d23715becc">

#### PR that does not introduce Security Risk

<img width="954" alt="290975013-79d7be96-a5fe-427d-9f4b-3392be8a8936" src="https://github.com/raxpd/threataware/assets/42084500/2881d670-a05b-4e88-b8d1-c7f5d9dd85e1">

## Installation

To install PR-ThreatAware within your GitHub organization, follow these steps:

1. **Clone the Repository:** Clone the PR-ThreatAware repository to a local environment or server that will host the application.
2. **Configure Webhooks:** Set up webhooks in your GitHub organization's repositories to trigger events that communicate with the PR-ThreatAware application. Configure these webhooks to point to the application's designated endpoint.
3. **Configure Permissions:** Ensure that PR-ThreatAware has appropriate permissions to access PR details and assign reviewers. Review and adjust permissions as needed within your GitHub organization settings.

## Configuration

Customize the risk threshold and reviewer assignment logic according to your organization's security policies and requirements. These configurations are adjustable within the designated configuration files provided with the application.

## Contributing

Contributions are welcome! Fork the PR-ThreatAware repository, make necessary changes or enhancements, and submit pull requests to contribute to the application's functionality or resolve identified issues.

## License

This project is licensed under the [MIT License](LICENSE.md).
