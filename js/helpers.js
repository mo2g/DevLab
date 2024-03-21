
// Example：
// const escapedString = escapeRegExp('hello-world');
// console.log(escapedString); // output: hello\-world
export const escapeRegExp = function (string) {
  return string.replace(/[-\/\\^$*+?.()|[\]{}]/g, '\\$&');
};

// Example：
// const arr = [];
// pushIfNonEmpty(arr, 'a', '', 'b', '   ', 'c');
// console.log(arr); // output: ['a', 'b', 'c']
export const pushIfNonEmpty = (arr, ...values) => {
  values.forEach((value) => {
    if (value && value.trim() !== '') {
      arr.push(value);
    }
  });
};

// Example：
// const highlighted = highlightedText('This is a test', 'is');
// console.log(highlighted); // output: ['Th', <b className="search-highlighted">is</b>, ' is a test']

// Highlight keyword
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

// Example：
// const truncated = truncateTextAroundKeyword('This is a test,one two three four five six', 'one', 2);
// console.log(truncated); // output: '...t,one t...'

// Extract N characters before and after the keyword from the text
export const truncateTextAroundKeyword = (text, keyword, maxLength = 10) => {
  if (!text || !keyword) {
    return '';
  }

  const keywordIndex = text.toLowerCase().indexOf(keyword.toLowerCase());
  if (keywordIndex === -1) {
    // Keyword does not exist in the text
    return text.substring(0, maxLength) + (text.length > maxLength ? '...' : '');
  }

  const start = Math.max(0, keywordIndex - maxLength);
  const end = Math.min(text.length, keywordIndex + keyword.length + maxLength);

  let truncatedText = text.substring(start, end);

  // Adding an ellipsis
  if (start > 0 || end < text.length) {
    truncatedText = (start > 0 ? '...' : '') + truncatedText + (end < text.length ? '...' : '');
  }

  return truncatedText;
}

// Example：
// const has = hasKeyword('This is a test', 'test');
// console.log(has); // output: true
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