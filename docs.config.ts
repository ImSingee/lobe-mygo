export default {
  atomDirs: [{ dir: 'catalog' }],
  homePage: './docs/home/home.tsx',
  publicDocs: ['docs/getting-started.mdx'],
  description: 'Lobe UI components for MyGo desktop apps, drawn on the GPU in pure Go.',
  siteUrl: 'https://mygo.lobehub.com',
  themeConfig: {
    apiHeader: {
      github: 'https://github.com/lobehub/lobe-mygo',
    },
    metadata: {
      openGraph: {
        image: 'https://mygo.lobehub.com/og.png',
      },
    },
    navItems: [{ href: '/getting-started', icon: 'rocket', label: 'Getting Started' }],
    prefersColor: 'auto',
    socialLinks: [{ href: 'https://github.com/lobehub/lobe-mygo', icon: 'github', label: 'GitHub' }],
  },
  title: 'Lobe MyGo',
};
