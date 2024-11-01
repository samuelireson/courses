// @ts-check
import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';
import rehypeKatex from 'rehype-katex';
import remarkMath from 'remark-math';
import starlightHeadingBadgesPlugin from 'starlight-heading-badges';
import react from '@astrojs/react';

// https://astro.build/config
export default defineConfig({
	site: 'https://samuelireson.github.io',
	base: '/courses',
	integrations: [starlight({
		plugins: [starlightHeadingBadgesPlugin()],
		title: 'Courses',
		social: {
			github: 'https://github.com/samuelireson/courses'
		},
		sidebar: [
			{
				label: 'Machine Learning',
				items: [
					{ label: 'Introduction', slug: 'mlnn/chapters/introduction' },
					{ label: 'Linear Algebra', slug: 'mlnn/chapters/linear-algebra' },
					{ label: 'Multivariable Calculus', slug: 'mlnn/chapters/multivariable-calculus' },
					{ label: 'Probability', slug: 'mlnn/chapters/probability' },
					{ label: 'Learning Problems', slug: 'mlnn/chapters/learning-problems' },
					{ label: 'Optimisation', slug: 'mlnn/chapters/optimisation' },
					{ label: 'Linear Regression', slug: 'mlnn/chapters/linear-regression' },
					{ label: 'Logistic Regression', slug: 'mlnn/chapters/logistic-regression' },
				],
			},
		],
		customCss: [
			'./node_modules/katex/dist/katex.css',
			'./src/styles/custom.css',
			'./src/fonts/font-face.css'
		]
	}), react()],
	markdown: {
		remarkPlugins: [remarkMath],
		rehypePlugins: [rehypeKatex],
	}
});
