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
			social: {
				github: 'https://github.com/samuelireson/courses'
			},
			sidebar: [
				//{
				//	label: 'Foundations',
				//	items: [
				//		{ label: 'Introduction', slug: 'foundations/introduction' },
				//		{ label: 'Logic', slug: 'foundations/logic' },
				//		{ label: 'Set Theory', slug: 'foundations/set-theory' },
				//		{ label: 'Real Analysis', slug: 'foundations/real-analysis' },
				//	],
				//},
				{
					label: 'Machine Learning',
					items: [
						{ label: 'Introduction', slug: 'mlnn/introduction' },
						{ label: 'Linear Algebra', slug: 'mlnn/linear-algebra' },
						{ label: 'Multivariable Calculus', slug: 'mlnn/multivariable-calculus' },
						{ label: 'Probability', slug: 'mlnn/probability' },
						{ label: 'Optimisation', slug: 'mlnn/optimisation' },
						{ label: 'Linear Regression', slug: 'mlnn/linear-regression' },
						{ label: 'Logistic Regression', slug: 'mlnn/logistic-regression' },
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
