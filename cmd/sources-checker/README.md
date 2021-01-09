# sources-checker

The utility for getting a list of the latest commits of repos from a specific source.

## Usage

```
$ sources-checker -h | -help | --help
$ sources-checker [options]
```

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit;
- `-source STRING` &mdash; source name (allowed: `github`, `gitlab`, `bitbucket`, `file-system`, and `external`);
- for `github`, `gitlab`, and `bitbucket` sources:
  - `-owner STRING` &mdash; repo owner; defaults:
    - for `github` source: [irenicaa](https://github.com/irenicaa/);
    - for `gitlab` source: [dzaporozhets](https://gitlab.com/dzaporozhets);
    - for `bitbucket` source: [MartinFelis](https://bitbucket.org/MartinFelis/);
  - `-group` &mdash; flag requiring the username to be treated as a group name (only for `gitlab` source);
  - `-pageSize INTEGER` &mdash; page size for API requests (default: `100`);
- for `file-system` source:
  - `-path STRING` &mdash; base path containing repos (default: `..`);
- for `external` source:
  - `-command STRING` &mdash; external program call in the form `command arg1 arg2 ...` returning a source state in JSON (default: `./tools/test_tool.bash ..`).

Environment variables:

- to access GitHub:
  - `GITHUB_USERNAME` &mdash; GitHub username;
  - `GITHUB_TOKEN` &mdash; GitHub [personal access token](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token);
- to access GitLab:
  - `GITLAB_TOKEN` &mdash; GitLab [personal access token](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html);
- to access Bitbucket:
  - `BITBUCKET_USERNAME` &mdash; Bitbucket username;
  - `BITBUCKET_PASSWORD` &mdash; Bitbucket [app password](https://support.atlassian.com/bitbucket-cloud/docs/app-passwords/).
