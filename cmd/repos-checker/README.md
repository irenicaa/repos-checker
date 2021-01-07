# repos-checker

## Usage

```
$ repos-checker -h | -help | --help
$ repos-checker [options]
```

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit;
- `-config STRING` &mdash; path to a config file of sources;
- `-reference STRING` &mdash; forced name of a reference source.

Environment variables:

- to access GitHub:
  - `GITHUB_USERNAME` &mdash; GitHub username;
  - `GITHUB_TOKEN` &mdash; GitHub [personal access token](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token);
- to access GitLab:
  - `GITLAB_TOKEN` &mdash; GitLab [personal access token](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html);
- to access Bitbucket:
  - `BITBUCKET_USERNAME` &mdash; Bitbucket username;
  - `BITBUCKET_PASSWORD` &mdash; Bitbucket [app password](https://support.atlassian.com/bitbucket-cloud/docs/app-passwords/).
