# List your pull request's
---

pull-requests is a small command line tool that prints all you'r github pull requests
within you'r organization
can also output results in a json format for [alfred](https://www.alfredapp.com/).
 

usage is straight forward.
```bash
pull-requeats
# will print your pull requests
pull-requeats -alfred
#will print pull requests in alfred format
```
the util requires some env vars to work 

export REPOS_DIR=[absolute path to folder with all you'r repos]
export GITHUB_USER=viggin543
export GITHUB_PASS=[token](https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line)
export GITHUB_ORG=[github org name]

