import type { DefaultTheme, LocaleSpecificConfig } from 'vitepress'

export const META_URL = 'https://www.golancet.cn/en/'
export const META_TITLE = 'Lancet'
export const META_DESCRIPTION = 'A powerful util function library of Go'

export const enConfig: LocaleSpecificConfig<DefaultTheme.Config> = {
    description: META_DESCRIPTION,

    head: [
        ['meta', { property: 'og:url', content: META_URL }],
        ['meta', { property: 'og:description', content: META_DESCRIPTION }],
    ],

    themeConfig: {
        editLink: {
            pattern: 'https://github.com/duke-git/lancet/edit/v2/docs/:path',
            text: 'Suggest changes to this page',
        },
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
                        text: 'Discussion',
                        link: 'https://github.com/duke-git/lancet/discussions',
                    },
                    {
                        text: 'Changelog',
                        link: 'https://github.com/duke-git/lancet/releases',
                    },
                    {
                        text: 'Contribution',
                        link: 'https://github.com/duke-git/lancet/blob/main/CONTRIBUTION.md',
                    },
                ],
            },
        ],

        sidebar: {
            '/en/guide/': [
                {
                    text: 'Introduction',
                    collapsed: false,
                    items: [
                        {
                            text: 'What is Lancet？',
                            link: '/en/guide/introduction',
                        },
                        {
                            text: 'Getting started',
                            link: '/en/guide/getting_started',
                        },
                    ],
                },
                {
                    text: 'Contribute Code',
                    collapsed: false,
                    items: [
                        {
                            text: 'Contribution guide',
                            link: '/en/guide/contribution_guide',
                        },
                        {
                            text: 'Contributors',
                            link: '/en/guide/contributors',
                        },
                    ],
                },
                {
                    text: 'API Reference',
                    link: '/en/api/overview'
                },
            ],
            '/en/api/': [
                {
                    text: 'Overview',
                    items: [{ text: 'API overview', link: '/en/api/overview' }],
                },
                {
                    text: 'Packages',
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
                        { text: 'enum', link: '/en/api/packages/enum' },
                        { text: 'eventbus', link: '/en/api/packages/eventbus' },
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
                        { text: 'tuple', link: '/en/api/packages/tuple' },
                        { text: 'validator', link: '/en/api/packages/validator' },
                        { text: 'system', link: '/en/api/packages/system' },
                        { text: 'xerror', link: '/en/api/packages/xerror' },
                    ],
                },
            ],
        },
    },
}
