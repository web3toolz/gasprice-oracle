/** @type {import('next').NextConfig} */
const nextConfig = {
    output: "export",

    images: {
        loader: 'akamai',
        path: '',
    },

    assetPrefix: '/gasprice-oracle/',
    basePath: '/gasprice-oracle',

    experimental: {
        optimizePackageImports: ['@mantine/core', '@mantine/hooks'],
    },
}

module.exports = nextConfig

