

export const escapeRegExp = function (string) {
  return string.replace(/[-\/\\^$*+?.()|[\]{}]/g, '\\$&');
};

export const pushIfNonEmpty = (arr, ...values) => {
  values.forEach((value) => {
    if (value && value.trim() !== '') {
      arr.push(value);
    }
  });
};

// 高亮关键字
export const highlightedText = (text, keyword) => {
  if (keyword == '' || typeof keyword !== 'string') {
    return text;
  }

  try {
    const escapedKeyword = escapeRegExp(keyword);
    const regex = new RegExp(`(${escapedKeyword})`, 'gi');

    return text.split(regex).map(function (item, idx) {
      if (regex.test(item)) {
        return <b className="search-highlighted">{item}</b>;
      }
      return item;
    });
  } catch (error) {
    console.log(error);
    return text;
  }
};

// 从文本中截取关键字前后N个字符
export const truncateTextAroundKeyword = (text, keyword, maxLength = 10) => {
  if (text == '' || !!!text) {
    return '';
  }

  if (keyword == '' || typeof keyword !== 'string') {
    const len = maxLength * 2;
    const hasMore = len < text.length;
    return text.slice(0, len) + (hasMore ? '...' : '');
  }

  const keywordIndex = text.toLowerCase().indexOf(keyword.toLowerCase());
  if (keywordIndex === -1) {
    // 关键词不存在于文本中
    return text;
  }

  const start = Math.max(0, keywordIndex - maxLength);
  const end = Math.min(text.length, keywordIndex + keyword.length + maxLength);

  let truncatedText = text.slice(start, end);

  // 添加省略号
  if (start > 0) {
    truncatedText = '...' + truncatedText;
  }
  if (end < text.length) {
    truncatedText += '...';
  }
  return truncatedText;
}

export const hasKeyword = (text, keyword) => {
  return text.toLowerCase().indexOf(keyword.toLowerCase()) > -1;
};


/*
*  Example
*
let a = '/v1/user?a=AAA&id=180028';
const notAllowedQuery = ['id'];
a = clearNotAllowedQuery(a, notAllowedQuery);
console.log(a);
*
* output: /v1/user?a=AAA
* */
export const clearNotAllowedQuery = (url, notAllowedQuery) => {
  const urlObject = new URL('http://example.com' + url); // Use placeholder base URL
  const params = new URLSearchParams(urlObject.search);

  // Remove not allowe query parameters
  notAllowedQuery.forEach(param => {
    params.delete(param);
  });

  // Update the query parameters of the URL object
  urlObject.search = params.toString();

  // Returns the updated path and query string
  return urlObject.pathname + urlObject.search;
};