/*
 * options.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-07 20:21
 * Distributed under terms of the MIT license.
 */


import { CherryOptions } from 'cherry-markdown/types/cherry';
const basicConfig: CherryOptions = {
  id: '',
  value: '',
  externals: {
    // echarts: window.echarts,
    // katex: window.katex,
    // MathJax: window.MathJax,
  },
  /** 预览区域跟随编辑器光标自动滚动 */
  autoScrollByCursor: true,
  forceAppend: false,
  locale: 'zh_CN',
  previewer: {
    dom: false,
    className: 'cherry-markdown',
    // Whether to enable the editing ability of preview area (currently supports editing picture size and table content)
    enablePreviewerBubble: true,
    // 配置图片懒加载的逻辑
    lazyLoadImg: {
      // 加载图片时如果需要展示loaing图，则配置loading图的地址
      loadingImgPath: '',
      // 同一时间最多有几个图片请求，最大同时加载6张图片
      maxNumPerTime: 1,
      // 不进行懒加载处理的图片数量，如果为0，即所有图片都进行懒加载处理， 如果设置为-1，则所有图片都不进行懒加载处理
      noLoadImgNum: 0,
      // 首次自动加载几张图片（不论图片是否滚动到视野内），autoLoadImgNum = -1 表示会自动加载完所有图片
      autoLoadImgNum: 3,
      // 针对加载失败的图片 或 beforeLoadOneImgCallback 返回false 的图片，最多尝试加载几次，为了防止死循环，最多5次。以图片的src为纬度统计重试次数
      maxTryTimesPerSrc: 1,
      // 加载一张图片之前的回调函数，函数return false 会终止加载操作
      beforeLoadOneImgCallback: (img: HTMLImageElement) => true,
      // 加载一张图片失败之后的回调函数
      failLoadOneImgCallback: (img: HTMLImageElement) => { },
      // 加载一张图片之后的回调函数，如果图片加载失败，则不会回调该函数
      afterLoadOneImgCallback: (img: HTMLImageElement) => { },
      // 加载完所有图片后调用的回调函数
      afterLoadAllImgCallback: () => { },
    }
  },
  theme: [],
  callback: {
    afterChange: () => { },
    /** 编辑器完成初次渲染后触发 */
    afterInit: () => { },
    /** img 标签挂载前触发，可用于懒加载等场景 */
    beforeImageMounted: (srcProp: string, src: string) => {
      return { srcProp: srcProp, src: src }
    },
    onClickPreview: () => { },
    onCopyCode: (e: ClipboardEvent, code: string) => code,
    changeString2Pinyin: (s) => s,
  },
  isPreviewOnly: false,
  fileUpload: (f) => { console.log('uploading file' + f.name) },
  fileTypeLimitMap: {
    video: "",
    audio: "",
    image: "",
    word: "",
    pdf: "",
    file: "",
  },
  openai: false,
  engine: {
    global: {
      urlProcessor(url, srcType) {
        // console.log(`url-processor`, url, srcType);
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
    showToolbar: true,
    theme: 'light',
    toolbar: [
      'bold',
      'italic',
      // {
      //   strikethrough: ['strikethrough', 'underline', 'sub', 'sup', 'ruby'],
      // },
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
    // toolbarRight: [],
    bubble: ['bold', 'italic', 'underline', 'strikethrough', 'sub', 'sup', 'quote', 'ruby', '|', 'size', 'color'], // array or false
    // sidebar: false,
    // float: false
  },
  drawioIframeUrl: '/cherry/drawio.html',
  editor: {
    defaultModel: 'edit&preview',
  },
};

export default basicConfig
