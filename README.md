# GitHub Action Prototype

A simple Go program that copy JSON files into qa env using github action automated CI/CD pipeline.

## Manual Trigger

The workflow can also be triggered manually using the **workflow_dispatch** feature:

### Via GitHub UI:
1. Go to your repository on GitHub
2. Click on **Actions** tab
3. Select **Go Service CI/CD** workflow
4. Click **Run workflow** button
5. Optionally provide a reason for the manual trigger
6. Click **Run workflow**

### Via GitHub CLI:
```bash
# Trigger workflow manually
gh workflow run cicd.yaml


### Via API:
```bash
curl -X POST \
  -H "Accept: application/vnd.github.v3+json" \
  -H "Authorization: token YOUR_TOKEN" \
  https://api.github.com/repos/OWNER/REPO/actions/workflows/cicd.yaml/dispatches \
  -d '{"ref":"main","inputs":{"reason":"Manual test run"}}'
```

## Cron Syntax Reference

GitHub Actions uses standard cron syntax:
```
┌───────────── minute (0 - 59)
│ ┌───────────── hour (0 - 23)
│ │ ┌───────────── day of the month (1 - 31)
│ │ │ ┌───────────── month (1 - 12)
│ │ │ │ ┌───────────── day of the week (0 - 6) (Sunday to Saturday)
│ │ │ │ │
│ │ │ │ │
* * * * *
```

# local testing - act 

github -  gh workflow run cicd.yaml  [brew install gh]
local  -  act 'workflow_dispatch' -W .github/workflows/cicd.yaml  [brew install act] (~github local simulator)


# File copy 
 act 'workflow_dispatch' -W .github/workflows/teleport-deploy.yaml --container-options "-v $HOME/.tsh:/root/.tsh:ro" --container-architecture linux/amd64 --eventpath event.json

ls file 

act 'workflow_dispatch' -W .github/workflows/teleport-ls.yaml --container-options "-v $HOME/.tsh:/root/.tsh:ro" --container-architecture linux/amd64 --eventpath event.json