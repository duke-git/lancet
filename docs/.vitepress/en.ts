import type { DefaultTheme, LocaleSpecificConfig } from 'vitepress'

export const META_URL = 'https://golancet.dev'
export const META_TITLE = 'Lancet'
export const META_DESCRIPTION = 'A powerful util function library of Go'

export const enConfig: LocaleSpecificConfig<DefaultTheme.Config> = {
    description: META_DESCRIPTION,

    head: [
        ['meta', { property: 'og:url', content: META_URL }],
        ['meta', { property: 'og:description', content: META_DESCRIPTION }],
        ['meta', { property: 'twitter:url', content: META_URL }],
        ['meta', { property: 'twitter:title', content: META_TITLE }],
        ['meta', { property: 'twitter:description', content: META_DESCRIPTION }],
    ],

    themeConfig: {
        nav: [
            {
                text: 'Home',
                link: '/en/',
                activeMatch: '^/en/',
            },
            {
                text: 'Guide',
                link: '/en/guide/introduction',
                activeMatch: '^/en/guide/',
            },
            { text: 'API', link: '/en/api/overview', activeMatch: '^/en/api/' },
            {
                text: 'Links',
                items: [
                    {
                        text: 'Releaselog',
                        link: 'https://github.com/duke-git/lancet/releases',
                    },
                ],
            },
        ],

        sidebar: {
            '/en/': [
                {
                    text: 'Introduction',
                    collapsed: false,
                    items: [
                        {
                            text: 'What is Lancet？',
                            link: '/en/guide/introduction',
                        },
                        {
                            text: 'getting started',
                            link: '/en/guide/getting_started',
                        },
                    ],
                },
            ],
            '/en/api/': [
                {
                    text: 'overview',
                    items: [{ text: 'overview of API', link: '/en/api/overview' }],
                },
                {
                    text: 'packages',
                    collapsed: false,
                    items: [
                        { text: 'algorithm', link: '/en/api/packages/algorithm' },
                        { text: 'compare', link: '/en/api/packages/compare' },
                        { text: 'concurrency', link: '/en/api/packages/concurrency' },
                        { text: 'condition', link: '/en/api/packages/condition' },
                        { text: 'convertor', link: '/en/api/packages/convertor' },
                        { text: 'cryptor', link: '/en/api/packages/cryptor' },
                        {
                            text: 'datastructure',
                            collapsed: true,
                            items: [
                                { text: 'list', link: '/en/api/packages/datastructure/list' },
                                { text: 'safelist', link: '/en/api/packages/datastructure/copyonwritelist' },
                                { text: 'link', link: '/en/api/packages/datastructure/link' },
                                { text: 'stack', link: '/en/api/packages/datastructure/stack' },
                                { text: 'queue', link: '/en/api/packages/datastructure/queue' },
                                { text: 'heap', link: '/en/api/packages/datastructure/heap' },
                                { text: 'tree', link: '/en/api/packages/datastructure/tree' },
                                { text: 'set', link: '/en/api/packages/datastructure/set' },
                                { text: 'hashmap', link: '/en/api/packages/datastructure/hashmap' },
                            ],
                        },
                        { text: 'datetime', link: '/en/api/packages/datetime' },
                        { text: 'fileutil', link: '/en/api/packages/fileutil' },
                        { text: 'formatter', link: '/en/api/packages/formatter' },
                        { text: 'function', link: '/en/api/packages/function' },
                        { text: 'mathutil', link: '/en/api/packages/mathutil' },
                        { text: 'maputil', link: '/en/api/packages/maputil' },
                        { text: 'netutil', link: '/en/api/packages/netutil' },
                        { text: 'pointer', link: '/en/api/packages/pointer' },
                        { text: 'random', link: '/en/api/packages/random' },
                        { text: 'retry', link: '/en/api/packages/retry' },
                        { text: 'slice', link: '/en/api/packages/slice' },
                        { text: 'stream', link: '/en/api/packages/stream' },
                        { text: 'struct', link: '/en/api/packages/struct' },
                        { text: 'strutil', link: '/en/api/packages/strutil' },
                    ],
                },
            ],
        },
    },
}
