/*
 * options.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-07 20:21
 * Distributed under terms of the MIT license.
 */


import { CherryOptions } from 'cherry-markdown/types/cherry';
const basicConfig: CherryOptions = {
  id: '',
  externals: {
    // echarts: window.echarts,
    // katex: window.katex,
    // MathJax: window.MathJax,
  },
  isPreviewOnly: false,
  engine: {
    global: {
      urlProcessor(url, srcType) {
        console.log(`url-processor`, url, srcType);
        return url;
      },
    },
    syntax: {
      codeBlock: {
        theme: 'twilight',
      },
      table: {
        enableChart: false,
        // chartEngine: Engine Class
      },
      fontEmphasis: {
        allowWhitespace: false, // 是否允许首尾空格
      },
      strikethrough: {
        needWhitespace: false, // 是否必须有前后空格
      },
      mathBlock: {
        engine: 'MathJax', // katex或MathJax
        // src: 'https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-svg.js', // 如果使用MathJax plugins，则需要使用该url通过script标签引入
      },
      inlineMath: {
        engine: 'MathJax', // katex或MathJax
      },
      emoji: {
        useUnicode: false,
        customResourceURL: 'https://github.githubassets.com/images/icons/emoji/unicode/${code}.png?v8',
        upperCase: true,
      },
      // toc: {
      //     tocStyle: 'nested'
      // }
      // 'header': {
      //   strict: false
      // }
    },
  },
  toolbars: {
    toolbar: [
      'bold',
      'italic',
      {
        strikethrough: ['strikethrough', 'underline', 'sub', 'sup', 'ruby'],
      },
      'size',
      '|',
      'color',
      'header',
      '|',
      'drawIo',
      '|',
      'ol',
      'ul',
      'checklist',
      'panel',
      'justify',
      'detail',
      '|',
      'formula',
      {
        insert: ['image', 'audio', 'video', 'link', 'hr', 'br', 'code', 'formula', 'toc', 'table', 'pdf', 'word', 'ruby'],
      },
      'graph',
      'togglePreview',
      'export',
    ],
    toolbarRight: [],
    bubble: ['bold', 'italic', 'underline', 'strikethrough', 'sub', 'sup', 'quote', 'ruby', '|', 'size', 'color'], // array or false
    sidebar: false,
    float: false
  },
  drawioIframeUrl: '/cherry/drawio.html',
  editor: {
    defaultModel: 'edit&preview',
  },
};

export default basicConfig
