# Test Module: GitHub Actions + SonarCloud Integration Demo

This repository contains a minimal Go module created to test and demonstrate the integration between **GitHub Actions** and **SonarCloud**.

The primary goal of this project is to validate the workflow for automated code analysis and quality checks triggered by CI/CD pipelines.

## 📦 Module Overview

The module itself is intentionally simple, serving as a test subject. Its core logic is a basic "summator" (adder). The code is structured to allow for unit tests, code coverage reporting, and subsequent analysis by SonarCloud.

*   **Core Functionality**: The application accepts a set of numbers (probably via command line arguments), adds them up, and outputs the result.
*   **Project Structure**:
    *   `cmd/summator/`: Likely contains the main application entry point.
    *   `internal/add/`: Holds the internal `Add` function package.
    *   `.github/workflows/`: Contains the GitHub Actions CI/CD configuration.
    *   `sonar-project.properties`: Configuration file for SonarCloud analysis.
*   **Purpose**: The code commits and history (like "test: coverage increased") show the process of modifying code and tests to see how SonarCloud detects and reports changes in code coverage and quality.

## 🚀 What This Demo Proves

1.  **Automated Trigger**: A `git push` to the repository triggers a pre-configured GitHub Actions workflow.
2.  **CI Execution**: The workflow (defined in `.github/workflows/`) likely runs the Go tests and generates a code coverage report.
3.  **SonarCloud Analysis**: After tests pass, the workflow calls SonarCloud, sending the code and coverage data for analysis based on the rules in `sonar-project.properties`.
4.  **Feedback Loop**: The results (quality gate status, new issues, coverage changes) are reported back as a comment on the pull request or commit in GitHub.

## 🔧 Usage Example

```bash
# Build the application
go build -o summator ./cmd/summator

# Run with arguments (example)
./summator 5 10 15
# Output: 30