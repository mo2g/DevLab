# python3 issue_to_md.py -u user -r repo -t token -o out_dir
# explain: https://github.com/mo2g/OpenWrt-Action
# python3 issue_to_md.py -u mo2g -r OpenWrt-Action -t abcdefg -o /tmp/OpenWrt-Action
# https://docs.github.com/en/rest/issues/issues?apiVersion=2022-11-28#get-an-issue
# https://docs.github.com/en/rest/issues/issues?apiVersion=2022-11-28#list-repository-issues
# https://docs.github.com/en/rest/issues/comments?apiVersion=2022-11-28#list-issue-comments

from tqdm import tqdm
import os
import argparse
import codecs
import json
import re
import requests

requests.session().keep_alive = False


def get_info(comments_url, token, params={}):
    headers = {'Content-Type': 'application/json',
               'Authorization': 'token %s' % token}
    r = requests.get(comments_url, headers=headers, params=params)
    ret = json.loads(r.text)
    if r.status_code > 300:
        print('error: ', r.text)
        return False
    return ret


def to_markdown(issue, comments):
    mk = '# ' + issue['title'] + '\n\n'
    mk += "user: " + issue['user']['login'] + "\n"
    mk += "created_at: " + issue['created_at'] + "\n"
    mk += "updated_at: " + issue['updated_at'] + "\n"

    labels = "label: "
    for idx, item in enumerate(issue['labels']):
        if (idx + 1) == len(issue['labels']):
            labels += item['name'] + '\n'
        else:
            labels += item['name'] + ','
    mk += labels + '\n'

    if issue['body']:
        mk += issue['body'] + '\n\n'
    else:
        print(f'issue {issue["number"]} has no body')

    for c in comments:
        mk += '# ' + c['user']['login'] + ' commented on ' + c['created_at'] + '\n\n'
        mk += c['body'] + "\n\n"
    return mk


def fetch_issue(token, issue_number):
    api_url = 'https://api.github.com/repos/%s/%s/issues/%s' % (
        args.username, args.repo, issue_number)
    return get_info(api_url, token, params)


def fetch_issues(token, params):
    api_url = 'https://api.github.com/repos/%s/%s/issues' % (
        args.username, args.repo)
    return get_info(api_url, token, params)


def fetch_all_comments(url, token):
    """Fetches all comments for an issue, handling pagination."""
    all_comments = []
    page = 1
    per_page = 100

    while True:
        params = {'page': page, 'per_page': per_page}
        comments = get_info(url, token, params)
        if not comments:
            break

        all_comments.extend(comments)
        page += 1

        if len(comments) < per_page:
            break

    return all_comments


def save_issue(issues, token):
    for issue in tqdm(issues):
        filename = f"{issue['number']}_{issue['title']}.md"
        filename = re.sub(r'[/:*?"<>|]', " ", filename)  # check filename
        filename = os.path.join(args.dir, filename)

        if os.path.exists(filename):
            continue

        comments = []
        if issue['comments'] > 0:
            comments = fetch_all_comments(issue['comments_url'], token)
        mk = to_markdown(issue, comments)

        # save
        with codecs.open(filename, 'w', 'utf-8') as file:
            file.write(mk)


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('-u', '--username', help="github account username")
    parser.add_argument('-r', '--repo', help="github account repo")
    parser.add_argument('-i', '--issue', help="github account repo issue number")
    parser.add_argument('-t', '--token', help="github personal access token")
    parser.add_argument('-o', '--dir', help="out put dir")

    args = parser.parse_args()

    # args.username = ""
    # args.repo = ""
    # args.token = ""
    # args.dir = "github_out"

    if not os.path.exists(args.dir):
        os.makedirs(args.dir)

    page = 1
    per_page = 100
    issue = args.issue

    while True:
        params = {'page': page, 'per_page': per_page, 'state': 'all'}
        issues = []
        if issue:
            issue = fetch_issue(args.token, issue)
            issues.append(issue)
        else:
            issues = fetch_issues(args.token, params)

        if not issues:
            break

        save_issue(issues, args.token)

        page += 1
        if len(issues) < per_page:
            break

        if issue:
            break
