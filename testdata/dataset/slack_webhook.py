import requests

SLACK_WEBHOOK_URL = "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"
SLACK_BOT_TOKEN = "xoxb-123456789012-123456789012-ABCDEFabcdefABCDEFabcdef"

def notify(msg):
    requests.post(SLACK_WEBHOOK_URL, json={"text": msg})
