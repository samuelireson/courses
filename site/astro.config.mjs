// @ts-check
import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';
import rehypeKatex from 'rehype-katex';
import remarkMath from 'remark-math';
import starlightHeadingBadgesPlugin from 'starlight-heading-badges';

// https://astro.build/config
export default defineConfig({
	integrations: [
		starlight({
			plugins: [starlightHeadingBadgesPlugin()],
			title: 'Courses',
			sidebar: [
				{
					label: 'Machine Learning',
					items: [
						{ label: 'Introduction', slug: 'mlnn/introduction' },
					],
				},
			],
			customCss: [
				'./node_modules/katex/dist/katex.css',
				'./src/styles/custom.css',
				'./src/fonts/font-face.css'
			]
		}),
	],
	markdown: {
		remarkPlugins: [remarkMath],
		rehypePlugins: [rehypeKatex],
	}
});
