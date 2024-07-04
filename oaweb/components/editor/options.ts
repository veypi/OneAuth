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
    changeString2Pinyin: (s: any) => s,
  },
  isPreviewOnly: false,
  fileUpload: (f: any) => { console.log('upload file: ' + f) },
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
      urlProcessor(url: any, srcType: any) {
        // console.log(`url-processor`, url, srcType);
        return url;
      },
    },
    syntax: {
      autoLink: {
        /** default open short link display */
        enableShortLink: true,
        /** default display 20 characters */
        shortLinkLength: 20,
      },
      list: {
        listNested: false, // The sibling list type becomes a child after conversion
        indentSpace: 2, // Default 2 space indents
      },
      table: {
        enableChart: false,
        // chartRenderEngine: EChartsTableEngine,
        // externals: ['echarts'],
      },
      inlineCode: {
        theme: 'red',
      },
      codeBlock: {
        theme: 'twilight', // Default to dark theme
        wrap: true, // If it exceeds the length, whether to wrap the line. If false, the scroll bar will be displayed
        lineNumber: true, // Default display line number
        customRenderer: {
          // Custom syntax renderer
        },
        /**
         * indentedCodeBlock Is the switch whether indent code block is enabled
         *
         *    this syntax is not supported by default in versions before 6.X.
         *    Because cherry's development team thinks the syntax is too ugly (easy to touch by mistake)
         *    The development team hopes to completely replace this syntax with ` ` code block syntax
         *    However, in the subsequent communication, the development team found that the syntax had better display effect in some scenarios
         *    Therefore, the development team in 6 This syntax was introduced in version X
         *    if you want to upgrade the following versions of services without users' awareness, you can remove this syntax:
         *        indentedCodeBlock：false
         */
        indentedCodeBlock: true,
      },
      fontEmphasis: {
        allowWhitespace: false, // 是否允许首尾空格
      },
      strikethrough: {
        needWhitespace: false, // 是否必须有前后空格
      },
      mathBlock: {
        engine: 'MathJax', // katex或MathJax
        src: 'https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-svg.js', // 如果使用MathJax plugins，则需要使用该url通过script标签引入
        // src: '/deps/mathjax/tex-svg.js',
        plugins: true,
      },
      inlineMath: {
        engine: 'MathJax', // katex或MathJax
      },
      emoji: {
        useUnicode: false,
        // customResourceURL: 'https://github.githubassets.com/images/icons/emoji/unicode/${code}.png?v8',
        upperCase: true,
      },
      toc: {
        /** By default, only one directory is rendered */
        allowMultiToc: false,
      },
      header: {
        /**
         * Style of title：
         *  - default       Default style with anchor in front of title
         *  - autonumber    There is a self incrementing sequence number anchor in front of the title
         *  - none          Title has no anchor
         */
        anchorStyle: 'autonumber',
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
      'saveMenu',
      'backMenu'
    ],
    // toolbarRight: [],
    bubble: ['bold', 'italic', 'underline', 'strikethrough', 'sub', 'sup', 'quote', 'ruby', '|', 'size', 'color'], // array or false
    // sidebar: false,
    // float: false
    customMenu: {
    } as any,
  },
  drawioIframeUrl: '/cherry/drawio.html',
  editor: {
    defaultModel: 'edit&preview',
  },
};

export default basicConfig
