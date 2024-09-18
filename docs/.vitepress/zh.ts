import type { DefaultTheme, LocaleSpecificConfig } from 'vitepress'

export const META_URL = 'https://www.golancet.cn'
export const META_TITLE = 'Lancet'
export const META_DESCRIPTION = '一个强大的Go语言工具函数库'

export const zhConfig: LocaleSpecificConfig<DefaultTheme.Config> = {
    description: META_DESCRIPTION,

    head: [
        ['meta', { property: 'og:url', content: META_URL }],
        ['meta', { property: 'og:description', content: META_DESCRIPTION }],
    ],

    themeConfig: {
        editLink: {
            pattern: 'https://github.com/duke-git/lancet/edit/v2/docs/:path',
            text: '对本页提出修改建议',
        },
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
                        text: '论坛',
                        link: 'https://github.com/duke-git/lancet/discussions',
                    },
                    {
                        text: '更新日志',
                        link: 'https://github.com/duke-git/lancet/releases',
                    },
                    {
                        text: '参与贡献',
                        link: 'https://github.com/duke-git/lancet/blob/main/CONTRIBUTION.zh-CN.md',
                    },
                ],
            },
        ],

        sidebar: {
            '/guide/': [
                {
                    text: '介绍',
                    collapsed: false,
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
                {
                    text: '贡献代码',
                    collapsed: false,
                    items: [
                        {
                            text: '贡献指南',
                            link: '/guide/contribution_guide',
                        },
                        {
                            text: '贡献者',
                            link: '/guide/contributors',
                        },
                    ],
                },
                {
                    text: 'API手册',
                    link: '/api/overview'
                },
            ],

            '/api/': [
                {
                    text: '概览',
                    items: [{ text: 'API概述', link: '/api/overview' }],
                },
                {
                    text: 'API文档',
                    collapsed: false,
                    items: [
                        { text: '算法', link: '/api/packages/algorithm' },
                        { text: '比较器', link: '/api/packages/compare' },
                        { text: '并发处理', link: '/api/packages/concurrency' },
                        { text: '条件判断', link: '/api/packages/condition' },
                        { text: '类型转换', link: '/api/packages/convertor' },
                        { text: '加密&解密', link: '/api/packages/cryptor' },
                        {
                            text: '数据结构',
                            collapsed: true,
                            items: [
                                { text: '线性表', link: '/api/packages/datastructure/list' },
                                {
                                    text: '线性表(线程安全)',
                                    link: '/api/packages/datastructure/copyonwritelist',
                                },
                                { text: '链表', link: '/api/packages/datastructure/link' },
                                { text: '栈', link: '/api/packages/datastructure/stack' },
                                { text: '队列', link: '/api/packages/datastructure/queue' },
                                { text: '堆', link: '/api/packages/datastructure/heap' },
                                { text: '树', link: '/api/packages/datastructure/tree' },
                                { text: '集合', link: '/api/packages/datastructure/set' },
                                { text: 'HashMap', link: '/api/packages/datastructure/hashmap' },
                            ],
                        },
                        { text: '日期&时间', link: '/api/packages/datetime' },
                        { text: '文件', link: '/api/packages/fileutil' },
                        { text: '格式化工具', link: '/api/packages/formatter' },
                        { text: '函数', link: '/api/packages/function' },
                        { text: '数学工具', link: '/api/packages/mathutil' },
                        { text: 'Map', link: '/api/packages/maputil' },
                        { text: '网络', link: '/api/packages/netutil' },
                        { text: '指针', link: '/api/packages/pointer' },
                        { text: '随机数', link: '/api/packages/random' },
                        { text: '重试', link: '/api/packages/retry' },
                        { text: '切片', link: '/api/packages/slice' },
                        { text: '流', link: '/api/packages/stream' },
                        { text: '结构体', link: '/api/packages/struct' },
                        { text: '字符串', link: '/api/packages/strutil' },
                        { text: '系统', link: '/api/packages/system' },
                        { text: '元组', link: '/api/packages/tuple' },
                        { text: '验证器', link: '/api/packages/validator' },
                        { text: '错误处理', link: '/api/packages/xerror' },
                    ],
                },
            ],
        },
    },
}
