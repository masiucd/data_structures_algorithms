/** @type {import('next').NextConfig} */
const nextConfig = {
  swcMinify: true,
  reactStrictMode: true,
  experimental: {
    serverActions: true,
  },
};

module.exports = nextConfig;
