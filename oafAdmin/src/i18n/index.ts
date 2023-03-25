/*
 * index.ts
 * Copyright (C) 2022 veypi <i@veypi.com>
 * 2022-08-15 17:39
 * Distributed under terms of the Apache license.
 */


import util from '@/libs/util';
import { createI18n, useI18n } from 'vue-i18n'
// import { createI18n, useI18n } from 'vue-i18n/dist/vue-i18n.esm-bundler.js'
import en from './en'
import zh from './zh'

const messages = {
    en, zh
}

export { useI18n }
export default createI18n({
    locale: util.getCookie('language') || 'zh',
    legacy: false,
    silentTranslationWarn: true,
    fallbackLocale: 'en', // set fallback locale
    messages,
})
