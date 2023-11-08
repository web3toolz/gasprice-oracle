/** @type {import('next').NextConfig} */
const nextConfig = {
    output: "export",

    experimental: {
        optimizePackageImports: ['@mantine/core', '@mantine/hooks'],
    },
}

module.exports = nextConfig

