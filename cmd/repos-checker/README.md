# repos-checker

The utility for checking that repo mirrors are up to date. Config of sources is described in an external file in JSON. (For details about its format, see the `README.md` file at the root of the project.)

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

## Output Example

```json
[
  {
    "NameOfLeft": "file-system:/home/irenicaa/go/src/github.com/irenicaa",
    "NameOfRight": "github:irenicaa|bitbucket:irenicaa",
    "Equal": [
      {
        "Name": "go-calculator",
        "LastCommit": "a5ef8b610cb0a025b2e4d1dcb584229d515819b7"
      },
      {
        "Name": "go-life",
        "LastCommit": "a2c9edb8e47c98a23fe520d3ce80150003eada9a"
      },
      {
        "Name": "go-weasel",
        "LastCommit": "b264afbc6415a3b345d16fc23691a77ad615c9a4"
      },
      {
        "Name": "go-wc",
        "LastCommit": "1ff30adf2fd5edae9f413e3499bf038dc72ad203"
      }
    ],
    "Diff": [
      {
        "Name": "repos-checker",
        "LastCommitInLeft": "be7ec730a107a3ba8f4e07083a498b851fd4f698",
        "LastCommitInRight": "e4697d20f11c7e4f5fcdef4433d3e679fe194a9e"
      }
    ],
    "MissedInLeft": null,
    "MissedInRight": null
  },
  {
    "NameOfLeft": "file-system:/media/irenicaa/external-hdd/go-projects",
    "NameOfRight": "github:irenicaa|bitbucket:irenicaa",
    "Equal": [
      {
        "Name": "go-life",
        "LastCommit": "a2c9edb8e47c98a23fe520d3ce80150003eada9a"
      },
      {
        "Name": "go-weasel",
        "LastCommit": "b264afbc6415a3b345d16fc23691a77ad615c9a4"
      },
      {
        "Name": "go-wc",
        "LastCommit": "1ff30adf2fd5edae9f413e3499bf038dc72ad203"
      }
    ],
    "Diff": null,
    "MissedInLeft": [
      {
        "Name": "repos-checker",
        "LastCommit": "e4697d20f11c7e4f5fcdef4433d3e679fe194a9e"
      },
      {
        "Name": "go-calculator",
        "LastCommit": "a5ef8b610cb0a025b2e4d1dcb584229d515819b7"
      }
    ],
    "MissedInRight": null
  }
]
```
