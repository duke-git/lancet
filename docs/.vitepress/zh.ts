import type { DefaultTheme, LocaleSpecificConfig } from 'vitepress'

export const META_URL = 'https://lancet.go.dev'
export const META_TITLE = 'Lancet'
export const META_DESCRIPTION = '一个强大的Go语言工具函数库'

export const zhConfig: LocaleSpecificConfig<DefaultTheme.Config> = {
    description: META_DESCRIPTION,

    head: [
        ['meta', { property: 'og:url', content: META_URL }],
        ['meta', { property: 'og:description', content: META_DESCRIPTION }],
        ['meta', { property: 'twitter:url', content: META_URL }],
        ['meta', { property: 'twitter:title', content: META_TITLE }],
        ['meta', { property: 'twitter:description', content: META_DESCRIPTION }],
    ],

    themeConfig: {
        outline: {
            label: '本页内容',
        },

        docFooter: {
            prev: '上一页',
            next: '下一页',
        },

        nav: [
            {
                text: '首页',
                link: '/',
                activeMatch: '^/',
            },
            {
                text: '指南',
                link: '/guide/introduction',
                activeMatch: '^/guide/',
            },
            { text: 'API', link: '/api/overview', activeMatch: '^/api/' },
            {
                text: '相关链接',
                items: [
                    {
                        text: '更新日志',
                        link: 'https://github.com/duke-git/lancet/releases',
                    },
                ],
            },
        ],

        sidebar: {
            '/': [
                {
                    text: '介绍',
                    items: [
                        {
                            text: 'Lancet是什么？',
                            link: '/guide/introduction',
                        },
                        {
                            text: '开始',
                            link: '/guide/getting_started',
                        },
                    ],
                },
            ],

            '/api/': [
                {
                    text: '概览',
                    items: [{ text: 'API概述', link: '/api/overview' }],
                },
                {
                    text: 'API文档',
                    items: [
                        { text: '算法', link: '/api/packages/algorithm' },
                        { text: '比较器', link: '/api/packages/compare' },
                        { text: '并发处理', link: '/api/packages/concurrency' },
                        { text: '条件判断', link: '/api/packages/condition' },
                        { text: '类型转换', link: '/api/packages/convertor' },
                        { text: '加密&解密', link: '/api/packages/cryptor' },
                        {
                            text: '数据结构',
                            items: [
                                { text: '线性表', link: '/api/packages/datastructure/list' },
                                { text: '线性表(线程安全)', link: '/api/packages/datastructure/copyonwritelist' },
                                { text: '链表', link: '/api/packages/datastructure/link' },
                                { text: '栈', link: '/api/packages/datastructure/stack' },
                                { text: '队列', link: '/api/packages/datastructure/queue' },
                                { text: '堆', link: '/api/packages/datastructure/heap' },
                                { text: '树', link: '/api/packages/datastructure/tree' },
                                { text: '集合', link: '/api/packages/datastructure/set' },
                                { text: 'HashMap', link: '/api/packages/datastructure/hashmap' },
                            ]
                        },
                    ],
                },
            ],
        },
    },
}
