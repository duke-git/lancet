import { defineConfig, HeadConfig } from 'vitepress'

export const META_IMAGE = '/lancet_logo.png'
export const isProduction = process.env.NETLIFY && process.env.CONTEXT === 'production'

if (process.env.NETLIFY) {
    console.log('Netlify build', process.env.CONTEXT)
}

const productionHead: HeadConfig[] = [
    [
        'script',
        {
            src: 'https://unpkg.com/thesemetrics@latest',
            async: '',
            type: 'text/javascript',
        },
    ],
]

const rControl = /[\u0000-\u001f]/g
const rSpecial = /[\s~`!@#$%^&*()\-_+=[\]{}|\\;:"'“”‘’<>,.?/]+/g
const rCombining = /[\u0300-\u036F]/g

/**
 * Default slugification function
 */
export const slugify = (str: string): string =>
    str
        .normalize('NFKD')
        // Remove accents
        .replace(rCombining, '')
        // Remove control characters
        .replace(rControl, '')
        // Replace special characters
        .replace(rSpecial, '-')
        // ensure it doesn't start with a number
        .replace(/^(\d)/, '_$1')

export const commonConfig = defineConfig({
    title: 'Lancet',
    appearance: true,
    ignoreDeadLinks: true,

    markdown: {
        theme: {
            dark: 'dracula-soft',
            light: 'vitesse-light',
        },

        attrs: {
            leftDelimiter: '%{',
            rightDelimiter: '}%',
        },

        anchor: {
            slugify,
        },
    },

    head: [
        // ['link', { rel: 'icon', type: 'image/svg+xml', href: '/logo.svg' }],
        ['link', { rel: 'icon', type: 'image/png', href: '/lancet_logo_mini.png' }],
        ['meta', { name: 'theme-color', content: '#5f67ee' }],
        ['meta', { name: 'og:type', content: 'website' }],
        ['meta', { name: 'og:locale', content: 'zh' }],

        ...(isProduction ? productionHead : []),
    ],

    themeConfig: {
        logo: { src: '/lancet_logo_mini.png', width: 24, height: 24 },
        outline: [2, 3],

        search: {
            provider: 'local',
        },
        socialLinks: [
            {
                icon: 'github',
                link: 'https://github.com/duke-git/lancet',
            },
        ],

        footer: {
            copyright: 'Copyright © 2023-present Duke Du',
            message: '备案号: 京ICP备2023022770号',
        },
    },
})
