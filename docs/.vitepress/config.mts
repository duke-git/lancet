import { defineConfig } from 'vitepress'
import { commonConfig } from './common'
import { zhConfig } from './zh'
import { enConfig } from './en'

// https://vitepress.dev/reference/site-config
export default defineConfig({
    ...commonConfig,

    locales: {
        root: { label: '简体中文', lang: 'zh-CN', link: '/', ...zhConfig },
        en: { label: 'English', lang: 'en-US', link: '/en/', ...enConfig },
    },
})
