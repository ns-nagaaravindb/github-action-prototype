# GitHub Actions Schedule Configuration

## Scheduled Runs

The CI/CD pipeline is configured to run automatically at the following times (UTC):

| Day | Time (UTC) | Description |
|-----|------------|-------------|
| Monday | 8:00 AM | `0 8 * * 1` |
| Tuesday | 8:00 PM | `0 20 * * 2` |
| Wednesday | 10:00 PM | `0 22 * * 3` |

### Time Zone Conversions

**Note:** GitHub Actions schedules use UTC time. Convert these to your local timezone:

**UTC to Common Timezones:**

| UTC Time | PST/PDT | EST/EDT | CST/CDT | IST |
|----------|---------|---------|---------|-----|
| Mon 8:00 AM | Mon 12:00 AM / 1:00 AM | Mon 3:00 AM / 4:00 AM | Mon 2:00 AM / 3:00 AM | Mon 1:30 PM |
| Tue 8:00 PM | Tue 12:00 PM / 1:00 PM | Tue 3:00 PM / 4:00 PM | Tue 2:00 PM / 3:00 PM | Wed 1:30 AM |
| Wed 10:00 PM | Wed 2:00 PM / 3:00 PM | Wed 5:00 PM / 6:00 PM | Wed 4:00 PM / 5:00 PM | Thu 3:30 AM |

*Times vary based on Daylight Saving Time*

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

# Trigger with a custom reason
gh workflow run cicd.yaml -f reason="Testing new feature"
```

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

### Day of Week Reference:
- 0 or 7 = Sunday
- 1 = Monday
- 2 = Tuesday
- 3 = Wednesday
- 4 = Thursday
- 5 = Friday
- 6 = Saturday

## Workflow Triggers

The complete list of triggers for this workflow:

1. **Push** - Triggers on push to `develop`, `main`, or `release/**` branches
2. **Pull Request** - Triggers on PR to `develop` or `main` branches
3. **Schedule** - Runs on the configured schedule (3 times per week)
4. **Manual** - Can be triggered manually via workflow_dispatch

## Notes

- Scheduled workflows run on the default branch (usually `main`)
- Scheduled runs may be delayed during periods of high GitHub Actions usage
- Scheduled workflows are automatically disabled if there is no repository activity for 60 days
- You can re-enable disabled workflows from the Actions tab
