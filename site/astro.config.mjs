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
					autogenerate: { directory: 'mlnn' },
				},
			],
			customCss: [
				'./src/styles/props.css'
			]
		}),
	],
	markdown: {
		remarkPlugins: [remarkMath],
		rehypePlugins: [rehypeKatex],
	}
});
