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

github -  gh workflow run cron-job.yaml  [brew install gh]
local  -  act 'workflow_dispatch' -W .github/workflows/cron-job.yaml  [brew install act] (~github local simulator)


# Teleport File copy 
 act 'workflow_dispatch' -W .github/workflows/teleport-deploy.yaml --container-options "-v $HOME/.tsh:/root/.tsh:ro" --container-architecture linux/amd64 --eventpath event.json

# Teleport ls file 

act 'workflow_dispatch' -W .github/workflows/teleport-ls.yaml --container-options "-v $HOME/.tsh:/root/.tsh:ro" --container-architecture linux/amd64 --eventpath event.json



# Run only one job at a time
gh workflow run manual-trigger.yaml
gh workflow run manual-trigger.yaml

# Manual trigger with force deploy (GitHub)
gh workflow run teleport-scheduled-rollout.yaml -f force_deploy=true

# Deploy specific group (GitHub)
gh workflow run teleport-scheduled-rollout.yaml -f deployment_group=group1

# Retry specific deployment (GitHub)
gh workflow run retry-failed-deployments.yaml -f original_deployment_id=deploy-20260120-123456

# Scheduled Rollout to multiple server (local test)
act 'workflow_dispatch' -W .github/workflows/teleport-scheduled-rollout.yaml --container-options "-v $HOME/.tsh:/root/.tsh:ro" --container-architecture linux/amd64  [ runs on offpeak hours]

act 'workflow_dispatch' -W .github/workflows/teleport-scheduled-rollout.yaml --container-options "-v $HOME/.tsh:/root/.tsh:ro" --container-architecture linux/amd64 --eventpath event_config.json [force trigger]

# Retry Failed Deployments (local test)
act 'workflow_dispatch' -W .github/workflows/retry-failed-deployments.yaml --container-options "-v $HOME/.tsh:/root/.tsh:ro" --container-architecture linux/amd64 



# Certificate Creation & GitHub Actions Setup

## Generate Teleport Certificate

```bash
# Generate certificate with 10 hour TTL
tctl auth sign --format=openssh --ttl=10h --user=nagaaravindb@netskope.com -o nagaaravindb_teleport

# This creates two files:
# -rw-------   1 nagaaravindb  staff        6786 20 Jan 18:51 nagaaravindb_teleport-cert.pub
# -rw-------   1 nagaaravindb  staff        1679 20 Jan 18:51 nagaaravindb_teleport
```

## Setup GitHub Secrets for Certificate-Based Authentication

To use certificate authentication in GitHub Actions, add these secrets to your repository:

### 1. Go to Repository Settings
- Navigate to: Settings → Secrets and variables → Actions → New repository secret

### 2. Add Required Secrets

**TELEPORT_KEY** (Private Key)
```bash
cat nagaaravindb_teleport
# Copy the entire content and paste as secret
```

**TELEPORT_CERT** (Certificate)
```bash
cat nagaaravindb_teleport-cert.pub
# Copy the entire content and paste as secret
```

**TELEPORT_PROXY** (Proxy Address)
```
teleport.netskope.io:443
```

**TELEPORT_USER** (Your Email)
```
nagaaravindb@netskope.com
```

## Using Certificate-Based Workflow

### Via GitHub UI:
1. Go to **Actions** → **Teleport Deploy with Certificate Auth**
2. Click **Run workflow**
3. Configure inputs (optional):
   - Target server
   - Cluster name (default: iad0)
   - Target path (default: /tmp)
   - Files to copy

### Via GitHub CLI:
```bash
# Deploy using certificates
gh workflow run teleport-cert-deploy.yaml

# Deploy with custom server
gh workflow run teleport-cert-deploy.yaml -f target_server=my-server.example.com

# Deploy with custom path
gh workflow run teleport-cert-deploy.yaml -f target_path=/opt/myapp
```

### Local Testing with act:
```bash
# Test locally (requires existing tsh session)
act 'workflow_dispatch' -W .github/workflows/teleport-cert-deploy.yaml \
  --container-options "-v $HOME/.tsh:/root/.tsh:ro" \
  --container-architecture linux/amd64 \
  --eventpath event.json
```

## Certificate Management

### Regenerate Certificates
Certificates expire based on the TTL. To regenerate:

```bash
# Generate new certificate with desired TTL
tctl auth sign --format=openssh --ttl=24h --user=nagaaravindb@netskope.com -o nagaaravindb_teleport

# Update GitHub secrets with new certificate content
cat nagaaravindb_teleport-cert.pub  # Update TELEPORT_CERT
cat nagaaravindb_teleport           # Update TELEPORT_KEY
```

### Check Certificate Expiry
```bash
# View certificate details
ssh-keygen -L -f nagaaravindb_teleport-cert.pub

# Look for "Valid: from X to Y" in output
```

### Security Best Practices
- Use shortest practical TTL (e.g., 24h for CI/CD)
- Rotate certificates regularly
- Never commit certificate files to git
- Use GitHub Secrets for storing credentials
- Limit certificate principals/roles to minimum required
 ```