# python3 issue_to_md.py -u user -r repo -t token -o out_dir
# explain: https://github.com/mo2g/OpenWrt-Action
# python3 issue_to_md.py -u mo2g -r OpenWrt-Action -t abcdefg -o /tmp/OpenWrt-Action


from tqdm import tqdm
import os
import argparse
import codecs
import json
import re
import requests
requests.session().keep_alive = False


def get_info(comments_url, token):
    headers = {'Content-Type': 'application/json',
               'Authorization': 'token %s' % token}
    r = requests.get(comments_url, headers=headers)
    ret = json.loads(r.text)
    if r.status_code > 300:
        print('error %s', r.text)
        return False
    return ret


def to_markdown(issue, comments):
    mk = '# ' + issue['title'] + '\n'
    mk += "created_at: "+issue['created_at']+"\n"
    mk += "updated_at: "+issue['updated_at']+"\n"

    labels = "label: "
    for idx, item in enumerate(issue['labels']):
        if (idx+1) == len(issue['labels']):
            labels += item['name']+'\n'
        else:
            labels += item['name']+','
    mk += labels+'\n'

    mk += issue['body']+'\n'

    for c in comments:
        mk += c['body']+"\n"
    return mk


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('-u', '--username', help="github account username")
    parser.add_argument('-r', '--repo', help="github account repo")
    parser.add_argument('-t', '--token', help="github personal access token")
    parser.add_argument('-o', '--dir', help="out put dir")

    args = parser.parse_args()

    # args.username = ""
    # args.repo = ""
    # args.token = ""
    # args.dir = "github_out"

    # todo
#     page = 1
#     per_page = 100

    if not os.path.exists(args.dir):
        os.makedirs(args.dir)

    api_url = 'https://api.github.com/repos/%s/%s/issues' % (
        args.username, args.repo)
    issue_ret = get_info(api_url, args.token)

    
    for issue in tqdm(issue_ret):
        comments = []
        if issue['comments'] > 0:
            comments = get_info(issue['comments_url'], args.token)
        mk = to_markdown(issue, comments)

        # save
        filename = issue['created_at'].split('T')[0]+','+issue['title']+'.md'
        filename = re.sub(r'[/:*?"<>|]', " ", filename)  # check filename
        filename = os.path.join(args.dir, filename)

        with codecs.open(filename, 'w', 'utf-8') as file:
            file.write(mk)
