// @ts-check
import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';
import rehypeKatex from 'rehype-katex';
import remarkMath from 'remark-math';

// https://astro.build/config
export default defineConfig({
	integrations: [
		starlight({
			title: 'Courses',
			sidebar: [
				{
					label: 'Machine Learning',
					items: [
						// Each item here is one entry in the navigation menu.
						{ label: 'Introduction', slug: 'mlnn/introduction' },
						{ label: 'Linear Algebra', slug: 'mlnn/linear-algebra' },
					],
				},
			],
			customCss: [
				'./node_modules/katex/dist/katex.min.css',
				'./src/styles/props.css'
			]
		}),
	],
	markdown: {
		remarkPlugins: [remarkMath],
		rehypePlugins: [rehypeKatex],
	}
});
